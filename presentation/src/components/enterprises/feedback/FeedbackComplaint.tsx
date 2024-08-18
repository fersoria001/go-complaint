'use client'
import ComplaintReplyBubble from "@/components/complaints/ComplaintReplyBubble"
import { ComplaintReply, Feedback, ReplyReview, User } from "@/gql/graphql"
import { useParams } from "next/navigation"
import { useEffect, useRef, useState } from "react"
import {
    addToMap, ColorAlreadyPickedError, ColorsFixedError, ColorSquare,
    CommentObject,
    commentsOverlapY,
    defaultSquares,
    FeedbackErrorType,
    load,
    removeReplyHighlight,
    replyAlreadyExists,
    selectAColor,
} from "./feedback"
import { useMutation, useSuspenseQueries } from "@tanstack/react-query"
import getGraphQLClient from "@/graphql/graphQLClient"
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery"
import ColorSelectorItem from "./ColorSelectorItem"
import FeedbackComment from "./FeedbackComment"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import createFeedbackMutation from "@/graphql/mutations/createFeedbackMutation"
import graphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"
import { addFeedbackComment, addFeedbackReply, endFeedback, removeFeedbackComment, removeFeedbackReply } from "@/lib/actions/graphqlActions"
import feedbackSubscription from "@/graphql/subscriptions/feedbackSubscription"
import getGraphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"



const FeedbackComplaint: React.FC = () => {
    const { enterpriseId, complaintId } = useParams()
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    const gqlClient = getGraphQLClient()
    const [
        { data: { createFeedback: feedback } },
        { data: { complaintById: complaint } },
        { data: { userDescriptor: currentReviewer } }
    ] = useSuspenseQueries({
        queries: [
            {
                queryKey: ["create-feedback-mutation", enterpriseName, complaintId],
                queryFn: async () => gqlClient.request(createFeedbackMutation, {
                    input: { complaintId: complaintId as string, enterpriseId: enterpriseName }
                })
            },
            {
                queryKey: ["complaint-by-id", complaintId],
                queryFn: async () => gqlClient.request(complaintByIdQuery, { id: complaintId as string })
            },
            {
                queryKey: ["user-descriptor"],
                queryFn: async () => gqlClient.request(userDescriptorQuery),
                staleTime: Infinity,
                gcTime: Infinity
            }
        ]
    })
    const [colorSquares, setColorSquares] = useState<ColorSquare[]>(defaultSquares);
    const [selectedColor, setSelectedColor] = useState<string | undefined>()
    const [comments, setComments] = useState<CommentObject[]>([]);
    const [commentsMap, setCommentsMap] = useState<Map<string, React.RefObject<HTMLFormElement>>>(new Map());
    const [repliesMap, setRepliesMap] = useState<Map<string, React.RefObject<HTMLLIElement>>>(new Map());
    const [errors, setErrors] = useState<FeedbackErrorType>({})
    const windowRef = useRef<HTMLDivElement>(null)

    const registerReplyPosition = (key: string, ref: React.RefObject<HTMLLIElement>) => {
        setRepliesMap(prev => {
            prev.set(key, ref)
            return new Map(prev)
        });
    }

    const registerCommentPosition = (commentObject: CommentObject, commentRef: React.RefObject<any>) => {
        setCommentsMap(prev => {
            return addToMap(prev, commentObject, commentRef)
        });
    }

    const fixSquares = (v: boolean) => {
        const newSquares = colorSquares.map((square) => {
            return { ...square, fixed: v }
        })
        setColorSquares(newSquares);
    }

    const handleColorSelect = (selected: string) => {
        if (feedback.isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        try {
            setErrors({});
            const newSquares = selectAColor(selected, colorSquares)
            setColorSquares(newSquares);
            setSelectedColor(selected)
        } catch (e: any) {
            if (e instanceof ColorAlreadyPickedError) {
                const err: any = {}
                err["colors"] = e.message
                setErrors({ square: err })
                return
            }
            if (e instanceof ColorsFixedError) {
                const err: any = {}
                err[selected] = e.message
                setErrors({ form: err })
                return
            }
        }
    }

    const handleAddReply = async (reply: ComplaintReply) => {
        if (feedback.isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        if (!selectedColor) {
            return setErrors({ feedback: { all: "error: pick a color first" } });
        }
        const comment = comments.find((c) => c.color === selectedColor)
        if (comment) {
            if (replyAlreadyExists(comment, reply)) {
                const commentReplyRef = comment.repliesRefs.get(reply.id!)
                if (commentReplyRef) {
                    removeReplyHighlight(commentReplyRef)
                }
                return await removeFeedbackReply({
                    color: comment.color,
                    feedbackId: feedback.id,
                    repliesIds: [reply.id!]
                })
            }
        }
        return await addFeedbackReply({
            color: selectedColor,
            feedbackId: feedback.id,
            repliesIds: [reply.id!],
            reviewerId: currentReviewer.id
        })
    }

    const handleAddComment = async (itemKey: string, comment: string) => {
        await addFeedbackComment({
            color: itemKey,
            comment: comment,
            feedbackId: feedback.id
        })
        fixSquares(false)
        setSelectedColor("")
    }

    const handleEditComment = (itemKey: string) => {
        if (feedback.isDone) {
            setErrors({ feedback: { all: "feedback is already submitted" } })
        }
        fixSquares(true)
        setSelectedColor(itemKey)
    }

    const handleDeleteComment = async (itemKey: string) => {
        await removeFeedbackComment({
            color: itemKey,
            feedbackId: feedback.id,
        })
        const comment = comments.find((c) => c.color === itemKey)
        if (comment) {
            comment.repliesRefs.forEach((r) => removeReplyHighlight(r))
        }
    }

    const submitFeedbackMutation = useMutation({
        mutationFn: () => endFeedback({
            feedbackId: feedback.id,
            reviewerId: currentReviewer.id
        })
    })

    useEffect(() => {

        async function subscribe() {
            const subscription = getGraphQLSubscriptionClient().iterate({
                query: feedbackSubscription(feedback.id),
            });
            for await (const event of subscription) {
                const f = event.data?.feedback as Feedback
                const user: Partial<User> = {
                    userName: currentReviewer.userName
                }
                const { newComments, newSquares } = load(
                    user,
                    f.replyReview as ReplyReview[],
                    defaultSquares,
                    repliesMap);
                setComments(newComments);
                setColorSquares(newSquares);
            }
        }
        const qty = complaint.replies?.filter((r) => r?.isEnterprise == true)
        if (repliesMap.size == qty?.length) {
            subscribe()
        }
    }, [feedback.id, repliesMap.size])

    useEffect(() => {
        commentsOverlapY(commentsMap, windowRef)
    }, [commentsMap, commentsMap.size])

    return (
        <div id="feedback-page"
            ref={windowRef}
            className="relative min-h-full">
            <div className="border border-b flex align-center px-2 py-5 sm:py-3 sm:px-5">
                <span
                    onClick={() => submitFeedbackMutation.mutate()}
                    className={`${feedback.isDone ? 'opacity-50 pointer-events-none' : ''}`}>
                    <button className="bg-blue-500 px-7 py-1 text-white font-bold rounded-md hover:bg-blue-600">
                        Submit Feedback
                    </button>
                </span>
                {
                    errors.feedback &&
                    <span className="text-red-500 mt-[19px] ml-5 sm:left-0 sm:inset-x-1/3 sm:top-12 ">
                        {errors.feedback.all}
                    </span>
                }
                <div className="flex self-center justify-end ms-auto sm:relative sm:min-w-[200px]">
                    {
                        colorSquares.map((square, index) => {
                            return <ColorSelectorItem
                                key={index}
                                isDone={feedback.isDone}
                                colorSquare={square}
                                callback={handleColorSelect}
                                errors={errors}
                            />
                        })
                    }
                    {
                        errors.square && errors.square["colors"] &&
                        <span className="absolute left-0 text-red-500 whitespace-nowrap">{errors.square["colors"]}
                        </span>
                    }
                </div>
            </div>
            {
                comments.map((comment, _) => {
                    return <FeedbackComment
                        key={comment.color}
                        commentObject={comment}
                        registerPosition={registerCommentPosition}
                        reviewer={currentReviewer}
                        isDone={feedback.isDone}
                        logErrors={(e: FeedbackErrorType) => setErrors(e)}
                        addComment={handleAddComment}
                        editComment={handleEditComment}
                        deleteComment={handleDeleteComment}

                    />
                })
            }
            <ul
                id="feedback-bubble-list"
                className="flex flex-col-reverse gap-2.5 p-2.5 w-full overflow-y-auto bg-white">
                {
                    complaint.replies!.length > 0 && complaint.replies!.map((msg, index) => {
                        return <ComplaintReplyBubble
                            key={msg?.id}
                            reply={msg as ComplaintReply}
                            enterpriseName={enterpriseName}
                            callback={msg?.isEnterprise ? registerReplyPosition : undefined}
                            onClick={msg?.isEnterprise ? () => { handleAddReply(msg!) } : undefined}
                        />
                    })
                }
            </ul>
        </div >
    )
}

export default FeedbackComplaint
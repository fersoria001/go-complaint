/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/no-unused-vars */
import { useEffect, useRef, useState } from "react";
import { FeedbackType, Reply } from "../../../lib/types";
import PrimaryButton from "../../buttons/PrimaryButton";
import ChatBubble from "../../complaint/ChatBubble";
import ReplyBubble from "../../feedback/ReplyBubble";
import { Route } from "../../../routes/$enterpriseID/feedbacks/";
import SquareFillIcon, { ColorSquare } from "../../icons/SquareFillIcon";
import CommentForm, { CommentObject } from "../../feedback/CommentForm";
import {
    addToMap,
    ColorAlreadyPickedError, ColorsFixedError, createComment,
    FeedbackErrorType, generateSquares, load, markCommentAsDone, positionComment, removeComment,
    removeReply,
    replyAlreadyExists, selectAColor,
} from "../../feedback/feedback";
import { deleteComment, removeNewReply, saveNewComment, saveNewReply, setCommentBody, submitFeedback } from "../../feedback/async";
import useSubscriber from "../../../lib/hooks/useSubscriber";
import { useRouter } from "@tanstack/react-router";
const EmployeeFeedback: React.FC = () => {
    const { subscriptionID, reviewer, complaint, feedback } = Route.useLoaderData()
    const { enterpriseID } = Route.useParams();
    const [color, setColor] = useState<string>("");
    const [colorSquares, setColorSquares] = useState<ColorSquare[]>(generateSquares);
    const [comments, setComments] = useState<CommentObject[]>([]);
    const [commentsMap, setCommentsMap] = useState<Map<string, React.RefObject<any>>>(new Map());
    const [repliesMap, setRepliesMap] = useState<Map<string, React.RefObject<HTMLLIElement>>>(new Map());
    const [bubblesMap, setBubblesMap] = useState<Map<string, React.RefObject<HTMLLIElement>>>(new Map());
    const [errors, setErrors] = useState<FeedbackErrorType>({});
    const { incomingMessage } = useSubscriber(subscriptionID);
    const [isDone, setIsDone] = useState<boolean>(feedback?.isDone || false);
    const feedbackRef = useRef<HTMLDivElement>(null);
    const router = useRouter();
    const fixSquares = (v: boolean) => {
        if (isDone) return
        const newSquares = colorSquares.map((square) => {
            return { ...square, fixed: v }
        })
        setColorSquares(newSquares);
    }

    const registerAllReplies = (id: string, ref: React.RefObject<HTMLLIElement>) => {
        if (bubblesMap.has(id)) return
        const newMap = new Map(bubblesMap.entries());
        newMap.set(id, ref);
        setBubblesMap(newMap);
    }

    const registerReplyPosition = (key: string, ref: React.RefObject<HTMLLIElement>) => {
        if (repliesMap.has(key)) return
        setRepliesMap(prev => {
            return new Map([...prev, [key, ref]])
        });
    }

    const registerCommentPosition = (commentObject: CommentObject, commentRef: React.RefObject<any>) => {
        const index = comments.findIndex((comment) => comment.color == color);
        if (index === -1) {
            console.error("comment not found in addNewComment")
        }
        setCommentsMap(prev => {
            return addToMap(prev, commentObject, commentRef)
        });
    }


    const selectColor = (selected: string): void => {
        if (isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        try {
            setErrors({});
            const newSquares = selectAColor(selected, colorSquares)
            setColorSquares(newSquares);
            setColor(selected);
        } catch (e: any) {
            if (e instanceof ColorAlreadyPickedError) {
                const err: any = {}
                err[selected] = e.message
                setErrors({ square: err })
                return
            }
            if (e instanceof ColorsFixedError) {
                const err: any = {}
                err[color] = e.message
                setErrors({ form: err })
                return
            }
        }
    }

    const editComment = (colorKey: string): void => {
        if (isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        const newComments = comments.map((comment) => {
            if (comment.color == colorKey) {
                return { ...comment, done: false }
            }
            return comment
        })
        fixSquares(true)
        setColor(colorKey);
        setComments(newComments);
    }


    const addToFeedback = async (colorKey: string, body: string): Promise<boolean> => {
        if (isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return false
        }
        try {
            setErrors({});
            const newComments = markCommentAsDone(comments, colorKey, body);
            const comment = newComments.find((comment) => comment.color == colorKey);
            if (!comment) {
                console.error("comment not found in addToFeedback")
                return false
            }
            const ok = await setCommentBody(comment, enterpriseID, feedback.id)
            if (!ok) {
                console.error("not ok")
                return ok
            }
            setComments(newComments);
            setColor("");
            fixSquares(!ok);
            return ok
        } catch (e: any) {
            console.error("add to feedback error", e.message)
            const err: any = {}
            err[colorKey] = e.message
            setErrors({ form: err })
            return false
        }
    }


    const deleteFromFeedback = async (itemKey: string) => {
        if (isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        const comment = comments.find((comment) => comment.color == itemKey);
        if (!comment) {
            console.error("comment not found in deleteFromFeedback")
            return
        }
        try {
            const { objects, refs } = removeComment(comments, commentsMap, itemKey);
            const ok = await deleteComment(comment!, enterpriseID, feedback.id)
            if (!ok) {
                console.error("delete comment error")
                return
            }
            setComments(objects);
            setCommentsMap(refs);
            fixSquares(false);
        } catch (e: any) {
            console.error("delete from feedback error", e.message)
        }
    }


    const selectAReply = async (reply: Reply, ref: React.RefObject<HTMLLIElement>) => {
        if (isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        setErrors({});
        if (!color || color === "") {
            setErrors({ feedback: { all: "error: pick a color first" } });
            return
        }
        const index = comments.findIndex((comment) => comment.color == color);
        if (index === -1) {
            const newComment = createComment(color, reviewer);
            if (!feedback) {
                await saveNewComment(newComment, complaint.id, enterpriseID)
                router.invalidate()
            } else {
                await saveNewReply(reply, newComment, enterpriseID, feedback.id, reviewer.email)
                const value = commentsMap.get(color);
                if (value) {
                    const p2 = positionComment(value, Array.from(comments[index].repliesRefs.values()))
                    value.current.style.top = `${p2.top}px`
                    value.current.style.bottom = `${p2.bottom}px`
                }
            }
        } else {
            const exists = replyAlreadyExists(comments[index], reply);
            if (!exists) {
                await saveNewReply(reply, comments[index], enterpriseID, feedback.id, reviewer.email)
                const value = commentsMap.get(color);
                if (value) {
                    const p2 = positionComment(value, Array.from(comments[index].repliesRefs.values()))
                    value.current.style.top = `${p2.top}px`
                    value.current.style.bottom = `${p2.bottom}px`
                }
                fixSquares(true);
            } else {
                if (comments[index].selectedReplies.length > 0) {
                    removeReply(ref)
                    await removeNewReply(reply, comments[index], enterpriseID, feedback.id)
                    const value = commentsMap.get(color);
                    if (value) {
                        const p2 = positionComment(value, Array.from(comments[index].repliesRefs.values()))
                        value.current.style.top = `${p2.top}px`
                        value.current.style.bottom = `${p2.bottom}px`
                    }
                } else {
                    await deleteComment(comments[index], enterpriseID, feedback.id)
                    fixSquares(false);
                }
            }
        }
    }

    const handleSubmit = async () => {
        if (isDone) {
            const error = { feedback: { all: "feedback is already submitted" } }
            setErrors(error)
            return
        }
        try {
            await submitFeedback(enterpriseID, feedback.id)
        } catch (e: any) {
            console.error("submit feedback error", e.message)
            const error = { feedback: { all: "complete all 3 reviews first" } }
            setErrors(error)
        }
    }

    useEffect(() => {
        const keys = Array.from(commentsMap.keys());
        if (keys.length === 0) return
        const commentRef = commentsMap.get(keys[keys.length - 1]);
        if (!commentRef) {
            console.error("commentRef not found in useEffect")
            return
        }
    }, [commentsMap.size])

    useEffect(() => {
        if (!incomingMessage) return
        try {
            const wsf = incomingMessage as FeedbackType;
            setIsDone(wsf.isDone);
            const { newComments, newSquares, actualColor } = load(
                reviewer,
                wsf.replyReview,
                colorSquares,
                repliesMap,);
            setColorSquares(newSquares);
            setColor(actualColor);
            setComments(newComments);
            setColorSquares(newSquares);
        } catch (e: any) {
            console.error("load feedback from incoming message error", e.message)
        }
    }, [incomingMessage])
    useEffect(() => {
        const filtered = complaint.replies!.filter((reply) => reply.isEnterprise == true);
        if (feedback && repliesMap.size === filtered.length - 1) {
            try {
                const { newComments, newSquares, actualColor } = load(
                    reviewer,
                    feedback.replyReview,
                    colorSquares,
                    repliesMap,);
                setIsDone(feedback.isDone);
                setColorSquares(newSquares);
                setColor(actualColor);
                setComments(newComments);
                setColorSquares(newSquares);
            } catch (e: any) {
                console.error("load feedback error", e.message)
            }
        }
    }, [repliesMap.size])


    useEffect(() => {
        console.log(feedback)
        console.log("isDone?", isDone)
    })

    return (
        <div
            ref={feedbackRef}
            id="feedback-page" className="relative min-h-full">
            {errors.feedback &&
                <span className="absolute text-red-500 mt-[19px] ml-5
                 sm:left-0 sm:inset-x-1/3  sm:top-12 ">
                    {errors.feedback.all}
                </span>
            }
            <div className="border border-b flex align-center  px-2 py-5 sm:py-3 sm:px-5 ">
                <span onClick={handleSubmit} className={`${isDone ? 'opacity-50 pointer-events-none' : ''}`}>
                    <PrimaryButton text="Submit Feedback" />
                </span>
                <div className="flex self-center justify-end ms-auto  sm:relative   sm:min-w-[200px] 
               ">
                    {
                        colorSquares.map((square, index) => {
                            return <SquareFillIcon
                                end={feedback.isDone}
                                errors={errors}
                                key={index}
                                colorSquare={square}
                                callback={selectColor}
                            />
                        })
                    }

                </div>
            </div>
            {
                comments.map((comment, _) => {
                    return <CommentForm
                        end={isDone}
                        reviewer={reviewer}
                        key={comment.color}
                        registerPosition={registerCommentPosition}
                        errors={errors}
                        commentObject={comment}
                        addComment={addToFeedback}
                        editComment={editComment}
                        deleteComment={deleteFromFeedback} />
                })
            }
            <ul id="feedback-bubble-list" className="flex flex-col w-full bg-white">
                {complaint.replies!.length > 0 && complaint.replies!.map((msg, index) => {
                    if (index != complaint.replies!.length - 1) {
                        if (msg.senderID == complaint.authorID) {
                            return (
                                <ChatBubble
                                    registerPosition={registerAllReplies}
                                    key={msg.id}
                                    direction={msg.senderID == complaint.authorID ? "ltr" : "rtl"}
                                    fullName={msg.senderName}
                                    profileIMG={msg.senderIMG}
                                    body={msg.body}
                                    createdAt={msg.createdAt}
                                    seen={msg.read}
                                    seenAt={msg.readAt}
                                    isEnterprise={false}
                                    enterpriseID=""
                                />
                            )
                        } else {
                            return (
                                <ReplyBubble
                                    registerPosition={registerReplyPosition}
                                    key={msg.id}
                                    callback={selectAReply}
                                    direction={msg.senderID == complaint.authorID ? "ltr" : "rtl"}
                                    msg={msg}
                                    end={feedback.isDone}
                                />
                            )
                        }
                    }
                })}
            </ul>
        </div >
    )
}

export default EmployeeFeedback;
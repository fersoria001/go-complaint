/* eslint-disable react-hooks/exhaustive-deps */
'use client'

import DeleteIcon from "@/components/icons/DeleteIcon";
import EditIcon from "@/components/icons/EditIcon";
import { UserDescriptor } from "@/gql/graphql";
import { useState, useRef, useEffect } from "react";
import { CommentObject, FeedbackErrorType, removeReplyHighlight } from "./feedback";
import { useMutation } from "@tanstack/react-query";
import { addFeedbackComment, removeFeedbackComment } from "@/lib/actions/graphqlActions";
import MinimizeIcon from "@/components/icons/MinimizeIcon";
import clsx from "clsx";
import ExpandIcon from "@/components/icons/ExpandIcon";

interface Props {
    reviewer: UserDescriptor;
    commentObject: CommentObject;
    addComment: (itemKey: string, comment: string) => void;
    editComment: (itemKey: string) => void;
    deleteComment: (itemKey: string) => void;
    callback?: (...args: any[]) => void;
    registerPosition: (commentObject: CommentObject, ref: React.RefObject<any>) => void;
    logErrors: (e: FeedbackErrorType) => void
    isDone: boolean;
}

const FeedbackComment: React.FC<Props> = ({
    registerPosition, commentObject,
    reviewer, isDone, addComment, editComment, deleteComment,
    logErrors }) => {
    const [editable, setEditable] = useState(commentObject.comment === "");
    const [confirmation, setConfirmation] = useState(false);
    const [comment, setComment] = useState(commentObject.comment);
    const [minimized, setMinimized] = useState<boolean>(false)
    const commentRef = useRef<HTMLFormElement>(null);

    const handleAddComment = async () => {
        if (isDone) {
            return logErrors({ feedback: { all: "feedback is already submitted" } })
        }
        addComment(commentObject.color, comment)
        setEditable(false)
    }
    const handleDelete = async () => {
        deleteComment(commentObject.color)
    }

    const handleDeleteConfirmation = () => {
        if (isDone) {
            return logErrors({ feedback: { all: "feedback is already submitted" } })
        }
        setConfirmation(true)
    }

    const handleEditComment = () => {
        if (isDone) {
            return logErrors({ feedback: { all: "feedback is already submitted" } })
        }
        editComment(commentObject.color)
        setEditable(true);
    }

    useEffect(() => {
        registerPosition(commentObject, commentRef)
    }, [])
    return (
        <>
            <form
                ref={commentRef}
                style={{ borderColor: commentObject.color }}
                className={clsx(`z-10 absolute xl:left-[300px] h-[270px] w-1/2 xl:w-[320px] bg-white  shadow rounded-md border ${commentObject.color}`, {
                    "h-[48px]": minimized
                })}>
                <div className="flex flex-col h-full border-t border-gray-200 rounded-lg bg-gray-50">
                    <div className="border-b p-2 flex justify-between">
                        <p className="text-sm md:text-xl text-gray-700 font-bold">
                            {reviewer.fullName}
                        </p>
                        {minimized ?
                            <ExpandIcon className="w-6 h-6 cursor-pointer" onClick={() => setMinimized(false)} /> :
                            <MinimizeIcon className="w-6 h-6 cursor-pointer" onClick={() => setMinimized(true)} />}
                    </div>
                    {!minimized &&
                        <>
                            <div className="px-2 bg-white rounded-t-lg grow">
                                <textarea
                                    id="comment"
                                    value={comment}
                                    maxLength={500}
                                    onChange={(e) => setComment(e.target.value)}
                                    rows={4}
                                    className={`h-full w-full px-0 text-sm text-gray-900
                     bg-white border-0 resize-none focus:ring-0 focus:outline-none`}
                                    placeholder="Write a review about this specific answer"
                                    disabled={editable ? false : true}
                                    required >
                                </textarea>
                            </div>

                            <div className="flex items-center justify-between px-3 py-2 border-t relative">
                                <button
                                    onClick={handleAddComment}
                                    type="button"
                                    className={
                                        !editable ?
                                            `opacity-50 cursor-not-allowed relative inline-flex items-center justify-center p-0.5 mb-2 me-2
                                overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500` :
                                            `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900
                                rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500
                            group-hover:from-cyan-500 group-hover:to-blue-500 focus:ring-4 focus:outline-none focus:ring-cyan-200` }
                                    disabled={editable ? false : true}>
                                    <span className={
                                        comment != "" ?
                                            `relative px-5 py-2.5 transition-all ease-in duration-75 bg-white rounded-md` :
                                            `relative px-5 py-2.5 transition-all ease-in duration-75 bg-white rounded-md group-hover:bg-opacity-0`}>
                                        {"Add to feedback"}
                                    </span>
                                </button>
                                <span
                                    onMouseUp={handleEditComment}
                                    className={`${false ? 'cursor-default' : 'cursor-pointer'}   mb-2 shadow shadow-gray-300 hover:bg-gray-100`/* end */}>
                                    <EditIcon fill="#374151" className="w-6 h-6" />
                                </span>
                                <span className={`${false ? 'cursor-default' : 'cursor-pointer'}`} onMouseUp={handleDeleteConfirmation}>
                                    <DeleteIcon fill="#335784" className="w-6 h-6" />
                                </span>
                                {/* {errors.form && errors.form[commentObject.color] &&
                            <span className={`bg-white border absolute top-[74px] left-0 py-2 px-3 rounded-md `}>
                                <p className="text-red-500 italic text-sm md:text-md">{errors.form[commentObject.color]}</p>
                            </span>} */}
                                {confirmation &&
                                    <div className={`absolute flex flex-col justify-center align-items top-1/2 right-11 md:left-28 w-[260px] p-3
                                        z-50 bg-white border border-gray-100 shadow rounded-md`}>
                                        <p className="text-red-500 text-xs mb-3">
                                            Are you sure you want to delete this comment?
                                        </p>
                                        <div className="flex">
                                            <button
                                                type="button"
                                                className="bg-blue-500 hover:bg-blue-600 px-7 py-1 text-white font-bold rounded-md"
                                                onClick={() => handleDelete()}>
                                                Delete
                                            </button>
                                            <button
                                                type="button"
                                                className="bg-blue-500 hover:bg-blue-600 px-7 py-1 text-white font-bold rounded-md"
                                                onClick={() => setConfirmation(false)}>
                                                Cancel
                                            </button>
                                        </div>
                                    </div>}
                            </div>
                        </>
                    }
                </div>
            </form >
        </>
    )
}

export default FeedbackComment;
/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useEffect, useRef, useState } from "react";
import DeleteIcon from "../icons/DeleteIcon";
import EditIcon from "../icons/EditIcon";
import ThinBtn from "../buttons/ThinBtn";
import { FeedbackErrorType } from "./feedback";
import { Reply, User } from "../../lib/types";
export type CommentObject = {
    comment: string;
    color: string;
    done: boolean;
    repliesRefs: Map<string, React.RefObject<HTMLLIElement>>,
    selectedReplies: Reply[],
    reviewer: User
}
interface Props {
    reviewer: User;
    commentObject: CommentObject;
    addComment: (itemKey: string, comment: string) => Promise<boolean>;
    editComment: (itemKey: string) => void;
    deleteComment: (itemKey: string) => void;
    callback?: (...args: any[]) => void;
    errors: FeedbackErrorType;
    registerPosition: (commentObject: CommentObject, ref: React.RefObject<any>) => void;
    end: boolean;
}
const CommentForm: React.FC<Props> = ({ end, reviewer, errors, commentObject, registerPosition, addComment, editComment, deleteComment, callback = (..._: any[]) => { } }: Props) => {
    const [editable, setEditable] = useState(!commentObject.done);
    const [confirmation, setConfirmation] = useState(false);
    const [comment, setComment] = useState(commentObject.comment);

    const commentRef = useRef<HTMLFormElement>(null);
    const handleDelete = () => {
        if (end) return
        setComment("");
        callback(false)
        setConfirmation(false);
        deleteComment(commentObject.color);
        setEditable(true);
    }
    const handleDeleteConfirmation = () => {
        if (end) return
        setConfirmation(true)
    }
    const handleAddComment = async () => {
        if (commentObject.reviewer.email != reviewer.email) {
            return;
        }
        const ok = await addComment(commentObject.color, comment);
        setEditable(!ok);
    }
    const handleEditComment = () => {
        if (end) return
        if (commentObject.reviewer.email != reviewer.email) {
            return;
        }
        editComment(commentObject.color);
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
                className={`absolute left-[300px] h-[270px]  w-[320px] bg-white  shadow rounded-md border ${commentObject.color}`}>
                <div className="flex flex-col h-full border-t border-gray-200 rounded-lg bg-gray-50">

                    <div className="border-b p-2">
                        <p className="text-sm md:text-xl text-gray-700 font-bold">
                            {commentObject.reviewer.firstName} {" "} {commentObject.reviewer.lastName}
                        </p>
                    </div>
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
                            className={`${end ? 'cursor-default' : 'cursor-pointer'}   mb-2 shadow shadow-gray-300 hover:bg-gray-100`}>
                            <EditIcon fill="#374151" />
                        </span>
                        <span className={`${end ? 'cursor-default' : 'cursor-pointer'}`} onMouseUp={handleDeleteConfirmation}>
                            <DeleteIcon fill="#335784" />
                        </span>
                        {errors.form && errors.form[commentObject.color] &&
                            <span className={`bg-white border absolute top-[74px] left-0 py-2 px-3 rounded-md `}>
                                <p className="text-red-500 italic text-sm md:text-md">{errors.form[commentObject.color]}</p>
                            </span>}
                        {confirmation &&
                            <div className={`absolute flex flex-col justify-center align-items top-1/2 right-11 md:left-28 w-[260px] p-3
                                            z-50 bg-white border border-gray-100 shadow rounded-md`}>
                                <p className="text-red-500 text-xs mb-3">
                                    Are you sure you want to delete this comment?
                                </p>
                                <div className="flex">
                                    <span onMouseUp={() => handleDelete()}>
                                        <ThinBtn text="Delete" variant="red" />
                                    </span>
                                    <span onMouseUp={() => setConfirmation(false)}>
                                        <ThinBtn text="Cancel" />
                                    </span>
                                </div>
                            </div>}
                    </div>
                </div>
            </form >
        </>
    )
}

export default CommentForm;
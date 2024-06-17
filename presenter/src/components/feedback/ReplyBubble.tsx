import { useEffect, useRef, useState } from "react";
import EditIcon from "../icons/EditIcon";
import DeleteIcon from "../icons/DeleteIcon";
import ThinButton from "../buttons/ThinButton";
interface ChatBubbleProps {
    body: string;
    fullName: string;
    profileIMG: string;
    createdAt: string;
    direction: string;
    seen: boolean;
    seenAt: string;
    itemKey: string;
    addComment: (itemKey: string, comment: string) => void;
    deleteComment: (itemKey: string) => void;
}
function ReplyBubble({ body, direction, profileIMG, createdAt, fullName, seen, seenAt, addComment, deleteComment, itemKey }: ChatBubbleProps) {
    const [showAddFeedback, setShowAddFeedback] = useState(false);
    const [comment, setComment] = useState("");
    const [editable, setEditable] = useState(true);
    const [confirmation, setConfirmation] = useState(false);
    const createdTime = new Date(parseInt(createdAt)).toLocaleTimeString().slice(0, 5);
    const seenTime = "Seen at " + new Date(parseInt(seenAt)).toLocaleTimeString().slice(0, 5);
    const editComment = () => {
        setEditable(true);
    }
    const addFeedback = () => {
        setShowAddFeedback(true);
    }
    const handleDelete = () => {
        setComment("");
        setShowAddFeedback(false);
        setConfirmation(false);
        deleteComment(itemKey);
    }
    const feedbackRef = useRef<HTMLFormElement>(null);
    const handleAddFeedback = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const inputValue = e.currentTarget.comment.value;
        addComment(itemKey, inputValue);
        setComment(inputValue);
        setEditable(false);
    }
    useEffect(() => {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        function handleClickOutside(event: any) {
            if (feedbackRef.current && !feedbackRef.current.contains(event.target)
                && comment == ""
            ) {
                setShowAddFeedback(false);
            }
        }
        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, [feedbackRef, comment]);
    return (
        <div>
            
            <div
                onClick={addFeedback}
                className={direction == "ltr" ? "relative flex self-start items-start gap-2.5" :
                    "relative flex items-start  flex-row-reverse self-end"}>
                <img className="w-8 h-8 rounded-full" src={profileIMG} alt="Jese image" />
                <div className="flex flex-col gap-1 w-full max-w-[320px]">
                    <div className="flex items-center space-x-2 rtl:space-x-reverse">
                        <span className="text-sm font-semibold text-gray-900">
                            {fullName}
                        </span>
                        <span className="text-sm font-normal text-gray-500">
                            {createdTime}
                        </span>
                    </div>
                    <div
                        className="flex flex-col leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-e-xl rounded-es-xl">
                        <p className="text-sm font-normal text-gray-900">
                            {body}
                        </p>
                    </div>
                    <span className="text-sm font-normal text-gray-500">
                        {seen ? seenTime : ""}</span>
                    {
                        <form
                            onSubmit={handleAddFeedback}
                            ref={feedbackRef}
                            className={
                                showAddFeedback ?
                                    `w-[260px] bg-white border-gray-100 shadow rounded-md ` :
                                    `hidden absolute z-50 bg-white border-gray-100 shadow rounded-md `
                            }>
                            <div
                                className="border
                                     border-gray-200 rounded-lg bg-gray-50">
                                <div
                                    className="px-4 py-3 bg-white rounded-t-lg">
                                    <textarea
                                        id="comment"
                                        rows={4}
                                        className={
                                            `w-full  px-0 text-sm text-gray-900
                                             bg-white border-0 resize-none focus:ring-0 focus:outline-none`
                                        }
                                        placeholder="Write a review about this specific answer"
                                        disabled={editable ? false : true}
                                        required >
                                    </textarea>
                                </div>
                                <div
                                    className="flex items-center justify-between px-3 py-2 border-t">
                                    <button
                                        type="submit"
                                        className={
                                            !editable ?
                                                `opacity-50 cursor-not-allowed relative inline-flex items-center justify-center p-0.5 mb-2 me-2
                                                overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500` :
                                                `relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900
                                            rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500 group-hover:from-cyan-500 group-hover:to-blue-500
                                        hover:text-white  focus:ring-4 focus:outline-none focus:ring-cyan-200` }
                                        disabled={editable ? false : true}>
                                        <span className={
                                            comment != "" ?
                                                `relative px-5 py-2.5 transition-all ease-in duration-75 bg-white rounded-md` :
                                                `relative px-5 py-2.5 transition-all ease-in duration-75 bg-white rounded-md group-hover:bg-opacity-0`}>
                                            {"Add to feedback"}
                                        </span>
                                    </button>
                                    <span
                                        onMouseUp={editComment}
                                        className="mb-2 shadow shadow-gray-300 cursor-pointer
                                    hover:bg-gray-100">
                                        <EditIcon fill="#374151" />
                                    </span>
                                    <span onMouseUp={() => setConfirmation(true)}>
                                        <DeleteIcon fill="#335784" />
                                    </span>
                                    {confirmation &&
                                        <div className={
                                            `absolute flex flex-col justify-center align-items
                                    top-1/2 right-11 md:left-28 w-[260px] p-3
                                     z-50 bg-white border border-gray-100 shadow rounded-md`}>
                                            <p className="text-red-500 text-xs mb-3">
                                                Are you sure you want to delete this comment?
                                            </p>
                                            <div className="flex">
                                                <span onMouseUp={() => handleDelete()}>
                                                    <ThinButton text="Delete" variant="red" />
                                                </span>
                                                <span onMouseUp={() => setConfirmation(false)}>
                                                    <ThinButton text="Cancel" />
                                                </span>
                                            </div>
                                        </div>}
                                </div>
                            </div>
                        </form>
                    }
                </div>
            </div>
        </div>
    )
}
export default ReplyBubble;
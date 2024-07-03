/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useRef } from "react";
import { Reply } from "../../lib/types";

interface ChatBubbleProps {
    msg: Reply
    direction: string;
    callback?: (...args: any[]) => void;
    registerPosition: (key: string, ref: React.RefObject<HTMLLIElement>) => void;
    end?:boolean
}
function ReplyBubble({ end, registerPosition, callback = (..._: any[]) => { }, direction, msg }: ChatBubbleProps) {
    const createdTime = new Date(parseInt(msg.createdAt)).toLocaleTimeString().slice(0, 5);
    const seenTime = "Seen at " + new Date(parseInt(msg.readAt)).toLocaleTimeString().slice(0, 5);
    const ref = useRef<HTMLLIElement>(null);
    const addFeedback = () => {
        if (end)return
        callback(msg, ref);
    }
    useEffect(() => {
        registerPosition(msg.id, ref);
    }, []);
    return (
        <li
            ref={ref}
            onClick={() => addFeedback()}
            className={`${end ? 'cursor-default' : 'cursor-pointer'} bg-white  ${direction == "ltr" ? "flex self-start items-start gap-2.5" :
             "flex items-start  flex-row-reverse self-end"} `}>
            <img className="w-8 h-8 rounded-full" src={msg.senderIMG} alt="image" />
            <div className={`flex flex-col gap-1 w-full max-w-[320px]`}>
                <div className={`flex items-center space-x-2 rtl:space-x-reverse`}>
                    <span className="text-sm font-semibold text-gray-900">
                        {msg.senderName}
                    </span>
                    <span className="text-sm font-normal text-gray-500">
                        {createdTime}
                    </span>
                </div>
                <div className="flex flex-col leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-e-xl rounded-es-xl">
                    <p className="text-sm font-normal text-gray-900">
                        {msg.body}
                    </p>
                </div>
                <span className="text-sm font-normal text-gray-500">{msg.read ? seenTime : ""}</span>
            </div>
        </li>

    )
}
export default ReplyBubble;
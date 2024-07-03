/* eslint-disable react-hooks/exhaustive-deps */
import { useRef } from "react";

interface ChatBubbleProps {
    body: string;
    fullName: string;
    profileIMG: string;
    createdAt: string;
    direction: string;
    seen: boolean;
    seenAt: string;
    isEnterprise: boolean;
    enterpriseID: string;
    registerPosition?: (id: string, ref: React.RefObject<HTMLLIElement>) => void;
}
function ChatBubble({  body, direction, profileIMG, createdAt, fullName, seen, seenAt, isEnterprise, enterpriseID }: ChatBubbleProps) {
    const createdTime = new Date(parseInt(createdAt)).toTimeString().slice(0, 5)
    const seenTime = "Seen at " + new Date(parseInt(seenAt)).toTimeString().slice(0, 5)
    const ref = useRef<HTMLLIElement>(null);

    return (
        <li
            ref={ref}
            className={direction == "ltr" ? "flex self-start items-start gap-2.5" : "flex items-start  flex-row-reverse self-end"}>
            <img className="w-8 h-8 rounded-full" src={profileIMG} alt="" />
            <div className="flex flex-col gap-1 w-full max-w-[320px]">
                <div className="flex items-center space-x-2 rtl:space-x-reverse">
                    <span className="text-sm font-semibold text-gray-900">
                        {isEnterprise ? fullName + " from " + enterpriseID : fullName}
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
            </div>
        </li>
    )
}
export default ChatBubble;
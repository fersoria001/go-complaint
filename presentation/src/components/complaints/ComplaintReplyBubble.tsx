import { ComplaintReply } from "@/gql/graphql";
import { dateFromMsString } from "@/lib/dateFromMsString";
import clsx from "clsx";
import CheckIcon from "../icons/CheckIcon";
import { useEffect, useRef } from "react";

interface Props {
    enterpriseName: string
    reply: ComplaintReply
    callback?: (id: string, ref: React.RefObject<HTMLLIElement>) => void
    onClick?: () => void
}
const ComplaintReplyBubble = ({ reply, enterpriseName, callback, onClick }: Props) => {
    let name = reply!.sender!.subjectName
    if (reply.isEnterprise!) {
        name += ` from ${enterpriseName}`
    }
    const ref = useRef<HTMLLIElement>(null)
    useEffect(() => {
        if (callback) {
            callback(reply.id!, ref)
        }
    }, [])
    return (
        <li
            onClick={() => onClick ? onClick() : undefined}
            ref={ref}
            className={clsx("bg-gray-50 border w-4/5 md:w-1/2 p-3 rounded-xl flex flex-col", {
                'self-end': reply.isEnterprise!,
                'cursor-pointer': onClick != undefined
            })}>
            <div className="flex">
                <p className="text-gray-700 text-xs md:text-sm font-bold">{name}</p>
                <p className="ms-2 text-gray-700 font-bold text-xs">{dateFromMsString(reply.readAt!).toLocaleTimeString()}</p>
            </div>
            <p className="text-gray-700 text-xs md:text-sm mb-2">{reply.body}</p>
            <CheckIcon className={clsx("w-5 h-5 self-end", {
                "fill-red-300": reply.read == false,
                "fill-blue-300": reply.read == true,
            })} />
        </li>
    )
}
export default ComplaintReplyBubble;
import { EnterpriseChatReplyType} from "../../../lib/types";

interface Props {
    direction: string;
    reply: EnterpriseChatReplyType
}
function EnterpriseChatReply({ direction, reply }: Props) {
    const seenAt = new Date(parseInt(reply.updatedAt)).toTimeString().slice(0,5)
    return (
        <li className={`${direction === "ltr" ? "flex flex-col items-start" :
            "flex flex-col items-end "}`}>
            <span className="text-xs font-semibold text-gray-900">
                {reply.user.firstName} {' '} {reply.user.lastName}
            </span>
            <p className="text-xs font-normal text-gray-900 w-2/3 bg-gray-100 rounded-md">
                {reply.content}
            </p>
            <p className="text-xs font-normal text-gray-800">{seenAt}</p>
        </li>
    )
}
export default EnterpriseChatReply;
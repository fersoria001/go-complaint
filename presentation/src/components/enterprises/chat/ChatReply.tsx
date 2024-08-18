import CheckIcon from "@/components/icons/CheckIcon";
import { ChatReply } from "@/gql/graphql";
import clsx from "clsx";


interface Props {
    direction: string;
    reply: ChatReply
}
function EnterpriseChatReply({ direction, reply }: Props) {
    const seenAt = new Date(parseInt(reply.updatedAt)).toTimeString().slice(0, 5)
    const createdAt = new Date(parseInt(reply.createdAt)).toTimeString().slice(0, 5)

    return (
        <li className={clsx("flex flex-col self-start mb-2 border w-2/3 px-2 rounded-xl", {
            "self-end": direction != "ltr",
        })}>
            <span className="text-xs font-semibold text-gray-700 mb-0.5">
                {reply.sender.subjectName} {createdAt}
            </span>
            <p className="text-xs font-normal text-gray-700  bg-gray-100 rounded-md mb-0.5">
                {reply.content}
            </p>
            <CheckIcon className={clsx("w-4 h-4", {
                "self-end": direction != "ltr",
                "self-start": direction === "ltr",
                "fill-red-300": reply.seen == false,
                "fill-blue-300": reply.seen == true,
            })} />
        </li>
    )
}
export default EnterpriseChatReply;

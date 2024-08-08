import { ComplaintReply, UserDescriptor } from "@/gql/graphql";
import { dateFromMsString } from "@/lib/dateFromMsString";
import clsx from "clsx";
import CheckIcon from "../icons/CheckIcon";

interface Props {
    rightSided: boolean
    currentUser: UserDescriptor
    reply: ComplaintReply
}
const ComplaintReplyBubble = ({ rightSided, currentUser, reply }: Props) => {
    let name = reply.sender.subjectName
    if (reply.sender.isEnterprise) {
        name = `${currentUser.fullName} from ${reply.sender.subjectName}`
    }
    return (
        <li className={clsx("bg-gray-50 border w-1/2 p-3 rounded-xl flex flex-col", {
            'self-end': rightSided
        })}>
            <div className="flex">
                <p className="text-gray-700 text-sm font-bold">{name}</p>
                <p className="ms-2 text-gray-700 font-bold text-xs">{dateFromMsString(reply.readAt).toLocaleTimeString()}</p>
            </div>
            <p className="text-gray-700 text-sm mb-2">{reply.body}</p>
            <CheckIcon className={clsx("w-5 h-5 self-end", {
                "fill-red-300": reply.read == false,
                "fill-blue-300": reply.read == true,
            })} />
        </li>
    )
}
export default ComplaintReplyBubble;
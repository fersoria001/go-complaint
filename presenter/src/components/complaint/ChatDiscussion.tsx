import { useEffect, useRef } from "react";
import { Reply } from "../../lib/types";
import ChatBubble from "./ChatBubble";

interface Props {
    replies: Reply[];
    currentAuthorID: string;
}
function ChatDiscussion({ replies, currentAuthorID }: Props) {
    const bottomRef = useRef<HTMLLIElement>(null);
    const scrollToBottom = () => {
        bottomRef.current?.scrollIntoView({
            behavior: "instant",
            block: "start",
        });
    };
    useEffect(() => {
        scrollToBottom();
    });
    scrollToBottom();
    return (
        <ul
            className="flex flex-col h-64 md:h-80 px-2 md:px-6 w-full
                    border-b overflow-y-auto overflow-x-hidden">
            {replies.length > 0 && replies!.map((msg) => {
                return (<ChatBubble
                    key={msg.id}
                    direction={
                        msg.senderID == currentAuthorID ? "ltr" : "rtl"
                    }
                    fullName={msg.senderName}
                    profileIMG={msg.senderIMG}
                    body={msg.body}
                    createdAt={msg.createdAt}
                    seen={msg.read}
                    seenAt={msg.readAt}
                    isEnterprise={msg.isEnterprise}
                    enterpriseID={msg.enterpriseID}
                />)


            })}
            <li ref={bottomRef}></li>
        </ul>
    )
}

export default ChatDiscussion;
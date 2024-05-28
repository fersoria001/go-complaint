import { useLoaderData } from "react-router-dom";
import { AuthMsg, Complaint, MarkAsReviewable, newReply, Sender } from "../lib/types";
import ProfileCard from "../components/complaint/ProfileCard";
import ComplaintCard from "../components/complaint/ComplaintCard";
import { useEffect, useRef, useState } from "react";
import ChatBubble from "../components/complaint/ChatBubble";
import ChatButton from "../components/buttons/ChatButton";
import useChat from "../lib/hooks/useChat";
import RateReviewIcon from "../components/icons/RateReviewIcon";
import ThinButton from "../components/buttons/ThinButton";
import { MarkAsReviewableMutation, Mutation } from "../lib/mutations";

function ComplaintPage() {
    const { complaints, id, sender, authMsg } = useLoaderData() as {
        complaints: Complaint,
        id: string,
        sender: Sender;
        authMsg: AuthMsg
    };
    const [tab, setTab] = useState("complaint");
    const [showConfirm, setShowConfirm] = useState(false);
    const [mar, setMar] = useState<MarkAsReviewable | null>(null);
    const { webSocket, webSocketReady, sortedMessages } = useChat(id, authMsg, tab);
    const selectComplaintTab = () => {
        setTab("complaint");
    };
    const selectChatTab = () => {
        switch (complaints.status) {
            case "OPEN":
                setTab("chat");
                break;
            case "STARTED":
                setTab("chat");
                break;
            case "IN_DISCUSSION":
                setTab("chat");
                break;
            case "IN_REVIEW":
                setTab("in_review");
                break;
            case "CLOSED":
                setTab("closed");
                break;
            case "IN_HISTORY":
                setTab("history");
                break;
        }

    };
    const send = () => {
        if (!webSocketReady || !webSocket) {
            console.error("trying to send to a closed WebSocket");
            return;
        }
        const input = document.getElementById("chat-input") as HTMLTextAreaElement;
        if (!input) {
            console.error("Chat input element not found");
            return;
        }
        const messageText = input.value.trim();
        if (messageText === "") {
            console.warn("Cannot send an empty message");
            return;
        }

        const message = newReply(sender.thumbnail, sender.fullName, messageText);
        try {
            console.log("sending message:", message)
            webSocket.send(JSON.stringify({ content: "reply", reply: message }));
            input.value = "";
        } catch (error) {
            console.error("Failed to send message:", error);
        }
    };
    const bottomRef = useRef<HTMLLIElement>(null);
    const scrollToBottom = () => {
        bottomRef.current?.scrollIntoView({
            behavior: "smooth",
            block: "start",
        });
    };
    useEffect(() => {
        scrollToBottom();
    }, [sortedMessages]);
    useEffect(() => {
        if (!mar) return;
        Mutation<MarkAsReviewable>(MarkAsReviewableMutation, mar).then(() => {
            setTab("in_review");
        });
    }, [mar])
    const confirmRef = useRef<HTMLDivElement>(null);
    const handleAskForReview = () => {
        let newMar = {} as MarkAsReviewable;
        if (authMsg.enterprise_id) {
            newMar = {
                complaintID: complaints.id,
                enterpriseID: authMsg.enterprise_id,
                assistantID: ""
            }
        }
        setMar(newMar);
    }
    return (
        <div className="flex flex-col pt-4 md:pt-6 relative">
            <div className="flex flex-col w-full items-center pb-2 md:pb-6">
                <ProfileCard
                    userFullName={complaints.authorFullName}
                    profileIMG={complaints.authorProfileIMG}
                    sub={complaints.industry ? complaints.industry : complaints.authorID}
                />
                <span
                    onMouseUp={() => setShowConfirm(true)}
                    className="flex  rounded-md shadow mt-2 p-2 hover:bg-gray-100 cursor-pointer">
                    <RateReviewIcon fill="#3242bf" />
                    <p>Ask for review</p>
                </span>
                {showConfirm && <div ref={confirmRef} className="absolute z-50 bg-white
                    p-2.5 w-[200px] md:w-[360px] 
                    left-1/2 -translate-x-1/2 md:top-1/2 
                    border rounded-md shadow flex flex-col ">
                    <p> Are you sure ?</p>
                    <p className="text-xs text-red-500 italic">
                        This will close the complaint,
                        no replies could be sent after.
                        It will ask for a review from the client,
                        later a manager may review the discussion.</p>
                    <div className="flex gap-2.5 self-center">
                        <span onMouseUp={() => handleAskForReview()}>
                            <ThinButton text="Yes" variant="red" />
                        </span>
                        <span onMouseUp={() => setShowConfirm(false)}>
                            <ThinButton text="No" variant="gray" />
                        </span>
                    </div>
                </div>}

            </div>
            <div className="w-full bg-white ">
                <ul className="
                text-sm font-medium text-center text-gray-500
                md:divide-x divide-y divide-gray-200 rounded-lg sm:flex rtl:divide-x-reverse">
                    <li className="w-full">
                        <button
                            onClick={selectComplaintTab}
                            type="button"
                            className="inline-block w-full p-4 rounded-ss-lg bg-gray-50 hover:bg-gray-100 focus:outline-none">
                            Complaint
                        </button>
                    </li>
                    <li className="w-full">
                        <button
                            onClick={selectChatTab}
                            type="button"
                            className="inline-block w-full p-4 bg-gray-50 hover:bg-gray-100 focus:outline-none">
                            Chat
                        </button>
                    </li>
                </ul>
                <div className="border-t border-gray-200">
                    <div className=" bg-white rounded-lg">
                        {tab == "complaint" ? <ComplaintCard
                            reason={complaints.message.title}
                            description={complaints.message.description}
                            body={complaints.message.body}
                            status={complaints.status}
                            createdAt={complaints.createdAt}
                        /> :
                            tab == "chat" ?
                                <div>
                                    <ul
                                        className="flex flex-col  h-72 md:h-96 px-2 md:px-6 w-full
                                border-b
                                 overflow-y-auto overflow-x-hidden">
                                        {sortedMessages.map((msg) => {
                                            console.log(msg.sender_id, complaints.authorID);
                                            return <li key={msg.id} className="first:pt-4">
                                                <ChatBubble
                                                    direction={
                                                        msg.sender_id == complaints.authorID ? "ltr" : "rtl"
                                                    }
                                                    fullName={msg.sender_name}
                                                    profileIMG={msg.sender_img}
                                                    body={msg.body}
                                                    createdAt={msg.created_at}
                                                    seen={msg.read}
                                                    seenAt={msg.read_at}
                                                /></li>

                                        })}
                                        <li ref={bottomRef}></li>
                                    </ul>
                                    <div className="relative bg-white">
                                        <textarea
                                            onKeyUpCapture={(e) => {
                                                if (e.key === "Enter") {
                                                    send();
                                                }
                                            }}
                                            id="chat-input"
                                            maxLength={120}
                                            rows={4}
                                            className="
                                    text-lg text-gray-500 md:text-xl
                                    h-auto w-full resize-none focus:outline-none pb-6 px-2 
                                    md:px-6
                                    md:py-4
                                    " />
                                        <span

                                            onClick={send}
                                            className="absolute right-0 bottom-0 md:right-24 bottom:8">
                                            <ChatButton />
                                        </span>
                                    </div>
                                </div>
                                : tab == "in_review" ?
                                    <div className="flex flex-col items-center">
                                        <h1 className="text-lg font-bold text-gray-500">In review</h1>
                                    </div>
                                    : <div> case ${tab} not implemented</div>
                        }
                    </div>
                </div>
            </div>
        </div>
    );
}

export default ComplaintPage;
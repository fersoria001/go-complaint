import { useLoaderData, useRevalidator } from "react-router-dom";
import { AuthMsg, Complaint, Office, Sender } from "../lib/types";
import ProfileCard from "../components/complaint/ProfileCard";
import ComplaintCard from "../components/complaint/ComplaintCard";
import { useEffect, useRef, useState } from "react";
import ChatBubble from "../components/complaint/ChatBubble";
import ChatButton from "../components/buttons/ChatButton";
import useChat from "../lib/hooks/useChat";
import RateReviewIcon from "../components/icons/RateReviewIcon";
import ThinButton from "../components/buttons/ThinButton";

//fix the revalidation and re-rendering issue
function ComplaintPage() {
    const { complaints, complaintID, sender, authMsg } = useLoaderData() as {
        complaints: Complaint,
        complaintID: string,
        sender: Sender,
        authMsg: AuthMsg,
        office: Office,
    };
    const [tab, setTab] = useState("complaint");
    const [showConfirm, setShowConfirm] = useState(false);
    const { closed, incomingMessages, open, close, send, closeRoom } = useChat(complaintID, authMsg);
    const revalidator = useRevalidator();
    const selectComplaintTab = () => {
        close();
        setTab("complaint");
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
        if (closed) {
            console.log("closed is true")
            revalidator.revalidate();
            setTab("in_review");
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [closed, complaints, incomingMessages]);
    const confirmRef = useRef<HTMLDivElement>(null);
    const selectChatTab = () => {
        if ((complaints.status === "IN_REVIEW" ||
            complaints.status === "CLOSED" ||
            complaints.status === "IN_HISTORY")
        ) {
            close();
            setTab("in_review");
        } else {
            open();
            setTab("chat");
        }
    }

    return (
        <div className="flex flex-col pt-4 md:pt-6 relative">
            <div className="flex flex-col w-full items-center pb-2 md:pb-6">
                <ProfileCard
                    userFullName={complaints.authorFullName}
                    profileIMG={complaints.authorProfileIMG}
                    sub={complaints.industry ? complaints.industry : complaints.authorID}
                />
                {
                    complaints.status === "IN_DISCUSSION" &&
                    sender.fullName != complaints.authorFullName &&
                    <span
                        onMouseUp={() => setShowConfirm(true)}
                        className="flex rounded-md shadow mt-2 p-2 hover:bg-gray-100 cursor-pointer">
                        <RateReviewIcon fill="#3242bf" />
                        <p>Ask for review</p>
                    </span>
                }
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
                        <span onMouseUp={() => {
                            setTab("in_review");
                            closeRoom(sender);
                            setShowConfirm(false);
                        }}>
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
                            className="inline-block w-full p-4 rounded-ss-lg bg-gray-50
                            hover:bg-gray-100 focus:outline-none">
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
                                        {incomingMessages.length > 0 && incomingMessages.map((msg) => {
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
                                                    isEnterprise={msg.is_enterprise}
                                                    enterpriseID={msg.enterprise_id}
                                                />
                                                </li>

                                        })}
                                        <li ref={bottomRef}></li>
                                    </ul>
                                    <div className="relative bg-white">
                                        <textarea
                                            onKeyUpCapture={(e) => {
                                                if (e.key === "Enter") {
                                                    send(e.currentTarget.value, sender);
                                                    e.currentTarget.value = "";
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
                                            onClick={() => {
                                                send((document.getElementById("chat-input") as HTMLTextAreaElement).value, sender)
                                                    ; (document.getElementById("chat-input") as HTMLTextAreaElement).value = "";
                                            }}
                                            className="absolute right-0 bottom-0 md:right-24 bottom:8">
                                            <ChatButton />
                                        </span>
                                    </div>
                                </div>
                                : tab == "in_review" ?
                                    <div className="flex flex-col items-center">
                                        <h1 className="text-lg font-bold text-gray-500">{complaints.status}</h1>
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
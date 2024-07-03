/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useEffect, useContext } from "react";
import ChatDiscussion from "../../complaint/ChatDiscussion";
import ChatInput from "../../complaint/ChatInput";
import ProfileCard from "../../complaint/ProfileCard";
import ComplaintTabs from "../../complaint/ComplaintTabs";
import { Route } from "../../../routes/_profile/sent/$complaintId/chat";
import { useRouter } from "@tanstack/react-router";
import useSubscriber from "../../../lib/hooks/useSubscriber";
import { replyComplaint } from "../../../lib/reply_complaint";
import { Reply } from "../../../lib/types";
import { isReplies, isReply, markAsSeen } from "../../../lib/mark_reply_complaint_as_seen";
import { SideBarContext } from "../../../react-context/SideBarContext";




const Chat: React.FC = () => {
    const { descriptor, complaint, id } = Route.useLoaderData();
    const { complaintId } = Route.useParams();
    const { setReload } = useContext(SideBarContext)
    const [replies, setReplies] = useState(complaint.replies);
    const [status, setStatus] = useState(complaint.status);
    const { incomingMessage } = useSubscriber(id);
    const router = useRouter();
    useEffect(() => {
        if (incomingMessage) {
            if (isReply(incomingMessage)) {
                const msg = incomingMessage as Reply;
                if (msg.complaintStatus === "IN_REVIEW") {
                    router.invalidate();
                    return;
                }
                if (msg.senderID !== descriptor!.email && msg.read === false) {
                    markAsSeen(complaintId, [msg]);
                }
                setStatus(msg.complaintStatus);
                setReplies(replies => [...replies!, incomingMessage]);
            }
            if (isReplies(incomingMessage)) {
                const msgs = incomingMessage as Reply[];
                setReplies(p => p?.map((reply) => {
                    const found = msgs.find((msg) => msg.id === reply.id)
                    return found ? found : reply
                }))
            }
        }
    }, [incomingMessage, router]);
    useEffect(() => {
        const received = complaint.replies?.filter((reply) => reply.senderID !== descriptor!.email && !reply.read)
        if (received?.length) {
            const load = async () => {
                markAsSeen(complaintId, received).then(() => setReload(true))
            }
            load()
        }
    }, [])
    const send = async (body: string) => {
        await replyComplaint(complaint.id, descriptor!.email, body, "")
    }
    return (
        <div className="flex flex-col relative min-h-[315px] md:min-h-[460px]">
            <div className="flex flex-col w-full items-center pb-2 md:pb-6">
                <ProfileCard
                    userFullName={complaint.receiverFullName}
                    profileIMG={complaint.receiverProfileIMG}
                    sub={complaint.industry || ''}
                />
            </div>

            <div className="w-full bg-white ">
                <ComplaintTabs
                    selected="chat"
                    complaintLink={`/sent/${complaintId}`}
                    chatLink={`/sent/${complaintId}/chat`} />
                <div className="border-t border-gray-200">
                    {status === "IN_REVIEW" ?
                        <div className="w-full bg-white ">
                            <div className="text-sm md:text-xl text-gray-700 mb-4">
                                The current complaint is solved and waiting for a review of the attention
                            </div>
                        </div> :
                        <div className=" bg-white rounded-lg">
                            <div>
                                <ChatDiscussion replies={replies!}
                                    currentAuthorID={descriptor!.email} />
                                <ChatInput callback={send} />
                            </div>
                        </div>
                    }
                </div>

            </div>

        </div>
    );
}

export default Chat;
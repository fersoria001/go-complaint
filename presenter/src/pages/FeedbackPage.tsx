import { useLoaderData } from "react-router-dom";
import { Complaint, CreateAFeedback, UserDescriptor } from "../lib/types";
import ReplyBubble from "../components/feedback/ReplyBubble";
import { useState } from "react";
import ChatBubble from "../components/complaint/ChatBubble";
import PrimaryButton from "../components/buttons/PrimaryButton";
function FeedbackPage() {
    const { user, complaint } = useLoaderData() as { user: UserDescriptor, complaint: Complaint };
    const [feedbackBatch, setFeedbackBatch] = useState<CreateAFeedback[]>([]);
    const findReply = (itemKey: string) => {
        return complaint.replies!.find((msg) => msg.id == itemKey);
    }
    const addToFeedback = (itemKey: string, comment: string) => {
        const reply = findReply(itemKey);
        if (!reply) {
            return;
        }
        const newFeedback: CreateAFeedback = {
            complaintID: complaint.id,
            reviewerID: user.email,
            reviewedID: reply.senderID,
            reviewerIMG: user.profileIMG,
            reviewerName: user.fullName,
            senderID: reply.senderID,
            senderIMG: reply.senderIMG,
            senderName: reply.senderName,
            body: reply.body,
            createdAt: reply.createdAt,
            read: reply.read,
            readAt: reply.readAt,
            updatedAt: reply.updatedAt,
            comment: comment,
        }
        feedbackBatch.push(newFeedback);
        setFeedbackBatch([...feedbackBatch]);
    }
    const deleteFromBatch = (itemKey: string) => {
        const reply = findReply(itemKey);
        if (!reply) {
            return;
        }
        const index = feedbackBatch.findIndex((msg) => msg.senderID == reply.senderID);
        if (index != -1) {
            feedbackBatch.splice(index, 1);
            setFeedbackBatch([...feedbackBatch]);
        }
    }
    return (
        <div>
            <div className="border border-b flex align-center py-3 px-5">
                <PrimaryButton text="Submit Feedback" />
            </div>
            <ul
                className="flex flex-col px-2 md:px-6 w-full
                overflow-y-auto
                h-screen
                bg-white shadow rounded-md border-b">
                {complaint.replies!.length > 0 && complaint.replies!.map((msg, index) => {
                    if (index != complaint.replies!.length - 1) {
                        if (msg.senderID == complaint.authorID) {
                            return <li
                                key={msg.id
                                }
                                className="first:pt-4" >
                                <ChatBubble
                                    direction={
                                        msg.senderID == complaint.authorID ? "ltr" : "rtl"
                                    }
                                    fullName={msg.senderName}
                                    profileIMG={msg.senderIMG}
                                    body={msg.body}
                                    createdAt={msg.createdAt}
                                    seen={msg.read}
                                    seenAt={msg.readAt}
                                    isEnterprise={false}
                                    enterpriseID=""
                                />
                            </li>
                        } else {
                            return <li
                                key={msg.id}
                                className="first:pt-4">
                                <ReplyBubble
                                    addComment={addToFeedback}
                                    deleteComment={deleteFromBatch}
                                    direction={
                                        msg.senderID == complaint.authorID ? "ltr" : "rtl"
                                    }
                                    itemKey={msg.id}
                                    fullName={msg.senderName}
                                    profileIMG={msg.senderIMG}
                                    body={msg.body}
                                    createdAt={msg.createdAt}
                                    seen={msg.read}
                                    seenAt={msg.readAt}
                                />
                            </li>
                        }

                    }
                })}
            </ul>

        </div >
    )
}

export default FeedbackPage;
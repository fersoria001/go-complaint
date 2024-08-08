'use client'
import { useSuspenseQueries } from "@tanstack/react-query";
import ComplaintInput from "./ComplaintInput";
import Image from "next/image";
import getGraphQLClient from "@/graphql/graphQLClient";
import { useParams } from "next/navigation";
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { dateFromMsString } from "@/lib/dateFromMsString";
import clsx from "clsx";
import ComplaintReplyBubble from "./ComplaintReplyBubble";
import { Complaint, ComplaintReply } from "@/gql/graphql";
import { useEffect, useRef, useState } from "react";
import useChat, { ChatSubProtocols, decodeFromBinary, encodeToBinary } from "@/lib/hooks/useChat";
import { getCookie } from "@/lib/actions/cookies";

enum complaintSubProtocolDataType {
    ReplyComplaint = "reply_complaint",
    MarkAsRead = "mark_as_read",
    ComplaintReply = "complaint_reply",
    Complaint = "complaint"
}
type complaintSubProtocolResult = {
    subProtocolDataType: complaintSubProtocolDataType;
    result: any;
}
type complaintSubProtocolPayload = {
    subProtocolDataType: complaintSubProtocolDataType;
    command: any;
}

type replyComplaintData = {
    senderId: string;
    complaintId: string;
    body: string;
}

type markReplyAsReadData = {
    complaintId: string
    replyId: string
}

const ComplaintChat: React.FC = () => {
    const params = useParams()
    const gqlClient = getGraphQLClient()
    const [{ data: { userDescriptor } }, { data: { complaintById } }, { data: jwt }, { data: alias }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await gqlClient.request(userDescriptorQuery)
            },
            {
                queryKey: ["complaintById", params.complaintId as string],
                queryFn: async () => await gqlClient.request(complaintByIdQuery, { id: params.complaintId as string })
            },
            {
                queryKey: ["serverSideJwtCookie"],
                queryFn: async () => await getCookie("jwt")
            },
            {
                queryKey: ["serverSideAliasCookie"],
                queryFn: async () => await getCookie("alias")
            }
        ]
    })
    const [item, setItem] = useState<Complaint>(complaintById as Complaint)
    const { isReady, send, incomingMsg } = useChat(params.complaintId as string, ChatSubProtocols.COMPLAINT, jwt!)
    const subject = alias != item.author!.id ? item.author! : item.receiver!
    const windowRef = useRef<HTMLDivElement>(null)
    useEffect(() => {
        const data: markReplyAsReadData[] = []
        for (let i = 0; i < item.replies.length; i++) {
            if (item.replies[i].sender.id != alias && item.replies[i].read === false) {
                data.push({
                    complaintId: params.complaintId as string,
                    replyId: item.replies[i].id
                })
            }
        }
        if (data.length > 0) {
            const payload: complaintSubProtocolPayload = {
                subProtocolDataType: complaintSubProtocolDataType.MarkAsRead,
                command: encodeToBinary(JSON.stringify(data))
            }
            send(payload)
        }
        if (incomingMsg) {
            const parsed = JSON.parse(decodeFromBinary(incomingMsg)) as complaintSubProtocolResult
            switch (parsed.subProtocolDataType) {
                case complaintSubProtocolDataType.ComplaintReply: {
                    const newR = JSON.parse(decodeFromBinary(parsed.result)) as ComplaintReply
                    setItem(prev => {
                        const newReplies = prev.replies.map((r) => r.id === newR.id ? newR : r)
                        const index = newReplies.findIndex((v) => v.id === newR.id)
                        if (index < 0) {
                            newReplies.unshift(newR)
                        }
                        const newItem = { ...prev, replies: newReplies }
                        return newItem
                    })
                    break;
                }
                case complaintSubProtocolDataType.Complaint: {
                    const newC = JSON.parse(decodeFromBinary(parsed.result)) as Complaint
                    setItem(prev => {
                        return { ...prev, replies: newC.replies }
                    })
                    break;
                }
            }

        }
        if (windowRef.current) {
            windowRef.current.scrollIntoView({
                behavior: "instant",
                block: "end"
            })
        }
    }, [incomingMsg, item])
    const handleReply = (b: string) => {
        const data: replyComplaintData = {
            senderId: alias!,
            complaintId: params.complaintId as string,
            body: b
        }
        const payload: complaintSubProtocolPayload = {
            subProtocolDataType: complaintSubProtocolDataType.ReplyComplaint,
            command: encodeToBinary(JSON.stringify(data))
        }
        if (isReady) {
            send(payload)
        }
    }
    return (
        <div ref={windowRef} className="w-full flex flex-col border">
            <div className="flex w-full my-2.5 py-2.5">
                <div className='relative mx-2 rounded-full h-10 w-10 bg-gray-300 self-center'>
                    <Image
                        src={subject.subjectThumbnail!}
                        alt={subject.subjectName!}
                        className="rounded-full"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                    />
                </div>
                <div className="px-2 self-center">
                    <h3 className="text-gray-700 font-bold text-md lg:text-lg xl:text-xl">{subject.subjectName}</h3>
                </div>
                <div className="ml-auto mr-4 my-auto flex items-center gap-2.5">
                    <button
                        type="button"
                        className="text-white bg-blue-500 rounded-xl px-2.5 hover:bg-blue-600">
                        Mark for review
                    </button>
                    <div className={clsx("rounded-full h-2 w-2 ml-auto mr-4 my-auto", {
                        "bg-blue-300": subject?.isOnline == true,
                        "bg-red-300": subject?.isOnline == false
                    })}></div>
                </div>
            </div>
            <div className="px-2 xl:px-0 xl:ps-12 xl:flex xl:justify-between">
                <div>
                    <div className="flex">
                        <label className="text-gray-700 text-sm xl:text-md font-medium" htmlFor="reason">Reason:</label>
                        <p className="ms-2 text-gray-700 text-sm xl:text-md">{item.title}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-sm xl:text-md font-medium" htmlFor="description">Description:</label>
                        <p className="ms-2 text-gray-700 text-sm xl:text-md">{item.description}</p>
                    </div>
                </div>
                <div className="pr-6">
                    <div className="flex">
                        <label className="text-gray-700 text-sm xl:text-md font-medium">Status:</label>
                        <p className="ms-2 text-gray-700 text-sm xl:text-md">{item.status}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-sm xl:text-md font-medium">Created at:</label>
                        <p className="ms-2 text-gray-700 text-sm xl:text-md">{dateFromMsString(item.createdAt).toUTCString()}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-sm xl:text-md font-medium">Last update:</label>
                        <p className="ms-2 text-gray-700 text-sm xl:text-md">{dateFromMsString(item.updatedAt).toUTCString()}</p>
                    </div>
                </div>
            </div>
            <ul className="overflow-y-auto p-2 h-[21.375rem] md:h-[21.525rem] xl:h-[24.925rem] flex gap-2.5 py-2.5 flex-col-reverse">
                {
                    item.replies.map((reply) => {
                        return (
                            <ComplaintReplyBubble
                                key={reply.id}
                                currentUser={userDescriptor}
                                reply={reply as ComplaintReply}
                                rightSided={alias != reply.sender.id} />
                        )
                    })
                }
            </ul>
            <ComplaintInput sendCallback={handleReply} />
        </div>
    )
}
export default ComplaintChat;
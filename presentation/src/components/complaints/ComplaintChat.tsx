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
import Link from "next/link";

enum complaintSubProtocolDataType {
    ReplyComplaint = "reply_complaint",
    MarkAsRead = "mark_as_read",
    ComplaintReply = "complaint_reply",
    Complaint = "complaint",
    SendToReview = "send_to_review"
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
    aliasId: string;
}

type markReplyAsReadData = {
    id: string
    replyId: string
}

type sendComplaintToReviewData = {
    receiverId: string;
    complaintId: string;
    currentUserId: string;
}

const validStatus = ["OPEN", "STARTED", "IN_DISCUSSION"]
const ComplaintChat: React.FC = () => {
    const params = useParams()
    const complaintId = params.complaintId as string
    const gqlClient = getGraphQLClient()
    const [{ data: { userDescriptor } }, { data: { complaintById } }, { data: jwt }, { data: alias }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await gqlClient.request(userDescriptorQuery)
            },
            {
                queryKey: ["complaintById", complaintId as string],
                queryFn: async () => await gqlClient.request(complaintByIdQuery, { id: complaintId as string })
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
    const { isReady, send, incomingMsg } = useChat(complaintId as string, ChatSubProtocols.COMPLAINT, jwt!)
    const subject = alias != item.author!.id ? item.author! : item.receiver!
    const windowRef = useRef<HTMLDivElement>(null)

    useEffect(() => {
        const data: markReplyAsReadData[] = []
        for (let i = 0; i < item.replies!.length; i++) {
            if (item.replies![i]!.sender!.id != userDescriptor.id && item.replies![i]!.read === false) {
                data.push({
                    id: complaintId as string,
                    replyId: item.replies![i]!.id!
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
    }, [alias, complaintId, item.replies, send, userDescriptor.id])

    useEffect(() => {
        if (incomingMsg) {
            const msg = incomingMsg as complaintSubProtocolResult
            switch (msg.subProtocolDataType) {
                case complaintSubProtocolDataType.ComplaintReply: {
                    const newR = JSON.parse(decodeFromBinary(msg.result)) as ComplaintReply
                    setItem(prev => {
                        const newReplies = prev.replies!.map((r) => r!.id === newR.id ? newR : r)
                        const index = newReplies.findIndex((v) => v!.id === newR.id)
                        if (index < 0) {
                            newReplies.unshift(newR)
                        }
                        const newItem = { ...prev, replies: newReplies }
                        return newItem
                    })
                    break;
                }
                case complaintSubProtocolDataType.Complaint: {
                    const newC = JSON.parse(decodeFromBinary(msg.result)) as Complaint
                    setItem(prev => {
                        return { ...prev, replies: newC.replies, status: newC.status }
                    })
                    break;
                }
            }

        }
    }, [alias, incomingMsg, complaintId])

    useEffect(() => {
        if (windowRef.current) {
            windowRef.current.scrollIntoView({
                behavior: "instant",
                block: "end"
            })
        }
    }, [])

    const handleMarkForReview = () => {
        const data: sendComplaintToReviewData = {
            receiverId: alias!,
            complaintId: complaintId,
            currentUserId: userDescriptor.id,
        }
        const payload: complaintSubProtocolPayload = {
            subProtocolDataType: complaintSubProtocolDataType.SendToReview,
            command: encodeToBinary(JSON.stringify(data))
        }
        send(payload)
    }

    const handleReply = (b: string) => {
        const data: replyComplaintData = {
            senderId: userDescriptor.id,
            complaintId: complaintId as string,
            body: b,
            aliasId: alias != userDescriptor.id ? alias! : ""
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
            <div className="flex w-full my-2 py-2.5">
                <div className='relative mx-2 rounded-full h-8 w-8 sm:h-10 sm:w-10 bg-gray-300 self-center'>
                    <Image
                        src={subject.subjectThumbnail!}
                        alt={subject.subjectName!}
                        className="rounded-full"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                    />
                </div>
                <div className="px-2 self-center">
                    <h3 className="text-gray-700 font-bold text-sm lg:text-md xl:text-lg">{subject.subjectName}</h3>
                </div>
                <div className="ml-auto mr-4 my-auto flex items-center gap-2.5">
                    {item.receiver?.id === alias && validStatus.findIndex((s) => s === item.status) > 0 && <button
                        type="button"
                        onMouseUp={() => handleMarkForReview()}
                        className="text-white text-sm lg:text-md bg-blue-500 rounded-xl px-2.5 hover:bg-blue-600">
                        Mark for review
                    </button>}
                    <div className={clsx("rounded-full h-2 w-2 ml-auto mr-4 my-auto shrink-0", {
                        "bg-red-300": subject!.isOnline! == false,
                        "bg-blue-300": subject!.isOnline! != true,
                    })}></div>
                </div>
            </div>
            <div className="px-0.5 xl:px-0 xl:ps-12 xl:flex xl:justify-between pb-1">
                <div>
                    <div className="flex">
                        <label className="text-gray-700 text-xs md:text-sm xl:text-md font-medium" htmlFor="reason">Reason:</label>
                        <p className="ms-0.5 text-gray-700 text-xs md:text-sm xl:text-md">{item.title}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-xs md:text-sm xl:text-md font-medium" htmlFor="description">Description:</label>
                        <p className="ms-0.5 text-gray-700 text-xs md:text-sm xl:text-md">{item.description}</p>
                    </div>
                </div>
                <div className="pr-6">
                    <div className="flex">
                        <label className="text-gray-700 text-xs md:text-sm xl:text-md font-medium">Status:</label>
                        <p className="ms-0.5 text-gray-700 text-xs md:text-sm xl:text-md">{item.status}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-xs md:text-sm xl:text-md font-medium">Created at:</label>
                        <p className="ms-0.5 text-gray-700 text-xs md:text-sm xl:text-md">{dateFromMsString(item.createdAt!).toUTCString()}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-xs md:text-sm xl:text-md font-medium">Last update:</label>
                        <p className="ms-0.5 text-gray-700 text-xs md:text-sm xl:text-md">{dateFromMsString(item.updatedAt!).toUTCString()}</p>
                    </div>
                </div>
            </div>

            {
                validStatus.findIndex((s) => s === item.status) >= 0 ?
                    <>
                        <ul className="overflow-y-auto p-2 h-[21.375rem] md:h-[21.525rem] xl:h-[24.925rem] flex gap-2.5 py-2.5 flex-col-reverse">
                            {
                                item!.replies!.map((reply) => {
                                    return (
                                        <ComplaintReplyBubble
                                            key={reply!.id}
                                            reply={reply as ComplaintReply}
                                            enterpriseName={
                                                complaintById.author?.id == alias ?
                                                    complaintById.author?.subjectName! :
                                                    complaintById.receiver?.subjectName!
                                            }
                                        />
                                    )
                                })
                            }
                        </ul>
                        <ComplaintInput sendCallback={handleReply} />
                    </>
                    :
                    <div className="overflow-y-auto p-2 h-[26.475rem] md:h-[27.525rem] xl:h-[30.925rem] flex gap-2.5 py-2.5 items-center">
                        <p className="text-gray-700 text-xs md:text-sm xl:text-md font-medium">
                            This complaint has been sent for review you cannot reply to its discussion anymore but you can track it and perform additional steps in the
                            {" "}
                            <Link href={"/reviews"} className="text-blue-300 underline">
                                reviews section
                            </Link>.
                            {" "}
                        </p>
                    </div>
            }

        </div >
    )
}
export default ComplaintChat;
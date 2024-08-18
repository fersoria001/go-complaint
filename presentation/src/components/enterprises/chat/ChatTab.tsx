'use client'
import { Chat, ChatReply, Employee, UserDescriptor } from "@/gql/graphql";
import EmojiPicker, { EmojiStyle, Theme } from "emoji-picker-react";
import { PreviewConfig } from "emoji-picker-react/dist/config/config";
import { useState, useRef, useEffect } from "react";
import EnterpriseChatReply from "./ChatReply";
import { useMutation, useQuery } from "@tanstack/react-query";
import { createEnterpriseChat } from "@/lib/actions/graphqlActions";
import getGraphQLClient from "@/graphql/graphQLClient";
import findEnterpriseChatQuery from "@/graphql/queries/findEnterpriseChatQuery";
import useChat, { ChatSubProtocols, decodeFromBinary, encodeToBinary } from "@/lib/hooks/useChat";
import { enterpriseChatSubProtocolDataType, enterpriseSubProtocolPayload, enterpriseSubProtocolResult, markAsSeenData, replyCommandData } from "./SideChat";
import useClickOutside from "@/lib/hooks/useClickOutside";
import useScroll from "@/lib/hooks/useScroll";
import useWindowSize from "@/lib/hooks/useWindowsSize";

export type TabType = {
    user: Employee;
    enterpriseID: string;
    index: number;
    isActive: boolean;
};
interface Props {
    jwt: string;
    isRightBarOpen: boolean;
    closeTab: (id: string) => void;
    minimizeTab: (id: string, v: boolean) => void;
    descriptor: UserDescriptor;
    tab: TabType;
}

const ChatTab: React.FC<Props> = ({ jwt, tab, descriptor, isRightBarOpen, closeTab, minimizeTab }: Props) => {
    const [showEmoji, setShowEmoji] = useState<boolean>(false)
    const emojiRef = useRef(null)
    const initial = 230 * tab.index
    const [ml, setMl] = useState<number>(initial)

    useEffect(() => {
        if (isRightBarOpen) {
            setMl(p => initial + 315);
        } else {
            setMl(p => initial + 74);
        }
    }, [initial, isRightBarOpen])
    const { data, isError, isSuccess, refetch } = useQuery({
        queryKey: ["find-enterprise-chat", tab.enterpriseID, descriptor.id, tab.user.user.id],
        queryFn: async () => await getGraphQLClient().request(findEnterpriseChatQuery, {
            input: {
                enterpriseId: tab.enterpriseID,
                recipientOneId: descriptor.id,
                recipientTwoId: tab.user.user.id,
            }
        }),
        throwOnError: false,
    })
    const [replies, setReplies] = useState<ChatReply[]>([])
    const createMutation = useMutation({
        mutationFn: () => createEnterpriseChat({
            enterpriseId: tab.enterpriseID,
            senderId: descriptor.id,
            receiverId: tab.user.user.id,
        }),
        onSuccess: () => refetch()
    })

    if (isError) {
        createMutation.mutate()
    }
    useEffect(() => {
        if (isSuccess) {
            setReplies(data.findEnterpriseChat.replies as ChatReply[])
        }
    }, [data?.findEnterpriseChat.replies, isSuccess])

    const { isReady, incomingMsg, send } = useChat(descriptor.id, ChatSubProtocols.ENTERPRISE_CHAT, jwt)
    const [content, setContent] = useState<string>("")

    const handleReply = () => {
        const commandData: replyCommandData = {
            senderId: descriptor.id,
            chatId: data?.findEnterpriseChat.id!,
            content: content,
        }
        const payload: enterpriseSubProtocolPayload = {
            subProtocolDataType: enterpriseChatSubProtocolDataType.ReplyCommand,
            command: encodeToBinary(JSON.stringify(commandData))
        }

        if (isReady) {
            send(payload)
            setContent("")
        }
    }


    useEffect(() => {
        const toMarkAsSeen: markAsSeenData[] = []

        for (let i = 0; i < replies.length; i++) {
            if (replies[i].sender.id != descriptor.id && replies[i].seen === false) {
                toMarkAsSeen.push({
                    id: data?.findEnterpriseChat.id!,
                    replyId: replies[i].id
                })
            }
        }

        if (toMarkAsSeen.length > 0) {
            const payload: enterpriseSubProtocolPayload = {
                subProtocolDataType: enterpriseChatSubProtocolDataType.MarkAsSeenCommand,
                command: encodeToBinary(JSON.stringify(toMarkAsSeen))
            }
            send(payload)
        }
    }, [data, descriptor.id, replies, send])

    useEffect(() => {
        if (incomingMsg) {
            const msg = incomingMsg as enterpriseSubProtocolResult
            switch (msg.subProtocolDataType) {
                case enterpriseChatSubProtocolDataType.ChatReply: {
                    const newR = JSON.parse(decodeFromBinary(msg.result)) as ChatReply
                    setReplies(prev => {
                        prev.unshift(newR)
                        return [...prev]
                    })
                    break;
                }
                case enterpriseChatSubProtocolDataType.Chat: {
                    const newC = JSON.parse(decodeFromBinary(msg.result)) as Chat
                    setReplies(newC.replies as ChatReply[])
                    break;
                }
            }
        }
    }, [incomingMsg])

    const { scrollTop } = useScroll()
    const tabRef = useRef<HTMLDivElement>(null)
    useEffect(() => {
        const footer = document.getElementById("footer-element")
        const body = document.getElementById("body-element")
        if (!body || !footer || !tabRef.current) {
            return
        }
        const bodyRect = body.getBoundingClientRect()
        const footerRect = footer.getBoundingClientRect()
        const tabRect = tabRef.current.getBoundingClientRect()
        const a = bodyRect.height - footerRect.height
        const b = tabRect.y + scrollTop
        const c = a - b
        if (c < 56) {
            minimizeTab(tab.user.id, true)
        }
    }, [minimizeTab, scrollTop, tab.user.id])


    useClickOutside(emojiRef, () => { setShowEmoji(false) })
    const windowSize = useWindowSize()
    if (windowSize.width && windowSize.width >= 768) {
        return (
            <>
                {tab.isActive ?
                    <div
                        ref={tabRef}
                        className={`fixed -bottom-1 px-2`} style={{ right: ml }}>
                        <div className="flex flex-col h-[216px] w-52 md:bg-white md:border" >
                            <div className="flex justify-end pr-1 gap-1">
                                <button onMouseUp={() => minimizeTab(tab.user.id, true)} type="button" className="text-xs text-gray-700">-</button>
                                <button onMouseUp={() => closeTab(tab.user.id)} type="button" className="text-xs text-gray-700">x</button>
                            </div>
                            <ul className="max-h-2/3 p-0.5 overflow-y-auto flex flex-col-reverse">
                                {replies.map((reply) => {
                                    return (
                                        <EnterpriseChatReply
                                            key={reply!.id}
                                            direction={reply!.sender.id === tab.user.id ? 'ltr' : 'rtl'}
                                            reply={reply!} />
                                    )
                                })}
                            </ul>
                            <div className="flex mt-auto py-2 px-1 bg-white border">
                                <button
                                    onMouseUp={() => { setShowEmoji(!showEmoji) }}
                                    type="button" className=" text-gray-700 rounded-lg cursor-pointer pr-1">
                                    <svg className="w-4 h-4 fill-yellow-500" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                                        <path
                                            stroke="currentColor"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="2"
                                            d="M13.408 7.5h.01m-6.876 0h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0ZM4.6 11a5.5 5.5 0 0 0 10.81 0H4.6Z" />
                                    </svg>
                                    <span className="sr-only">Add emoji</span>
                                </button>
                                {showEmoji && <div
                                    ref={emojiRef}
                                    className="absolute -top-[80px] -left-64 z-50">
                                    <EmojiPicker
                                        onEmojiClick={(emojiData, _: MouseEvent) => { setContent(p => p + emojiData.emoji); return; }}
                                        theme={Theme.LIGHT}
                                        emojiStyle={EmojiStyle.GOOGLE}
                                        searchDisabled={true}
                                        skinTonesDisabled={true}
                                        previewConfig={{ showPreview: false } as PreviewConfig}
                                        width={265}
                                        height={250} />
                                </div>}
                                <input
                                    value={content}
                                    onChange={(e) => setContent(e.target.value)}
                                    type="text" className="text-xs px-0.5 w-full h-8  bg-white focus:outline-none" />
                                <button
                                    onMouseUp={() => handleReply()}
                                    id="submit-btn"
                                    type="button"
                                    className="px-1 w-5 h-5 my-auto ms-1">
                                    <svg
                                        className="w-3 h-3 rotate-90 rtl:-rotate-90 fill-gray-700"
                                        aria-hidden="true"
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 18 20">
                                        <path d="m17.914 18.594-8-18a1 1 0 0 0-1.828 0l-8 18a1 1 0 0 0 1.157 1.376L8 18.281V9a1 1 0 0 1 2 0v9.281l6.758 1.689a1 1 0 0 0 1.156-1.376Z" />
                                    </svg>
                                </button>
                            </div>
                        </div>
                    </div>
                    : <div className={`fixed bottom-0 px-2  `} style={{ right: ml }}>
                        <div className="flex  bg-blue-500 w-52 h-[45px]">
                            <button
                                onMouseUp={() => minimizeTab(tab.user.id, false)}
                                className="cursor-pointer flex flex-col w-full ">
                                <p className="text-center text-white my-auto text-md ms-2">
                                    {tab.user.user.person.firstName} {' '} {tab.user.user.person.lastName}
                                </p>
                            </button>
                            <button onMouseUp={() => closeTab(tab.user.id)} type="button" className="ms-auto  mr-2 mb-0.5 text-md text-white">x</button>
                        </div>
                    </div>
                }
            </>
        )
    } else {
        return (
            < div >
                <div className="flex flex-col h-[19rem] w-full md:bg-white md:border relative border-t" >
                    <div className="flex justify-end pr-1 gap-1">
                        <button onMouseUp={() => closeTab(tab.user.id)} type="button" className=" text-gray-700">x</button>
                    </div>
                    <ul className="max-h-2/3 p-0.5 overflow-y-auto px-2 flex flex-col-reverse">
                        {replies.map((reply) => {
                            return (
                                <EnterpriseChatReply
                                    key={reply.id}
                                    direction={reply.sender.id === tab.user.user.id ? 'ltr' : 'rtl'}
                                    reply={reply} />
                            )
                        })}
                    </ul>
                    <div className="flex mt-auto py-2 px-1 group bg-gradient-to-br from-cyan-500 to-blue-500">
                        <button
                            onMouseUp={() => { setShowEmoji(!showEmoji) }}
                            type="button" className=" text-gray-900 rounded-lg cursor-pointer p-2">
                            <svg className="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M13.408 7.5h.01m-6.876 0h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0ZM4.6 11a5.5 5.5 0 0 0 10.81 0H4.6Z" />
                            </svg>
                            <span className="sr-only">Add emoji</span>
                        </button>
                        <input
                            value={content}
                            onChange={(e) => setContent(e.target.value)}
                            type="text" className="text-xs px-0.5 w-full h-8  bg-gray-100 focus:outline-none" />
                        <button
                            onMouseUp={() => handleReply()}
                            id="submit-btn"
                            type="button"
                            className="px-1 w-5 h-5 my-auto ms-1">
                            <svg className="w-3 h-3 rotate-90 rtl:-rotate-90" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="#ffffff" viewBox="0 0 18 20">
                                <path d="m17.914 18.594-8-18a1 1 0 0 0-1.828 0l-8 18a1 1 0 0 0 1.157 1.376L8 18.281V9a1 1 0 0 1 2 0v9.281l6.758 1.689a1 1 0 0 0 1.156-1.376Z" />
                            </svg>
                        </button>
                        {showEmoji && <div
                            ref={emojiRef}
                            className="absolute -top-[83px] -left-[1px] z-10">
                            <EmojiPicker
                                onEmojiClick={(emojiData, _: MouseEvent) => { setContent(p => p + emojiData.emoji); return; }}
                                theme={Theme.LIGHT}
                                emojiStyle={EmojiStyle.GOOGLE}
                                searchDisabled={true}
                                skinTonesDisabled={true}
                                previewConfig={{ showPreview: false } as PreviewConfig}
                                width={265}
                                height={250} />
                        </div>}
                    </div>
                </div>
            </div >
        )
    }
}


export default ChatTab;


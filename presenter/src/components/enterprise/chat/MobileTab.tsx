/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState, useEffect, useRef } from "react";
import useSubscriber from "../../../lib/hooks/useSubscriber";
import { Query, EnterpriseChatQuery, EnterpriseChatTypeCast } from "../../../lib/queries";
import { EnterpriseChatReplyType, EnterpriseChatType, UserDescriptor } from "../../../lib/types";
import { replyEnterpriseChat } from "./enterprise_chat";
import EnterpriseChatReply from "./EnterpriseChatReply";
import { TabType } from "./enterprise_chat";
import EmojiPicker, { Theme, EmojiStyle } from "emoji-picker-react";
import { PreviewConfig } from "emoji-picker-react/dist/config/config";
import useOutsideDenier from "../../../lib/hooks/useOutsideDenier";
import { isReplies, markAsSeen } from "../../../lib/mark_enterprise_chat_reply_as_seen";

interface Props {
    closeTab: (id: string) => void;
    descriptor: UserDescriptor;
    tab: TabType;
    callback: (v: boolean) => void
}


function MobileTab({ callback, tab, descriptor, closeTab }: Props) {
    const chatID = `chat:${tab.enterpriseID}=${descriptor.email}#${tab.user.email}`
    const [replies, setReplies] = useState<EnterpriseChatReplyType[]>([])
    const [content, setContent] = useState<string>("")
    const { incomingMessage } = useSubscriber(chatID)
    const Reply = async () => {
        await replyEnterpriseChat(tab.enterpriseID, descriptor, tab.user, content).catch((e) => console.error(e))
        setContent("")
    }
    useEffect(() => {
        if (incomingMessage) {
            if (isReplies(incomingMessage)) {
                const msgs = incomingMessage as EnterpriseChatReplyType[];
                setReplies(p => p?.map((reply) => {
                    const found = msgs.find((msg) => msg.id === reply.id)
                    return found ? found : reply
                }))
                return
            }
            const msg = incomingMessage as EnterpriseChatReplyType
            if (msg.user.email !== descriptor!.email && msg.seen === false) {
                markAsSeen(tab.enterpriseID, [msg]);
            }
            setReplies(prev => { return [...prev, msg] })
        }
    }, [incomingMessage])

    useEffect(() => {
        const fetchChat = async () => {
            const chat = await Query<EnterpriseChatType>(
                EnterpriseChatQuery,
                EnterpriseChatTypeCast,
                [tab.enterpriseID, chatID]
            )
            const received = chat.replies?.filter((reply) => reply.user.email !== descriptor!.email && !reply.seen)
            if (received?.length) {
                const load = async () => {
                    markAsSeen(tab.enterpriseID, received).then(() => callback(true))
                }
                load()
            }
            setReplies(chat.replies)
        }
        fetchChat()
    }, [])

    const lastRef = useRef<HTMLLIElement>(null)
    useEffect(() => {
        if (lastRef.current) {
            lastRef.current.scrollIntoView({ behavior: "instant" })
        }
    }, [replies])
    const emojiRef = useRef(null)
    const [showEmoji, setShowEmoji] = useState<boolean>(false)
    useOutsideDenier(emojiRef, () => { setShowEmoji(false) })
    return (
        <div>
            <div className="flex flex-col h-[216px] w-full md:bg-white md:border relative" >
                <div className="flex justify-end pr-1 gap-1">
                    <button onMouseUp={() => closeTab(tab.user.email)} type="button" className="text-xs text-gray-700">x</button>
                </div>
                <ul className="max-h-2/3 p-0.5 overflow-y-auto px-2">
                    {replies.map((reply) => {
                        return (
                            <EnterpriseChatReply
                                key={reply.id}
                                direction={reply.user.email === tab.user.email ? 'ltr' : 'rtl'}
                                reply={reply} />
                        )
                    })}
                    <li className="" ref={lastRef}></li> {/* Last element for scrolling */}
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
                        onMouseUp={() => Reply()}
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
        </div>
    )
}


export default MobileTab;


/* eslint-disable @typescript-eslint/no-explicit-any */
import { useContext, useEffect, useState } from "react";
import CircleFillIcon from "../../icons/CircleFillIcon";
import useWindowDimensions from "../../../lib/hooks/useWindowsDimensions";
import ArrowForwardIcon from "../../icons/ArrowForwardIcon";
import ChatIcon from "../../icons/ChatIcon";
import { Route } from "../../../routes/$enterpriseID";
import useSubscriber from "../../../lib/hooks/useSubscriber";
import { EnterpriseChatReplyType, User } from "../../../lib/types";
import Tab from "./Tab";
import { isPair, isReply, TabType } from "./enterprise_chat";
import MobileTab from "./MobileTab";
import { SideBarContext } from "../../../react-context/SideBarContext";
import { useRouter } from "@tanstack/react-router";
const SideChat: React.FC = () => {
    const { descriptor, onlineUsers } = Route.useLoaderData()
    const { enterpriseID } = Route.useParams()
    const router = useRouter()
    const { rightBarOpen, setRightBarOpen } = useContext(SideBarContext)
    const [online, setOnline] = useState<User[]>(onlineUsers || [])
    const [tabs, setTabs] = useState<TabType[]>([])
    //const [isRightbarOpen, setIsRightbarOpen] = useState<boolean>(false)
    const { width } = useWindowDimensions()
    const { incomingMessage } = useSubscriber(`chat:${enterpriseID}?chat:${enterpriseID}=${descriptor!.email}`)//

    const handleReload = (v: boolean) => {
        if (v) {
            router.invalidate()
        }
    }

    const handleRightBar = (v: boolean) => {
        setRightBarOpen(v)
    }

    const closeTab = (id: string) => {
        setTabs(prev => prev.filter(tab => tab.user.email !== id))
        setOnline(prev => prev.map(user => {
            if (user.email === id) {
                return { ...user, msgs: 0 }
            }
            return user
        }))
    }

    const minimizeTab = (id: string, v: boolean) => {
        const newTabs = tabs.map(tab => {
            if (tab.user.email === id) {
                return { ...tab, isActive: !v }
            }
            return tab
        })
        setTabs(newTabs)
    }

    const openTab = (user: User) => {
        const prev = tabs
        if (width >= 760 && prev.length >= 3) {
            prev.shift()
        } else if (width < 760 && prev.length >= 1) {
            prev.shift()
        }
        const exists = prev.find(tab => tab.user.email === user.email)
        if (exists) {
            return prev
        }
        const existUser = online.find(u => u.email === user.email)
        if (!existUser) return prev
        const emptyMsgsUser = { ...existUser, msgs: 0 }
        const newTabs = [...prev, { user: emptyMsgsUser, enterpriseID, index: prev.length, isActive: true }]
        const newUsers = online.map(u => {
            if (u.email === user.email) {
                return { ...u, msgs: 0 }
            }
            return u
        })
        setTabs(newTabs)
        setOnline(newUsers)
    }

    useEffect(() => {
        if (isPair(incomingMessage)) {
            setOnline(prev => {
                return prev.map(user => {
                    if (user.email === incomingMessage.one) {
                        return { ...user, status: incomingMessage.two }
                    }
                    return user
                })
            })
        } else if (isReply(incomingMessage)) {
            const reply = incomingMessage as EnterpriseChatReplyType
            if (reply.user.email === descriptor?.email) return
            setOnline(prev => {
                return prev.map(user => {
                    const existsTab = tabs.find(tab => tab.user.email === reply.user.email && tab.isActive)
                    if (existsTab) {
                        return user;
                    }
                    const existUser = prev.find(user => user.email === reply.user.email)
                    if (existUser) {
                        return { ...existUser, msgs: existUser.msgs ? existUser.msgs + 1 : 1 }
                    }
                    return user
                })
            })
        }
    }, [descriptor?.email, incomingMessage, tabs])

    return (
        <>
            {rightBarOpen ?
                <div id="drawer-right-example"
                    className="fixed top-20 right-0 z-100 h-screen overflow-y-auto border-l bg-white w-full md:w-80">
                    <button
                        onMouseUp={() => handleRightBar(false)}
                        type="button"
                        className="mt-6 w-12 items-end rounded-r-xl cursor-pointer flex flex-col
                         group bg-gradient-to-br from-cyan-500 to-blue-500">
                        <ArrowForwardIcon fill="#ffffff" />
                    </button>
                    <div className="divide-y-1">
                        <ul>
                            {
                                online.map((user) => {
                                    return (
                                        <li
                                            key={user.email}
                                            onMouseUp={() => openTab(user)}
                                            className="flex p-5 h-[200px] md:h-full cursor-pointer">
                                            <CircleFillIcon fill={user.status === "ONLINE" ? '#93c5fd' : '#fca5a5'} />
                                            <p className="ms-10 text-gray-700 text-sm md:text-xl"> {user.firstName} {' '} {user.lastName} </p>
                                            {user.msgs > 0 ?
                                                <div className="ms-auto mr-2 h-6 w-6  bg-blue-300 flex flex-col justify-center  rounded-full">
                                                    <p className="text-gray-500 text-center text-sm md:text-xl">
                                                        {user.msgs}
                                                    </p>
                                                </div> : null
                                            }
                                        </li>
                                    )
                                })
                            }

                        </ul>
                        {
                            width < 768 ? tabs.map((tab) => {
                                return (
                                    <MobileTab
                                        callback={handleReload}
                                        tab={tab}
                                        key={tab.index}
                                        closeTab={closeTab}
                                        descriptor={descriptor!}
                                    />
                                )
                            }) : null
                        }
                    </div>
                </div>
                :
                <div className="fixed z-200 top-[430px] right-0 md:top-[550px] md:right-2  text-center">
                    <button
                        onMouseUp={() => handleRightBar(true)}
                        className="text-white  group bg-gradient-to-br from-cyan-500 to-blue-500  font-medium
                        rounded-lg text-sm px-5 py-2.5 mb-2"
                        type="button">
                        <ChatIcon fill="#ffffff" />
                    </button>
                </div>
            }
            {

                width >= 768 ? tabs.map((tab) => {
                    return (
                        <Tab
                            callback={handleReload}
                            minimizeTab={minimizeTab}
                            descriptor={descriptor!}
                            key={tab.index}
                            closeTab={closeTab}
                            isRightBarOpen={rightBarOpen}
                            tab={tab}
                        />
                    )
                }) : null
            }

        </>
    )


}


export default SideChat;
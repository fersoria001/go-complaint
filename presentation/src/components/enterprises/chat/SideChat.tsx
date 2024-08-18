'use client'
import ChatIcon from "@/components/icons/ChatIcon"
import KeyboardArrowRightIcon from "@/components/icons/KeyboardArrowRightIcon"
import { Employee, Enterprise } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import graphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { useSuspenseQueries } from "@tanstack/react-query"
import { useEffect, useState } from "react"
import OnlineEmployeesList from "./OnlineEmployeesList"
import useWindowSize from "@/lib/hooks/useWindowsSize"
import ChatTab, { TabType } from "./ChatTab"
import { getCookie } from "@/lib/actions/cookies"
import enterpriseByIdSubscription from "@/graphql/subscriptions/enterpriseByIdSubscription"
import clsx from "clsx"
import getGraphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"

export enum enterpriseChatSubProtocolDataType {
    MarkAsSeenCommand = "mark_as_seen_command",
    ReplyCommand = "reply_command",
    ChatReply = "chat_reply",
    Chat = "chat",
}

export type enterpriseSubProtocolResult = {
    subProtocolDataType: enterpriseChatSubProtocolDataType;
    result: any;
}

export type enterpriseSubProtocolPayload = {
    subProtocolDataType: enterpriseChatSubProtocolDataType;
    command: any;
}

export type createEnterpriseChatCommandData = {
    enterpriseId: string;
    senderId: string;
    receiverId: string;
}

export type replyCommandData = {
    chatId: string;
    senderId: string;
    content: string;
}

export type markAsSeenData = {
    id: string
    replyId: string
}

const SideChat: React.FC = () => {
    const [{ data }, { data: jwtToken }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await getGraphQLClient().request(userDescriptorQuery),
                staleTime: Infinity,
                gcTime: Infinity,
            },
            {
                queryKey: ['server-side-jwt'],
                queryFn: async () => {
                    const jwt = await getCookie("jwt")
                    if (!jwt) return ""
                    return jwt
                }
            }
        ],
    })

    const [tabs, setTabs] = useState<TabType[]>([])
    const [rightBarOpen, setRightBarOpen] = useState<boolean>(false)
    const windowSize = useWindowSize()

    const handleRightBar = (v: boolean) => {
        setRightBarOpen(v)
    }

    const openTab = (user: Employee) => {
        setTabs(prev => {
            if (windowSize.width! >= 760 && prev.length >= 3) {
                prev.shift()
            } else if (windowSize.width! < 760 && prev.length >= 1) {
                prev.shift()
            }
            const exists = prev.find(tab => tab.user.id === user.id)
            if (exists) {
                return prev
            }
            return [...prev, { user: user, enterpriseID: user.enterpriseId, index: prev.length, isActive: true }]
        })
    }

    const minimizeTab = (id: string, v: boolean) => {
        setTabs(p => p.map(tab => {
            if (tab.user.id === id) {
                return { ...tab, isActive: !v }
            }
            return tab
        }))
    }

    const closeTab = (id: string) => {
        setTabs(prev => prev.filter(tab => tab.user.id !== id))
    }

    const [enterprise, setEnterprise] = useState<Enterprise>()
    const [selectedEnterpriseId, setSelectedEnterpriseId] = useState<string>("")
    useEffect(() => {
        async function subscribe() {
            const subscription = getGraphQLSubscriptionClient().iterate({
                query: enterpriseByIdSubscription(selectedEnterpriseId, data.userDescriptor.id),
            });
            for await (const event of subscription) {
                const c = event.data?.enterpriseById as Enterprise
                setEnterprise(c)
            }
        }
        if (selectedEnterpriseId != "") {
            subscribe()
        } else {
            if (data) {
                setSelectedEnterpriseId(data.userDescriptor.authorities![0]!.enterpriseId)
            }
        }
    }, [selectedEnterpriseId, data])
    return (
        <>
            {rightBarOpen ?
                <div id="drawer-right-example"
                    className="absolute top-20
                    md:fixed md:top-0 right-0 md:pt-[82px] z-10 h-screen min-h-screen overflow-y-auto border-l bg-white w-full md:w-80">
                    <button
                        onMouseUp={() => handleRightBar(false)}
                        type="button"
                        className="mb-6 mt-6 w-12 items-end rounded-r-xl cursor-pointer flex flex-col outline outline-gray-700">
                        <KeyboardArrowRightIcon className="h-6 w-6 fill-gray-700" />
                    </button>
                    <div className="divide divide-x-1">

                        <div className="pl-2 overflow-x-auto">
                            {
                                data.userDescriptor.authorities?.map((enterprise) => {
                                    return (
                                        <button
                                            key={enterprise?.enterpriseId}
                                            type="button"
                                            onClick={() => { setSelectedEnterpriseId(enterprise?.enterpriseId!) }}
                                            className={clsx("bg-blue-500 rounded-lg px-2.5 py-0.5 text-white font-bold", {
                                                "scale 110 bg-blue-700": selectedEnterpriseId === enterprise?.enterpriseId,
                                                "hover:bg-blue-600": selectedEnterpriseId != enterprise?.enterpriseId
                                            })}>
                                            {enterprise?.principal}
                                        </button>
                                    )
                                })
                            }
                        </div>
                        {
                            enterprise &&
                            <OnlineEmployeesList
                                openTab={openTab}
                                items={enterprise.employees as Employee[]} />
                        }

                        {
                            windowSize.width! < 768 ? tabs.map((tab) => {
                                return (
                                    <ChatTab
                                        jwt={jwtToken as string}
                                        minimizeTab={minimizeTab}
                                        descriptor={data.userDescriptor}
                                        key={tab.index}
                                        closeTab={closeTab}
                                        isRightBarOpen={rightBarOpen}
                                        tab={tab}
                                    />
                                )
                            }) : null
                        }

                    </div>
                </div>
                :
                <div className="fixed z-200 top-[430px] right-0 md:top-[500px] md:right-4  text-center">
                    <button
                        onMouseUp={() => handleRightBar(true)}
                        className="text-white  group bg-gradient-to-br from-cyan-500 to-blue-500  font-medium
                        rounded-lg text-sm px-5 py-2.5 mb-2"
                        type="button">
                        <ChatIcon fill="#ffffff" className="w-6 h-6" />
                    </button>
                </div>
            }
            {

                windowSize.width! >= 768 ? tabs.map((tab) => {
                    return (
                        <ChatTab
                            jwt={jwtToken as string}
                            minimizeTab={minimizeTab}
                            descriptor={data.userDescriptor}
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
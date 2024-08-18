'use client'
import getGraphQLClient from "@/graphql/graphQLClient"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { useSuspenseQuery } from "@tanstack/react-query"
import ComplaintItem from "./ComplaintItem"
import Link from "next/link"
import FilterBy from "../search/FilterBy"
import { useEffect, useState } from "react"
import clsx from "clsx"
import { Complaint } from "@/gql/graphql"
import graphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"
import complaintsSubscription from "@/graphql/subscriptions/complaintsSubscription"
import { getCookie, setCookie } from "@/lib/actions/cookies"
import getGraphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"


enum ComplaintsFilter {
    ALL = "ALL",
    RECEIVED = "RECEIVED",
    SENT = "SENT"
}
const options = [
    {
        id: "0",
        name: "All",
        value: ComplaintsFilter.ALL
    },
    {
        id: "1",
        name: "Received",
        value: ComplaintsFilter.RECEIVED
    },
    {
        id: "2",
        name: "Sent",
        value: ComplaintsFilter.SENT
    },
]

const ComplaintsMain: React.FC = () => {
    const { data } = useSuspenseQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => getGraphQLClient().request(userDescriptorQuery),
        staleTime: Infinity,
        gcTime: Infinity
    })
    const [alias, setAlias] = useState<string>("")
    const [filter, setFilter] = useState<ComplaintsFilter>(ComplaintsFilter.ALL)
    const [msgIfEmpty, setMsgIfEmpty] = useState<string>("You have no complaints in your draft, send one by clicking the symbol in the bottom right corner.")
    const [complaints, setComplaints] = useState<Complaint[]>([])
    const [filtered, setFiltered] = useState<Complaint[]>([])
    const handleAliasSelect = async (value: string) => {
        const ok = await setCookie("alias", value)
        if (!ok) {
            console.error("couln't set the alias")
            return
        }
        setAlias(ok)
    }
    useEffect(() => {
        let filtered = [...complaints]
        switch (filter) {
            case ComplaintsFilter.ALL:
                setMsgIfEmpty("You have no complaints in your draft, send one by clicking the symbol in the bottom right corner.")
                break
            case ComplaintsFilter.RECEIVED:
                setMsgIfEmpty("You have not receiving any complaint yet.")
                filtered = filtered.filter((c) => c.receiver!.id === alias)
                break
            case ComplaintsFilter.SENT:
                setMsgIfEmpty("You have not send any complaint yet, start by sending one! click the symbol in the bottom right corner.")
                filtered = filtered.filter((c) => c.author!.id === alias)
                break
        }
        setFiltered(filtered)
    }, [complaints, alias, filter])
    
    useEffect(() => {
        async function subscribe() {
            const id = await getCookie("alias")
            if (!alias) {
                setAlias(id!)
            }
            const subscription = getGraphQLSubscriptionClient().iterate({
                query: complaintsSubscription(id!, data.userDescriptor.id),
            });
            for await (const event of subscription) {
                const c = event.data?.complaints as Complaint
                setComplaints(prev => {
                    const copy = prev.map((m) => m.id != c.id ? m : c)
                    const index = copy.findIndex((m) => m.id === c.id)
                    if (index < 0) copy.push(c)
                    return copy
                })
            }
        }
        subscribe()
    }, [alias, data.userDescriptor.id])
    return (
        <>
            <div className="flex flex-col mt-4">
                <div className="flex gap-5 overflow-x-auto w-full self-center my-2 whitespace-nowrap px-2.5">
                    <button
                        type="button"
                        onClick={() => { handleAliasSelect(data.userDescriptor.id) }}
                        className={clsx("bg-blue-500 rounded-lg px-2.5 py-0.5 text-white font-bold", {
                            "scale 110 bg-blue-700": alias === data.userDescriptor.id,
                            "hover:bg-blue-600": alias != data.userDescriptor.id
                        })}>
                        {data.userDescriptor.fullName.split(" ")[0]}(you)
                    </button>
                    {
                        data.userDescriptor.authorities?.map((enterprise) => {
                            return (
                                <button
                                    key={enterprise?.enterpriseId}
                                    type="button"
                                    onClick={() => { handleAliasSelect(enterprise?.enterpriseId!) }}
                                    className={clsx("bg-blue-500 rounded-lg px-2.5 py-0.5 text-white font-bold", {
                                        "scale 110 bg-blue-700": alias === enterprise?.enterpriseId,
                                        "hover:bg-blue-600": alias != enterprise?.enterpriseId
                                    })}>
                                    {enterprise?.principal}
                                </button>
                            )
                        })
                    }
                </div>
                <FilterBy options={options} callback={(v: string) => { setFilter(v as ComplaintsFilter) }} />
                <div className="w-full mt-4">
                    {
                        filtered.length > 0 ? filtered.map((complaint) => {
                            return (
                                <Link href={`/complaints/${complaint?.id}`} key={complaint?.id}>
                                    <ComplaintItem currentId={alias} item={complaint!} />
                                </Link>
                            )
                        }) :
                            <div>{msgIfEmpty}</div>
                    }
                </div>
            </div>
        </>
    )
}
export default ComplaintsMain
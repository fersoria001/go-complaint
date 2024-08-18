'use client'
import { ComplaintDataType, ComplaintData, UserDescriptor, Roles } from "@/gql/graphql";
import graphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient";
import { useState, useEffect } from "react";
import { TimeData } from "../charts/ChartBuilder";
import TimeDataChart from "../charts/TimeDataChart";
import complaintDataByOwnershipSubscription from "@/graphql/subscriptions/complaintDataByOwnership";
import { getCookie, setCookie } from "@/lib/actions/cookies";
import { useRouter } from "next/navigation";
import clsx from "clsx";

interface Props {
    currentUser: UserDescriptor
}

const ComplaintDataCharts: React.FC<Props> = ({ currentUser }: Props) => {
    const [complaintData, setComplaintdata] = useState<Map<ComplaintDataType, ComplaintData[]>>(
        new Map<ComplaintDataType, ComplaintData[]>([
            [ComplaintDataType.Received, []],
            [ComplaintDataType.Resolved, []],
            [ComplaintDataType.Reviewed, []],
            [ComplaintDataType.Sent, []]
        ])
    )
    const [alias, setAlias] = useState<string>("")
    const router = useRouter()
    const handleAliasSelect = async (value: string) => {
        const ok = await setCookie("alias", value)
        if (!ok) {
            console.error("couln't set the alias")
            return
        }
        setAlias(ok)
    }
    useEffect(() => {
        setComplaintdata(new Map<ComplaintDataType, ComplaintData[]>([
            [ComplaintDataType.Received, []],
            [ComplaintDataType.Resolved, []],
            [ComplaintDataType.Reviewed, []],
            [ComplaintDataType.Sent, []]
        ]))
        async function subscribeToComplaintData() {
            let localAlias = alias
            if (!alias) {
                const ok = await getCookie("alias")
                if (!ok) {
                    return router.push("/")
                }
                localAlias = ok
            }
            const subscription = graphQLSubscriptionClient.iterate({
                query: complaintDataByOwnershipSubscription(localAlias),
            });
            for await (const event of subscription) {
                const c = event.data?.complaintDataByOwnership as ComplaintData
                setComplaintdata(prev => {
                    const arr = prev.get(c.dataType)
                    prev.set(c.dataType, [...arr!, c])
                    return new Map(prev)
                })
            }
        }
        subscribeToComplaintData()
    }, [router, alias])
    return (
        <div className="lg:p-5 flex flex-col lg:items-center">
            <div className="flex gap-5 overflow-x-auto w-full self-center my-2 whitespace-nowrap px-2.5">
                <button
                    type="button"
                    onClick={() => { handleAliasSelect(currentUser.id) }}
                    className={clsx("bg-blue-500 rounded-lg px-2.5 py-0.5 text-white font-bold", {
                        "scale 110 bg-blue-700": alias === currentUser.id,
                        "hover:bg-blue-600": alias != currentUser.id
                    })}>
                    {currentUser.fullName.split(" ")[0]}(you)
                </button>
                {
                    currentUser.authorities?.map((enterprise) => {
                        if (enterprise?.authority === Roles.Owner) {
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
                        }
                    })
                }
            </div>
            <div className="flex flex-col md:flex-row mb-4 gap-2">
                <div className="w-full lg:w-[480px] xl:w-[540px]">
                    <TimeDataChart chartLabel={"Complaints sent"} yLabel="complaints" data={complaintData.get(ComplaintDataType.Sent) as TimeData[]} />
                </div>
                <div className="w-full lg:w-[480px] xl:w-[540px]">
                    <TimeDataChart chartLabel={"Complaints received"} yLabel="complaints" data={complaintData.get(ComplaintDataType.Received) as TimeData[]} />
                </div>
            </div>
            <div className="flex flex-col md:flex-row mb-4 gap-2">
                <div className="w-full lg:w-[480px] xl:w-[540px]">
                    <TimeDataChart chartLabel={"Complaints resolved"} yLabel="complaints" data={complaintData.get(ComplaintDataType.Resolved) as TimeData[]} />
                </div>
                <div className="w-full lg:w-[480px] xl:w-[540px]">
                    <TimeDataChart chartLabel={"Complaints reviewed"} yLabel="complaints" data={complaintData.get(ComplaintDataType.Reviewed) as TimeData[]} />
                </div>
            </div>
        </div>
    )
}
export default ComplaintDataCharts;
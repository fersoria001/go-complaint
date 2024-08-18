'use client'
import ArrowSortIcon from "@/components/icons/ArrowSortIcon"
import KeyboardArrowDownIcon from "@/components/icons/KeyboardArrowDownIcon"
import KeyboardArrowRightIcon from "@/components/icons/KeyboardArrowRightIcon"
import { EnterpriseActivity } from "@/gql/graphql"
import getGraphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"
import graphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient"
import employeesActivityLogSubscription from "@/graphql/subscriptions/employeesActivityLogSubscription"
import { dateFromMsString } from "@/lib/dateFromMsString"
import Link from "next/link"
import { useParams } from "next/navigation"
import { useEffect, useState } from "react"

const EmployeesActivityLog: React.FC = () => {
    const [show, setShow] = useState<boolean>(true)
    const { enterpriseId } = useParams()
    const [log, setLog] = useState<EnterpriseActivity[]>([])
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    useEffect(() => {
        async function subscribe() {
            const subscription = getGraphQLSubscriptionClient().iterate({
                query: employeesActivityLogSubscription(enterpriseName),
            });
            for await (const event of subscription) {
                const c = event.data?.employeesActivityLog as EnterpriseActivity
                setLog(prev => [...prev, c])
            }
        }
        subscribe()
    }, [enterpriseName])
    return (
        <div>
            <div className="flex items-center">
                <div className="flex w-full bg-gray-200 h-0.5"></div>
                <h3 className="text-gray-700 text-md lg:text-xl whitespace-nowrap px-2.5 font-bold cursor-default">
                    Employees activity
                </h3>
                <div className="flex w-1/6 shrink bg-gray-200 h-0.5 ps-10 md:ps-[5.5rem] lg:ps-[7rem] xl:ps-[9.1rem]"></div>
                {
                    show ?
                        <span className="py-3">
                            <KeyboardArrowDownIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700 cursor-pointer" />
                        </span>
                        :
                        <span className="py-3">
                            <KeyboardArrowRightIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700 cursor-pointer" />
                        </span>
                }
            </div>

            {show && <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
                <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" className="px-6 py-3">
                                Person name
                            </th>
                            <th scope="col" className="px-6 py-3">
                                <div className="flex items-center">
                                    Enterprise Action

                                    <ArrowSortIcon className="w-3 h-3 ms-1.5 cursor-pointer" />

                                </div>
                            </th>
                            <th scope="col" className="px-6 py-3">
                                <div className="flex items-center">
                                    Occurred on

                                    <ArrowSortIcon className="w-3 h-3 ms-1.5 cursor-pointer" />

                                </div>
                            </th>
                            <th scope="col" className="px-6 py-3">
                                <span className="sr-only">Details</span>
                            </th>
                        </tr>
                    </thead>

                    <tbody>
                        {
                            log.length > 0 && log.map((v) => {
                                return (
                                    <tr
                                        key={v.id}
                                        className="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                                        <th
                                            scope="row"
                                            className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                            {v.user.subjectName}
                                        </th>
                                        <td className="px-6 py-4">
                                            {v.activityType}
                                        </td>
                                        <td className="px-6 py-4">
                                            {dateFromMsString(v.occurredOn).toUTCString()}
                                        </td>
                                        <td className="px-6 py-4 text-right">
                                            <Link
                                                href={`/enterprises/${enterpriseName}/employees/activity?id=${v.user.id}`}
                                                className="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                                Details
                                            </Link>
                                        </td>
                                    </tr>
                                )
                            })
                        }

                    </tbody>
                </table>
            </div>
            }


        </div>
    )
}

export default EmployeesActivityLog;
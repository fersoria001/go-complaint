'use client'

import KeyboardArrowDownIcon from "@/components/icons/KeyboardArrowDownIcon";
import KeyboardArrowRightIcon from "@/components/icons/KeyboardArrowRightIcon";
import getGraphQLClient from "@/graphql/graphQLClient";
import enterpriseByNameQuery from "@/graphql/queries/enterpriseByNameQuery";
import { useSuspenseQuery } from "@tanstack/react-query";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useState } from "react";

const EmployeesMain: React.FC = () => {
    const params = useParams()
    const { data } = useSuspenseQuery({
        queryKey: ['enterprise', params.enterpriseId as string],
        queryFn: async ({ queryKey }) => {
            try {
                return await getGraphQLClient().request(enterpriseByNameQuery, { name: decodeURIComponent(queryKey[1]) })
            } catch (e:any) {
                console.log(e.response.errors[0])
            }
        }
    })
    const [show, setShow] = useState<boolean>(true)
    return (
        <div>
            <div className="flex flex-col px-5 mt-[2.275rem]">
                <Link
                    href={`/enterprises/${params.enterpriseId}/employees/hire-new`}
                    className="text-gray-700 mb-4 underline text-md">Hire new employees</Link>
                <Link
                    href={`/enterprises/${params.enterpriseId}/employees/feedback`}
                    className="text-gray-700 mb-4 underline text-md">Feedback employees</Link>
            </div>
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
        </div>
    )
}
export default EmployeesMain;
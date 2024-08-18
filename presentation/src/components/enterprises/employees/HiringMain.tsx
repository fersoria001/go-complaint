'use client'
import getGraphQLClient from "@/graphql/graphQLClient"
import hiringProcessByEnterpriseNameQuery from "@/graphql/queries/hiringProcessByEnterpriseNameQuery"
import { useSuspenseQueries, useSuspenseQuery } from "@tanstack/react-query"
import Link from "next/link"
import { useParams } from "next/navigation"
import HiringProccessCard from "./HiringProcessCard"
import { HiringProcess } from "@/gql/graphql"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"

const HiringMain = () => {
    const params = useParams()
    const enterpriseName = decodeURIComponent(params?.enterpriseId as string)
    const [{ data: { userDescriptor } }, { data: { hiringProcessByEnterpriseName: hiringProcesses } }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await getGraphQLClient().request(userDescriptorQuery),
                staleTime: Infinity,
                gcTime: Infinity
            },
            {
                queryKey: ['hiring-process-by-enterprise-id', enterpriseName],
                queryFn: async () => await getGraphQLClient().request(hiringProcessByEnterpriseNameQuery, { name: enterpriseName })
            }
        ]
    })
    return (
        <div>
            <div className="flex flex-col px-5 mt-[2.275rem]">
                <Link
                    href={`/enterprises/${params.enterpriseId}/employees/hiring/hire-new`}
                    className="text-gray-700 mb-4 underline text-md">
                    Hire a new employee
                </Link>
            </div>
            <div>
                {
                    hiringProcesses.map((h) => <HiringProccessCard 
                    key={h?.id!} 
                    currentUser={userDescriptor}
                    hiringProccess={h as HiringProcess} />)
                }
            </div>
        </div>
    )
}

export default HiringMain
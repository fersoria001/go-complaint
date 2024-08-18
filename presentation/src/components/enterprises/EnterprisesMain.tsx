'use client'

import Link from "next/dist/client/link"
import EnterprisesList from "./EnterprisesList"
import OfficesList from "./OfficesList"
import { useSuspenseQueries } from "@tanstack/react-query"
import getGraphQLClient from "@/graphql/graphQLClient"
import enterprisesByAuthenticatedUserQuery from "@/graphql/queries/enterprisesByAuthenticatedUserQuery"
import { EnterpriseByAuthenticatedUser } from "@/gql/graphql"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"

const EnterprisesMain: React.FC = () => {
    const [{ data }, { data: { userDescriptor: user } }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['enterprisesByAuthenticatedUser'],
                queryFn: async () => {
                    const r = await getGraphQLClient().request(enterprisesByAuthenticatedUserQuery)
                    return r.enterprisesByAuthenticatedUser
                },
            },
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await getGraphQLClient().request(userDescriptorQuery),
                staleTime: Infinity,
                gcTime: Infinity
            }
        ]
    })
    return (
        <div className="h-screen relative">
            <div className="flex flex-col px-5 mt-[2.275rem]">
                <Link href="/enterprises/register" className="text-gray-700 mb-4 underline text-md">
                    Register an enterprise.
                </Link>
                <Link href="/hiring" className="text-gray-700 mb-4 underline text-md">
                    Check your invitations to enterprises.
                </Link>
            </div>
            <EnterprisesList enterprises={data.enterprises as EnterpriseByAuthenticatedUser[]} />
            <OfficesList currentUser={user} offices={data.offices as EnterpriseByAuthenticatedUser[]} />
        </div>
    )
}

export default EnterprisesMain
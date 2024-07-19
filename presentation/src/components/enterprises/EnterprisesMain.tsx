'use client'

import Link from "next/dist/client/link"
import EnterprisesList from "./EnterprisesList"
import OfficesList from "./OfficesList"
import { useSuspenseQuery } from "@tanstack/react-query"
import getGraphQLClient from "@/graphql/graphQLClient"
import enterprisesByAuthenticatedUserQuery from "@/graphql/queries/enterprisesByAuthenticatedUserQuery"
import { EnterpriseByAuthenticatedUser } from "@/gql/graphql"

const EnterprisesMain: React.FC = () => {
    const { data } = useSuspenseQuery({
        queryKey: ['enterprisesByAuthenticatedUser'],
        queryFn: async () => (await getGraphQLClient().request(enterprisesByAuthenticatedUserQuery)).enterprisesByAuthenticatedUser,
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
            <EnterprisesList enterprises={data.enterprises as EnterpriseByAuthenticatedUser[]}/>
            <OfficesList />
        </div>
    )
}
export default EnterprisesMain
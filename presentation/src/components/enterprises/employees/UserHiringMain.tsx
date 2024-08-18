'use client'
import getGraphQLClient from "@/graphql/graphQLClient"
import { useSuspenseQuery } from "@tanstack/react-query"
import { HiringProcess } from "@/gql/graphql"
import hiringProcessByAuthenticatedUserQuery from "@/graphql/queries/hiringInvitationsByAuthenticatedUserQuery"
import UserHiringProccessCard from "./UserHiringProcessCard"

const UserHiringMain = () => {
    const { data: { hiringProcessByAuthenticatedUser: hiringProcesses } } = useSuspenseQuery({
        queryKey: ['hiring-process-by-authenticated-user'],
        queryFn: async () => await getGraphQLClient().request(hiringProcessByAuthenticatedUserQuery)
    })
    return (
        <div>
            <div className="flex flex-col px-5 mt-[2.275rem]">
                {
                    hiringProcesses.map((h) => 
                    <UserHiringProccessCard 
                    key={h?.id!} 
                    hiringProccess={h as HiringProcess} />)
                }
            </div>
        </div>
    )
}

export default UserHiringMain
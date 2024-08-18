import UserHiringMain from "@/components/enterprises/employees/UserHiringMain"
import getGraphQLClient from "@/graphql/graphQLClient"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"
import hiringProcessByAuthenticatedUserQuery from "@/graphql/queries/hiringInvitationsByAuthenticatedUserQuery"

const Hiring: React.FC = async () => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['hiring-process-by-authenticated-user'],
        queryFn: async () => await gqlClient.request(hiringProcessByAuthenticatedUserQuery)
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <UserHiringMain />
        </HydrationBoundary>
    )
}
export default Hiring
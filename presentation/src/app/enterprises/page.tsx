import EnterprisesMain from "@/components/enterprises/EnterprisesMain"
import getGraphQLClient from "@/graphql/graphQLClient"
import enterprisesByAuthenticatedUserQuery from "@/graphql/queries/enterprisesByAuthenticatedUserQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

async function Enterprises() {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['enterprisesByAuthenticatedUser'],
        queryFn: async () => {
            const r = await gqlClient.request(enterprisesByAuthenticatedUserQuery)
            return r.enterprisesByAuthenticatedUser.enterprises
        },
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <EnterprisesMain />
        </HydrationBoundary >
    )
}
export default Enterprises;
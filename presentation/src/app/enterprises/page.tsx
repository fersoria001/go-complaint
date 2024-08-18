import EnterprisesMain from "@/components/enterprises/EnterprisesMain"
import getGraphQLClient from "@/graphql/graphQLClient"
import enterprisesByAuthenticatedUserQuery from "@/graphql/queries/enterprisesByAuthenticatedUserQuery"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"

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
            return r.enterprisesByAuthenticatedUser
        },
    })
    await queryClient.prefetchQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => {
            try {
                return await gqlClient.request(userDescriptorQuery)
            } catch (e: any) {
                console.log("error: ", e)
                return null
            }
        },
        staleTime: Infinity,
        gcTime: Infinity
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <EnterprisesMain />
        </HydrationBoundary >
    )
}
export default Enterprises;
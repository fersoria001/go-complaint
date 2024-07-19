import EnterprisesMain from "@/components/enterprises/EnterprisesMain"
import getGraphQLClient from "@/graphql/graphQLClient"
import enterprisesByAuthenticatedUserQuery from "@/graphql/queries/enterprisesByAuthenticatedUserQuery"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"

const Enterprises: React.FC = async () => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['enterprisesByAuthenticatedUser'],
        queryFn: async () => (await gqlClient.request(enterprisesByAuthenticatedUserQuery)).enterprisesByAuthenticatedUser,
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <EnterprisesMain />
        </HydrationBoundary >
    )
}
export default Enterprises;
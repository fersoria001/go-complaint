import Settings from "@/components/profile/settings/Settings"
import getGraphQLClient from "@/graphql/graphQLClient"
import countriesQuery from "@/graphql/queries/countriesQuery"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"

const ProfileSettings: React.FC = async () => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()

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
    queryClient.prefetchQuery({
        queryKey: ['countries'],
        queryFn: async () => getGraphQLClient().request(countriesQuery),
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <Settings />
        </HydrationBoundary>
    )
}

export default ProfileSettings
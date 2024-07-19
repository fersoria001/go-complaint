import ProfileDashboard from "@/components/profile/ProfileDashboard";
import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { cookies } from "next/headers";

const Profile: React.FC = async () => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => (await gqlClient.request(userDescriptorQuery)).userDescriptor,
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <ProfileDashboard />
        </HydrationBoundary>
    )
}
export default Profile;
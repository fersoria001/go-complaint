import PageProps from "@/app/pageProps"
import HiringMain from "@/components/enterprises/employees/HiringMain"
import getGraphQLClient from "@/graphql/graphQLClient"
import hiringProcessByEnterpriseNameQuery from "@/graphql/queries/hiringProcessByEnterpriseNameQuery"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"

const Hiring: React.FC<PageProps> = async ({ params }: PageProps) => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => await gqlClient.request(userDescriptorQuery),
        staleTime: Infinity,
        gcTime: Infinity
    })
    await queryClient.prefetchQuery({
        queryKey: ['hiring-process-by-enterprise-id', decodeURIComponent(params?.enterpriseId as string)],
        queryFn: async ({ queryKey }) => await gqlClient.request(hiringProcessByEnterpriseNameQuery, { name: queryKey[1] })
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <HiringMain />
        </HydrationBoundary>
    )
}
export default Hiring;
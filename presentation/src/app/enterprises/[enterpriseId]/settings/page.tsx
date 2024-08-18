import PageProps from "@/app/pageProps";
import Settings from "@/components/enterprises/settings/Settings";
import getGraphQLClient from "@/graphql/graphQLClient";
import countriesQuery from "@/graphql/queries/countriesQuery";
import enterpriseByNameQuery from "@/graphql/queries/enterpriseByNameQuery";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers";

const EnterpriseSettings: React.FC<PageProps> = async ({ params, searchParams }: PageProps) => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    const enterpriseName = decodeURIComponent(params!.enterpriseId! as string)
    await queryClient.prefetchQuery({
        queryKey: ["enterpriseByName", enterpriseName],
        queryFn: async ({ queryKey }) => await gqlClient.request(enterpriseByNameQuery, { name: queryKey[1] })
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

export default EnterpriseSettings;
import PageProps from "@/app/pageProps";
import UsersForHiring from "@/components/enterprises/employees/UsersForHiring";
import getGraphQLClient from "@/graphql/graphQLClient";
import usersForHiringQuery from "@/graphql/queries/usersForHiringQuery";

import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const HireNew: React.FC<PageProps> = async ({ params, searchParams }: PageProps) => {
    const cookie = cookies().get("jwt")
    if (!cookie) {
        redirect("/sign-in")
    }
    if (!params?.enterpriseId) {
        redirect("/enterprises")
    }
    const query = searchParams?.query as string || ""
    const page = searchParams?.page ? parseInt(searchParams.page as string, 10) : 0
    const strCookie = `${cookie.name}=${cookie.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchInfiniteQuery({
        queryKey: ['users-for-hiring', params.enterpriseId, query],
        queryFn: async ({ pageParam, queryKey }) => gqlClient.request(usersForHiringQuery, {
            input: {
                id: queryKey[1],
                query: queryKey[2],
                limit: 10,
                offset: pageParam * 10,
            }
        }),
        initialPageParam: page,
        getNextPageParam: (lastPage: any, pages: any) => {
            if (lastPage.nextCursor < 0) {
                return undefined
            }
            return lastPage.nextCursor
        },
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <UsersForHiring />
        </HydrationBoundary>
    )
}
export default HireNew;
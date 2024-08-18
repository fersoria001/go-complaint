import PageProps from "@/app/pageProps"
import HireUser from "@/components/enterprises/employees/HireUser"
import getGraphQLClient from "@/graphql/graphQLClient"
import userByIdQuery from "@/graphql/queries/userByIdQuery"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

const Hire: React.FC<PageProps> = async ({ searchParams, params }: PageProps) => {
    if (!searchParams?.userId) {
        redirect(`/enterprises/${params?.enterpriseId}/employees/hire-new`)
    }
    const cookie = cookies().get("jwt")
    if (!cookie) {
        redirect("/sign-in")
    }
    const strCookie = `${cookie.name}=${cookie.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['user-by-id', searchParams.userId],
        queryFn: async ({ queryKey }) => gqlClient.request(userByIdQuery, { id: queryKey[1] as string })
    })
    await queryClient.prefetchQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => {
          try {
            return await gqlClient.request(userDescriptorQuery)
          } catch (e: any) {
            console.log("error: ",e)
            return null
          }
        },
        staleTime: Infinity,
        gcTime: Infinity
      })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <HireUser />
        </HydrationBoundary>
    )
}

export default Hire
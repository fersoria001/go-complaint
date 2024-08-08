import PageProps from "@/app/pageProps"
import EmployeesMain from "@/components/enterprises/employees/EmployeesMain"
import getGraphQLClient from "@/graphql/graphQLClient"
import enterpriseByNameQuery from "@/graphql/queries/enterpriseByNameQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

const Employees: React.FC<PageProps> = async ({ params }: PageProps) => {
    const cookie = cookies().get('jwt')
    if (!cookie) {
        redirect('/sign-in')
    }
    const strCookie = `${cookie.name}=${cookie.value}`
    if (!params?.enterpriseId) {
        redirect('/enterprises')
    }
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['enterprise', params.enterpriseId as string],
        queryFn: async ({ queryKey }) => {
            try {
                return await gqlClient.request(enterpriseByNameQuery, { name: decodeURIComponent(queryKey[1]) })
            }
            catch (e: any) {
                console.log(e.response.errors[0])
            }
        }
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)} >
            <EmployeesMain />
        </HydrationBoundary>
    )
}
export default Employees
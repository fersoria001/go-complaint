import PageProps from "@/app/pageProps"
import EmployeeActivityDetails from "@/components/enterprises/employees/EmployeeActivityDetails"
import EmployeesActivityLog from "@/components/enterprises/employees/EmployeesActivityLog"
import getGraphQLClient from "@/graphql/graphQLClient"
import userByIdQuery from "@/graphql/queries/userByIdQuery"
import { QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"

const EmployeesActivity: React.FC<PageProps> = async ({ searchParams }: PageProps) => {
    if (searchParams?.id) {
        const jwtCookie = cookies().get("jwt")
        const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
        const gqlClient = getGraphQLClient()
        gqlClient.setHeader("Cookie", strCookie)
        const queryClient = new QueryClient()
        const id = searchParams.id as string
        await queryClient.prefetchQuery({
            queryKey: ["user-by-id", id],
            queryFn: async ({ queryKey }) => await gqlClient.request(userByIdQuery, { id: queryKey[1] })
        })
        return <EmployeeActivityDetails />
    }
    return (
        <EmployeesActivityLog />
    )
}

export default EmployeesActivity
import PageProps from "@/app/pageProps";
import RegisterEnterpriseForm from "@/components/enterprises/RegisterEnterpriseForm";
import getGraphQLClient from "@/graphql/graphQLClient";
import countriesQuery from "@/graphql/queries/countriesQuery";
import industriesQuery from "@/graphql/queries/industriesQuery";
import { QueryClient, HydrationBoundary, dehydrate } from "@tanstack/react-query";

const RegisterEnterprise: React.FC<PageProps> = () => {
    const queryClient = new QueryClient()
    const gqlClient = getGraphQLClient()
    queryClient.prefetchQuery({
        queryKey: ['countries'],
        queryFn: async () => gqlClient.request(countriesQuery),
    })
    queryClient.prefetchQuery({
        queryKey:['industries'],
        queryFn: async () => gqlClient.request(industriesQuery)
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <RegisterEnterpriseForm />
        </HydrationBoundary>
    )
}
export default RegisterEnterprise;
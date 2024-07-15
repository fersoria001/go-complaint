import SignUpForm from "@/components/sign-up/SignUpForm";
import SignUpSucceed from "@/components/sign-up/SignUpSucceed";
import getGraphQLClient from "@/graphql/graphQLClient";
import countriesQuery from "@/graphql/queries/countriesQuery";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import PageProps from "../pageProps";

const SignUp: React.FC<PageProps> = ({
    searchParams,
}: PageProps) => {
    const queryClient = new QueryClient()
    queryClient.prefetchQuery({
        queryKey: ['countries'],
        queryFn: async () => getGraphQLClient().request(countriesQuery),
    })
    if (searchParams?.success) {
        return <SignUpSucceed />
    }
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <SignUpForm />
        </HydrationBoundary>
    )
}
export default SignUp;

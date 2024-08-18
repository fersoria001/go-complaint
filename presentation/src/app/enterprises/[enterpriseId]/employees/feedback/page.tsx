import { getQueryClient } from "@/app/get-query-client";
import PageProps from "@/app/pageProps";
import FeedbacksMain from "@/components/enterprises/feedback/FeedbacksMain";
import getGraphQLClient from "@/graphql/graphQLClient";
import complaintsForFeedbackByEmployeeIdQuery from "@/graphql/queries/complaintsForFeedbackByEmployeeIdQuery";
import complaintsOfResolvedFeedbackByEmployeeIdQuery from "@/graphql/queries/complaintsOfResolvedFeedbackByEmployeeIdQuery";
import { dehydrate, HydrationBoundary } from "@tanstack/react-query"
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const Feedback: React.FC<PageProps> = async ({ searchParams, params }: PageProps) => {
    if (!params?.enterpriseId) {
        redirect("/enterprises")
    }
    if (!searchParams?.id) {
        redirect(`/enterprises/${params.enterpriseId}/employees`)
    }
    const id = searchParams.id as string
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = getQueryClient()
    await queryClient.prefetchQuery({
        queryKey: ["complaints-for-feedback-by-employee-id", id],
        queryFn: async ({ queryKey }) => getGraphQLClient().request(complaintsForFeedbackByEmployeeIdQuery,
            { id: queryKey[1] })
    })
    await queryClient.prefetchQuery({
        queryKey: ["complaints-of-resolved-feedback-by-employee-id", id],
        queryFn: async ({ queryKey }) => getGraphQLClient().request(complaintsOfResolvedFeedbackByEmployeeIdQuery,
            { id: queryKey[1] })
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <FeedbacksMain />
        </HydrationBoundary>
    )

}

export default Feedback;
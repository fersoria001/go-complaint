import ReviewsMain from "@/components/reviews/ReviewsMain";
import getGraphQLClient from "@/graphql/graphQLClient";
import complaintsRatedByAuthorIdQuery from "@/graphql/queries/complaintsRatedByAuthorIdQuery";
import complaintsRatedByReceiverIdQuery from "@/graphql/queries/complaintsRatedByReceiverIdQuery";
import complaintsSentForReviewByReceiverIdQuery from "@/graphql/queries/complaintsSentForReviewByReceiverIdQuery";
import pendingReviewsByAuthorIdQuery from "@/graphql/queries/pendingReviewsByAuthorIdQuery";
import { QueryClient, HydrationBoundary, dehydrate } from "@tanstack/react-query";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const Review: React.FC = async () => {
    const alias = cookies().get("alias")?.value as string
    if (!alias) {
        redirect("/")
    }
    const jwt = cookies().get("jwt")?.value as string
    if (!jwt) {
        redirect("/sign-in")
    }
    const strCookie = `jwt=${jwt};`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ["pendingReviewsByAuthorIdQuery", alias, ""],
        queryFn: async ({ queryKey }) => await gqlClient.request(
            pendingReviewsByAuthorIdQuery,
            { id: queryKey[1], term: queryKey[2] },
        )
    })
    await queryClient.prefetchQuery({
        queryKey: ["complaintsSentForReviewByReceiverIdQuery", alias, ""],
        queryFn: async ({ queryKey }) => await gqlClient.request(
            complaintsSentForReviewByReceiverIdQuery,
            { id: queryKey[1], term: queryKey[2] },
        )
    })
    await queryClient.prefetchQuery({
        queryKey: ["complaintsRatedByReceiverIdQuery", alias, ""],
        queryFn: async ({ queryKey }) => await gqlClient.request(
            complaintsRatedByReceiverIdQuery,
            { id: queryKey[1], term: queryKey[2] },
        )
    })
    await queryClient.prefetchQuery({
        queryKey: ["complaintsRatedByAuthorIdQuery", alias, ""],
        queryFn: async ({ queryKey }) => await gqlClient.request(
            complaintsRatedByAuthorIdQuery,
            { id: queryKey[1], term: queryKey[2] }
        )
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)}>
            <ReviewsMain />
        </HydrationBoundary>
    )
}
export default Review;
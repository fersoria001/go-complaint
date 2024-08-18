import { getQueryClient } from "@/app/get-query-client"
import PageProps from "@/app/pageProps"
import FeedbackComplaint from "@/components/enterprises/feedback/FeedbackComplaint"
import getGraphQLClient from "@/graphql/graphQLClient"
import createFeedbackMutation from "@/graphql/mutations/createFeedbackMutation"
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { HydrationBoundary, dehydrate } from "@tanstack/react-query"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

const FeedbackEditor: React.FC<PageProps> = async ({ searchParams, params }: PageProps) => {
    if (!params?.enterpriseId) {
        redirect("/enterprises")
    }
    if(!params?.complaintId) {
        redirect(`/enterprises/${params.enterpriseId}/employees`)
    }
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = getQueryClient()
    const complaintId = params.complaintId as string
    const enterpriseName = decodeURIComponent(params.enterpriseId)
    queryClient.prefetchQuery({
        queryKey: ["create-feedback-mutation", enterpriseName, complaintId],
        queryFn: async ({ queryKey }) => gqlClient.request(createFeedbackMutation, {
            input: { complaintId: queryKey[2], enterpriseId: queryKey[1] }
        })
    })
    queryClient.prefetchQuery({
        queryKey: ["complaint-by-id", complaintId],
        queryFn: async ({ queryKey }) => gqlClient.request(complaintByIdQuery, { id: queryKey[1] })
    })
    await queryClient.prefetchQuery({
        queryKey: ['user-descriptor'],
        queryFn: async () => {
            try {
                return await gqlClient.request(userDescriptorQuery)
            } catch (e: any) {
                console.log("error: ", e)
                return null
            }
        },
        staleTime: Infinity,
        gcTime: Infinity
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)} >
            <FeedbackComplaint />
        </HydrationBoundary>
    )
}
export default FeedbackEditor;
import PageProps from "@/app/pageProps";
import ComplaintChat from "@/components/complaints/ComplaintChat";
import getGraphQLClient from "@/graphql/graphQLClient";
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { getCookie } from "@/lib/actions/cookies";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const Complaint: React.FC<PageProps> = async ({ params }: PageProps) => {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    await queryClient.prefetchQuery({
        queryKey: ['userDescriptor'],
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
    await queryClient.prefetchQuery({
        queryKey: ["complaintById", params?.complaintId!],
        queryFn: async ({ queryKey }) => {
            try {
                return await gqlClient.request(complaintByIdQuery, { id: queryKey[1] })
            } catch (e: any) {
                console.log("error: ", e)
                redirect("/complaints")
            }
        }
    })
    await queryClient.prefetchQuery({
        queryKey: ["serverSideJwtCookie"],
        queryFn: async () => {
            try {
                return await getCookie("jwt")
            } catch (e: any) {
                console.log("error: ", e)
                //redirect("/complaints")
            }
        },
        staleTime: Infinity,
        gcTime: Infinity
    })
    await queryClient.prefetchQuery({
        queryKey: ["serverSideAliasCookie"],
        queryFn: async () => {
            try {
                return await getCookie("alias")
            } catch (e: any) {
                console.log("error: ", e)
                //redirect("/complaints")
            }
        },
        staleTime: Infinity,
        gcTime: Infinity
    })
    return (
        <HydrationBoundary state={dehydrate(queryClient)} >
            <ComplaintChat />
        </HydrationBoundary>
    )
}

export default Complaint;
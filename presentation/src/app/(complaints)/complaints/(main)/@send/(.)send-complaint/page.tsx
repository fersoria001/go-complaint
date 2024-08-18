import PageProps from "@/app/pageProps"
import Complain from "@/components/complaints/Complain"
import DescribeComplaint from "@/components/complaints/DescribeComplaint"
import FindReceiver from "@/components/complaints/FindReceiver"
import Modal from "@/components/modal/Modal"
import getGraphQLClient from "@/graphql/graphQLClient"
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery"
import recipientsByNameLikeQuery from "@/graphql/queries/recipientsByNameLikeQuery"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

const SendComplaintPage: React.FC<PageProps> = async ({ searchParams }: PageProps) => {
    const modalClassName = "absolute flex flex-col p-2 bg-white border-t shadow-md rounded-md w-full h-screen top-20 mt-[1px] md:h-[390px] md:w-[320px] md:right-0 md:mr-16 md:bottom-0 md:inset-y-40"
    if (!searchParams?.step) {
        redirect('/complaints')
    }
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    const queryClient = new QueryClient()
    if (searchParams.step === "1") {
        await queryClient.prefetchQuery({
            queryKey: ['recipientsByNameLike', ""],
            queryFn: async ({ queryKey }) => {
                try {
                    return await gqlClient.request(recipientsByNameLikeQuery, { term: queryKey[1] })
                } catch (e: any) {
                    console.log("error: ", e)
                    return null
                }
            },
        })
        return (
            <HydrationBoundary state={dehydrate(queryClient)}>
                <Modal className={modalClassName}>
                    <FindReceiver />
                </Modal>
            </HydrationBoundary>
        )
    }
    if (searchParams.step === "2") {
        if (!searchParams?.id) {
            redirect("/complaints/send-complaint?step=1")
        }
        await queryClient.prefetchQuery({
            queryKey: ['complaintById', searchParams.id as string],
            queryFn: async ({ queryKey }) => {
                try {
                    return await gqlClient.request(complaintByIdQuery, { id: queryKey[1] })
                } catch (e: any) {
                    console.log("error: ", e)
                    return null
                }
            },
        })
        return (
            <HydrationBoundary state={dehydrate(queryClient)}>
                <Modal className={modalClassName}>
                    <DescribeComplaint />
                </Modal>
            </HydrationBoundary>
        )
    }
    if (searchParams.step === "3") {
        if (!searchParams?.id) {
            redirect("/complaints/send-complaint?step=1")
        }
        await queryClient.prefetchQuery({
            queryKey: ['complaintById', searchParams.id as string],
            queryFn: async ({ queryKey }) => {
                try {
                    return await gqlClient.request(complaintByIdQuery, { id: queryKey[1] })
                } catch (e: any) {
                    console.log("error: ", e)
                    return null
                }
            },
        })
        return (
            <HydrationBoundary state={dehydrate(queryClient)}>
                <Modal className={modalClassName}>
                    <Complain />
                </Modal>
            </HydrationBoundary>
        )
    }
}
export default SendComplaintPage
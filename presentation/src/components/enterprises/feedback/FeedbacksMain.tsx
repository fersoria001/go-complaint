'use client'
import { Complaint } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import complaintsForFeedbackByEmployeeIdQuery from "@/graphql/queries/complaintsForFeedbackByEmployeeIdQuery"
import { useSuspenseQueries, useSuspenseQuery } from "@tanstack/react-query"
import { useSearchParams } from "next/navigation"
import FeedbackList from "./FeedbackList"
import complaintsOfResolvedFeedbackByEmployeeIdQuery from "@/graphql/queries/complaintsOfResolvedFeedbackByEmployeeIdQuery"


const FeedbacksMain = () => {
    const id = useSearchParams().get("id") as string
    const [
        { data: { complaintsForFeedbackByEmployeeId: feedbacksPending } },
        { data: { complaintsOfResolvedFeedbackByEmployeeId: feedbacksResolved } }
    ] = useSuspenseQueries({
        queries: [
            {
                queryKey: ["complaints-for-feedback-by-employee-id", id],
                queryFn: async () => getGraphQLClient().request(complaintsForFeedbackByEmployeeIdQuery, { id: id })
            },
            {
                queryKey: ["complaints-of-resolved-feedback-by-employee-id", id],
                queryFn: async () => getGraphQLClient().request(complaintsOfResolvedFeedbackByEmployeeIdQuery, { id: id })
            }
        ]
    })

    return (
        <div>
            <FeedbackList
                label={"Feedback solved complaints"}
                fallbackLabel={"This employee has not provided any assistance yet."}
                complaints={feedbacksPending as Complaint[]} />
            <FeedbackList
                label={"Read feedbacks"}
                fallbackLabel={"This employee has not received any feedback yet."}
                complaints={feedbacksResolved as Complaint[]} />
        </div>
    )
}

export default FeedbacksMain
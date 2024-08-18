'use client'
import { TimeData } from "@/components/charts/ChartBuilder";
import TimeDataChart from "@/components/charts/TimeDataChart";
import { EnterpriseActivity, EnterpriseActivityType } from "@/gql/graphql";
import getGraphQLClient from "@/graphql/graphQLClient";
import getGraphQLSubscriptionClient from "@/graphql/graphQLSubscriptionClient";
import userByIdQuery from "@/graphql/queries/userByIdQuery";
import employeeActivitySubscription from "@/graphql/subscriptions/employeeActivitySubscription";
import { useSuspenseQuery } from "@tanstack/react-query";
import { useParams, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

const EmployeeActivityDetails: React.FC = () => {
    const searchParams = useSearchParams()
    const employeeUserId = searchParams.get('id')!
    const { enterpriseId } = useParams()
    const { data: { userById: employee } } = useSuspenseQuery({
        queryKey: ["user-by-id", employeeUserId],
        queryFn: async ({ queryKey }) => await getGraphQLClient().request(userByIdQuery, { id: queryKey[1] })
    })
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    const [activity, setActivity] = useState<Map<EnterpriseActivityType, EnterpriseActivity[]>>(
        new Map<EnterpriseActivityType, EnterpriseActivity[]>([
            [EnterpriseActivityType.EmployeesFired, []],
            [EnterpriseActivityType.EmployeesHired, []],
            [EnterpriseActivityType.FeedbacksReceived, []],
            [EnterpriseActivityType.FeedbacksStarted, []],
            [EnterpriseActivityType.JobProposalsSent, []],
            [EnterpriseActivityType.ComplaintResolved, []],
            [EnterpriseActivityType.ComplaintReviewed, []],
            [EnterpriseActivityType.ComplaintSent, []]
        ])
    )
    useEffect(() => {

        async function subscribeToEnterpriseActivity() {
            const subscription = getGraphQLSubscriptionClient().iterate({
                query: employeeActivitySubscription(employeeUserId),
            });
            for await (const event of subscription) {
                const c = event.data?.employeeActivity as EnterpriseActivity
                if (c.activityType) {
                    setActivity(prev => {
                        const arr = prev.get(c.activityType)
                        if (arr) {
                            prev.set(c.activityType, [...arr, c])
                            return new Map(prev)
                        }
                        return prev
                    })
                }
            }
        }
        subscribeToEnterpriseActivity()
        return () => {
            setActivity(new Map<EnterpriseActivityType, EnterpriseActivity[]>([
                [EnterpriseActivityType.EmployeesFired, []],
                [EnterpriseActivityType.EmployeesHired, []],
                [EnterpriseActivityType.FeedbacksReceived, []],
                [EnterpriseActivityType.FeedbacksStarted, []],
                [EnterpriseActivityType.JobProposalsSent, []],
                [EnterpriseActivityType.ComplaintResolved, []],
                [EnterpriseActivityType.ComplaintReviewed, []],
                [EnterpriseActivityType.ComplaintSent, []]
            ]))
        }
    }, [employeeUserId])
    return (
        <div className="my-5">
            <h3 className="text-gray-700 font-bold text-center">
                {employee.person.firstName} {" "} {employee.person.lastName}{" "}activity in {enterpriseName}
            </h3>
            <div className="lg:p-5 flex flex-col lg:items-center">
                <div className="flex flex-col xl:flex-row mb-4 gap-2">
                    <div className="w-full xl:w-[420px]">
                        <TimeDataChart
                            chartLabel={`Complaints sent as ${enterpriseName}`}
                            data={activity.get(EnterpriseActivityType.ComplaintSent) as TimeData[]} yLabel={"complaints"} />
                    </div>
                    <div className="w-full xl:w-[420px]">
                        <TimeDataChart chartLabel={`Complaints resolved for ${enterpriseName}`}
                            data={activity.get(EnterpriseActivityType.ComplaintResolved) as TimeData[]} yLabel={"complaints"} />
                    </div>
                    <div className="w-full xl:w-[420px]">
                        <TimeDataChart
                            chartLabel={`Complaints reviewed for ${enterpriseName}`}
                            data={activity.get(EnterpriseActivityType.ComplaintReviewed) as TimeData[]} yLabel={"complaints"} />
                    </div>
                </div>
            </div>

            <div className="lg:p-5 flex flex-col lg:items-center">
                <div className="flex flex-col xl:flex-row mb-4 gap-2">
                    <div className="w-full lg:w-[480px] xl:w-[540px]">
                        <TimeDataChart
                            chartLabel={`Feedback to employees`}
                            data={activity.get(EnterpriseActivityType.FeedbacksStarted) as TimeData[]} yLabel={"feedbacks"} />
                    </div>
                    <div className="w-full lg:w-[480px] xl:w-[540px]">
                        <TimeDataChart chartLabel={`Feedbacks received`}
                            data={activity.get(EnterpriseActivityType.FeedbacksReceived) as TimeData[]} yLabel={"feedbacks"} />
                    </div>
                </div>
            </div>
            <div className="w-full lg:px-8">
                <TimeDataChart
                    chartLabel={`Job proposals sent`}
                    data={activity.get(EnterpriseActivityType.JobProposalsSent) as TimeData[]} yLabel={"hiring proposal"} />
            </div>
            <div className="lg:p-5 flex flex-col lg:items-center">
                <div className="flex flex-col xl:flex-row mb-4 gap-2">
                    <div className="w-full lg:w-[480px] xl:w-[540px]">
                        <TimeDataChart
                            chartLabel={`Employees hired`}
                            data={activity.get(EnterpriseActivityType.EmployeesHired) as TimeData[]} yLabel={"employees"} />
                    </div>
                    <div className="w-full lg:w-[480px] xl:w-[540px]">
                        <TimeDataChart chartLabel={`Employees fired`}
                            data={activity.get(EnterpriseActivityType.FeedbacksReceived) as TimeData[]} yLabel={"employees"} />
                    </div>
                </div>
            </div>
        </div>
    )
}
export default EmployeeActivityDetails
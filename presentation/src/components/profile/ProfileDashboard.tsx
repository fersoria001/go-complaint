'use client'

import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { useSuspenseQuery } from "@tanstack/react-query";
import ComplaintDataCharts from "./ComplaintDataCharts";


const ProfileDashboard = () => {
    const { data: { userDescriptor: user } } = useSuspenseQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => await getGraphQLClient().request(userDescriptorQuery),
        staleTime: Infinity,
        gcTime: Infinity
    })
    return (
        <section className="flex flex-col my-5">
            <h3 className="text-gray-700 font-bold text-center">My Go Complaint Activity</h3>
            <ComplaintDataCharts currentUser={user} />
        </section>
    )
}
export default ProfileDashboard;
'use client'

import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { useSuspenseQuery } from "@tanstack/react-query";
import ComplaintsInfo from "./ComplaintsInfo";

const ProfileDashboard = () => {
    const { data: user } = useSuspenseQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => (await getGraphQLClient().request(userDescriptorQuery)).userDescriptor
    })
    return (
        <section>
            <ComplaintsInfo id={user.userName} />
        </section>
    )
}
export default ProfileDashboard;
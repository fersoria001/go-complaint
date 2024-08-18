'use client'
import getGraphQLClient from "@/graphql/graphQLClient";
import enterpriseByNameQuery from "@/graphql/queries/enterpriseByNameQuery";
import { useSuspenseQueries, useSuspenseQuery } from "@tanstack/react-query";
import Link from "next/link";
import { useParams } from "next/navigation";
import EmployeeCard from "./EmployeeCard";
import { Employee } from "@/gql/graphql";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";

const EmployeesMain: React.FC = () => {
    const { enterpriseId } = useParams()
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    const [{ data: { enterpriseByName } }, { data: { userDescriptor } }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['enterprise', enterpriseName],
                queryFn: async () => await getGraphQLClient().request(enterpriseByNameQuery, { name: enterpriseName })
            },
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await getGraphQLClient().request(userDescriptorQuery),
                staleTime: Infinity,
                gcTime: Infinity
            },
        ]
    })

    return (
        <div>
            <div className="flex flex-col px-5 mt-[2.275rem]">
                <Link
                    href={`/enterprises/${enterpriseName}/employees/hiring`}
                    className="text-gray-700 mb-4 underline text-md">
                    Hiring
                </Link>
                <Link
                    href={`/enterprises/${enterpriseName}/employees/activity`}
                    className="text-gray-700 mb-4 underline text-md">
                    Employees activity
                </Link>
            </div>
            <div className="xl:px-5">
                {
                    enterpriseByName.employees.map((e) => {
                        return <EmployeeCard
                            key={e?.id}
                            employee={e as Employee}
                            currentUser={userDescriptor}
                            enterpriseName={enterpriseByName.name} />
                    })
                }
            </div>
        </div>
    )
}
export default EmployeesMain;
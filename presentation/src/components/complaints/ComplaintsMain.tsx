'use client'
import getGraphQLClient from "@/graphql/graphQLClient"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { useSuspenseQuery } from "@tanstack/react-query"
import ComplaintItem from "./ComplaintItem"
import Link from "next/link"
import SelectIcon from "../icons/SelectIcon"
import FilterBy from "../search/FilterBy"

const options = [
    {
        id: "0",
        name: "all",
        value: "all"
    },
    {
        id: "1",
        name: "received",
        value: "received"
    },
    {
        id: "2",
        name: "sent",
        value: "sent"
    },
]


const ComplaintsMain: React.FC = () => {
    const { data } = useSuspenseQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => getGraphQLClient().request(userDescriptorQuery),
        staleTime: Infinity,
        gcTime: Infinity
    })
    return (
        <>
            <div className="flex flex-col mt-4">
                <div className="flex gap-5 overflow-x-auto w-full self-center my-2 whitespace-nowrap px-2.5">
                    <button
                        type="button"
                        className="bg-blue-500 hover:bg-blue-600 rounded-lg px-2.5 py-0.5 text-white font-bold">
                        UserFirstName(you)
                    </button>
                    {
                        data.userDescriptor.authorities?.map((enterprise) => {
                            return (
                                <button
                                    key={enterprise?.enterpriseId}
                                    type="button"
                                    className="bg-blue-500 hover:bg-blue-600 rounded-lg px-2.5 py-0.5 text-white font-bold">
                                    {enterprise?.enterpriseId}
                                </button>
                            )
                        })
                    }
                </div>
                <FilterBy options={options} />
                <div className="w-full mt-4">
                    <Link href={`/complaints/1`}>
                        <ComplaintItem />
                    </Link>
                </div>
            </div>
        </>
    )
}
export default ComplaintsMain
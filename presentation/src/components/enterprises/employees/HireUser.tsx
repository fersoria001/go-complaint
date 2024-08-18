'use client'
import InlineAlert from "@/components/error/InlineAlert";
import getGraphQLClient from "@/graphql/graphQLClient";
import userByIdQuery from "@/graphql/queries/userByIdQuery";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { inviteToProject } from "@/lib/actions/graphqlActions";
import { useMutation, useSuspenseQueries, useSuspenseQuery } from "@tanstack/react-query";
import clsx from "clsx";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";
import { z } from "zod";

const HireUser: React.FC = () => {
    const userId = useSearchParams().get("userId") || ""
    const { enterpriseId } = useParams()
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    const [selectedPosition, setSelectedPosition] = useState<string>("")
    const [{ data: { userDescriptor } }, { data: { userById: proposeTo } }] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['userDescriptor'],
                queryFn: async () => await getGraphQLClient().request(userDescriptorQuery)
            },
            {
                queryKey: ['user-by-id', userId],
                queryFn: async () => getGraphQLClient().request(userByIdQuery, { id: userId })
            },
        ]
    })
    const router = useRouter()
    const mutation = useMutation({
        mutationFn: () => inviteToProject({
            enterpriseId: enterpriseName,
            role: selectedPosition,
            proposeTo: userId,
            proposedBy: userDescriptor.id,
        }),
        onMutate: () => {
            router.prefetch(`/enterprises/${enterpriseId}/employees/hiring`)
        },
        onSuccess: (data, variables, context) => {
            router.push(`/enterprises/${enterpriseId}/employees/hiring`)
        },
    })
    return (
        <div className="flex flex-col justify-around items-center p-5 bg-white  border-gray-200 rounded-lg shadow-md">
            <h1 className="mb-3 text-lg text-gray-700 md:text-xl xl:text-2xl font-bold">
                You are about to invite {proposeTo.person.firstName} {proposeTo.person.lastName} {" "}
                to {enterpriseName}
            </h1>
            <div className="mr-auto px-4">
                <p className="text-gray-700 text-md">
                    There are a few things you should know before inviting a user to your enterprise
                </p>
                <p className="text-gray-700 text-md">
                    The current available positions are, choose one from the list:
                </p>
                <ul className="max-w-md space-y-1 p-4 text-gray-700 text-md list-disc list-inside">
                    <li
                        onClick={() => { setSelectedPosition("ASSISTANT"); }}
                        className={
                            clsx("rounded-xl", {
                                'bg-gray-100 p-2 cursor-default': selectedPosition === "ASSISTANT",
                                'hover:bg-gray-100 cursor-pointer': selectedPosition !== "ASSISTANT"
                            })
                        }
                    >
                        <b>Assistant:</b> An assistant is a person who helps in the management of the enterprise.
                        He/she is responsible for the day-to-day running of the enterprise
                        and has permissions to view and answer the enterprise&apos;s complaints.
                    </li>
                    <li
                        onClick={() => { setSelectedPosition("MANAGER"); }}
                        className={
                            clsx("rounded-xl", {
                                'bg-gray-100 p-2 cursor-default': selectedPosition === "MANAGER",
                                'hover:bg-gray-100 cursor-pointer': selectedPosition !== "MANAGER"
                            })
                        }
                    >
                        <b>Manager:</b> A manager is a person who is responsible for the management of the enterprise
                        employees. While he/she can view and answer the enterprise&apos;s complaints, the main responsibility
                        of them is to invite new users, propose promotions to other positions, and review solved complaints to provide
                        feedback to the employees replies to the complaints.
                    </li>
                </ul>
            </div>
            <div className="flex flex-col relative md:static">
                <p className="p-4 text-gray-700 text-md">
                    Note that: every employee can be rated by the users if he/she has resolved a complaint and receive feedback
                    and private messages from the enterprise owner. While managers can invite new users and propose position changes,
                    it is the responsibility of the owner to accept or decline the changes.
                    And only the owner can modify the enterprise&apos;s information.
                </p>
                {mutation.isError && <InlineAlert errors={mutation.error.message.split(",")} />}
                <span onClick={() => { }} className="self-center">
                    <button
                        onClick={() => { mutation.mutate() }}
                        type="button"
                        className="bg-blue-500 text-white px-7 py-3 rounded-md font-bold hover:bg-blue-600 text-md"
                    >Invite to project
                    </button>
                </span>
            </div>
        </div >
    )
}
export default HireUser;
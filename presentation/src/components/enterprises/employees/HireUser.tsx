'use client'
import getGraphQLClient from "@/graphql/graphQLClient";
import userByIdQuery from "@/graphql/queries/userByIdQuery";
import { useSuspenseQuery } from "@tanstack/react-query";
import clsx from "clsx";
import { useParams, useSearchParams } from "next/navigation";
import { useState } from "react";

const HireUser: React.FC = () => {
    const userId = useSearchParams().get("userId") || ""
    const { enterpriseId: enterpriseName } = useParams()
    const [selectedPosition, setSelectedPosition] = useState<string>("")
    const { data } = useSuspenseQuery({
        queryKey: ['user-by-id', userId],
        queryFn: async ({ queryKey }) => getGraphQLClient().request(userByIdQuery, { id: queryKey[1] as string })
    })
    return (
        <div className="flex flex-col justify-around items-center p-5 bg-white  border-gray-200 rounded-lg shadow-md">
            <h1 className="mb-3 text-lg text-gray-700 md:text-xl xl:text-2xl font-bold">
                You are about to invite {data.userById.person.firstName} {data.userById.person.firstName} {" "}
                to {enterpriseName}
            </h1>
            <div className="mr-auto px-4">
                <p className="text-gray-700 text-md lg:text-lg xl:text-xl">
                    There are a few things you should know before inviting a user to your enterprise
                </p>
                <p className="text-gray-700 text-md lg:text-lg xl:text-xl">
                    The current available positions are, choose one from the list:
                </p>
                <ul className="max-w-md space-y-1 p-4 text-gray-700 text-md lg:text-lg xl:text-xl list-disc list-inside">
                    <li
                        onClick={() => { setSelectedPosition("Assistant"); }}
                        className={
                            clsx("rounded-xl", {
                                'bg-gray-100 p-2 cursor-default': selectedPosition === "Assistant",
                                'hover:bg-gray-100 cursor-pointer': selectedPosition !== "Assistant"
                            })
                        }
                    >
                        Assistant: An assistant is a person who helps in the management of the enterprise.
                        He/she is responsible for the day-to-day running of the enterprise
                        and has permissions to view and answer the enterprise&apos;s complaints.
                    </li>
                    <li
                        onClick={() => { setSelectedPosition("Manager"); }}
                        className={
                            clsx("rounded-xl", {
                                'bg-gray-100 p-2 cursor-default': selectedPosition === "Manager",
                                'hover:bg-gray-100 cursor-pointer': selectedPosition !== "Manager"
                            })
                        }
                    >
                        Manager: A manager is a person who is responsible for the management of the enterprise
                        employees. While he/she can view and answer the enterprise&apos;s complaints, the main responsibility
                        of them is to invite new users, propose promotions to other positions, and review solved complaints to provide
                        feedback to the employees replies to the complaints.
                    </li>
                </ul>
            </div>
            <div className="flex flex-col relative md:static">
                <p className="p-4 text-gray-700 text-md lg:text-lg xl:text-xl">
                    Note that: every employee can be rated by the users if he has resolved a complaint and receive feedback
                    and private messages from the enterprise owner. While managers can invite new users and propose position changes,
                    it is the responsibility of the owner to accept or decline the changes.
                    And only the owner can modify the enterprise&apos;s information.
                </p>
                {/* {errors.position && <p className="self-center text-red-500 italic text-xs">{errors.position}</p>} */}
                <span onClick={() => { }} className="self-center">
                    <button
                        type="button"
                        className="bg-blue-500 text-white px-7 py-3 rounded-md font-bold hover:bg-blue-600 text-md lg:text-lg xl:text-xl"
                    >Invite to project</button>
                </span>

                {/* {showModal && (
                    <Confirm
                        id="confirm-complaint-modal"
                        show={showModal}
                        userFullName={`${user.firstName} ${user.lastName}`}
                        enterpriseName={enterprise.name}
                        position={selectedPosition}
                        proposedTo={user.email}
                        callbackFn={modalCallback}
                        closeFn={() => { setShowModal(false); }}
                    />
                )} */}
            </div>
        </div >
    )
}
export default HireUser;
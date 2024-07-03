/* eslint-disable @typescript-eslint/no-explicit-any */

import {  InviteToProjectSchema } from "../lib/types";
import { useEffect, useState } from "react";

import PrimaryButton from "../components/buttons/PrimaryButton";
import Confirm from "../components/hiring/Confirm";

import { Route } from "../routes/$enterpriseID/hire";

import { InviteToProjectMutation } from "../lib/mutations";
import { useRouter } from "@tanstack/react-router";
import { useForm } from "../lib/hooks/useForm";

function Hire() {
    const router = useRouter();
    const { enterprise, user } = Route.useLoaderData();
    const [selectedPosition, setSelectedPosition] = useState<string>("");
    const [showModal, setShowModal] = useState<boolean>(false);
    const [formData, setFormData] = useState<FormData | null>(null);
    const { success, errors } = useForm(formData, InviteToProjectSchema, InviteToProjectMutation);
    const handleInvite = () => {
        setShowModal(true);
    };

    const modalCallback = (formData: FormData) => {
        setFormData(formData);
    };

    useEffect(() => {
        if (success) {
            router.navigate({
                to: `/${enterprise.name}/success`,
                search: { content: { message: "User invited successfully", link: `/${enterprise.name}`, to: enterprise.name } }
            });
        }
    }, [enterprise.name, router, success]);




    return (
        <div className="flex flex-col justify-around items-center p-5 bg-white border border-gray-200 rounded-lg shadow">
            <h1 className="mb-3 text-lg text-gray-500 md:text-xl">
                You are about to invite {user.firstName} {user.lastName} {" "}
                to your current enterprise {enterprise.name}
            </h1>
            <div className="mr-auto px-4">
                <p className="text-gray-500 dark:text-gray-400">
                    There are a few things you should know before inviting a user to your enterprise
                </p>
                <p className="text-gray-500 dark:text-gray-400">
                    The current available positions are, choose one from the list:
                </p>
                <ul className="max-w-md space-y-1 p-4 text-gray-500 list-disc list-inside">
                    <li
                        onClick={() => { setSelectedPosition("Assistant"); }}
                        className={
                            selectedPosition === "Assistant" ?
                                "rounded-xl bg-cyan-500 text-gray-800 shadow p-2" :
                                "rounded-xl hover:bg-gray-100 hover:cursor-pointer"
                        }
                    >
                        Assistant: An assistant is a person who helps in the management of the enterprise.
                        He/she is responsible for the day-to-day running of the enterprise
                        and has permissions to view and answer the enterprise's complaints.
                    </li>
                    <li
                        onClick={() => { setSelectedPosition("Manager"); }}
                        className={
                            selectedPosition === "Manager" ?
                                "rounded-xl bg-cyan-500 text-gray-800 shadow p-2" :
                                "rounded-xl hover:bg-gray-100 hover:cursor-pointer"
                        }
                    >
                        Manager: A manager is a person who is responsible for the management of the enterprise
                        employees. While he/she can view and answer the enterprise's complaints, the main responsibility
                        of them is to invite new users, propose promotions to other positions, and review solved complaints to provide
                        feedback to the employees' replies to the complaints.
                    </li>
                </ul>
            </div>
            <div className="flex flex-col relative md:static">
                <p className="p-4 text-gray-500 dark:text-gray-400">
                    Note that: every employee can be rated by the users if he has resolved a complaint and receive feedback
                    and private messages from the enterprise owner. While managers can invite new users and propose position changes,
                    it is the responsibility of the owner to accept or decline the changes.
                    And only the owner can modify the enterprise's information.
                </p>
                {errors.position && <p className="self-center text-red-500 italic text-xs">{errors.position}</p>}
                <span onClick={handleInvite} className="self-center">
                    <PrimaryButton text="Invite" />
                </span>

                {showModal && (
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
                )}
            </div>
        </div>
    );
}

export default Hire;

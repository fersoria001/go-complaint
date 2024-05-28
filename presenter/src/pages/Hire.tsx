import { useLoaderData, useNavigate } from "react-router-dom";
import { Enterprise, InviteToProject, User } from "../lib/types";
import { useEffect, useRef, useState } from "react";
import useOutsideDenier from "../lib/hooks/useOutsideDenier";
import PrimaryButton from "../components/buttons/PrimaryButton";
import Confirm from "../components/hiring/Confirm";
import { InviteToProjectMutation, Mutation } from "../lib/mutations";

function Hire() {
    const { enterprise, user } = useLoaderData() as { enterprise: Enterprise | null, user: User | null };
    const [selectedPosition, setSelectedPosition] = useState<string>("");
    const [showModal, setShowModal] = useState<boolean>(false);
    const [accepted, setAccepted] = useState<boolean>(false);
    const [errors, setErrors] = useState<{ [key: string]: string }>({});
    const [validatedObject, setValidatedObject] = useState<InviteToProject | null>(null);
    const listRef = useRef<HTMLUListElement>(null);
    const navigate = useNavigate();

    const handleInvite = () => {
        setErrors({});
        if (!selectedPosition) {
            setErrors({ position: "You must select a position" });
            return;
        }
        if (selectedPosition !== "Assistant" && selectedPosition !== "Manager") {
            setErrors({ position: "Invalid position" });
            return;
        }
        const invite: InviteToProject = {
            enterpriseName: enterprise!.name,
            position: selectedPosition,
            userEmail: user!.email,
            userFullName: `${user!.firstName} ${user!.lastName}`,
        };
        setValidatedObject(invite);
        setShowModal(true);
    };

    useEffect(() => {
        if (!accepted) return;
        if (!validatedObject) {
            setErrors({ position: "You must select a position" });
            return;
        }
        Mutation<InviteToProject>(InviteToProjectMutation, validatedObject)
            .then(() => {
                return navigate("/success/invite%20to%20enterprise");
            })
            .catch((error) => {
                console.error(error);
            });
    }, [accepted, validatedObject, navigate]);

    useOutsideDenier(listRef, () => setSelectedPosition(""));

    if (!enterprise || !user) {
        return <div>loading...</div>;
    }

    return (
        <div className="flex flex-col justify-around items-center bg-white border border-gray-200 rounded-lg shadow">
            <h1 className="mb-3 text-lg text-gray-500 md:text-xl">
                You are about to invite {user.firstName} {user.lastName} {" "}
                to your current enterprise {enterprise.name}
            </h1>
            <div className="mr-auto p-4">
                <p className="text-gray-500 dark:text-gray-400">
                    There are a few things you should know before inviting a user to your enterprise
                </p>
                <p className="text-gray-500 dark:text-gray-400">
                    The current available positions are, choose one from the list:
                </p>
                <ul ref={listRef} className="max-w-md space-y-1 p-4 text-gray-500 list-disc list-inside">
                    <li
                        onClick={() => { console.log("Assistant clicked"); setSelectedPosition("Assistant"); }}
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
                        onClick={() => { console.log("Manager clicked"); setSelectedPosition("Manager"); }}
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
                <span ref={listRef} onClick={handleInvite} className="self-center mt-4">
                    <PrimaryButton text="Invite" />
                </span>

                {showModal && validatedObject && (
                    <Confirm
                        id="confirm-complaint-modal"
                        show={showModal}
                        validatedObject={validatedObject}
                        callbackFn={() => { setAccepted(true); }}
                        closeFn={() => { setShowModal(false); }}
                    />
                )}
            </div>
        </div>
    );
}

export default Hire;

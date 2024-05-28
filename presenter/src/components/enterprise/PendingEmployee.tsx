import { useEffect, useState } from "react";
import { Employee, EndHiringProcess } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import MaleFaceIcon from "../icons/MaleFaceIcon";
import WorkIcon from "../icons/WorkIcon";
import { EndHiringProcessMutation, Mutation } from "../../lib/mutations";
import LoadingSpinner from "../icons/LoadingSpinner";
import DoneIcon from "../icons/DoneIcon";
import ErrorIcon from "../icons/ErrorIcon";
import ThinButton from "../buttons/ThinButton";

interface Props {
    eventID: string;
    employee: Employee;
}
function PendingEmployee({ eventID, employee }: Props) {
    const [accepted, setAccepted] = useState<boolean | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const [success, setSuccess] = useState<boolean | null>(null);
    useEffect(() => {
        if (accepted == null) {
            return
        }
        const enterpriseID = employee.ID.split("-")[0];

        if (accepted) {
            setLoading(true)
            Mutation<EndHiringProcess>(EndHiringProcessMutation,
                { pendingEventID: eventID, enterpriseID, employeeID: employee.ID, accepted: true }
            ).then(() => {
                setSuccess(true)
                setLoading(false)
            }).catch(() => {
                setSuccess(false)
                setLoading(false)
            })
        } else {
            setLoading(true)
            Mutation<EndHiringProcess>(EndHiringProcessMutation,
                { pendingEventID: eventID, enterpriseID, employeeID: employee.ID, accepted: false }
            ).then(() => {
                setSuccess(true)
                setLoading(false)
            }).catch(() => {
                setSuccess(false)
                setLoading(false)
            })
        }
    }, [accepted, employee, eventID])
    return (
        <div className="flex flex-col md:flex-row justify-around items-center
        bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
            <div className="flex flex-col align-center justify-center" >
                <img className="w-full h-48 object-scale-down rounded-t-lg" src={employee.profileIMG} alt="avatar" />
            </div>
            <div className="flex flex-col">
                <div className="flex self-center md:self-auto flex-col md:flex-row  md:justify-around">
                    <h5 className="pb-2 mb-2 text-2xl text-center font-bold tracking-tight text-gray-900">
                        {employee.firstName} {employee.lastName}
                    </h5>
                </div>
                <div className="flex flex-col self-center">
                    <div className="self-start mb-3 ">
                        <div className="flex mb-3">
                            <MaleFaceIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Age: {employee.age}</p>
                        </div>
                        <div className="flex mb-3">
                            <WorkIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Position: {employee.position}</p>
                        </div>

                    </div>
                    <div className="self-start mb-3 ">
                        <div className="flex mb-3">
                            <ContactMailIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">{employee.email}</p>
                        </div>
                    </div>
                    <div className="flex align-center justify-between w-full">
                        <p className="mb-3  mr-2 font-normal text-gray-700">
                            Waiting since: {new Date(parseInt(employee.hiringDate)).toLocaleDateString()}
                        </p>
                        <div className="flex gap-1" >
                            {loading && <LoadingSpinner />}
                            {success != null &&
                                success ? <DoneIcon fill="#06b6d4" /> :
                                success != null && !success ? <ErrorIcon fill="#ef4444" />
                                    :
                                    <></>
                            }
                            <span onClick={() => setAccepted(true)}>
                                <ThinButton text="Approve" />
                            </span>
                            <span onClick={() => setAccepted(false)}>
                                <ThinButton text="Reject" />
                            </span>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    )
}

export default PendingEmployee;
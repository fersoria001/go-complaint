import { useEffect, useState } from "react";
import { EndHiringProcess, ErrorType, User } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import MaleFaceIcon from "../icons/MaleFaceIcon";
import WorkIcon from "../icons/WorkIcon";
import { EndHiringProcessMutation, Mutation } from "../../lib/mutations";
import LoadingSpinner from "../icons/LoadingSpinner";
import DoneIcon from "../icons/DoneIcon";
import ThinButton from "../buttons/ThinButton";
import ErrorOnHoverIcon from "../error/ErrorOnHoverIcon";
import { useParams } from "react-router-dom";
interface Props {
    eventID: string;
    employee: User;
    position: string;
    pendingDate: string;
}
function PendingEmployee({ eventID, employee, position, pendingDate }: Props) {
    const [loading, setLoading] = useState<boolean>(false);
    const [success, setSuccess] = useState<boolean | null>(null);
    const [error, setError] = useState<ErrorType | null>(null);
    const { id } = useParams();
    const handleSubmit = (accept: boolean) => {
        setLoading(true)
        Mutation<EndHiringProcess>(EndHiringProcessMutation,
            { pendingEventID: eventID, enterpriseID: id!, accepted: accept }
        ).then(() => {
            setSuccess(true)
            setLoading(false)
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
        }).catch((error: any) => {
            setSuccess(false)
            setError({ message: error.message, code: error.code })
            setLoading(false)
        })
    }
    useEffect(() => {
    }, [error])
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
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Position: {position}</p>
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
                            Waiting since: {pendingDate}
                        </p>
                        <div className="flex gap-1" >
                            {loading && <LoadingSpinner />}
                            {
                                success != null && success ? <DoneIcon fill="#06b6d4" /> :
                                    error ? <ErrorOnHoverIcon error={error} /> : null
                            }
                            <span onClick={() => handleSubmit(true)}>
                                <ThinButton text="Approve" />
                            </span>
                            <span onClick={() => handleSubmit(false)}>
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
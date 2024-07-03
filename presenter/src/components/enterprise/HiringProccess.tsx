/* eslint-disable @typescript-eslint/no-explicit-any */
import { EnterpriseEventType, HiringProccessType } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import MaleFaceIcon from "../icons/MaleFaceIcon";
import WorkIcon from "../icons/WorkIcon";
import { timeLeft } from "../../lib/time_left";
import CircleFillIcon from "../icons/CircleFillIcon";
import AcceptBtn from "../buttons/AcceptBtn";
import DeclineBtn from "../buttons/DeclineBtn";
import FemaleFaceIcon from "../icons/FemaleFaceIcon";
import { CancelHiringProccessMutation, HireEmployeeMutation, Mutation } from "../../lib/mutations";
import { Route } from "../../routes/$enterpriseID/hiring-procceses";
import { useCallback, useState } from "react";
import { useRouter } from "@tanstack/react-router";
import { timeAgo } from "../../lib/time_ago";
interface Props {
    hiringProccess: HiringProccessType
}
function HiringProccess({ hiringProccess }: Props) {
    const [popUp, setPopUp] = useState<boolean>(false)
    const [reason, setReason] = useState<string>("")
    const [cancelBlock, setCancelBlock] = useState<boolean>(false)
    const [acceptBlock, setAcceptBlock] = useState<boolean>(false)
    const router = useRouter()
    const params = Route.useParams()
    let color = "";
    switch (hiringProccess.status) {
        case "pending":
            color = "#99CCFF";
            break;
        case "rejected":
            color = "#fde68a";
            break;
        case "canceled":
            color = "#FFCCCC";
            break;
        case 'user_accepted':
            color = "#99FFCC"
            break;
        case `hired`:
            color = "#dbeafe"
            break;
        case `leaved`:
            color = "#fde68a"
            break;
        case `fired`:
            color = "#FFCCCC"
            break;
    }
    const handlePopUp = useCallback((): Promise<boolean> => {
        setPopUp(true)
        return new Promise((resolve) => resolve(true))
    }, [])

    const modalCleanUp = useCallback((): void => {
        setPopUp(false)
        setReason("")
    }, [])
    const handleAccept = useCallback((): Promise<boolean> => {
        const res = Mutation<EnterpriseEventType>(HireEmployeeMutation,
            {
                enterpriseName: params.enterpriseID,
                eventID: hiringProccess.eventID
            }).then((res: any) => {
                if (res) {
                    setCancelBlock(true)
                    return res
                }
                throw new Error("Error accepting invitation")
            }).catch((err) => {
                console.error(err)
                router.navigate({ to: "/errors" })
            })
        return res
    }, [params.enterpriseID, hiringProccess.eventID, router])
    const handleDecline = useCallback((reason: string): Promise<boolean> => {
        const res = Mutation<EnterpriseEventType>(CancelHiringProccessMutation,
            {
                enterpriseName: params.enterpriseID,
                eventID: hiringProccess.eventID,
                reason: reason
            }).then((res: any) => {
                if (res) {
                    setAcceptBlock(true)
                    return res
                }
                throw new Error("Error declining invitation")
            }).catch((err) => {
                console.error(err)
                router.navigate({ to: "/errors" })
            })
        return res
    }, [params.enterpriseID, hiringProccess.eventID, router])
    return (
        <div className="relative flex flex-col md:flex-row justify-around items-center
        bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
            <div className="flex flex-col align-center justify-center" >
                <img className="w-full h-48 object-scale-down rounded-t-lg" src={hiringProccess.user.profileIMG} alt="avatar" />
            </div>
            <div className="flex flex-col">
                <div className="flex self-center md:self-auto flex-col md:flex-row  md:justify-around">
                    <h5 className="pb-2 mb-2 text-2xl text-center font-bold tracking-tight text-gray-900">
                        {hiringProccess.user.firstName} {" "} {hiringProccess.user.lastName}
                    </h5>
                </div>
                <div className="flex flex-col self-center">
                    <div className="self-start   mb-3 py-2">
                        <div className="flex mb-3 py-2">{
                            hiringProccess.user.gender === 'female' ?
                                <FemaleFaceIcon fill="#5f6368" /> :
                                <MaleFaceIcon fill="#5f6368" />
                        }
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Age: {hiringProccess.user.age}</p>
                        </div>
                        <div className="flex mb-3 py-2">
                            <ContactMailIcon fill="#5f6368" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">{hiringProccess.user.email}</p>
                        </div>
                        <div className="flex mb-3 py-2">
                            <WorkIcon fill="#5f6368" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Position: {hiringProccess.position}</p>
                        </div>
                        <div className="flex mb-3 py-2">
                            <img src={hiringProccess.emitedBy.profileIMG} alt="flag" className="w-5 h-5 rounded-full" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">
                                Updated by: {hiringProccess.emitedBy.firstName} {" "} {hiringProccess.emitedBy.lastName}  </p>
                        </div>
                    </div>

                    <div className="flex flex-col md:flex-row mb-2 align-center justify-between w-full">
                        <div className="mb-3  mr-2 font-normal text-gray-700">
                            <div className="flex mb-3">
                                <CircleFillIcon fill={color} />
                                <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Status:
                                    {" "}   {hiringProccess.status === 'user_accepted' ? hiringProccess.status.slice(5) : hiringProccess.status}</p>
                            </div>

                        </div>
                        <div className="flex gap-1 pl-2 " >
                            <AcceptBtn
                                variant="thin"
                                text="Accept"
                                callback={handleAccept}
                                status={hiringProccess.status}
                                blocked={acceptBlock}
                            />
                            <DeclineBtn
                                variant="thin"
                                text="Decline"
                                callback={handlePopUp}
                                status={hiringProccess.status}
                                animate={false}
                                blocked={cancelBlock}
                            />
                        </div>
                    </div>

                    <div className="md:flex md:gap-5 self-center p-2 mb-2">
                        <p> Since: {timeAgo(hiringProccess.occurredOn)} </p>
                        {
                            hiringProccess.status === "pending" && (
                                <p> Time left: {timeLeft(new Date(parseInt(hiringProccess.occurredOn)))} </p>
                            )
                        }
                        <p> Last update: {new Date(parseInt(hiringProccess.lastUpdate)).toLocaleDateString()}  </p>
                    </div>
                </div>
            </div>
            {popUp &&
                <div className="absolute top-6  translate-x-1/2 px-5 py-2 flex flex-col rounded-2xl border bg-white z-50">
                    <button
                        type="button"
                        className="self-end text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900
                                        rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center"
                        onClick={() => { setPopUp(false) }}>
                        <svg
                            className="w-3 h-3"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 14 14">
                            <path
                                stroke="currentColor"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth="2"
                                d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6" />
                        </svg>
                        <span
                            className="sr-only">
                            Close modal
                        </span>
                    </button>
                    <div>
                        <textarea rows={3} maxLength={80} placeholder="Reason for declining (optional) "
                            className="text-sm md:text-xl
                        border
                        h-full w-full"
                            onChange={(e) => setReason(e.target.value)}
                        >

                        </textarea>
                        <DeclineBtn
                            variant="thin"
                            text="Decline"
                            callback={handleDecline}
                            callbackArgs={[reason]}
                            status={hiringProccess.status}
                            cleanUp={modalCleanUp}
                        />
                    </div>
                </div>
            }
        </div>
    )
}

export default HiringProccess;
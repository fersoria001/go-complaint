'use client'
import { useEffect, useState } from "react";
import { HiringProccessStatus, HiringProcess } from "@/gql/graphql";
import ContactMailIcon from "@/components/icons/ContactMailIcon";
import WorkIcon from "@/components/icons/WorkIcon";
import CircleFillIcon from "@/components/icons/CircleFillIcon";
import { dateFromMsString } from "@/lib/dateFromMsString";
import timeAgo from "@/lib/timeAgo";
import Image from "next/image";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { acceptHiringInvitation, rejectHiringInvitation } from "@/lib/actions/graphqlActions";
import colorSwitch from "@/lib/colorSwitch";
import { mayHaveReason } from "./constants";
interface Props {
    hiringProccess: HiringProcess
}
function UserHiringProccessCard({ hiringProccess }: Props) {
    const [popUp, setPopUp] = useState<boolean>(false)
    const [reason, setReason] = useState<string>("")
    const [color, setColor] = useState<string>("")
    const queryClient = useQueryClient()
    const acceptMutation = useMutation({
        mutationFn: () => acceptHiringInvitation({
            hiringProcessId: hiringProccess.id,
            userId: hiringProccess.user.id
        }),
        onSuccess: () => queryClient.invalidateQueries({ queryKey: ['hiring-process-by-authenticated-user'] })
    })
    const rejectMutation = useMutation({
        mutationFn: () => rejectHiringInvitation({
            hiringProcessId: hiringProccess.id,
            userId: hiringProccess.user.id,
            rejectionReason: reason
        }),
        onSuccess: () => { queryClient.invalidateQueries({ queryKey: ['hiring-process-by-authenticated-user'] }); setPopUp(false) }
    })

    useEffect(() => {
        if (hiringProccess.status) {
            setColor(colorSwitch(hiringProccess.status))
        }
    }, [hiringProccess.status])
    return (
        <div className="relative flex flex-col md:flex-row justify-around items-center
        bg-white border border-gray-200 rounded-lg shadow py-2">
            <div className="flex flex-col align-center justify-center" >
                <div className='relative w-64 h-48'>
                    <Image
                        priority
                        src={hiringProccess.enterprise?.subjectThumbnail!}
                        className=""
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                        alt="user photo" />
                </div>
            </div>
            <div className="flex flex-col">
                <div className="flex self-center md:self-auto flex-col">
                    <h2 className="pb-2 mb-2 text-2xl text-center font-bold tracking-tight text-gray-900">
                        {hiringProccess.enterprise?.subjectName}
                    </h2>
                    <h3 className="pb-2 mb-2 text-md text-center font-bold tracking-tight text-gray-900">
                        {hiringProccess.industry?.name}
                    </h3>
                </div>
                <div className="flex flex-col self-center">
                    <div className="self-start   mb-3 py-2">
                        <div className="flex mb-3 py-2">
                            <ContactMailIcon fill="#5f6368" className="w-8 h-8" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">{hiringProccess.enterprise?.subjectEmail}</p>
                        </div>
                        <div className="flex mb-3 py-2">
                            <WorkIcon fill="#5f6368" className="w-8 h-8" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Position: {hiringProccess.role}</p>
                        </div>
                        <div className="flex mb-3 py-2">
                            <div className='relative w-8 h-8'>
                                <Image
                                    src={hiringProccess.emitedBy.subjectThumbnail!}
                                    className="rounded-full"
                                    sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                                    fill
                                    alt="user photo" />
                            </div>
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">
                                Updated by: {hiringProccess.emitedBy.subjectName} from {hiringProccess.enterprise?.subjectName}
                            </p>
                        </div>
                    </div>
                    <div className="flex flex-col md:flex-row mb-2 align-center justify-between w-full">
                        <div className="mb-3  mr-2 font-normal text-gray-700">
                            <div className="flex mb-3">
                                <CircleFillIcon fill={color} className="w-8 h-8" />
                                <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">
                                    Status:{" "}
                                    {hiringProccess.status === 'USER_ACCEPTED' ?
                                        hiringProccess.status.replace("_", " ") : hiringProccess.status}
                                </p>
                            </div>
                            {
                                mayHaveReason.findIndex((v) => v == hiringProccess.status) >= 0 &&
                                <div className="py-3  mr-2 font-normal text-gray-700">
                                    <div className="flex mb-3">
                                        <label
                                            htmlFor="reason"
                                            className="pl-2 font-normal text-gray-700 underline underline-offset-8">
                                            Rejection Reason:
                                        </label>
                                        <p
                                            id="reason"
                                            className="pl-2 font-normal text-gray-700">
                                            {hiringProccess.reason}
                                        </p>

                                    </div>
                                </div>
                            }
                        </div>
                        {hiringProccess.status === HiringProccessStatus.Pending &&
                            <div className="flex gap-1 pl-2 " >
                                <button
                                    onMouseUp={() => acceptMutation.mutate()}
                                    className="bg-blue-500 hover:bg-blue-600 font-bold text-white px-5 h-8 rounded-md">
                                    Accept
                                </button>
                                <button
                                    onMouseUp={() => setPopUp(true)}
                                    className="bg-blue-500 hover:bg-blue-600 font-bold text-white px-5 h-8 rounded-md">
                                    Decline
                                </button>
                            </div>}
                    </div>
                    <div className="md:flex md:gap-5 self-center p-2 mb-2">
                        <p className="text-gray-700">
                            Started: {timeAgo(hiringProccess.occurredOn)}
                        </p>
                        <p className="text-gray-700">
                            Last update: {dateFromMsString(hiringProccess.lastUpdate).toUTCString()}
                        </p>
                    </div>
                </div>
            </div>
            {popUp &&
                <div className="absolute translate-y-1/2 px-5 py-2 flex flex-col rounded-2xl border bg-white z-50">
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
                        <textarea
                            rows={3}
                            maxLength={80}
                            placeholder="Reason for declining (optional) "
                            className="text-sm md:text-xl border h-full w-full appearance-none focus:outline-none"
                            onChange={(e) => setReason(e.target.value)}>
                        </textarea>
                        <button
                            onMouseUp={() => rejectMutation.mutate()}
                            className="bg-blue-500 hover:bg-blue-600 font-bold text-white px-5 h-8 rounded-md">
                            Decline
                        </button>
                    </div>
                </div>
            }
        </div>
    )
}

export default UserHiringProccessCard;



{/* {
                            hiringProccess.status === "PENDING" && (
                                <p className="text-gray-700"> Time left: {timeLeft(hiringProccess.occurredOn)} </p>
                            )
                        } */}
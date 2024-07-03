/* eslint-disable @typescript-eslint/no-explicit-any */
import { useCallback, useState } from "react";
import { AcceptHiringInvitationMutation, DeclineHiringInvitationMutation, Mutation } from "../../lib/mutations";
import { DeclineHiringInvitation, HiringInvitationType } from "../../lib/types";
import { useRouter } from "@tanstack/react-router";
import { timeLeft } from "../../lib/time_left";
import AcceptBtn from "../buttons/AcceptBtn";
import DeclineBtn from "../buttons/DeclineBtn";

interface Props {
    invitation: HiringInvitationType
}
function HiringInvitation({ invitation }: Props) {
    const router = useRouter()
    const [popUp, setPopUp] = useState<boolean>(false)
    const [reason, setReason] = useState<string>("")
    const [blocked, setBlocked] = useState<boolean>(false)
    const invitationDate = new Date(parseInt(invitation.occurredOn))
    const handleAccept = useCallback((): Promise<boolean> => {
        const res = Mutation<string>(AcceptHiringInvitationMutation,
            invitation.eventID).then((res: any) => {
                if (res) {
                    const responseBtns = document.getElementById('response-btns')
                    responseBtns?.classList.add('pl-10')
                    setBlocked(true)
                    return res
                }
                throw new Error("Error accepting invitation")
            }).catch((err) => {
                console.error(err)
                router.navigate({ to: "/errors" })
            })
        return res
    }, [invitation.eventID, router])





    const handleDecline = useCallback((reason: string): Promise<boolean> => {
        const res = Mutation<DeclineHiringInvitation>(DeclineHiringInvitationMutation,
            { id: invitation.eventID, reason }).then((res: any) => {
                if (res) {
                    return res
                }
                throw new Error("Error declining invitation")
            }).catch((err) => {
                console.error(err)
                router.navigate({ to: "/errors" })
            })
        return res
    }, [invitation.eventID, router])

    const handlePopUp = useCallback((): Promise<boolean> => {
        setBlocked(true)
        setPopUp(true)
        return new Promise((resolve) => resolve(true))
    }, [])

    const modalCleanUp = useCallback((): void => {
        setPopUp(false)
        setBlocked(false)
        setReason("")
    }, [])
    return (
        <div className="relative md:grid md:grid-cols-3 md:grid-row-2 border-b items-center p-2">
            <div className="md:col-start-1 md:col-end-1 md:row-start-1 md:row-end-1">
                <div className="text-sm md:text-xl text-gray-700 mb-2">
                    {invitation.fullName} we are glad to announce</div>
                <div className="text-sm md:text-xl text-gray-700 mb-2">
                    You have been invited to be part of
                    <p className="underline underline-offset-8">{invitation.enterpriseID}</p>
                </div>
            </div>
            <div className=" flex flex-col md:flex-row border
                 md:border-none p-5 rounded-md">
                <img src={invitation.enterpriseLogoIMG} alt="logo" className="self-center md:self-start
                 md:mt-4 w-20 h-20  rounded-full" />

                <div className="self-center md:self-start md:mt-4 text-center md:whitespace-nowrap">
                    <h1 className="text-sm md:text-xl text-gray-700 mb-2"> {invitation.enterpriseEmail} </h1>
                    <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {invitation.enterprisePhone} </h2>
                </div>
            </div>
            <div className="md:row-start-2 md:col-start-3 md:self-end md:text-end  text-center p-2">{timeLeft(invitationDate)}  before expiration </div>
            <div id='response-btns' className="md:col-start-1 md:row-start-2 flex justify-center py-2">

                <AcceptBtn blocked={blocked} variant="primary" text="Accept" callback={handleAccept} status={invitation.status} />

                <DeclineBtn  blocked={blocked} variant="primary" text="Decline" callback={handlePopUp} status={invitation.status} animate={false} />

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
                        <DeclineBtn variant="primary"
                            text="Decline"
                            callback={handleDecline}
                            callbackArgs={[reason]}
                            status={invitation.status}
                            cleanUp={modalCleanUp}
                            blocked={blocked}
                        />
                    </div>
                </div>
            }
        </div>

    )

}

export default HiringInvitation;
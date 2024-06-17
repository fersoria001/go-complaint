import Realistic from "react-canvas-confetti/dist/presets/realistic";
import { useLoaderData, useNavigate, useParams } from "react-router-dom";
import { ErrorType, HiringInvitation, StringID } from "../lib/types";
import { useEffect, useState } from "react";
import { AcceptHiringInvitationMutation, Mutation } from "../lib/mutations";

function AcceptInvitation() {
    const { id } = useParams();
    const data = useLoaderData() as HiringInvitation | null;
    const [errors, setErrors] = useState<ErrorType>({});
    useEffect(() => {
    }, [errors]);
    const navigate = useNavigate();
    if (!data) {
        navigate("/error/invitation%20not%20found");
        return null;
    } else if (data.seen) {
        navigate("/error/invitation%20already%20solved");
        return null;
    }
    const msg = `You have been invited to be part of ${data.enterprise_id}!, ${data.position_proposal}!`;
    const handleSubmit = () => {
        setErrors({});
        if (!id) { setErrors({ id: "Invalid invitation ID!" }); return; }
        console.log("Accepting invitation...", id!);
        Mutation<StringID>(AcceptHiringInvitationMutation, { id: id! }).then(() => {
            navigate("/success/hiring%20invitation%20accepted");
        }).catch((err) => {
            setErrors({ mutation: err });
        })
    }
    return (
        <div className="relative">
            <div className="">
                <Realistic autorun={{ speed: 0.3 }} />
            </div>
            <div className="absolute transform translate-y-1/2 flex flex-col w-full justify-center">
                <h1
                    className="text-center mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl">
                    Congratulations!
                </h1>
                <p className="text-center mb-6 text-lg font-normal text-gray-500 lg:text-xl sm:px-16 xl:px-48">
                    {msg}
                </p>
                {errors && Object.keys(errors).length > 0 && <div className="text-red-500 text-center mb-6"> {Object.values(errors).join(" ")} </div>}
                <button
                    onClick={handleSubmit}
                    className="self-center inline-flex items-center justify-center px-5 py-3 text-base font-medium text-center text-white
                 bg-cyan-500 rounded-lg hover:bg-cyan-700 focus:ring-4 focus:ring-cyan-300">
                    Accept
                    <svg className="w-3.5 h-3.5 ms-2 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M1 5h12m0 0L9 1m4 4L9 9" />
                    </svg>
                </button>
            </div>
        </div>
    )
}

export default AcceptInvitation;
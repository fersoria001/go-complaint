import { useRef } from "react";
import AcceptBtn from "../buttons/AcceptBtn";
import DeclineBtn from "../buttons/DeclineBtn";
import CloseIcon from "../icons/CloseIcon";
interface Props {
    confirm: () => Promise<boolean>;
    successCleanUp: () => Promise<boolean>;
    cleanUp: () => Promise<boolean>;
}
function SendToReview({ confirm, successCleanUp, cleanUp }: Props) {
    const confirmRef = useRef<HTMLDivElement>(null);
    return (
        <div ref={confirmRef} className="absolute z-50 bg-white
                    p-2.5 w-[200px] md:w-[360px] 
                    left-1/2 -translate-x-1/2 md:top-1/2 
                    border rounded-md shadow flex flex-col ">
            <div className="self-end"><CloseIcon fill="#6b7280" /></div>
            <p className="text-sm md:text-xl text-gray-700 ps-2"> Are you sure ?</p>
            <p className="text-xs md:text-sm text-red-500 italic p-2">
                This will close the complaint,
                no replies could be sent after.
                It will ask for a review from the client,
                later a manager may review the discussion.</p>
            <div className="flex self-center">
                <AcceptBtn
                    variant='primary'
                    text="Yes"
                    callback={confirm}
                    cleanUp={successCleanUp}
                />
                <DeclineBtn
                    variant='primary'
                    text="No"
                    callback={cleanUp}
                    animate={false}
                />
            </div>
        </div>
    )
}

export default SendToReview;
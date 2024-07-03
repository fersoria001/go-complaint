import { useContext, useState } from "react";
import { useRouter } from "@tanstack/react-router";
import { Receiver, ErrorType } from "../../../lib/types";
import { ComplaintContext } from "../../../react-context/ComplaintContext";
import PrimaryButton from "../../buttons/PrimaryButton";

import Stepper from "../../send-complaint/Stepper";
import { Route } from "../../../routes/$enterpriseID";
import Search from "./Search";



function FindReceiver() {
    const { enterpriseID } = Route.useParams();
    const { setKeyValue } = useContext(ComplaintContext);
    const [result, setResult] = useState<Receiver>({} as Receiver);
    const [errors, setErrors] = useState<ErrorType>({});
    const childValue = (receiver: Receiver) => {
        setResult(receiver)
    }
    const router = useRouter()
    const handleClick = () => {
        const errors = setKeyValue("receiverID", result.id)
        if (Object.keys(errors).length > 0) {
            setErrors(errors)
        }
        if (!errors.receiverID) {
            return router.navigate({ to: `/${enterpriseID}/describe-complaint` })
        }

    }

    return (
        <div className="flex flex-col">
            <Search callback={childValue} />
            {<span
                className=" self-center text-red-500 text-xs italic">
                {errors.receiverID}
            </span>
            }
            <span
                onClick={handleClick}
                className="self-center">
                <PrimaryButton text="Next step" />
            </span>
            <div
                className="self-center px-2 md:px-6">
                <Stepper step={1} />
            </div>
        </div>
    )
}

export default FindReceiver
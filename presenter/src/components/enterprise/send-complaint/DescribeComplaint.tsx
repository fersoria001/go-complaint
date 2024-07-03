import { useContext, useState } from "react";

import { useRouter } from "@tanstack/react-router";
import { ErrorType } from "../../../lib/types";
import { ComplaintContext } from "../../../react-context/ComplaintContext";
import PrimaryButton from "../../buttons/PrimaryButton";
import Reason from "../../send-complaint/Reason";
import Stepper from "../../send-complaint/Stepper";
import Description from "./Description";
import { Route } from "../../../routes/$enterpriseID";

function DescribeComplaint() {
    const { enterpriseID } = Route.useParams()
    const { setKeyValue } = useContext(ComplaintContext);
    const [title, setTitle] = useState<string>("")
    const [errors, setErrors] = useState<ErrorType>({})
    const [description, setDescription] = useState<string>("")
    const router = useRouter()
    const childReasonValue = (reason: string) => {
        setTitle(reason)
    }
    const childDescriptionValue = (description: string) => {
        setDescription(description)
    }
    const handleClick = () => {
        let errors = setKeyValue("title", title)
        errors = setKeyValue("description", description)
        if (Object.keys(errors).length > 0) {
            setErrors(errors)
        }
        if (errors.title || errors.description) {
            return
        }
        return router.navigate({ to: `/${enterpriseID}/complain` })
    }
    return (
        <div className="flex flex-col">
            <div className="mb-2 w-full px-4 md:mb-[17px] md:w-2/3 md:px-0 self-center">
                <Reason callback={childReasonValue} />
                {
                    errors.title &&
                    <span
                        className=" self-center text-red-500 text-xs italic">
                        {errors.title}
                    </span>
                }
                <Description callback={childDescriptionValue} />
                {
                    errors.description &&
                    <span
                        className=" self-center text-red-500 text-xs italic">
                        {errors.description}
                    </span>
                }
            </div>
            <span
                onClick={handleClick}
                className="self-center">
                <PrimaryButton text="Next step" />
            </span>
            <div
                className="self-center px-2 md:px-6">
                <Stepper step={2} />
            </div>

        </div>
    );
}

export default DescribeComplaint;
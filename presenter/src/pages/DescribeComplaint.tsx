import { useState, useContext, useEffect } from "react";
import PrimaryButton from "../components/buttons/PrimaryButton";
import Description from "../components/complaint-writer/Description";
import Reason from "../components/complaint-writer/Reason";
import Stepper from "../components/complaint-writer/Stepper";
import { DescriptionValidationSchema } from "../lib/types";
import { useLocation, useNavigate } from "react-router-dom";
import { ComplaintContext } from "../react-context/user-create-complaint/ComplaintContext";

function DescribeComplaint() {
    const [errors, setErrors] = useState<{ [key: string]: string }>({});
    const { complaintData } = useContext(
        ComplaintContext
    );
    const location = useLocation();
    const navigate = useNavigate();
    const handleNextStep = () => {
        console.log(complaintData)
        setErrors({});
        const parsed = DescriptionValidationSchema.safeParse({
            reason: complaintData?.reason || "",
            description: complaintData?.description || ""
        })
        if (!parsed.success) {
            parsed.error.errors.forEach((error) => {
                errors[error.path.join("")] = error.message;
            });
            setErrors(errors);
            return
        }
        const url = location.pathname.split("/").slice(0, -1).join("/") + "/complain";
        return navigate(url);
    }
    useEffect(() => {
    }, [errors]);
    return (
        <div className="flex flex-col pt-4 md:pt-8">
            <div className="mb-20 md:mb-10 h-40 md:h-56 lg:h-96 w-full px-4 md:w-2/3 md:px-0 self-center ">
                <Reason />
                {errors.reason && <span
                    className=" self-center text-red-500 text-xs italic">
                    {errors.reason} </span>}
                <Description />
                {errors.description && <span
                    className=" self-center text-red-500 text-xs italic">
                    {errors.description} </span>}
            </div>

            <span
                onClick={() => handleNextStep()}
                className="self-center pt-2">
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
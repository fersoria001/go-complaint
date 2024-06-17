import Stepper from "../components/complaint-writer/Stepper";
import Search from "../components/complaint-writer/Search";
import PrimaryButton from "../components/buttons/PrimaryButton";
import { useContext, useEffect, useState } from "react";
import { ComplaintContext } from "../react-context/ComplaintContext";
import { ReceiverValidationSchema } from "../lib/types";
import { useLocation, useNavigate } from "react-router-dom";

function FindReceiver() {
    const { complaintData } = useContext(ComplaintContext);
    const [errors, setErrors] = useState<{ [key: string]: string }>({});
    const [showErrors, setShowErrors] = useState<boolean>(false);
    const navigate = useNavigate();
    const location = useLocation();
    useEffect(() => {
        ReceiverValidationSchema
            .spa({ term: complaintData?.receiverID || "" })
            .then((result) => {
                if (!result.success) {
                    result.error.errors.forEach((error) => {
                        errors[error.path.join("")] = error.message;
                    });
                    setErrors(errors);
                }
            });
    }, [complaintData, errors]);
    const cleanUpErrors = () => {
        setShowErrors(false);
        setErrors({});
    }
    const handleNextStep = () => {
        if (!errors.term) {
            const url = location.pathname;
            return navigate(url + "/describe");
        }
        setShowErrors(true);
        return
    }
    return (
        <div className="flex flex-col pt-6">
            <Search callbackFn={cleanUpErrors} />
            {showErrors &&
                errors.term &&
                <span
                    className=" self-center text-red-500 text-xs italic">
                    {errors.term}
                </span>
            }
            <span
                onClick={() => handleNextStep()}
                className="self-center py-2">
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
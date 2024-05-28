import { useContext, useState, useEffect } from "react";
import PrimaryButton from "../components/buttons/PrimaryButton";
import Stepper from "../components/complaint-writer/Stepper";
import useWindowDimensions from "../lib/hooks/useWindowsDimensions";
import { ComplaintContext } from "../react-context/user-create-complaint/ComplaintContext";
import { ComplaintBodyValidationSchema, Enterprise, isUserDescriptor, SendComplaint, UserDescriptor } from "../lib/types";
import Modal from "../components/complaint-writer/Modal";
import { useLoaderData, useNavigate } from "react-router-dom";
import { Mutation, SendComplaintMutation } from "../lib/mutations";

function Complain() {
    const loggedData = useLoaderData() as UserDescriptor | Enterprise;
    const [errors, setErrors] = useState<{ [key: string]: string }>({});
    const [showModal, setShowModal] = useState<boolean>(false);
    const [accepted, setAccepted] = useState<boolean>(false);
    const navigate = useNavigate();
    const { complaintData, updateState } = useContext(
        ComplaintContext
    );
    const { width } = useWindowDimensions();
    const [rows, setRows] = useState<number>(4);
    useEffect(() => {
        let senderID = "" 
        if (isUserDescriptor(loggedData)) {
            const casted = loggedData as UserDescriptor;
            senderID = casted.email;
        } else {
            const casted = loggedData as Enterprise;
            senderID = casted.name;
        }
        if (width >= 768) {
            setRows(8);
        } else {
            setRows(4);
        }
        if (!accepted) return
        Mutation<SendComplaint>(SendComplaintMutation, {
            reason: complaintData!.reason,
            description: complaintData!.description,
            body: complaintData!.body,
            senderID: senderID,
            receiverID: complaintData!.receiverID,
            fullName: complaintData!.fullName,
        } as SendComplaint).then(() => {
            return navigate("/success/sent%20complaint");
        }).catch((error) => {
            console.error(error);
        });
    }, [width, accepted, complaintData, loggedData, navigate,errors]);
    const placeholder = `Complain to ${complaintData?.fullName} about ${complaintData?.reason}`
    const handleOnChange = (input: string) => {
        updateState({
            complaintData: {
                fullName: complaintData?.fullName || "",
                senderID: complaintData?.senderID || "",
                receiverID: complaintData?.receiverID || "",
                reason: complaintData?.reason || "",
                description: complaintData?.description || "",
                body: input || ""
            }
        });
    }
    const handleSubmit = () => {
        const parsed = ComplaintBodyValidationSchema.safeParse({
            body: complaintData?.body || ""
        })
        if (!parsed.success) {
            parsed.error.errors.forEach((error) => {
                errors[error.path.join("")] = error.message;
            });
            setErrors(errors);
            console.log(errors);
            return
        }
        setShowModal(true);
        console.log("show modal", showModal);
    }
    return (
        <div className="flex flex-col pt-4 md:pt-8">
            <div className="mb-8 md:mb-10 h-40 md:h-56 lg:h-96 w-full px-4 md:px-0 md:w-2/3 self-center ">
                <label
                    htmlFor="complain"
                    className="block mb-2 text-sm font-medium text-gray-900"
                >Complain about it
                </label>
                <textarea
                    id="complaint"
                    rows={rows}
                    minLength={50}
                    maxLength={250}
                    onChange={(e) => handleOnChange(e.target.value)}
                    className="block p-2.5 w-full text-sm
               text-gray-900 bg-gray-50 rounded-lg border border-gray-300
                focus:ring-blue-500 focus:border-blue-500"
                    placeholder={placeholder}>
                </textarea>
            </div>
            {errors.body && <span
                className=" self-center text-red-500 text-xs italic">
                {errors.body} </span>}
            <span
                onClick={() => handleSubmit()}
                className="self-center pt-2">
                <PrimaryButton text="Complain!" />
            </span>
            <div
                className="self-center px-2 md:px-6">
                <Stepper step={3} />
            </div>
            {showModal && (<Modal
                id="confirm-complaint-modal"
                show={showModal}
                validatedObject={complaintData!}
                callbackFn={() => { setAccepted(true) }}
                closeFn={() => { setShowModal(false) }}
            />)}

        </div>
    );
}

export default Complain;
import { useContext } from "react";
import { ComplaintContext } from "../../react-context/user-create-complaint/ComplaintContext";
import ExclamationIcon from "../icons/ExclamationIcon"

function Reason() {
    const { complaintData, updateState } = useContext(
        ComplaintContext
    );
    const handleOnChange = (input: string) => {
        updateState({
            complaintData: {
                fullName: complaintData?.fullName || "",
                senderID: complaintData?.senderID || "",
                receiverID: complaintData?.receiverID || "",
                reason: input || "",
                description: complaintData?.description || "",
                body: complaintData?.body || ""
            }
        });
    }
    return (
        <div
            className="relative mb-6">
            <div
                className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none">
                <ExclamationIcon fill="#06b6d4" />
            </div>
            <input
                onChange={(e) => handleOnChange(e.target.value)}
                type="text"
                id="input-group-1"
                minLength={10}
                maxLength={80}
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg
               focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5 "
                placeholder="Why do you complain?"
            />
        </div>
    )
}

export default Reason
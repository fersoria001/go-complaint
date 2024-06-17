import { useContext, useEffect, useState } from "react";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";
import { ComplaintContext } from "../../react-context/ComplaintContext";

function Description() {
    const { complaintData, updateState } = useContext(
        ComplaintContext
    );
    const { width } = useWindowDimensions();
    const [rows, setRows] = useState<number>(4);
    useEffect(() => {
        if (width >= 768) {
            setRows(8);
        } else {
            setRows(4);
        }
    }, [width]);
    const handleOnChange = (input: string) => {
        updateState({
            complaintData: {
                fullName: complaintData?.fullName || "",
                senderID: complaintData?.senderID || "",
                receiverID: complaintData?.receiverID || "",
                reason: complaintData?.reason || "",
                description: input || "",
                body: complaintData?.body || ""
            }
        });
    }
    return (
        <>
            <label
                htmlFor="description"
                className="block mb-2 text-sm font-medium text-gray-900"
            >Description
            </label>
            <textarea
                onChange={(e) => handleOnChange(e.target.value)}
                id="description"
                rows={rows}
                minLength={3}
                maxLength={120}
                className="block p-2.5 w-full text-sm
               text-gray-900 bg-gray-50 rounded-lg border border-gray-300
                focus:ring-blue-500 focus:border-blue-500"
                placeholder="Shortly describe the problem here...">
            </textarea>
        </>
    )
}

export default Description
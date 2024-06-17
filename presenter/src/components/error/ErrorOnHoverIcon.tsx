import { useState } from "react";
import { ErrorType } from "../../lib/types";
import ErrorIcon from "../icons/ErrorIcon";

interface Props {
    error: ErrorType;
}

function ErrorOnHoverIcon({ error }: Props) {
    const [showError, setShowError] = useState<boolean>(false);
    const handleMouseEvents = () => {
        setShowError(!showError);
    };
    return (
        <div className="relative ">
            {showError && (
                <div
                    className="
                    pointer-events-none
                    p-2.5 absolute bg-white border rounded-md shadow-sm z-50 text-red-500 text-xs italic"
                >
                    {error.message}
                </div>
            )}
            <span
                
                onMouseEnter={handleMouseEvents}
                onMouseLeave={handleMouseEvents}>
                <ErrorIcon fill="#ef4444" />
            </span>
        </div>
    )
}

export default ErrorOnHoverIcon;
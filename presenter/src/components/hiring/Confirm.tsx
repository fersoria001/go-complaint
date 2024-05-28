import { InviteToProject } from "../../lib/types";
interface Props {
    id: string;
    validatedObject: InviteToProject;
    show: boolean;
    callbackFn: () => void;
    closeFn: () => void;
}
function ConfirmHiringModal({ id, closeFn, validatedObject, callbackFn }: Props) {
    const { enterpriseName, userFullName, position } = validatedObject;
    const handleAccept = () => {
        callbackFn();
        closeFn();
    }
    const handleClose = () => {
        closeFn();
    }

    return (
        <>
            <div
                id={id}
                tabIndex={-1}
                className="overflow-x-hidden w-full
                absolute bottom-0 md:top-1/4 md:left-1/3 md:translate-x--1/4 z-50 
              justify-center items-center  md:h-full">
                <div
                    className=" relative p-4 w-full max-w-2xl max-h-full">
                    <div
                        className="flex flex-col relative bg-white rounded-lg shadow">
                        <div
                            className="flex items-center justify-between p-4 md:p-5 border-b rounded-t">
                            <h3
                                className="text-center w-full text-xl font-semibold text-red-500 ">
                                Warning
                            </h3>
                            <button
                                type="button"
                                className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900
                              rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center"
                                onClick={handleClose}
                            >
                                <svg
                                    className="w-3 h-3"
                                    xmlns="http://www.w3.org/2000/svg"
                                    fill="none"
                                    viewBox="0 0 14 14">
                                    <path
                                        stroke="currentColor"
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                        strokeWidth="2"
                                        d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6" />
                                </svg>
                                <span
                                    className="sr-only">
                                    Close Confirm
                                </span>
                            </button>
                        </div>
                        <div
                            className="p-4 md:p-5 space-y-4 overflow-y-auto">
                            <p
                                className="text-base leading-relaxed text-gray-500">
                                You are about to invite {userFullName} to join {enterpriseName} as a {position}.
                            </p>
                            <p className="text-center text-red-500 text-xs italic">
                                Note that you can't undo this action and it will be sent to the receiver,
                                You will be notified if the receiver accepts or declines the invitation.
                                If him/her accepts, you will be prompt for confirmation to proceed with the hiring process.
                            </p>
                        </div>

                        <p className="text-center text-xl font-semibold text-gray-900">
                            Do you want to proceed?
                        </p>
                        <div
                            className="self-center flex items-center p-4 md:p-5 border-t border-gray-200 rounded-b">
                            <button
                                onClick={handleAccept}
                                type="button"
                                className="text-white bg-blue-700 hover:bg-blue-800 
                             focus:ring-4 focus:outline-none focus:ring-blue-300
                              font-medium rounded-lg text-sm px-5 py-2.5 text-center">
                                I accept
                            </button>
                            <button
                                onClick={handleClose}
                                type="button"
                                className="py-2.5 px-5 ms-3
                              text-sm font-medium text-gray-900
                               focus:outline-none bg-white rounded-lg border border-gray-200
                                hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4
                                 focus:ring-gray-100">
                                Decline
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}
export default ConfirmHiringModal;
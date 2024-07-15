'use client'
import { confirmSignIn } from "@/lib/actions/authentication";
import confirmationCodeSchema from "@/lib/validation/confirmationCodeSchema";
import { useFormState, useFormStatus } from "react-dom";
import { z } from "zod";
import InlineAlert from "../error/InlineAlert";
type stateType = z.inferFlattenedErrors<typeof confirmationCodeSchema>
const initialState: Partial<stateType> = {}
const ConfirmSignInForm: React.FC = () => {
    const { pending } = useFormStatus()
    const [state, formAction] = useFormState(confirmSignIn, initialState)
    return (
        <form
            action={formAction}
            className="my-10 lg:my-16 flex flex-col md:mt-12 border-t
             bg-white shadow-md rounded-md px-4 py-8 max-w-lg mx-auto">
            <div className="mb-4">
                {
                    state?.formErrors &&
                    <InlineAlert errors={state.formErrors} />
                }
                <label
                    className="block"
                    htmlFor="confirmationCode">
                    <h3 className="text-gray-700 mb-4 font-bold text-md md:text-lg">
                        We need to validate your identity before signing in.
                    </h3>
                    <p className="text-gray-700 text-sm md:text-md mb-4">
                        For security concerns we have sent you an email with a confirmation code
                        to the address you registered with to Go Complaint.
                    </p>
                    <p className="text-gray-700 text-sm md:text-md mb-4">
                        Enter the confirmation code you obtained below:
                    </p>
                </label>

                <input
                    className="appearance-none border rounded w-full py-2 px-3 text-md lg:text-lg
                 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="confirmationCode" name="confirmationCode"/>
                {
                    state?.fieldErrors?.confirmationCode &&
                    <InlineAlert
                        errors={state.fieldErrors.confirmationCode} />
                }
            </div>
            <button
                className="mx-auto bg-blue-500 px-12 py-2.5 rounded-md hover:bg-blue-700
                text-white text-md lg:text-lg"
                type="submit"
                disabled={false}>
                Confirm
            </button>
        </form>
    )
}
export default ConfirmSignInForm;
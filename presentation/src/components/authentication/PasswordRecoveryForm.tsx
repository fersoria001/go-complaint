'use client'
import { recoverPassword } from "@/lib/actions/authentication";
import signInSchema from "@/lib/validation/signInSchema";
import { useFormState } from "react-dom";
import { z } from "zod";
import InlineAlert from "../error/InlineAlert";
const justTheEmail = signInSchema.pick({ userName: true })
type formState = z.inferFlattenedErrors<typeof justTheEmail>
const initialState: Partial<formState> = {}
const PasswordRecoveryForm: React.FC = () => {
    const [state, formAction] = useFormState(recoverPassword, initialState)
    return (
        <form
            action={formAction}
            className="bg-white shadow-md rounded px-4 py-8 my-12 border-t max-w-md mx-auto flex flex-col">
            <div className="mb-4">
                {state.formErrors && <InlineAlert errors={state.formErrors} />}
                <label
                    className="block"
                    htmlFor="email">
                    <h3 className="text-gray-700 mb-4 font-bold text-md md:text-lg">
                        Recover your password
                    </h3>
                    <p className="text-gray-700 text-sm md:text-md mb-4">
                        First we need the email on wich you have registered on Go Complaint
                    </p>
                    <p className="text-gray-700 text-sm md:text-md mb-4">
                        Note that your password will be reset and a new one will be
                        delivered to your email.
                    </p>
                    <p className="text-gray-700 text-sm md:text-md mb-4">Enter the email below:</p>
                </label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="email" id="email" name="email" autoComplete="username" />
                {state.fieldErrors?.userName && <InlineAlert errors={state.fieldErrors.userName} />}
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
export default PasswordRecoveryForm;
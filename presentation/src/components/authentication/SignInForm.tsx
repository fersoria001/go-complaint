'use client'
import { userSignIn } from "@/lib/actions/authentication";
import signInSchema from "@/lib/validation/signInSchema";
import Link from "next/dist/client/link";
import { useFormStatus, useFormState } from "react-dom";
import { z } from "zod";
import InlineAlert from "../error/InlineAlert";
type stateType = z.inferFlattenedErrors<typeof signInSchema>
const initialState: Partial<stateType> = {}
const SignInForm: React.FC = () => {
    const { pending } = useFormStatus()
    const [state, formAction] = useFormState(userSignIn, initialState)
    return (
        <form
            action={formAction}
            className="my-10 border-t bg-white shadow-md rounded-md py-6 flex flex-col max-w-sm mx-auto">
            <div className="mb-4 px-4">
                {
                    state?.formErrors &&
                    <InlineAlert
                        className="w-full mb-0.5 mt-0.5 py-0.5 px-0.5 inline-flex items-center text-sm text-red-800 rounded-lg bg-red-50"
                        errors={state.formErrors} />
                }
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3 text-md lg:text-lg
                 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="email" id="email" name="email" autoComplete="username" />
                {
                    state?.fieldErrors?.email &&
                    <InlineAlert
                        className="w-full mt-0.5 py-0.5 px-0.5 inline-flex items-center text-sm text-red-800 rounded-lg bg-red-50"
                        errors={state.fieldErrors.email} />
                }
            </div>
            <div className="mb-4 px-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="password">Password</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="password" id="password" name="password" autoComplete="password" />
                {
                    state?.fieldErrors?.password &&
                    <InlineAlert
                        className="w-full mt-0.5 py-0.5 px-0.5 inline-flex items-center text-sm text-red-800 rounded-lg bg-red-50"
                        errors={state.fieldErrors.password} />
                }
            </div>
            <div className="px-4 mb-4 me-auto">
                <div className="flex items-start">
                    <div className="flex items-center h-5">
                        <input
                            className="w-4 h-4 border border-gray-300 rounded bg-gray-50
                            text-md lg:text-lg focus:ring-3 focus:ring-blue-300"
                            type="checkbox" id="rememberMe" name="rememberMe"
                        />
                    </div>
                    <label
                        className="block text-gray-700 text-sm lg:text-md font-bold mb-2 ms-2"
                        htmlFor="rememberMe">Remember me</label>
                </div>
                {
                    state?.fieldErrors?.rememberMe &&
                    <InlineAlert
                        className="w-full mt-0.5 py-0.5 px-0.5 inline-flex items-center text-sm text-red-800 rounded-lg bg-red-50"
                        errors={state.fieldErrors.rememberMe} />
                }
            </div>
            <div className="ms-auto mb-4 px-4">
                <div className="flex flex-col">
                    <Link
                        href="/password-recovery"
                        className="text-gray-700 text-sm lg:text-md font-medium mb-2 cursor-pointer
                    hover:underline">
                        Forgot your password?
                    </Link>
                    <Link
                        href="/sign-up"
                        className="self-start text-gray-700 text-sm lg:text-md font-medium mb-2 cursor-pointer
                    hover:underline">
                        Register a new account
                    </Link>
                </div>
            </div>
            <button
                className="mx-auto bg-blue-500 px-12 py-2.5 rounded-md hover:bg-blue-700
                text-white text-md lg:text-lg"
                type="submit"
                disabled={pending}>
                Sign in
            </button>
        </form>
    )
}
export default SignInForm;
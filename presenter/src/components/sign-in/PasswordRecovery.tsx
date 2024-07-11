/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";
import PrimaryButton from "../buttons/PrimaryButton";
import { ErrorType } from "../../lib/types";
import { Mutation, PasswordRecoveryMutation } from "../../lib/mutations";
import { z } from "zod";
import { useRouter } from "@tanstack/react-router";
const validEmail = z.string().email({ message: "Please enter a valid email" })
const PasswordRecovery: React.FC = () => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [email, setEmail] = useState<string>("")
    const router = useRouter()
    const handleRecovery = async () => {
        setErrors({})
        try {
            const result = validEmail.safeParse(email)
            if (!result.success) {
                setErrors({ email: result.error.errors[0].message })
                return
            }
            await Mutation<string>(PasswordRecoveryMutation, email)
            return router.navigate({ to: "/recovery-succeed" })
        } catch (error: any) {
            setErrors({ form: error.message })
        }
    }
    return (
        <div className="h-screen px-2 md:px-0">
            <form className="bg-white shadow-md rounded px-4 pt-6 pb-8 mb-4 mt-24 border max-w-lg mx-auto">
                <div className="mb-4">
                    {errors?.form && <span className="mx-auto text-red-500 text-xs italic" >{errors.form}</span>}
                    <label
                        className="block text-gray-700 text-sm font-bold mb-2"
                        htmlFor="email">Enter your email and we'll send you an email with the next steps to follow
                    </label>
                    <input
                        onChange={(e: any) => setEmail(e.target.value)}
                        className="shadow appearance-none border rounded w-full py-2 px-3
                         text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        type="email" id="email" name="email" />
                    {errors?.email && <span className="text-red-500 text-xs italic" >{errors.email}</span>}
                </div>
                <span className="text-center" onMouseUp={handleRecovery}>
                    <PrimaryButton text="Confirm" />
                </span>
            </form>
        </div>
    )
}
export default PasswordRecovery;
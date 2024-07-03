/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";
import { ErrorType, passwordRegex } from "../../../lib/types";
import AcceptBtn from "../../buttons/AcceptBtn";
import { z } from "zod";
import { updatePassword } from "./settings_lib";
const validatePassword = z.object({
    oldPassword: z
        .string()
        .regex(
            passwordRegex,
            "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
        ),
    newPassword: z
        .string()
        .regex(
            passwordRegex,
            "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
        ),
    confirmPassword: z
        .string()
        .regex(
            passwordRegex,
            "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
        )
}).superRefine(({ confirmPassword, newPassword }, ctx) => {
    if (confirmPassword !== newPassword) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: ["confirmPassword"],
            message: "The passwords did not match",
        });
    }
});
const UpdatePassword: React.FC = () => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [reset, setReset] = useState<boolean>(false)
    const [oldPassword, setOldPassword] = useState<string>("")
    const [newPassword, setNewPassword] = useState<string>("")
    const [confirmPassword, setConfirmPassword] = useState<string>("")
    const handleOldPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setOldPassword(e.target.value)
        setReset(true)
    }
    const handleNewPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setNewPassword(e.target.value)
        setReset(true)
    }

    const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setConfirmPassword(e.target.value)
        setReset(true)
    }

    const handleUpdate = async () => {
        const result = validatePassword.safeParse({
            oldPassword,
            newPassword,
            confirmPassword
        })
        if (!result.success) {
            setErrors({ lastName: result.error.errors[0].message })
            return false
        }
        try {
            return await updatePassword(oldPassword, newPassword)
        } catch (e: any) {
            if (e.message.includes("old password and new password must be different")) {
                setErrors({ password: "old password and new password must be different" })
                return false
            }
        }
        return false
    }
    const handleCleanUp = () => {
        setReset(false)
    }

    return (
        <div className="flex flex-col">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="password">Old Password</label>
                <input
                    onChange={handleOldPasswordChange}
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="password" type="password" placeholder="Password" />
                {errors?.password && <span className="text-red-500 text-xs italic">{errors.password}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="newPassword">New password</label>
                <input
                    onChange={handleNewPasswordChange}
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="newPassword" type="password" placeholder="Confirm Password" />
                {errors?.newPassword && <span className="text-red-500 text-xs italic">{errors.newPassword}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="confirmPassword">Confirm Password</label>
                <input
                    onChange={handleConfirmPasswordChange}
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="confirmPassword" type="password" placeholder="Confirm Password" />
                {errors?.confirmPassword && <span className="text-red-500 text-xs italic">{errors.confirmPassword}</span>}
            </div>
            <div className="self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdatePassword;
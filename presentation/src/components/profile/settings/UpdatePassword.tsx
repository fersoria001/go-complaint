'use client'
import InlineAlert from "@/components/error/InlineAlert";
import CheckIcon from "@/components/icons/CheckIcon";
import { ChangePassword, UserDescriptor } from "@/gql/graphql";
import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { updatePassword } from "@/lib/actions/graphqlActions";
import { useMutation, useSuspenseQuery } from "@tanstack/react-query";
import { useState } from "react";
import { z } from "zod";

export const passwordRegex = new RegExp(
    /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/
);
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
}).superRefine(({ oldPassword, newPassword }, ctx) => {
    if (oldPassword === newPassword) {
        ctx.addIssue({
            code: z.ZodIssueCode.custom,
            path: ["confirmPassword"],
            message: "old password and new password must be different",
        });
    }
})
interface Props {
    descriptor: UserDescriptor
}
const UpdatePassword: React.FC<Props> = ({ descriptor }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof validatePassword> | undefined>()
    const [oldPassword, setOldPassword] = useState<string>("")
    const [newPassword, setNewPassword] = useState<string>("")
    const [confirmPassword, setConfirmPassword] = useState<string>("")
    const handleOldPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setOldPassword(e.target.value)
    }
    const handleNewPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setNewPassword(e.target.value)
    }
    const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setConfirmPassword(e.target.value)
    }

    const updateMutation = useMutation({
        mutationFn: async ({ oldPassword, newPassword }: Omit<ChangePassword, "username">) => await updatePassword({
            username: descriptor.userName,
            oldPassword: oldPassword,
            newPassword: newPassword,
        }),
        onError: (e) => console.log(e),
    })

    const handleUpdate = async () => {
        const { success, data, error } = validatePassword.safeParse({
            oldPassword,
            newPassword,
            confirmPassword
        })
        if (!success) {
            setErrors(error.flatten())
            return
        }
        updateMutation.mutate({ oldPassword: data.oldPassword, newPassword: data.newPassword })
    }


    return (
        <form className="flex flex-col w-full md:w-2/3 mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="password">Old Password</label>
                <input
                    onChange={handleOldPasswordChange}
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="password" type="password" placeholder="Password" autoComplete="password" />
                {
                    errors?.fieldErrors.oldPassword &&
                    <InlineAlert
                        className="flex items-center p-4 mb-2 text-sm text-red-800 rounded-lg bg-red-50"
                        errors={errors.fieldErrors.oldPassword} />
                }
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
                {errors?.fieldErrors.newPassword &&
                    <InlineAlert
                        className="flex items-center p-4 mb-2 text-sm text-red-800 rounded-lg bg-red-50"
                        errors={errors.fieldErrors.newPassword} />
                }
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
                {errors?.fieldErrors.confirmPassword &&
                    <InlineAlert
                        className="flex items-center p-4 mb-2 text-sm text-red-800 rounded-lg bg-red-50"
                        errors={errors.fieldErrors.confirmPassword} />
                }
            </div>
            <div className="self-end flex">
                {updateMutation.isSuccess && <CheckIcon className="w-6 h-6 my-auto fill-blue-300" />}
                <button
                    type="button"
                    onClick={() => handleUpdate()}
                    className="px-6 py-3 bg-blue-500 hover:bg-blue-600 rounded-md text-white font-bold">
                    Submit
                </button>
            </div>
            {updateMutation.isError && updateMutation.error.message.includes("crypto/bcrypt: hashedPassword is not the hash of the given password")
                && <InlineAlert errors={["the old password field does not match your current password"]} />
            }
        </form>
    )
}

export default UpdatePassword;
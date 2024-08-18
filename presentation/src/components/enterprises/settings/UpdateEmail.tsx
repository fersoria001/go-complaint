'use client'
import InlineAlert from "@/components/error/InlineAlert"
import CheckIcon from "@/components/icons/CheckIcon"
import { Enterprise, ChangeEnterpriseEmail } from "@/gql/graphql"
import { changeEnterpriseEmail } from "@/lib/actions/graphqlActions"
import { useMutation } from "@tanstack/react-query"
import { useState } from "react"
import { z } from "zod"

const validemail = z.object({
    email: z.string().email({ message: "Please enter a valid email" })
})
interface Props {
    enterprise: Enterprise
}
const UpdateEmail: React.FC<Props> = ({ enterprise }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof validemail> | undefined>()
    const [email, setEmail] = useState<string>("")
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setEmail(e.target.value)
    }
    const updateMutation = useMutation({
        mutationFn: async ({ newEmail }: Omit<ChangeEnterpriseEmail, "enterpriseId">) => await changeEnterpriseEmail({
            enterpriseId: enterprise.id,
            newEmail
        }),
        onError: (e) => console.log(e),
    })
    const handleUpdate = async () => {
        const { success, data, error } = validemail.safeParse({
            email
        })
        if (!success) {
            setErrors(error.flatten())
            return
        }
        updateMutation.mutate({ newEmail: data.email })
    }
    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto ">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    onChange={handleChange}
                    value={email}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="email" type="text" />
                {errors?.fieldErrors.email && <InlineAlert errors={errors.fieldErrors.email} />}
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
            {
                updateMutation.isError
                && <InlineAlert errors={[updateMutation.error.message]} />
            }
        </div>
    )
}

export default UpdateEmail;
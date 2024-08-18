import InlineAlert from "@/components/error/InlineAlert";
import CheckIcon from "@/components/icons/CheckIcon";
import { ChangeUserLastName, UserDescriptor } from "@/gql/graphql";
import { changeLastName } from "@/lib/actions/graphqlActions";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { z } from "zod";

const validName = z.object({
    lastName: z.string()
        .min(2, { message: "Last name must be at least 2 characters long" })
        .max(50, { message: "Last name must be at most 50 characters long" })
})

interface Props {
    descriptor: UserDescriptor
}

const UpdateFirstName: React.FC<Props> = ({ descriptor }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof validName> | undefined>()
    const [lastName, setFirstName] = useState<string>("")
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setFirstName(e.target.value)
    }
    const updateMutation = useMutation({
        mutationFn: async ({ newLastName }: Omit<ChangeUserLastName, "userId">) => await changeLastName({
            userId: descriptor.id,
            newLastName
        }),
        onError: (e) => console.log(e),
    })
    const handleUpdate = async () => {
        const { success, data, error } = validName.safeParse({
            lastName
        })
        if (!success) {
            setErrors(error.flatten())
            return
        }
        updateMutation.mutate({ newLastName: data.lastName })
    }
    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="lastName">Last name
                </label>
                <input
                    onChange={handleChange}
                    value={lastName}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="lastName" type="text" />
                {
                    errors?.fieldErrors.lastName &&
                    <InlineAlert
                        className="flex items-center p-4 mb-2 text-sm text-red-800 rounded-lg bg-red-50"
                        errors={errors.fieldErrors.lastName} />
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
            {
                updateMutation.isError
                && <InlineAlert errors={[updateMutation.error.message]} />
            }
        </div>
    )
}

export default UpdateFirstName;
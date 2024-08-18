import InlineAlert from "@/components/error/InlineAlert";
import CheckIcon from "@/components/icons/CheckIcon";
import { ChangeUserFirstName, UserDescriptor } from "@/gql/graphql";
import { changeFirstName } from "@/lib/actions/graphqlActions";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { z } from "zod";

const validName = z.object({
    firstName: z.string()
        .min(2, { message: "First name must be at least 2 characters long" })
        .max(50, { message: "First name must be at most 50 characters long" })
})

interface Props {
    descriptor: UserDescriptor
}

const UpdateFirstName: React.FC<Props> = ({ descriptor }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof validName> | undefined>()
    const [firstName, setFirstName] = useState<string>("")
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setFirstName(e.target.value)
    }
    const updateMutation = useMutation({
        mutationFn: async ({ newFirstName }: Omit<ChangeUserFirstName, "userId">) => await changeFirstName({
            userId: descriptor.id,
            newFirstName
        }),
        onError: (e) => console.log(e),
    })
    const handleUpdate = async () => {
        const { success, data, error } = validName.safeParse({
            firstName
        })
        if (!success) {
            setErrors(error.flatten())
            return
        }
        updateMutation.mutate({ newFirstName: data.firstName })
    }
    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="firstName">First name
                </label>
                <input
                    onChange={handleChange}
                    value={firstName}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="firstName" type="text" />
                {
                    errors?.fieldErrors.firstName &&
                    <InlineAlert
                        className="flex items-center p-4 mb-2 text-sm text-red-800 rounded-lg bg-red-50"
                        errors={errors.fieldErrors.firstName} />
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
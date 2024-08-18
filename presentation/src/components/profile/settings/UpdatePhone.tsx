import InlineAlert from "@/components/error/InlineAlert";
import CheckIcon from "@/components/icons/CheckIcon";
import { ChangeUserPhone, UserDescriptor } from "@/gql/graphql";
import { changeUserPhone } from "@/lib/actions/graphqlActions";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { z } from "zod";

const validatePhone = z.object({
    phone: z.string({ message: "We could not validate your phone number" })
        .min(8, { message: "We could not validate your phone number" })
        .transform((val, ctx) => {
            const parsed = parseInt(val);
            if (isNaN(parsed)) {
                ctx.addIssue({
                    path: ["phone"],
                    code: z.ZodIssueCode.custom,
                    message: "Not a number",
                });
                return z.NEVER;
            }
            return val;
        })
})

interface Props {
    descriptor: UserDescriptor
}

const UpdatePhone: React.FC<Props> = ({ descriptor }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof validatePhone> | undefined>()
    const [phone, setPhone] = useState<string>("")

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setPhone(e.target.value)
    }

    const updateMutation = useMutation({
        mutationFn: async ({ newPhone }: Omit<ChangeUserPhone, "userId">) => await changeUserPhone({
            userId: descriptor.id,
            newPhone
        }),
        onError: (e) => console.log(e),
    })

    const handleUpdate = async () => {
        const { success, data, error } = validatePhone.safeParse({
            phone
        })
        if (!success) {
            setErrors(error.flatten())
            return
        }
        updateMutation.mutate({ newPhone: data.phone })
    }

    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="phone">Phone
                </label>

                <input
                    onChange={handleChange}
                    value={phone}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="phone" type="text" />
                    
                {
                    errors?.fieldErrors.phone &&
                    <InlineAlert
                        className="flex items-center p-4 mb-2 text-sm text-red-800 rounded-lg bg-red-50"
                        errors={errors.fieldErrors.phone} />
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

export default UpdatePhone;
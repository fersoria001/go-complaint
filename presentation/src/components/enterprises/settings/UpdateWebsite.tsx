'use client'
import InlineAlert from "@/components/error/InlineAlert"
import CheckIcon from "@/components/icons/CheckIcon"
import { ChangeEnterpriseWebsite, Enterprise } from "@/gql/graphql"
import { changeEnterpriseWebsite } from "@/lib/actions/graphqlActions"
import { useMutation } from "@tanstack/react-query"
import { useState } from "react"
import { z } from "zod"

const validWebsite = z.object({
    website: z.string().url({
        message: "Please enter a valid website e.g: http://www.mywebsite.com",
    })
})
interface Props {
    enterprise: Enterprise
}
const UpdateWebsite: React.FC<Props> = ({ enterprise }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof validWebsite> | undefined>()
    const [website, setWebsite] = useState<string>("")
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors(undefined)
        setWebsite(e.target.value)
    }
    const updateMutation = useMutation({
        mutationFn: async ({ newWebsite }: Omit<ChangeEnterpriseWebsite, "enterpriseId">) => await changeEnterpriseWebsite({
            enterpriseId: enterprise.id,
            newWebsite
        }),
        onError: (e) => console.log(e),
    })
    const handleUpdate = async () => {
        const { success, data, error } = validWebsite.safeParse({
            website
        })
        if (!success) {
            setErrors(error.flatten())
            return
        }
        updateMutation.mutate({ newWebsite: data.website })
    }
    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto ">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="website">Website</label>
                <input
                    onChange={handleChange}
                    value={website}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="website" type="text" />
                {errors?.fieldErrors.website && <InlineAlert errors={errors.fieldErrors.website} />}
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

export default UpdateWebsite;
'use client'
import Link from "next/link"
import ExclamationIcon from "../icons/ExclamationIcon"
import { useSuspenseQuery } from "@tanstack/react-query"
import { useSearchParams } from "next/navigation"
import getGraphQLClient from "@/graphql/graphQLClient"
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery"
import { describeComplaint } from "@/lib/actions/graphqlActions"
import { useFormState } from "react-dom"
import InlineAlert from "../error/InlineAlert"

const DescribeComplaint: React.FC = () => {
    const params = useSearchParams()
    const id = params.get("id")
    const [state, formAction] = useFormState(describeComplaint, undefined)
    const { data } = useSuspenseQuery({
        queryKey: ['complaintById', id as string],
        queryFn: async ({ queryKey }) => {
            try {
                return await getGraphQLClient().request(complaintByIdQuery, { id: queryKey[1] })
            } catch (e: any) {
                console.log("error: ", e)
                return null
            }
        },
    })
    return (
        <form className="relative flex flex-col h-full" action={formAction}>
            <div className="mb-4">
                <label
                    htmlFor="input-group-1"
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2">
                    Reason
                </label>
                <input className="hidden" value={id!} name="complaintId" readOnly />
                <div className="relative mb-0.5">
                    <input
                        name="title"
                        type="text"
                        id="input-group-1"
                        defaultValue={data?.complaintById.title != "" ? data?.complaintById.title : undefined}
                        minLength={10}
                        maxLength={80}
                        className="bg-gray-50 border border-gray-300 text-gray-700 text-sm md:text-md rounded-lg appearance-none focus:outline-none block w-full ps-8 py-2.5"
                        placeholder="Why do you complain?" />
                    <div
                        className="absolute inset-y-0 start-0 flex items-center ps-1 pointer-events-none">
                        <ExclamationIcon fill="#3b82f6" width={24} height={24} />
                    </div>
                </div>
                {
                    state?.fieldErrors.title && <InlineAlert
                        className={"flex items-center text-sm text-red-800 rounded-lg bg-red-50"}
                        errors={state.fieldErrors.title} />
                }
            </div>
            <div className="mb-4">
                <label
                    htmlFor="description"
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                >Description
                </label>
                <textarea
                    id="description"
                    name="description"
                    rows={5}
                    defaultValue={data?.complaintById.description != "" ? data?.complaintById.description : undefined}
                    minLength={3}
                    maxLength={120}
                    className="mb-0.5 resize-none block p-2.5 w-full text-sm md:text-md appearance-none focus:outline-none text-gray-700 bg-gray-50 rounded-lg border border-gray-300"
                    placeholder="Shortly describe the problem here...">
                </textarea>
                {state?.fieldErrors.description && <InlineAlert
                    className={"flex items-center text-sm text-red-800 rounded-lg bg-red-50"}
                    errors={state.fieldErrors.description} />}
            </div>
            {state?.formErrors && <InlineAlert
                className={"flex items-center text-sm text-red-800 rounded-lg bg-red-50"}
                errors={state.formErrors} />}
            <div className="absolute bottom-28 md:bottom-2 w-full flex justify-around">
                <Link
                    href={"/complaints/send-complaint?step=1"}
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Back
                </Link>
                <button
                    type="submit"
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Next
                </button>
            </div>
        </form>
    )
}
export default DescribeComplaint
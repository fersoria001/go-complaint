'use client'
import getGraphQLClient from "@/graphql/graphQLClient"
import complaintByIdQuery from "@/graphql/queries/complaintByIdQuery"
import { sendComplaint } from "@/lib/actions/graphqlActions"
import { useSuspenseQuery } from "@tanstack/react-query"
import Link from "next/link"
import { useSearchParams } from "next/navigation"
import { useFormState } from "react-dom"
import InlineAlert from "../error/InlineAlert"

const Complain = () => {
    const params = useSearchParams()
    const id = params.get("id")
    const [state, formAction] = useFormState(sendComplaint, undefined)
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
        <form action={formAction} className="flex flex-col relative h-full">
            <label
                htmlFor="complain"
                className="block text-gray-700 text-sm md:text-md font-bold mb-2">
                Complain about it
            </label>
            <input className="hidden" name="complaintId" value={id!} readOnly />
            <textarea
                id="complaint"
                name="body"
                rows={8}
                minLength={50}
                maxLength={250}
                className="resize-none block p-2.5 w-full text-sm md:text-md appearance-none focus:outline-none text-gray-700 bg-gray-50 rounded-lg border border-gray-300"
                placeholder="Write your complaint here..."
                value={data?.complaintById.replies![0] ? data?.complaintById.replies[0].body! : undefined}
            >
            </textarea>
            {state?.fieldErrors.body && <InlineAlert errors={state.fieldErrors.body} />}
            {state?.formErrors && <InlineAlert errors={state.formErrors} />}
            <div className="absolute bottom-28 md:bottom-2 w-full flex justify-around">
                <Link
                    href={`/complaints/send-complaint?step=2&id=${id}`}
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Back
                </Link>
                <button
                    type="submit"
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Complain!
                </button>
            </div>
        </form>
    )
}
export default Complain
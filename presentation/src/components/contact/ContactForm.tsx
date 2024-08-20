'use client'

import contactSchema from "@/lib/validation/contactSchema";
import { useFormState } from "react-dom";
import { z } from "zod";
import InlineAlert from "../error/InlineAlert";
import contact from "@/lib/actions/graphqlActions";

type state = z.inferFlattenedErrors<typeof contactSchema>
const initialState: Partial<state> = {}
const ContactForm: React.FC = () => {
    const [state, formAction] = useFormState(contact, initialState)
    return (
        <form
            action={formAction}
            className="flex flex-col px-3">
            {state.formErrors && <InlineAlert errors={state.formErrors} />}
            <input
                className={`px-3 py-2 bg-white border 
                        focus:outline-none
                        text-gray-700
                        text-md md:text-xl
                         focus:border-gray-500 focus:ring-0 rounded-md mb-4`}
                name="email"
                placeholder="enter your email"
                autoComplete="username" />
            {state.fieldErrors?.email && <InlineAlert errors={state.fieldErrors.email} />}
            <textarea
                rows={4}
                className={`w-full text-gray-700 text-md md:text-xl mb-4
                         border rounded-md bg-white resize-none 
                        px-2 py-2 focus:outline-none focus:border-gray-500 focus:ring-0`}
                name="text"
                placeholder="write a message" />
            {state.fieldErrors?.text && <InlineAlert errors={state.fieldErrors.text} />}
            <button
                type="submit"
                className="bg-blue-500 text-white font-medium text-lg md:text-xl rounded-md self-center px-9 py-2 md:self-end hover:bg-blue-600">
                Send
            </button>
        </form>
    )
}
export default ContactForm;
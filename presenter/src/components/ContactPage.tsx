import { useState } from "react";
import { ContactMutation, Mutation } from "../lib/mutations";
import { contactSchema, ContactType, ErrorType } from "../lib/types";
import { syncParseSchema } from "../lib/parse_schema";

const ContactPage: React.FC = () => {
    const [email, setEmail] = useState("")
    const [text, setText] = useState("")
    const [errors, setErrors] = useState<ErrorType>({})
    const send = async (): Promise<boolean> => {
        const obj = { email, text }
        const { data, errors } = syncParseSchema(obj, contactSchema)
        if (Object.keys(errors).length > 0) {
            setErrors(errors)
        }
        return await Mutation<ContactType>(
            ContactMutation,
            data,
        )
    }
    const handleTextChange = (v: string) => {
        setErrors({})
        setText(v)
    }
    const handleEmailChange = (v: string) => {
        setErrors({})
        setEmail(v)
    }
    return (
        <div className="flex flex-col md:pt-10 md:px-12  px-3 md:items-end ">
            <div className="flex flex-col md:w-1/3">
                <div className="">
                    <img className="h-24 w-24 rounded-full mb-2" src="./contact.jpg" />
                    <p className="text-gray-700 text-sm md:text-xl mb-2 md:mb-4"> Fernando Agustín Soria </p>
                </div>
                <div>
                    <p className="text-gray-700 text-sm md:text-xl mb-2 md:mb-4 font-medium"> Web developer </p>
                    <div className="flex">
                        <p className="text-gray-700 text-sm md:text-xl mb-2 md:mb-4 font-medium"> Number:</p>
                        <p className="pl-1 text-gray-700 text-sm md:text-xl mb-2 md:mb-4 "> +54 2944 7818 23</p>
                    </div>
                    <div className="flex">
                        <p className="text-gray-700 text-sm md:text-xl mb-2 md:mb-4 font-medium"> Email:</p>
                        <p className="pl-1 text-gray-700 text-sm md:text-xl mb-2 md:mb-4 ">bercho001@gmail.com</p>
                    </div>
                </div>
            </div>
            <div className="flex flex-col md:w-1/3 px-3">
                <input
                    onChange={(e) => handleEmailChange(e.currentTarget.value)}
                    className={`
                        ${errors.email ? 'mb-0 md:mb-0' : 'mb-2 md:mb-4'}
                        px-3 py-2 bg-white border  focus:outline-none focus:border-gray-500 focus:ring-0 rounded-md`}
                    placeholder="email" />
                {errors?.email && <span className="text-red-500 text-xs italic" >{errors.email}</span>}
                <textarea
                    onChange={(e) => handleTextChange(e.currentTarget.value)}
                    className=
                    {`${errors.text ? 'mb-0 md:mb-0' : 'mb-2 md:mb-4'}
                     w-full  text-gray-600 text-sm md:text-xl border rounded-md bg-white resize-none h-[74px] md:h-[89px] px-2 py-2 
                    focus:outline-none focus:border-gray-500 focus:ring-0`} />
                {errors?.text && <span className="text-red-500 text-xs italic" >{errors.text}</span>}
                <button
                    onMouseUp={send}
                    type="button"
                    className="border border-black rounded-md  self-center px-5 py-2 md:self-end ">
                    Send
                </button>

            </div>
        </div>
    )
}

export default ContactPage;
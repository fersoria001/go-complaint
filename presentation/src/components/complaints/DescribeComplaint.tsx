'use client'
import Link from "next/link"
import ExclamationIcon from "../icons/ExclamationIcon"

const DescribeComplaint: React.FC = () => {
    return (
        <form className="relative flex flex-col h-full">
            <div className="mb-4">
                <label
                    htmlFor="input-group-1"
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2">
                    Reason
                </label>
                <div className="relative">
                    <input
                        type="text"
                        id="input-group-1"
                        minLength={10}
                        maxLength={80}
                        className="bg-gray-50 border border-gray-300 text-gray-700 text-sm md:text-md rounded-lg appearance-none focus:outline-none block w-full ps-8 py-2.5"
                        placeholder="Why do you complain?" />
                    <div
                        className="absolute inset-y-0 start-0 flex items-center ps-1 pointer-events-none">
                        <ExclamationIcon fill="#3b82f6" width={24} height={24} />
                    </div>
                </div>
            </div>
            <div className="mb-4">
                <label
                    htmlFor="description"
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                >Description
                </label>
                <textarea
                    id="description"
                    rows={5}
                    minLength={3}
                    maxLength={120}
                    className="resize-none block p-2.5 w-full text-sm md:text-md appearance-none focus:outline-none text-gray-700 bg-gray-50 rounded-lg border border-gray-300"
                    placeholder="Shortly describe the problem here...">
                </textarea>
            </div>
            <div className="absolute bottom-28 md:bottom-2 w-full flex justify-around">
                <Link
                    href={"/complaints/send-complaint?step=1"}
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Back
                </Link>
                <Link
                    href={"/complaints/send-complaint?step=3"}
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Next
                </Link>
            </div>
        </form>
    )
}
export default DescribeComplaint
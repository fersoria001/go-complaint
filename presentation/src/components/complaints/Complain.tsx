'use client'
import Link from "next/link"

const Complain = () => {
    return (
        <form className="flex flex-col relative h-full">
            <label
                htmlFor="complain"
                className="block text-gray-700 text-sm md:text-md font-bold mb-2">
                Complain about it
            </label>
            <textarea
                id="complaint"
                rows={8}
                minLength={50}
                maxLength={250}
                // onChange={(e) => setContent(e.target.value)}
                className="resize-none block p-2.5 w-full text-sm md:text-md appearance-none focus:outline-none text-gray-700 bg-gray-50 rounded-lg border border-gray-300"
            //</div>placeholder={`Complain to ${receiver.fullName} about ${complaint.title}`}
            >
            </textarea>
            {/* {errors.content && <span
                    className="self-center text-red-500 text-xs italic">
                    {errors.content} </span>} */}
            <div className="absolute bottom-28 md:bottom-2 w-full flex justify-around">
                <Link
                    href={"/complaints/send-complaint?step=2"}
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Back
                </Link>
                <button
                    type="button"
                    onClick={() => { }}
                    className="px-7 py-3 bg-blue-500 hover:bg-blue-600 font-bold text-white rounded-md"
                >
                    Complain!
                </button>
            </div>
        </form>
    )
}
export default Complain
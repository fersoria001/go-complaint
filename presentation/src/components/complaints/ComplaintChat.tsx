'use client'
import ComplaintInput from "./ComplaintInput";
import ComplaintReply from "./ComplaintReply";

const ComplaintChat: React.FC = () => {
    return (
        <div className="w-full flex flex-col border">
            <div className="flex w-full my-2 py-2.5">
                <div className="mx-2 rounded-full h-10 w-10 bg-gray-300 self-center"></div>
                <div className="px-2 self-center">
                    <h3 className="text-gray-700 font-bold text-md lg:text-lg xl:text-xl">Name</h3>
                    <p className="text-gray-700 text-sm xl:text-md">subject</p>
                </div>
                <div className="ml-auto mr-4 my-auto flex items-center gap-2.5">
                    <button
                        type="button"
                        className="text-white bg-blue-500 rounded-xl px-2.5 hover:bg-blue-600">
                        Mark for review
                    </button>
                    <div className="rounded-full h-2 w-2 bg-blue-300 "></div>

                </div>
            </div>
            <div className="p-2 h-[21.375rem] md:h-[21.525rem] xl:h-[23.175rem] flex flex-col gap-2.5 py-2.5">
                <ComplaintReply rightSided={false} />
                <ComplaintReply rightSided={true} />
            </div>
            <ComplaintInput />
        </div>
    )
}
export default ComplaintChat;
const ComplaintItem = () => {
    return (
        <div className="w-full flex flex-col border">
            <div className="flex w-full my-2 py-0.5">
                <div className="mx-2 rounded-full h-10 w-10 bg-gray-300 self-center"></div>
                <div className="px-2 self-center">
                    <h3 className="text-gray-700 font-bold text-md lg:text-lg xl:text-xl">Name</h3>
                    <p className="text-gray-700 text-sm xl:text-md">subject</p>
                </div>
                <div className="rounded-full h-2 w-2 bg-blue-300 ml-auto mr-4 my-auto"></div>
            </div>
            <div className="bg-gray-100 p-2">
                <p className="text-gray-700 text-md xl:text-lg">Response</p>
                <div className="flex justify-between">
                    <p className="text-gray-700 font-bold text-sm xl:text-md">in discussion</p>
                    <p className="text-gray-700 font-bold text-sm xl:text-md">00:00</p>
                </div>
            </div>
        </div>
    )
}
export default ComplaintItem
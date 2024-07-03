interface ComplaintCardProps {
    reason: string;
    description: string;
    body: string;
    status: string;
    createdAt: string;
}
function ComplaintCard({ reason, description, body, status, createdAt }: ComplaintCardProps) {
    return (
        <div className="flex flex-col">
            <div className="flex flex-col md:flex-row w-full justify-between mb-2">
                <h5 id="status" className="mx-auto md:mx-0 px-2 mb-2 text-lg font-bold tracking-tight text-gray-500 md:text-xl">{
                    new Date(parseInt(createdAt)).toLocaleDateString()
                }</h5>
                <div className="flex">
                    <label htmlFor="status" className="px-2 mb-2 text-lg font-bold tracking-tight text-gray-700 md:text-xl" >Status: </label>
                    <h5 id="status" className="px-2 mb-2 text-lg font-bold tracking-tight text-gray-500 md:text-xl">{status}</h5>
                </div>
            </div>
            <div className="flex flex-col md:flex-row align-center ">
                <label htmlFor="reason" className="mb-2 text-center text-lg pr-2 font-bold tracking-tight text-gray-700 md:text-xl">Reason: </label>
                <p id="reason" className="mb-3 text-lg text-gray-500 md:text-xl">
                    {reason}
                </p>
            </div>
            <div className="flex flex-col md:flex-row align-center ">
                <label htmlFor="description" className="mb-2 text-center pr-2 text-lg font-bold tracking-tight text-gray-700 md:text-xl">Complaint: </label>
                <p id="description" className="mb-3 text-lg text-gray-500 md:text-xl">
                    {description}
                </p>
            </div>
            <div className="flex flex-col align-center md:pl-24">
                <p className="mb-3 text-lg text-gray-500 md:text-xl">
                    {body}
                </p>
            </div>
        </div>
    );
}

export default ComplaintCard;
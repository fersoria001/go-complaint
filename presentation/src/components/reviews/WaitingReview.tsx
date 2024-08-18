import { Complaint } from "@/gql/graphql";
import { dateFromMsString } from "@/lib/dateFromMsString";

interface Props {
    item: Complaint
}
const WaitingReview: React.FC<Props> = ({ item }: Props) => {
    return (
        <form className="w-full border-t border-b">
            <div className="max-w-xl flex flex-col p-5 bg-white">
                <p className="text-gray-500 text-sm md:text-md xl:text-md font-medium text-end mb-2">
                    {dateFromMsString(item.rating?.createdAt!).toDateString()}
                </p>
                <p className="text-sm md:text-md text-gray-700 mb-4">
                    You asked {item.author?.subjectName} to review your attention at this complaint:
                </p>
                <div className="mx-4 border rounded-md p-2">
                    <div className="flex">
                        <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="reason">Reason:</label>
                        <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md">{item.title}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="description">Description:</label>
                        <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md">{item.description}</p>
                    </div>
                    <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md text-end">
                        {dateFromMsString(item.createdAt!).toDateString()}
                    </p>
                </div>
            </div>
        </form>
    )
}

export default WaitingReview;
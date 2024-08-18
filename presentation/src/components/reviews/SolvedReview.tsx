import { Complaint } from "@/gql/graphql";
import { dateFromMsString } from "@/lib/dateFromMsString";
import clsx from "clsx";
import Image from "next/image";
import StarIcon from "../icons/StarIcon";

interface Props {
    item: Complaint
}
const SolvedReview: React.FC<Props> = ({ item }: Props) => {
    let ratedBy = item.rating?.sentToReviewBy?.subjectName
    if (item.receiver!.isEnterprise) {
        ratedBy += " " + "from" + " " + item.receiver!.subjectName
    }
    return (
        <div className="border-t border-b py-2">
            <article className="md:gap-8 max-w-lg">
                <div>
                    <div className="flex items-center mb-6">
                        <div className='relative w-10 h-10'>
                            <Image
                                src={item.author?.subjectThumbnail!}
                                className="rounded-full"
                                sizes='32px'
                                fill
                                alt="user photo" />
                        </div>
                        <div className="ms-4 font-medium">
                            <p>{ratedBy}</p>
                        </div>
                    </div>
                    <ul className="space-y-4 text-sm text-gray-500 dark:text-gray-400">
                        <li className="flex">
                            <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="reason">Reason:</label>
                            <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md">{item.title}</p>
                        </li>
                        <li className="flex">
                            <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="description">Description:</label>
                            <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md">{item.description}</p>
                        </li>
                        <li className="flex">
                            <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="body">Complaint:</label>
                            <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md ">{item.replies![item.replies!.length - 1]!.body}</p>
                        </li>
                    </ul>
                </div>
                <div className="col-span-2 mt-6">
                    <div className="flex flex-col md:flex-row justify-between mb-5 items-center">
                        <div className="pe-4">
                            <p className="mb-2 text-sm text-gray-500 dark:text-gray-400">
                                Reviewed: {" "} {dateFromMsString(item.rating?.lastUpdate!).toDateString()}
                            </p>
                        </div>
                        <div className="flex py-5">
                            {[...Array(5)].map((_, index) => {
                                index += 1;
                                return <StarIcon key={index} className={clsx('w-6 h-6 md:w-8 md:h-8 fill-gray-200 cursor-pointer', {
                                    "fill-yellow-500": index <= item.rating?.rate!
                                })} />

                            })}
                        </div>
                    </div>
                    <p className="mb-2 text-gray-500 dark:text-gray-400">{item.rating?.comment}</p>
                </div>
            </article>
        </div>
    )
}

export default SolvedReview;
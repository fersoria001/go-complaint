import { Complaint } from "@/gql/graphql"
import { dateFromMsString } from "@/lib/dateFromMsString"
import clsx from "clsx"
import Image from "next/image"

interface Props {
    currentId: string
    item: Complaint
}

const ComplaintItem: React.FC<Props> = ({ currentId, item }: Props) => {
    const subject = currentId != item.author!.id ? item.author : item.receiver
    return (
        <div className="w-full flex flex-col border">
            <div className="flex w-full my-2 py-0.5">
                <div className='relative mx-2 rounded-full h-10 w-10 bg-gray-300 self-center'>
                    <Image
                        src={currentId != item.author?.id ? item.author?.subjectThumbnail! : item.receiver.subjectThumbnail!}
                        alt={currentId != item.author?.id ? item.author?.subjectName! : item.receiver.subjectName!}
                        className="rounded-full"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                    />
                </div>
                <div className="px-2 self-center">
                    <h3 className="text-gray-700 font-bold text-md lg:text-lg xl:text-xl">
                        {currentId != item.author?.id ? item.author?.subjectName : item.receiver.subjectName}
                    </h3>
                    <p className="text-gray-700 text-sm xl:text-md">
                        {item.title}
                    </p>
                </div>
                <div className={clsx("rounded-full h-2 w-2 ml-auto mr-4 my-auto", {
                    "bg-blue-300": subject?.isOnline == true,
                    "bg-red-300": subject?.isOnline == false
                })}></div>
            </div>
            <div className="bg-gray-100 p-2">
                <div className="mb-2">
                    <p className="text-gray-700 text-md font-medium">{item.replies[item.replies.length - 1].sender.subjectName}</p>
                    <p className="text-gray-700 text-sm ps-2">
                        {item.replies[item.replies.length - 1].body}
                    </p>
                    <p className="text-gray-700 text-xs xl:text-sm text-end pe-2">
                        {dateFromMsString(item.replies[item.replies.length - 1].createdAt).toLocaleTimeString().slice(0, 5)}
                    </p>
                </div>
                <div className="flex justify-between">
                    <p className="text-gray-700 font-bold text-sm xl:text-md">{item.status}</p>
                    <p className="text-gray-700 font-bold text-sm xl:text-md">
                        {dateFromMsString(item.createdAt).toDateString()}
                    </p>
                </div>
            </div>
        </div>
    )
}
export default ComplaintItem
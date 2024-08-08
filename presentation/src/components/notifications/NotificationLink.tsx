'use client'
import Image from "next/image";
import timeAgo from "../../lib/timeAgo";
import Link from "next/dist/client/link";
import clsx from "clsx";
import { Notification } from "@/gql/graphql";


interface Props {
    notification: Notification
}


const NotificationLink: React.FC<Props> = ({ notification }: Props) => {
    return (
        <Link href={notification.link} key={notification.id}
            id={notification.id}
            className={clsx(
                'flex flex-col w-full justify-between rounded-xl pb-2 hover:cursor-pointer',
                {
                    'bg-gray-100': notification.seen,
                    'hover:bg-gray-100': !notification.seen,
                }
            )}>
            <div className="flex py-3">
                <div className="relative h-10 w-10 mx-3">
                    <Image
                        className="rounded-full"
                        src={notification.sender.subjectThumbnail!}
                        alt="notification"
                        fill />
                </div>
                <div>
                    <p className="text-xs sm:text-sm md:text-md text-gray-700">{notification.title}</p>
                    <p className="text-xs sm:text-sm md:text-md text-gray-700">{notification.content}</p>
                </div>
            </div>
            <p className="text-end text-xs sm:text-sm md:text-md text-gray-700 px-3">{timeAgo(notification.occurredOn)}</p>
        </Link>

    )
}

export default NotificationLink;
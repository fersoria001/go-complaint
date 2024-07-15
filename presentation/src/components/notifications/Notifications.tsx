'use client'
import { useRef, useState } from "react";
import NotificationIcon from "../icons/NotificationIcon";
import Notification from "./Notification";
import useClickOutside from "../../lib/hooks/useClickOutside";
import clsx from 'clsx';
import NotificationType from "../../lib/types/notificationType";
interface Props {
    notifications: NotificationType[]
}
const Notifications: React.FC<Props> = ({ notifications }: Props) => {
    const [show, setShow] = useState<boolean>(false)
    const notificationsRef = useRef<HTMLDivElement>(null)
    const unread = notifications.filter((n) => !n.seen).length;
    useClickOutside(notificationsRef, () => { setShow(false) })
    return (
        <div ref={notificationsRef} className="z-20">
            <button
                onClick={() => setShow(!show)}
                type="button"
                className="relative inline-flex items-center p-3 text-sm font-medium text-center text-white rounded-lg">
                <NotificationIcon />
                <span className="sr-only">Notifications</span>
                <div
                    className={clsx(`absolute inline-flex items-center justify-center w-6 h-6 text-xs font-bold text-white bg-red-500 border-2 border-white rounded-full top-1 end-0.5`,
                        {
                            'hidden': unread <= 0,
                            '': unread > 0,
                        })}>
                    {unread}
                </div>
            </button>
            {show &&
                <div
                    className="absolute w-full sm:max-w-md top-[82px] border-r border-l border-b overflow-y-auto rounded-md right-0 bg-white">
                    <ul className=" text-gray-500 list-none">
                        {
                            notifications.length > 0 ? notifications.map(n =>
                                <li key={n.id} className="first:pt-4 last:pb-0">
                                    <Notification notification={n} />
                                </li>
                            ) :
                                <li className="p-4">
                                    <p className="text-xs text-gray-700">No notifications yet.</p>
                                </li>
                        }
                    </ul>
                </div>
            }
        </div>
    )
}
export default Notifications;
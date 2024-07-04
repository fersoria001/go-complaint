/* eslint-disable @typescript-eslint/no-explicit-any */
import { useCallback, useEffect, useRef, useState } from "react";
import { Route } from "../../routes/__root";
import useOutsideDenier from "../../lib/hooks/useOutsideDenier";
import Notification from "./Notification";
import useSubscriber from "../../lib/hooks/useSubscriber";
import { Notifications } from "../../lib/types";
import NotificationIcon from "../icons/NotificationIcon";


function NotificationList() {
    const { notifications, id } = Route.useLoaderData()
    const [showNotifications, setShowNotifications] = useState(false)
    const [notificationList, setNotificationList] = useState<Notifications[]>(notifications)
    const [unread, setUnread] = useState(
        notifications ? notifications.filter(notification => notification.seen === false).length : 0
    )
    const ref = useRef<HTMLDivElement>(null)
    const { incomingMessage } = useSubscriber(id)
    const toggleNotifications = useCallback(() => {
        setShowNotifications(!showNotifications)
    }, [showNotifications]
    )
    useEffect(() => {
        if (incomingMessage) {
            const notification = incomingMessage as Notifications
            setNotificationList(prev => {
                const exists = prev.find(notif => notif.id === notification.id)
                if (exists) {
                    return prev.map(notif => notif.id === notification.id ? notification : notif)
                }
                return [...prev, notification]
            })
            setUnread(prev => notification.seen ? prev - 1 : prev + 1)
        }
    }, [incomingMessage])
    useOutsideDenier(ref, () => setShowNotifications(false))

    return (
        <div>
            <button
                onClick={() => setShowNotifications(!showNotifications)}
                type="button" className="relative inline-flex items-center p-3 text-sm font-medium text-center text-white rounded-lg">
                <NotificationIcon fill="#5f6368" />
                <span className="sr-only">Notifications</span>
                <div className={`${unread <= 0 ? 'hidden' : ''} absolute inline-flex items-center justify-center w-6 h-6 text-xs font-bold text-white bg-red-500 border-2 border-white rounded-full top-1 end-0.5`}>{unread}</div>
            </button>


            {showNotifications &&
                <div ref={ref} className="absolute top-20 p-2 border border-solid overflow-y-auto max-h-[220px]
                 border-cyan-500 rounded-md shadow right-0 bg-white">
                    <ul className=" max-w-md space-y-1 text-gray-500 list-none">
                        {
                            notificationList.length > 0 ? notificationList.map(notification =>
                                <li key={notification.id}>
                                    <Notification notification={notification} callback={toggleNotifications} />
                                </li>
                            ) :
                                <li className="p-2">
                                    <p>No notifications yet</p>
                                </li>
                        }
                    </ul>
                </div>
            }
        </div>
    )
}



export default NotificationList;

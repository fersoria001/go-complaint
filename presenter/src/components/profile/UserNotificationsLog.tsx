/* eslint-disable @typescript-eslint/no-explicit-any */
import { Link } from "react-router-dom";
import { UserNotifications, UserNotificationType } from "../../lib/types";
import { useEffect, useMemo, useRef, useState } from "react";
import NotificationIcon from "../icons/NotificationIcon";
import useOutsideDenier from "../../lib/hooks/useOutsideDenier";
interface Props {
    notifications: UserNotifications
}

function UserNotificationsLog({ notifications }: Props) {
    const [unread, setUnread] = useState(0);
    const [showNotifications, setShowNotifications] = useState(false);
    const ref = useRef<HTMLDivElement>(null);
    useOutsideDenier(ref, () => setShowNotifications(false));
    const timeAgo = (date: string): string => {
        const obj = new Date(parseInt(date))
        const now = new Date()
        const diff = now.getTime() - obj.getTime()
        const seconds = Math.floor(diff / 1000)
        let result = 0
        if (seconds < 3600) {
            result = Math.floor(seconds / 60)
            return `${result}m ago`
        }
        result = Math.floor(seconds / 3600)
        if (result > 24) {
            return `${Math.floor(result / 24)}d ago`
        }
        return `${result}h ago`
    }
    const renderNotification = (notification: any, type: UserNotificationType) => (
        <li
            onClick={() => { setShowNotifications(false) }}
            key={notification.id || notification.event_id} // Use a unique identifier for the key
            className={notification.seen ?
                `flex justify-between first:pt-0 last:mb-0 bg-gray-200 rounded-xl p-2 mb-2 hover:cursor-pointer`
                : `flex justify-between first:pt-0 last:mb-0 rounded-xl p-2 mb-2 hover:bg-gray-100 hover:cursor-pointer`}>
            <Link to={
                type === "hiring_invitation" ?
                    `/invitation/${type}/${notification.id}` :
                    `/profile/reviews`
            }>
                {type === "waiting_for_review" ? (
                    <p>You've been asked to review your complaint attention.</p>
                ) : (
                    <p>You have been invited to be part of {notification.enterprise_id}!</p>
                )}
                <p>{timeAgo(notification.occurred_on)}</p>
            </Link>
        </li>
    );
    const processedNotifications = useMemo(() => {
        const commonNotificationsMap: any[] = [];
        notifications.waiting_for_review.forEach(notification => commonNotificationsMap.push({ type: "waiting_for_review", notification }));
        notifications.hiring_invitation.forEach(notification => commonNotificationsMap.push({
            type: "hiring_invitation",
            notification
        }));
        commonNotificationsMap.sort((a, b) => {
            return parseInt(b.notification.occurred_on) - parseInt(a.notification.occurred_on)
        });
        return commonNotificationsMap;
    }, [notifications]);

    useEffect(() => {
        let unreads = 0;
        processedNotifications.forEach(notification => {
            if (!notification.notification.seen) unreads++;
        });
        setUnread(unreads);
    }, [processedNotifications]);



    return (
        <>
            <span
                onClick={() => setShowNotifications(!showNotifications)}
                className="relative  cursor-pointer rounded-full active:bg-gray-100">
                <span className="">
                    <NotificationIcon fill="#374151" />
                </span>
                {unread > 0 && <span className="absolute border border-white p-2
                flex items-center justify-center text-xs text-white font-bold
                 bottom-0 right-0 z-100 bg-red-700 rounded-full h-4 w-4">
                    <p className="text-white">{unread}</p>
                </span>}
            </span >
            {showNotifications &&
                <div ref={ref} className="absolute top-14 p-2 border border-solid border-cyan-500 rounded-md shadow right-0 bg-white">
                    <ul className="divide-y divide-gray-200 max-w-md space-y-1 text-gray-500 list-none">
                        {
                            processedNotifications.length > 0 ?
                                processedNotifications.map(({ type, notification }) => renderNotification(notification, type))
                                :
                                <li className="p-2">
                                    <p className="text-gray-500">No new notifications</p>
                                </li>
                        }
                    </ul>
                </div>
            }
        </>
    );
}



export default UserNotificationsLog;

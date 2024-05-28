import { Link } from "react-router-dom";
import { UserNotifications } from "../../lib/types";
interface Props {
    notifications: UserNotifications;
}

function Notifications({ notifications }: Props) {
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
        return `${result}h ago`
    }
    return (
        <div className="absolute top-14 p-2 border border-solid border-cyan-500 rounded-md shadow right-0 bg-white">
            <ul className="divide-y divide-gray-200 max-w-md space-y-1  text-gray-500 list-none">
                {
                    notifications.hiring_invitation.map((notification, index) => {
                        return (
                            <li key={index} className={notification.seen ?
                                `
                            flex justify-between
                            first:pt-0 last:mb-0
                            bg-gray-200
                            rounded-xl p-2
                            mb-2  hover:cursor-pointer`
                                : `
                            flex justify-between
                            first:pt-0 last:mb-0
                            rounded-xl p-2
                            mb-2 hover:bg-gray-100 hover:cursor-pointer`}>
                                <Link to={`/invitation/hiring/${notification.event_id}`} >
                                    <p>You have been invited to be part of {notification.enterprise_id}!</p>
                                    <p>{timeAgo(notification.occurred_on)}</p>
                                </Link>
                            </li>
                        );
                    })
                }
            </ul>
        </div>
    );
}



export default Notifications;
import { Link } from "@tanstack/react-router"
import { Notifications } from "../../lib/types"
import { Mutation, MarkNotificationAsRead } from "../../lib/mutations"
import { timeAgo } from "../../lib/time_ago"
interface Props {
    notification: Notifications
    callback: () => void
}
async function markAsRead(id: string) {
    const success = await Mutation<string>(MarkNotificationAsRead, id)
    return success
}
const Notification: React.FC<Props> = ({ notification, callback }: Props) => {
    const handleHover = async () => {
        if (!notification.seen) {
            const ok = await markAsRead(notification.id)
            console.log("markAsRead",ok)
        }
    }
    const style = notification.seen ? `flex justify-between first:pt-0 last:mb-0 bg-gray-200 rounded-xl p-2 mb-2 hover:cursor-pointer` :
        `flex justify-between first:pt-0 last:mb-0 rounded-xl p-2 mb-2 hover:bg-gray-100 hover:cursor-pointer`
        return (
            <div
                onMouseEnter={handleHover}
                onClick={callback}
                key={notification.id}
                id={notification.id}
                className={style}>
                <Link to={notification.link}>
                    <div className="flex">
                        <img
                            className="h-10 w-10 rounded-full shrink-0"
                            src={notification.thumbnail}
                            alt="notification" />
                        <div>
                            <p>{notification.title}</p>
                            <p>{notification.content}</p>
                        </div>
                    </div>
                    <p>{timeAgo(notification.occurredOn)}</p>
                </Link>
            </div>
        )
    }

    export default Notification
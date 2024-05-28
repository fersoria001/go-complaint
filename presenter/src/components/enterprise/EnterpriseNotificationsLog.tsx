import { Link } from "react-router-dom";
import { EnterpriseNotifications } from "../../lib/types";
interface Props {
    enterpriseName: string;
    notifications: EnterpriseNotifications;
}

function EnterpriseNotificationsLog({ enterpriseName, notifications }: Props) {
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
    const pending = (notifications: EnterpriseNotifications) => {
        const notSeen = notifications.employee_waiting_for_approval.filter(notification => !notification.seen)
        return { count: notSeen.length, occurred_on: timeAgo(notSeen[0].occurred_on) }
    }
    const { count, occurred_on } = pending(notifications)
return (
    <div className="absolute top-14 p-2 border border-solid border-cyan-500 rounded-md shadow right-0 bg-white">
        <ul className="divide-y divide-gray-200 max-w-md space-y-1  text-gray-500 list-none">
            {
                count > 0 &&
                <li key={"pendingsHiringInvitations"} className={`
                            flex justify-between
                            first:pt-0 last:mb-0
                            rounded-xl p-2
                            mb-2 hover:bg-gray-100 hover:cursor-pointer`}>
                    <Link to={`/enterprises/${enterpriseName}/pending`} >
                        <p>You have pending job invitations waiting for approval</p>
                        <p>{occurred_on}</p>
                    </Link>
                </li>
            }
        </ul>
    </div>
);
}



export default EnterpriseNotificationsLog;
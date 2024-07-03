import { Link } from "@tanstack/react-router";

interface Props {
    selected: string;//'complaint' | 'chat';
    complaintLink: string;
    chatLink: string;
}
function ComplaintTabs({ selected, complaintLink, chatLink } : Props) {
    let complaintClass
    let chatClass
    if (selected === 'complaint') {
        complaintClass = `inline-block w-full p-4 rounded-ss-lg bg-gray-100 outline-none pointer-events-none cursor-none`
        chatClass = 'inline-block w-full p-4 bg-gray-50 hover:bg-gray-100 focus:outline-none'
    } else if (selected === 'chat') {
        complaintClass = 'inline-block w-full p-4 rounded-ss-lg bg-gray-50 hover:bg-gray-100 focus:outline-none'
        chatClass = 'inline-block w-full p-4 bg-gray-50 bg-gray-100 outline-none pointer-events-none cursor-none'
    } else {
        complaintClass = 'inline-block w-full p-4 bg-gray-50 hover:bg-gray-100 focus:outline-none'
        chatClass = 'inline-block w-full p-4 bg-gray-50 hover:bg-gray-100 cursor-none pointer-events-none'
    }
    return (
        <ul className="
        text-sm font-medium text-center text-gray-500
        md:divide-x divide-y divide-gray-200 rounded-lg sm:flex rtl:divide-x-reverse">
            <li className="w-full">
                <Link
                    to={complaintLink}
                    type="button"
                    className={complaintClass}>
                    Complaint
                </Link>
            </li>
            <li className="w-full">
                <Link
                    to={chatLink}
                    type="button"
                    className={chatClass}>
                    Chat
                </Link>
            </li>
        </ul>
    )
}

export default ComplaintTabs;
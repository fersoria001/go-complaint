import clsx from "clsx";

interface Props {
    rightSided: boolean
}
const ComplaintReply = ({ rightSided }: Props) => {
    return (
        <div className={clsx("bg-gray-50 border w-1/2 p-3 rounded-xl", {
            'self-end': rightSided
        })}>
            <p className="text-gray-700 text-sm font-bold">Name LastName</p>
            <p className="text-gray-700 text-sm">Response</p>
            <p className="text-gray-700 font-bold text-xs text-end mr-2">00:00</p>
        </div>
    )
}
export default ComplaintReply;
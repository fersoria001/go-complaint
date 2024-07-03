import ExclamationIcon from "../icons/ExclamationIcon"
interface Props {
    callback: (input: string) => void;
}
function Reason({ callback }: Props) {
    return (
        <div>
            <label
                htmlFor="input-group-1"
                className="block mb-2 text-sm md:text-xl font-medium text-gray-900">
                Reason
            </label>
            <div className="relative">
                <input
                    onChange={(e) => callback(e.target.value)}
                    type="text"
                    id="input-group-1"
                    minLength={10}
                    maxLength={80}
                    className="bg-gray-50 border border-gray-300 text-gray-900 
                    text-sm md:text-xl rounded-lg focus:ring-blue-500
                     focus:border-blue-500 
                    block w-full ps-6 p-2.5"
                    placeholder="Why do you complain?"
                />
                <div
                    className="absolute inset-y-0 start-0 flex items-center
                 ps-1 pointer-events-none">
                    <ExclamationIcon fill="#06b6d4" />
                </div>
            </div>
        </div>
    )
}

export default Reason
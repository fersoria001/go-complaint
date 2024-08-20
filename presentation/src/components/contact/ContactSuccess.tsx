import CheckIcon from "../icons/CheckIcon";

const ContactSuccess: React.FC = () => {
    return (
        <div className="flex flex-col mt-12 px-2">
            <div className="flex flex-col">
                <div className="flex items-center mb-4">
                    <CheckIcon className="w-12 h-12 fill-blue-300" />
                    <p className="text-gray-700 text-md font-medium">Success!</p>
                </div>
                <article>
                    <p className="text-gray-700 text-md mb-2 md:mb-4 font-medium">
                        I will receive your email soon, you may or may not receive another email in response to this message.
                    </p>
                </article>
            </div>
        </div>
    )
}
export default ContactSuccess;
import Link from "next/link";

const PasswordRecoverySucceed: React.FC = () => {
    return (
        <div className="mx-auto bg-white shadow rounded-md 
        border max-w-md self-center py-8 mt-8 flex flex-col justify-center">
            <h3 className="self-center text-gray-700 font-bold pb-5 px-2 text-md md:text-lg">
                We have sent you an email with the next steps
                to recover your password.
            </h3>
            <p className="text-gray-700 mb-4 px-2 text-sm md:text-md">
                There&apos;s many reason you may not receive the email,
                if you didn&apos;t receive it within the next two hours consider:
            </p>
            <ul className="list-disc list-inside px-4 pb-4 space-y-1 text-gray-700 text-sm md:text-md">
                <li>
                    Looking for the email in the spam box, often
                    custom email filters can take it as a spam.
                </li>
                <li>
                    Try the option to recover your password from
                    the website again. Maybe the email sending
                    was delayed from external reasons and trying to trigger it
                    again could solve the problem.
                </li>
                <li>
                    If none of the above works contact Go Complaint
                    at{" "}
                    <a
                        className="font-bold"
                        href="emailto:owner@go-complaint.com">
                        owner@go-complaint.com
                    </a>.
                </li>
            </ul>
            <Link href="/sign-in" className="self-center underline text-gray-700 font-medium text-sm md:text-md pt-2.5">
                Sign-in to Go Complaint
            </Link>
        </div>
    )
}
export default PasswordRecoverySucceed;
import { Link } from "@tanstack/react-router";
import { Route } from "../routes/$enterpriseID/success";


function SuccessPage() {
    const { message, link, to } = Route.useLoaderData();
    return (
        <div className="flex flex-col rounded-md w-full shadow border mb-[190px] md:mb-[371px]">
            <p className="text-sm md:text-xl text-gray-700 p-4">
                {message}
            </p>
            <span className="self-center hover:underline hover:underline-offset-8 text-sm md:text-xl text-gray-700 p-4">
                <Link
                    to={link}
                >Return to {to}</Link>
            </span>
        </div>
    )
}
export default SuccessPage;
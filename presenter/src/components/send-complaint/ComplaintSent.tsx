import { Link } from "@tanstack/react-router";

function ComplaintSent() {
    return (
        <div className="flex flex-col rounded-md w-full shadow border mb-[190px] md:mb-[371px]">
            <p className="text-sm md:text-xl text-gray-700 p-4">
                Your complaint has been sent succesfully
            </p>
            <span className="self-center hover:underline hover:underline-offset-8 text-sm md:text-xl text-gray-700 p-4">
                <Link
                    to="/profile"
                >Return to profile</Link>
            </span>
        </div>
    )
}
export default ComplaintSent;
import { Link } from "@tanstack/react-router";

function ConfirmationSucceed() {
    return (
        <div className="h-screen bg-white flex flex-col">
            <div className="bg-white shadow rounded-md border h-1/2 w-1/2 self-center p-5 mt-12 flex flex-col justify-center">
                <h1 className="self-center text-2xl font-bold pb-5">Your email has been validated with success!</h1>
                <Link to="/sign-in" className="self-center text-gray-600">Sign-in to Go Complaint</Link>
            </div>
        </div>
    )
}

export default ConfirmationSucceed;
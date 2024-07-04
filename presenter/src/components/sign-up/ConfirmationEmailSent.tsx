import { Link } from "@tanstack/react-router";

function ConfirmationEmailSent() {
    <div className="h-screen bg-white flex flex-col">
        <div className="bg-white shadow rounded-md border h-1/2 w-1/2 self-center p-5 mt-12 flex flex-col justify-center">
            <h1 className="self-center text-2xl font-bold pb-5">Your have been register to Go Complaint!</h1>
            <p className="text-gray-700 text-md md:text-xl mb-4">
                The email must be verified before
                sign in. We sent you a confirmation a confirmation link.
            </p>
            <Link to="/sign-in" className="self-center text-gray-600">Sign-in to Go Complaint</Link>
        </div>
    </div>
}

export default ConfirmationEmailSent;
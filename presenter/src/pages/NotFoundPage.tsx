import { Link } from "@tanstack/react-router";

function NotFoundPage() {
    return (
        <div className="h-screen bg-white flex flex-col">
            <div className="bg-white shadow rounded-md border h-1/2 w-1/2 self-center p-5 mt-12 flex flex-col justify-center">
                <h1 className="self-center text-2xl font-bold text-gray-700">404 - Not Found</h1>
                <Link to="/" className="self-center text-blue-600">Go back to Home</Link>
            </div>
        </div>
    );
}

export default NotFoundPage;
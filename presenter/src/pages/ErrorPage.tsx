import { Link } from "@tanstack/react-router";


export function ErrorPage() {
  return (
    <div className="h-screen bg-white flex flex-col">
      <div className="bg-white shadow rounded-md border h-1/2 w-1/2 self-center p-5 mt-12 flex flex-col justify-center">
        <h1 className="self-center text-2xl font-bold text-red-600">We encountered an error</h1>
        <p className="self-center text-gray-500">We are working  hard to solve it</p>
        <Link to="/" className="self-center text-gray-600">Go back to Home</Link>
      </div>
    </div>
  );



}
export default ErrorPage;
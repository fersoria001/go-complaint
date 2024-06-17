import { useLoaderData } from "react-router-dom";
import PendingEmployee from "../components/enterprise/PendingEmployee";
import { User } from "../lib/types";

function PendingHires() {
    const pending = useLoaderData() as { eventID: string; employee: User, position: string, pendingDate: string }[] | null;
    return (
        <div className="">
            <h1
                className="text-3xl font-bold text-center p-4 text-gray-900 tracking-tight sm:text-4xl"
            >These employees are waiting for your approval</h1>
            <div>
                {pending ? pending.map((hire) => {
                    return (
                        <PendingEmployee
                            employee={hire.employee}
                            eventID={hire.eventID}
                            position={hire.position}
                            pendingDate={hire.pendingDate}
                            key={hire.eventID} />
                    )
                }) : <div className="flex justify-center items-center h-screen">
                    <p className="text-center text-gray-700">
                        Currently you have not any pending hiring proccess with any other user.
                    </p>
                </div>}
            </div>
        </div>
    );
}
export default PendingHires;
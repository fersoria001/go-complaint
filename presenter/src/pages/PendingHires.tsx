import { useLoaderData } from "react-router-dom";
import { Employee } from "../lib/types";
import PendingEmployee from "../components/enterprise/PendingEmployee";

function PendingHires() {
    const pending = useLoaderData() as { eventID: string; employee: Employee }[] | null;
    return (
        <div className="">
            <h1
                className="text-3xl font-bold text-center p-4 text-gray-900 tracking-tight sm:text-4xl"
            >These employees are waiting for your approval</h1>
            <div>
                {pending && pending.map((hire) => {
                    return (
                        <PendingEmployee employee={hire.employee} eventID={hire.eventID} key={hire.eventID} />
                    )
                })}
            </div>
        </div>
    );
}
export default PendingHires;
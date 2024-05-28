import { useLoaderData } from "react-router-dom";
import { Employee } from "../lib/types";
import EmployeeCard from "../components/employee/EmployeeCard";

function Employees() {
    const employees = useLoaderData() as Employee[];
    if (!employees) return <div>Loading...</div>;
    return (
        <div>
            {
                employees.map((employee) => (
                    <li key={employee.ID}>
                        <EmployeeCard employee={employee} />
                    </li>
                ))
            }
        </div>
    );
}

export default Employees;
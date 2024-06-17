import { useLoaderData, useParams } from "react-router-dom";
import { Employee } from "../lib/types";
import EmployeeCard from "../components/employee/EmployeeCard";

function Employees() {
    const employees = useLoaderData() as Employee[];
    const {id} = useParams();
    if (!employees) return <div>Loading...</div>;
    return (
        <ul>
            {
                employees.map((employee) => (
                    <li key={employee.ID}>
                        <EmployeeCard employee={employee} enterpriseID={id!}/>
                    </li>
                ))
            }
        </ul>
    );
}

export default Employees;
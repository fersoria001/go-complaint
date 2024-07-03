import { EmployeeType } from "../../../lib/types";
import { Route } from "../../../routes/$enterpriseID/employees";
import Employee from "./Employee";


function Employees() {
    const employees = Route.useLoaderData()
    const { enterpriseID } = Route.useParams()
    return (
        <div className="min-h-[315px] md:min-h-[460px]">
            {
                employees.map((employee: EmployeeType) => (
                    <div key={employee.id}>
                        <Employee employee={employee} enterpriseId={enterpriseID} />
                    </div>
                ))
            }
        </div>
    )
}

export default Employees;
import { createFileRoute, redirect } from '@tanstack/react-router'
import { EmployeesQuery, EmployeesTypeList, Query } from '../../../lib/queries'
import { EmployeeType } from '../../../lib/types'
import Employees from '../../../components/enterprise/employee/Employees'

export const Route = createFileRoute('/$enterpriseID/employees/')({
    beforeLoad: ({ context: { isLoggedIn } }) => {
        if (!isLoggedIn) {
            throw redirect({
                to: '/sign-in',
                search: {
                    redirect: location.href,
                },
            })
        }
    },
    loader: async ({ params, context: { fetchUserDescriptor } }) => {
        const descriptor = await fetchUserDescriptor()
        const result = await Query<EmployeeType[]>(
            EmployeesQuery,
            EmployeesTypeList,
            [params.enterpriseID]
        )
        const employees = result.filter((employee) => employee.email !== descriptor.email)
        return employees
    },
    component: Employees,
})
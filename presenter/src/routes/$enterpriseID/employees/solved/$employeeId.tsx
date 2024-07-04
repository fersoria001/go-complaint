import { createFileRoute, redirect } from '@tanstack/react-router'
import { ComplaintQuery, ComplaintQueryType, EmployeeQuery, EmployeeQueryType, Query } from '../../../../lib/queries'
import { ComplaintType, EmployeeType } from '../../../../lib/types'
import SolvedList from '../../../../components/enterprise/feedback/SolvedList'

export const Route = createFileRoute('/$enterpriseID/employees/solved/$employeeId')({
  beforeLoad: async ({ params: { enterpriseID }, context: { hasPermission, isLoggedIn } }) => {
    if (!isLoggedIn) {
      throw redirect({
        to: '/sign-in',
        search: {
          redirect: location.href,
        },
      })
    }
    const ok = await hasPermission("MANAGER", enterpriseID)
    if (!ok) {
      throw redirect({
        to: '/',
        search: {
          redirect: location.href,
        },
      })
    }
  },
  loader: async ({ params: { employeeId, enterpriseID } }) => {
    const employee = await Query<EmployeeType>(
      EmployeeQuery,
      EmployeeQueryType,
      [enterpriseID, employeeId]
    )
    const promises = []
    for (const solvedId of employee.complaintsSolvedIds) {
      promises.push(Query<ComplaintType>(
        ComplaintQuery,
        ComplaintQueryType,
        [solvedId]
      ))
    }
    const complaints = await Promise.all(promises)
    return complaints
  },
  component: SolvedList,
})
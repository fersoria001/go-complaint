import { createFileRoute, notFound } from '@tanstack/react-router'
import EnterpriseLayout from '../EnterpriseLayout'
import { ComplaintInfo, Enterprise, User } from '../lib/types'
import { Query, EnterpriseType, CompleteEnterpriseQuery, ComplaintsReceivedInfoQuery, ComplaintReceivedInfoType, OnlineUsersQuery, OnlineUsersType } from '../lib/queries'

async function fetchOnlineUsers(enterpriseID: string): Promise<User[]> {
  try {
    const onlineUsers = await Query<User[]>(OnlineUsersQuery, OnlineUsersType, [enterpriseID])
    return onlineUsers
  } catch (error) {
    return []
  }
}

export const Route = createFileRoute('/$enterpriseID')({
  loader: async ({ params, context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const onlineUsers : User[] = await fetchOnlineUsers(params.enterpriseID)
    const authorities = descriptor!.grantedAuthorities
    let enterprise
    let role
    for (const authority of authorities) {
      if (authority.enterpriseID === params.enterpriseID) {
        enterprise = await Query<Enterprise>(
          CompleteEnterpriseQuery,
          EnterpriseType, [authority.enterpriseID])
        role = authority.authority
      }
    }
    const complaintsInfo = await Query<ComplaintInfo>(ComplaintsReceivedInfoQuery,
      ComplaintReceivedInfoType, [enterprise!.name]
    )
    const employee = enterprise!.employees!.find(employee => employee.email === descriptor!.email)
    if (!enterprise) {
      throw notFound()
    }
    return {
      employee,
      onlineUsers,
      complaintsInfo,
      descriptor,
      role,
      enterprise
    }
  },
  component: EnterpriseLayout
})
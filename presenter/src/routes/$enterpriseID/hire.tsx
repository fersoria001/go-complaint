import { createFileRoute } from '@tanstack/react-router'
import { Enterprise, User } from '../../lib/types'
import { EnterpriseQuery, EnterpriseType, Query, UserQuery, UserType } from '../../lib/queries'
import Hire from '../../pages/Hire'

type UserFind = {
  id: EmailID,
}
type EmailID = { email: string }

export const Route = createFileRoute('/$enterpriseID/hire')({
  validateSearch: (search: Record<string, unknown>): UserFind => {
    return {
      id: search.id as EmailID,
    }
  },
  loaderDeps: ({ search: { id } }) => ({ id }),
  loader: async ({ params, deps: { id } }) => {
    const user = await Query<User>(
      UserQuery,
      UserType,
      [id.email]
    )
    const enterprise = await Query<Enterprise>(
      EnterpriseQuery,
      EnterpriseType,
      [params.enterpriseID]
    )
    return {user, enterprise}
  },
  component: Hire,
})
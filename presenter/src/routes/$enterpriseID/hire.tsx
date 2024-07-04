import { createFileRoute, redirect } from '@tanstack/react-router'
import { Enterprise, User } from '../../lib/types'
import { EnterpriseQuery, EnterpriseType, Query, UserQuery, UserType } from '../../lib/queries'
import Hire from '../../pages/Hire'

type UserFind = {
  id: EmailID,
}
type EmailID = { email: string }

export const Route = createFileRoute('/$enterpriseID/hire')({
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
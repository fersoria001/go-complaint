import { createFileRoute, redirect } from '@tanstack/react-router'
import Hiring from '../../pages/Hiring'
import { Query, UsersForHiringQuery, UsersForHiringType } from '../../lib/queries'
import { UsersForHiring } from '../../lib/types'

type UsersForHiringSearchType = {
  filter: UsersForHiringFilterType
}
type UsersForHiringFilterType = {
  query: string
  page: number
}
export const Route = createFileRoute('/$enterpriseID/hiring')({
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
  validateSearch: (search: Record<string, unknown>): UsersForHiringSearchType => {
    return {
      filter: search.filter as UsersForHiringFilterType || { query: "", page: 1 }
    }
  },
  loaderDeps: ({ search: { filter } }) => ({ filter }),
  loader: async ({ params: { enterpriseID }, deps: { filter } }) => {
    let usersForHiring: UsersForHiring = {
      users: [],
      count: 0,
      currentLimit: 10,
      currentOffset: 0,
    }
    const offset = (filter.page - 1) * 10
    const limit = 10
    usersForHiring = await Query<UsersForHiring>(
      UsersForHiringQuery,
      UsersForHiringType,
      [enterpriseID, limit, offset, filter.query])
    const search = { filter }
    return { usersForHiring, search }
  },
  component: Hiring,
})
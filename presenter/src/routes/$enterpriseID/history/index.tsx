import { createFileRoute, redirect } from '@tanstack/react-router'
import { daysAgoFilter } from '../../../lib/days_ago_filter'
import { Query, ComplaintHistoryQuery, ComplaintHistoryTypeList } from '../../../lib/queries'
import { ComplaintTypeList } from '../../../lib/types'
import { ComplaintSearchType, ComplaintSearchFilterType } from '../../_profile/sent'
import History from '../../../components/enterprise/history/History'

export const Route = createFileRoute('/$enterpriseID/history/')({
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
  validateSearch: (search: Record<string, unknown>): ComplaintSearchType => {
    return {
      page: search.page ? Number(search.page) : 1,
      filter: search.filter as ComplaintSearchFilterType || { query: '', date: 'Last year' }
    }
  },
  loaderDeps: ({ search: { page, filter } }) => ({ page, filter }),
  loader: async ({ params: { enterpriseID }, deps: { page, filter } }) => {
    let sent: ComplaintTypeList = { count: 0, currentOffset: 0, currentLimit: 0, complaints: [] }
      const offset = (page - 1) * 10
      const limit = 10
      const [after, before] = daysAgoFilter(filter.date)
      sent = await Query<ComplaintTypeList>(
        ComplaintHistoryQuery,
        ComplaintHistoryTypeList,
        [
          enterpriseID,
          filter.query,
          after,
          before,
          limit,
          offset,
        ])
    const search = { page, filter }
    return {  sent, search }
  },
  component: History,
})
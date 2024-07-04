import { createFileRoute, redirect } from '@tanstack/react-router'
import { ComplaintHistoryQuery, ComplaintHistoryTypeList, Query } from '../../../lib/queries'
import { daysAgoFilter } from '../../../lib/days_ago_filter'
import { ComplaintTypeList } from '../../../lib/types'
import { ComplaintSearchType, ComplaintSearchFilterType } from '../sent'
import History from '../../../components/profile/history/History'

export const Route = createFileRoute('/_profile/history/')({
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
  loader: async ({ context: { fetchUserDescriptor }, deps: { page, filter } }) => {
    const descriptor = await fetchUserDescriptor()
    let sent: ComplaintTypeList = { count: 0, currentOffset: 0, currentLimit: 0, complaints: [] }
    if (descriptor) {
      const offset = (page - 1) * 10
      const limit = 10
      const [after, before] = daysAgoFilter(filter.date)
      sent = await Query<ComplaintTypeList>(
        ComplaintHistoryQuery,
        ComplaintHistoryTypeList,
        [
          descriptor.email,
          filter.query,
          after,
          before,
          limit,
          offset,
        ])
    }
    const search = { page, filter }
    return { descriptor, sent, search }
  },
  component: History,
})
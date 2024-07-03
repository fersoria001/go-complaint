import { createFileRoute } from '@tanstack/react-router'
import Inbox from '../../../components/profile/inbox/Inbox'
import { Query, SearchInDraftQuery, SearchInDraftTypeList } from '../../../lib/queries'
import { ComplaintTypeList } from '../../../lib/types'
import { daysAgoFilter } from '../../../lib/days_ago_filter'
import { ComplaintSearchType, ComplaintSearchFilterType } from '../sent'

export const Route = createFileRoute('/_profile/inbox/')({
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
        SearchInDraftQuery,
        SearchInDraftTypeList,
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
  component: Inbox,
})
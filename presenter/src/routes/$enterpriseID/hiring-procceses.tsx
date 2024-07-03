import { createFileRoute } from '@tanstack/react-router'
import { HiringProccessesQuery, HiringProccessesTypeList, Query } from '../../lib/queries'
import { HiringProccessList } from '../../lib/types'
import HiringProcceses from '../../pages/HiringProcceses'

type HiringProccessSearchType = {
  id: string
  filter: HiringProccessSearchFilter
}
type HiringProccessSearchFilter = {
  query: string
  page: number
}
export const Route = createFileRoute('/$enterpriseID/hiring-procceses')({
  validateSearch: (search: Record<string, unknown>): HiringProccessSearchType => {
    return {
      id: search.id as string || '',
      filter: search.filter as HiringProccessSearchFilter || { query: '', page: 1 }
    }
  },
  loaderDeps: ({ search: { id, filter } }) => ({ id, filter }),
  loader: async ({ params, deps: { id, filter } }) => {
    const offset = (filter.page - 1) * 10
    const limit = 5
    const hiringProcceses = await Query<HiringProccessList>(
      HiringProccessesQuery,
      HiringProccessesTypeList,
      [
        params.enterpriseID,
        filter.query,
        offset,
        limit
      ]
    )
    const search = { id, filter }
    if (id && id != '') {
      hiringProcceses.hiringProccesses = hiringProcceses.hiringProccesses.filter(
        (hiringProccess) => {
          return hiringProccess.eventID === id
        })
    }
    return { hiringProcceses, search }
  },
  component: HiringProcceses,
})
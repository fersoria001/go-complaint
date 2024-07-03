import { createFileRoute } from '@tanstack/react-router'
import { Query, ComplaintReviews, ComplaintReviewsTypeList } from '../../../lib/queries'
import { ComplaintReviewType } from '../../../lib/types'
import Reviews from '../../../components/enterprise/Reviews'

export const Route = createFileRoute('/$enterpriseID/reviews/')({
  loader: async ({ params, context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const complaintReviews = await Query<ComplaintReviewType[]>(
      ComplaintReviews,
      ComplaintReviewsTypeList,
      [params.enterpriseID],
    )
    return { descriptor, complaintReviews }
  },
  component: Reviews,
})
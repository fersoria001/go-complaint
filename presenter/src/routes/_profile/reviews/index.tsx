import { createFileRoute, redirect } from '@tanstack/react-router'
import { ComplaintReviews, ComplaintReviewsTypeList, Query } from '../../../lib/queries'
import { ComplaintReviewType } from '../../../lib/types'
import Reviews from '../../../components/profile/Reviews'

export const Route = createFileRoute('/_profile/reviews/')({
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
  loader: async ({ context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const complaintReviews = await Query<ComplaintReviewType[]>(
      ComplaintReviews,
      ComplaintReviewsTypeList,
      [descriptor.email],
    )
    return { descriptor, complaintReviews }
  },
  component: Reviews,
})
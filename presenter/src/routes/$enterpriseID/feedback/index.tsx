/* eslint-disable @typescript-eslint/no-explicit-any */
import { createFileRoute, redirect } from '@tanstack/react-router'
import { FeedbackByRevieweeIDQuery, FeedbackByRevieweeIDType, Query } from '../../../lib/queries'
import { FeedbackType } from '../../../lib/types'
import FeedbacksDone from '../../../components/enterprise/feedback/FeedbacksDone'

export const Route = createFileRoute('/$enterpriseID/feedback/')({
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
    try {
      const descriptor = await fetchUserDescriptor()
      const feedbacks = await Query<FeedbackType[]>(
        FeedbackByRevieweeIDQuery,
        FeedbackByRevieweeIDType,
        [descriptor.email],
      )
      console.log(feedbacks)
      return feedbacks
    } catch (e: any) {
      console.error(e)
      return []
    }

  },
  component: FeedbacksDone
})
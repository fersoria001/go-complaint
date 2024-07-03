import { createFileRoute } from '@tanstack/react-router'
import { ComplaintQuery, ComplaintQueryType, FeedbackByComplaintIDQuery, FeedbackByComplaintIDType, Query, UserQuery, UserType } from '../../../../lib/queries'
import { ComplaintType, FeedbackType, User } from '../../../../lib/types'
import Feedback from '../../../../components/enterprise/feedback/Feedback'

type ComplaintId = {
  complaintId: string
}
export const Route = createFileRoute('/$enterpriseID/feedback/$employeeId/')({
  validateSearch: (search: Record<string, unknown>): ComplaintId => {
    return {
      complaintId: search.complaintId as string
    }
  },
  loaderDeps: ({ search: { complaintId } }) => ({ complaintId }),
  loader: async ({ deps: { complaintId }, context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const reviewer = await Query<User>(UserQuery, UserType, [descriptor.email])
    const complaint = await Query<ComplaintType>(
      ComplaintQuery,
      ComplaintQueryType,
      [complaintId]
    )
    const feedback = await Query<FeedbackType>(
      FeedbackByComplaintIDQuery,
      FeedbackByComplaintIDType,
      [complaintId]
    )
    let subscriptionID = null
    if (feedback) {
      subscriptionID = `feedback:${feedback.id}`
    }
    return { subscriptionID, reviewer, complaint, feedback }
  },
  component: Feedback,
})
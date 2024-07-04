import { createFileRoute, redirect } from '@tanstack/react-router'
import { Query, ComplaintQuery, ComplaintQueryType } from '../../../../lib/queries'
import { createSubscription, ComplaintLastReplySubscription, ComplaintLastReplyReturnType } from '../../../../lib/subscriptions'
import { ComplaintType, Reply } from '../../../../lib/types'
import Complaint from '../../../../components/enterprise/inbox/Complaint'

export const Route = createFileRoute('/$enterpriseID/inbox/$complaintId/')({
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
  loader: async ({ params, context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const complaint = await Query<ComplaintType>(
      ComplaintQuery,
      ComplaintQueryType,
      [params.complaintId],
    )
    const subscription = createSubscription<Reply>(
      `complaintLastReply:${complaint.id}`,
      ComplaintLastReplySubscription,
      [complaint.id],
      ComplaintLastReplyReturnType,
    )
    return { descriptor, complaint, subscription }
  },
  component: Complaint,
})
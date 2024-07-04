import { createFileRoute, redirect } from '@tanstack/react-router'
import Complaint from '../../../../components/enterprise/sent/Complaint'
import { Query, ComplaintQuery, ComplaintQueryType } from '../../../../lib/queries'
import { createSubscription, ComplaintLastReplySubscription, ComplaintLastReplyReturnType } from '../../../../lib/subscriptions'
import { ComplaintType, Reply } from '../../../../lib/types'

export const Route = createFileRoute('/$enterpriseID/sent/$complaintId/')({
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
  component: Complaint
})
import { createFileRoute } from '@tanstack/react-router'
import Chat from '../../../../components/enterprise/inbox/Chat'
import { Query, ComplaintQuery, ComplaintQueryType } from '../../../../lib/queries'
import { ComplaintType } from '../../../../lib/types'

export const Route = createFileRoute('/$enterpriseID/inbox/$complaintId/chat')({
  loader: async ({ params, context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const complaint = await Query<ComplaintType>(
      ComplaintQuery,
      ComplaintQueryType,
      [params.complaintId],
    )
    let id = null
    if (complaint) {
      id = `complaintLastReply:${complaint.id}`
    }
    return { descriptor, complaint, id }
  },
  component: Chat,
})
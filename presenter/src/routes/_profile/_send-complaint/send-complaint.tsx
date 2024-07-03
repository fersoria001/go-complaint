import { createFileRoute } from '@tanstack/react-router'
import { FindComplaintReceiversQuery, FindComplaintReceiversTypeList, Query } from '../../../lib/queries'
import { Receiver } from '../../../lib/types'
import FindReceiver from '../../../components/send-complaint/FindReceiver'


export const Route = createFileRoute('/_profile/_send-complaint/send-complaint')({
  loader: async ({ context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const receivers = await Query<Receiver[]>(FindComplaintReceiversQuery,
      FindComplaintReceiversTypeList, [descriptor.email, ""])
    return { descriptor, receivers }
  },
  component: FindReceiver
})
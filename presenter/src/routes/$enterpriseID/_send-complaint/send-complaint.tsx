import { createFileRoute } from '@tanstack/react-router'
import { FindComplaintReceiversQuery, FindComplaintReceiversTypeList, Query } from '../../../lib/queries'
import { Receiver } from '../../../lib/types'
import FindReceiver from '../../../components/enterprise/send-complaint/FindReceiver'


export const Route = createFileRoute('/$enterpriseID/_send-complaint/send-complaint')({
  loader: async ({ params: { enterpriseID } }) => {
    const receivers = await Query<Receiver[]>(FindComplaintReceiversQuery,
      FindComplaintReceiversTypeList, [enterpriseID, ""])
    return receivers
  },
  component: () => <FindReceiver />
})
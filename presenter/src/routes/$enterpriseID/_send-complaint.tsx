import { createFileRoute, redirect } from '@tanstack/react-router'
import SendComplaint from '../../pages/SendComplaint'
import { Query, FindComplaintReceiversQuery, FindComplaintReceiversTypeList } from '../../lib/queries';
import { Receiver } from '../../lib/types';


export const Route = createFileRoute('/$enterpriseID/_send-complaint')({
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
  loader: async ({params:{enterpriseID}}) => {
    const receivers = await Query<Receiver[]>(
      FindComplaintReceiversQuery,
      FindComplaintReceiversTypeList,
      [enterpriseID, ""]
    );
    return receivers
  },
  component: SendComplaint
})
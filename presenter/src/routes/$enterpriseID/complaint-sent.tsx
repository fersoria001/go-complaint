import { createFileRoute, redirect } from '@tanstack/react-router'
import ComplaintSent from '../../components/enterprise/send-complaint/ComplaintSent'

export const Route = createFileRoute('/$enterpriseID/complaint-sent')({
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
  component: ComplaintSent,
})
import { createFileRoute, redirect } from '@tanstack/react-router'
import ComplaintSent from '../../components/send-complaint/ComplaintSent'

export const Route = createFileRoute('/_profile/complaint-sent')({
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
  component: ComplaintSent
})
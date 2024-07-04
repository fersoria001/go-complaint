import { createFileRoute, redirect } from '@tanstack/react-router'
import SendComplaint from '../../pages/SendComplaint'


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
  component: () => <SendComplaint />
})
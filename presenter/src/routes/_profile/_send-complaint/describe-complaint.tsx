import { createFileRoute, redirect } from '@tanstack/react-router'
import DescribeComplaint from '../../../components/send-complaint/DescribeComplaint'

export const Route = createFileRoute('/_profile/_send-complaint/describe-complaint')({
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
  component: DescribeComplaint,
})
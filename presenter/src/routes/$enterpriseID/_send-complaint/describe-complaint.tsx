import { createFileRoute, redirect } from '@tanstack/react-router'
import DescribeComplaint from '../../../components/enterprise/send-complaint/DescribeComplaint'


export const Route = createFileRoute('/$enterpriseID/_send-complaint/describe-complaint')({
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
  component: () => <DescribeComplaint />
})
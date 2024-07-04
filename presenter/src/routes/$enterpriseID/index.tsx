import { createFileRoute, redirect } from '@tanstack/react-router'
import EnterprisePage from '../../pages/EnterprisePage'

export const Route = createFileRoute('/$enterpriseID/')({
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
  component: EnterprisePage
})
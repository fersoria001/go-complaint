import { createFileRoute, redirect } from '@tanstack/react-router'
import SignIn from '../pages/SignIn'

export const Route = createFileRoute('/sign-in')({
  beforeLoad: ({ context: { isLoggedIn } }) => {
    if (isLoggedIn()) {
      throw redirect({
        to: '/profile',
        search: {
          redirect: location.href,
        },
      })
    }
  },
  component: SignIn,
})
import { createFileRoute, redirect } from '@tanstack/react-router'
import Profile from '../../components/profile/Profile'
export const Route = createFileRoute('/_profile/profile')({
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
  component: () => <Profile />
})
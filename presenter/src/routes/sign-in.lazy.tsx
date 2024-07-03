import { createLazyFileRoute } from '@tanstack/react-router'
import SignIn from '../pages/SignIn'

export const Route = createLazyFileRoute('/sign-in')({
  component: () => <SignIn />
})
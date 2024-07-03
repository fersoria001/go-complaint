import { createFileRoute } from '@tanstack/react-router'
import Profile from '../../components/profile/Profile'
export const Route = createFileRoute('/_profile/profile')({
  component: () => <Profile />
})
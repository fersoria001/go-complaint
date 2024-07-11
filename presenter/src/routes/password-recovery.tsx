import { createFileRoute } from '@tanstack/react-router'
import PasswordRecovery from '../components/sign-in/PasswordRecovery'

export const Route = createFileRoute('/password-recovery')({
  component: PasswordRecovery,
})
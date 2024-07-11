import { createFileRoute } from '@tanstack/react-router'
import PasswordRecoverySucceed from '../components/sign-in/PasswordRecoverySucceed'

export const Route = createFileRoute('/recovery-succeed')({
  component: PasswordRecoverySucceed
})
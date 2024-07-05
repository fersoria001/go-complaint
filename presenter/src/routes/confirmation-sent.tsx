import { createFileRoute } from '@tanstack/react-router'
import ConfirmationEmailSent from '../components/sign-up/ConfirmationEmailSent'

export const Route = createFileRoute('/confirmation-sent')({
  component: ConfirmationEmailSent
})
import { createFileRoute } from '@tanstack/react-router'
import LoginForm from '../components/sign-in/LoginForm'

export const Route = createFileRoute('/confirmation')({
  component: LoginForm,
})
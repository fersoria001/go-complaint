import { createLazyFileRoute } from '@tanstack/react-router'
import PrivacyPolicyPage from '../components/PrivacyPolicyPage'

export const Route = createLazyFileRoute('/privacy')({
  component: PrivacyPolicyPage
})
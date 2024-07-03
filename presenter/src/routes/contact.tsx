import { createFileRoute } from '@tanstack/react-router'
import ContactPage from '../components/ContactPage'

export const Route = createFileRoute('/contact')({
  component: ContactPage
})
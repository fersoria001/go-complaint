import { createFileRoute } from '@tanstack/react-router'
import LicensingPage from '../components/LicensingPage'

export const Route = createFileRoute('/licensing')({
  component: LicensingPage
})
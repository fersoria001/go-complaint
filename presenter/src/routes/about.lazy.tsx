import { createLazyFileRoute } from '@tanstack/react-router'
import AboutPage from '../components/AboutPage'

export const Route = createLazyFileRoute('/about')({
  component: AboutPage,
})


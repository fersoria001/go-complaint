import { createFileRoute } from '@tanstack/react-router'
import SendComplaint from '../../pages/SendComplaint'

export const Route = createFileRoute('/_profile/_send-complaint')({
  component: () => <SendComplaint />
})
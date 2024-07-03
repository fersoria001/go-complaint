import { createFileRoute } from '@tanstack/react-router'
import SendComplaint from '../../pages/SendComplaint'


export const Route = createFileRoute('/$enterpriseID/_send-complaint')({
  component: () => <SendComplaint />
})
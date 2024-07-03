import { createFileRoute } from '@tanstack/react-router'
import ComplaintSent from '../../components/enterprise/send-complaint/ComplaintSent'

export const Route = createFileRoute('/$enterpriseID/complaint-sent')({
  component: ComplaintSent,
})
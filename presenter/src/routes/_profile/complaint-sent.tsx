import { createFileRoute } from '@tanstack/react-router'
import ComplaintSent from '../../components/send-complaint/ComplaintSent'

export const Route = createFileRoute('/_profile/complaint-sent')({
  component: ComplaintSent
})
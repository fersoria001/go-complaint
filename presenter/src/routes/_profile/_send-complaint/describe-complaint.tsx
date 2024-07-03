import { createFileRoute } from '@tanstack/react-router'
import DescribeComplaint from '../../../components/send-complaint/DescribeComplaint'

export const Route = createFileRoute('/_profile/_send-complaint/describe-complaint')({
  component: () => <DescribeComplaint />
})
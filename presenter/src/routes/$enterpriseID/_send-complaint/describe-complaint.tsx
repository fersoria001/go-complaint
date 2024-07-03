import { createFileRoute } from '@tanstack/react-router'
import DescribeComplaint from '../../../components/enterprise/send-complaint/DescribeComplaint'


export const Route = createFileRoute('/$enterpriseID/_send-complaint/describe-complaint')({
  component: () => <DescribeComplaint />
})
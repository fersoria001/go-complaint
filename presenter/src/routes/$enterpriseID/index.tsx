import { createFileRoute } from '@tanstack/react-router'
import EnterprisePage from '../../pages/EnterprisePage'

export const Route = createFileRoute('/$enterpriseID/')({
  component: EnterprisePage
})
/* eslint-disable @typescript-eslint/no-explicit-any */
import { createFileRoute, Navigate } from '@tanstack/react-router'
import Complain from '../../../components/enterprise/send-complaint/Complain'


export const Route = createFileRoute('/$enterpriseID/_send-complaint/complain')({
  component: () => <Complain />,
  errorComponent: () => Navigate({ to: `/errors` }),
})

/* eslint-disable @typescript-eslint/no-explicit-any */
import { createFileRoute, Navigate } from '@tanstack/react-router'
import Complain from '../../../components/send-complaint/Complain'

export const Route = createFileRoute('/_profile/_send-complaint/complain')({
  component: () => <Complain />,
  errorComponent: () => Navigate({ to: '/send-complaint' }),
})

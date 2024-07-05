/* eslint-disable @typescript-eslint/no-explicit-any */
import { createFileRoute, Navigate, redirect } from '@tanstack/react-router'
import Complain from '../../../components/enterprise/send-complaint/Complain'


export const Route = createFileRoute('/$enterpriseID/_send-complaint/complain')({
  beforeLoad: ({ context: { isLoggedIn } }) => {
    if (!isLoggedIn) {
      throw redirect({
        to: '/sign-in',
        search: {
          redirect: location.href,
        },
      })
    }
  },
  component: Complain,
  errorComponent: () => Navigate({ to: `/errors` }),
})

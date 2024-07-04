/* eslint-disable @typescript-eslint/no-explicit-any */
import { createFileRoute, redirect } from '@tanstack/react-router'
import { Mutation, VerifyEmailMutation } from '../lib/mutations'
import ConfirmationSucceed from '../components/ConfirmationSucceed'
export type ConfirmationLinkSearchType = {
  token: string
  success: boolean
}
export const Route = createFileRoute('/confirmation-link')({
  validateSearch: (search: Record<string, unknown>): ConfirmationLinkSearchType => {
    return {
      token: search.token as string | "",
      success: search.success as boolean | false,
    } as ConfirmationLinkSearchType
  },
  loaderDeps: ({ search: { token, success } }) => ({ token, success }),
  loader: async ({ deps: { token, success } }) => {
    if (success) return
    if (!token || token == "") return redirect({ to: "/" })
    try {
      await Mutation<string>(VerifyEmailMutation, token)
      return redirect({ to: "/confirmation-link", search: { success: true } })
    } catch (e : any ) {
      return redirect({ to: "/error" })
    }
  },
  component: ConfirmationSucceed,
})
import { createFileRoute, redirect } from '@tanstack/react-router'
import SuccessPage from '../../pages/SuccessPage'



type SuccessLoaderData = {
  content: Success
}

type Success = {
  message: string,
  link: string,
  to: string,
}

export const Route = createFileRoute('/$enterpriseID/success')({
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
  validateSearch: (search: Record<string, unknown>): SuccessLoaderData => {
    return {
      content: search.content as Success
    }
  },
  loaderDeps: ({ search: { content } }) => ({ content }),
  loader: async ({ deps: { content } }) => {
    return { message : content.message, link: content.link, to: content.to}
  },
  component: SuccessPage,
})
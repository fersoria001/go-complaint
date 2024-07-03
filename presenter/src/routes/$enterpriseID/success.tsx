import { createFileRoute } from '@tanstack/react-router'
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
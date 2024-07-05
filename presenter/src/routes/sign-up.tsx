import { createFileRoute, redirect } from '@tanstack/react-router'
import SignUp from '../pages/SignUp'
import { CountriesQuery, CountryListType, Query } from '../lib/queries'
import { Country } from '../lib/types'
export const Route = createFileRoute('/sign-up')({
  beforeLoad: ({ context: { isLoggedIn } }) => {
    if (isLoggedIn()) {
      throw redirect({
        to: '/profile',
        search: {
          redirect: location.href,
        },
      })
    }
  },
  loader: async () => Query<Country[]>(CountriesQuery, CountryListType, []),
  component: () => <SignUp />,
})
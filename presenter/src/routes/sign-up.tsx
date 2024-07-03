import { createFileRoute } from '@tanstack/react-router'
import SignUp from '../pages/SignUp'
import { CountriesQuery, CountryListType, Query } from '../lib/queries'
import { Country } from '../lib/types'
export const Route = createFileRoute('/sign-up')({
  loader: async () => Query<Country[]>(CountriesQuery, CountryListType, []),
  component: () => <SignUp />,
})
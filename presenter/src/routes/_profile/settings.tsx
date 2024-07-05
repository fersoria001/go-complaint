import { createFileRoute, redirect } from '@tanstack/react-router'
import { CountriesQuery, CountryListType, Query } from '../../lib/queries'
import { Country } from '../../lib/types'
import Settings from '../../components/profile/settings/Settings'

export const Route = createFileRoute('/_profile/settings')({
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
  loader: async ({ context: { fetchUserDescriptor } }) => {
    const descriptor = await fetchUserDescriptor()
    const countries = await Query<Country[]>(CountriesQuery,
      CountryListType,
      []
    )
    return { descriptor, countries }
  },
  component: Settings
})
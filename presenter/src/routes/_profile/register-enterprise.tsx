import { createFileRoute, redirect } from '@tanstack/react-router'
import { Query, CountriesQuery, CountryListType, IndustriesQuery, IndustryListType } from '../../lib/queries'
import { Country, Industry } from '../../lib/types'
import RegisterEnterprise from '../../components/enterprise/RegisterEnterprise'

export const Route = createFileRoute('/_profile/register-enterprise')({
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
  loader: async () => {
    const countries = await Query<Country[]>(CountriesQuery, CountryListType, [])
    const industries = await Query<Industry[]>(IndustriesQuery, IndustryListType, [])
    return { countries, industries }
  },
  component: RegisterEnterprise,
  onCatch: (e) => {
    console.error(e)
  }
})
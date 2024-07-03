import { createFileRoute } from '@tanstack/react-router'
import { Query, CountriesQuery, CountryListType, IndustriesQuery, IndustryListType } from '../../lib/queries'
import { Country, Industry } from '../../lib/types'
import RegisterEnterprise from '../../components/enterprise/RegisterEnterprise'

export const Route = createFileRoute('/_profile/register-enterprise')({
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
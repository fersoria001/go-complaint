import { createFileRoute } from '@tanstack/react-router'
import Settings from '../../components/enterprise/settings/Settings'
import { Query, CountriesQuery, CountryListType, EnterpriseQuery, EnterpriseType } from '../../lib/queries'
import { Country, Enterprise } from '../../lib/types'

export const Route = createFileRoute('/$enterpriseID/settings')({
  loader: async ({ params: { enterpriseID } }) => {
    const countries = await Query<Country[]>(CountriesQuery,
      CountryListType,
      []
    )
    const enterprise = await Query<Enterprise>(
      EnterpriseQuery,
      EnterpriseType,
      [enterpriseID]
    )
    return { enterprise, countries }
  },
  component: Settings,
})
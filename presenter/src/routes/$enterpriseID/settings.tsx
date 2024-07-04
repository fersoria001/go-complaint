import { createFileRoute, redirect } from '@tanstack/react-router'
import Settings from '../../components/enterprise/settings/Settings'
import { Query, CountriesQuery, CountryListType, EnterpriseQuery, EnterpriseType } from '../../lib/queries'
import { Country, Enterprise } from '../../lib/types'

export const Route = createFileRoute('/$enterpriseID/settings')({
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
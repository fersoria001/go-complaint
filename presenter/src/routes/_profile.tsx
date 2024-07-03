import { createFileRoute, redirect } from '@tanstack/react-router'
import ProfileLayout from '../ProfileLayout'
import { Query, EnterpriseQuery, EnterpriseType } from '../lib/queries'
import { Enterprise } from '../lib/types'

export const Route = createFileRoute('/_profile')({
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
    const authorities = descriptor.grantedAuthorities
    const ownerEnterprisesP = []
    const employeeEnterprisesP = []
    let ownerEnterprises = []
    let employeeEnterprises = []
    for (const authority of authorities) {
      if (authority.authority === 'OWNER') {
        ownerEnterprisesP.push(
          Query<Enterprise>(EnterpriseQuery, EnterpriseType, [authority.enterpriseID])
        )
      }
      if (authority.authority === 'MANAGER' ||
        authority.authority === 'ASSISTANT') {
        employeeEnterprisesP.push(Query<Enterprise>(EnterpriseQuery, EnterpriseType, [authority.enterpriseID]))
      }
    }
    try {
      ownerEnterprises = await Promise.all(ownerEnterprisesP)
      employeeEnterprises = await Promise.all(employeeEnterprisesP)
      const enterprises = ownerEnterprises.every((enterprise) => enterprise !== null) &&
        employeeEnterprises.every((enterprise) => enterprise !== null)
      if (!enterprises) {
        throw new Error('Error fetching enterprises')
      }
      return {
        descriptor,
        ownerEnterprises,
        employeeEnterprises
      }
    } catch (error) {
      return {
        descriptor: descriptor,
        ownerEnterprises: [],
        employeeEnterprises: []
      }
    }
  },
  component: ProfileLayout,
})
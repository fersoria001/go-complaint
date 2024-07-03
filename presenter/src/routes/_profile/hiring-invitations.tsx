import { createFileRoute } from '@tanstack/react-router'
import { HiringInvitationsQuery, HiringInvitationsTypeList, Query } from '../../lib/queries'
import { HiringInvitationType } from '../../lib/types'
import HiringInvitations from '../../pages/HiringInvitations'

export const Route = createFileRoute('/_profile/hiring-invitations')({
  loader: async () => {
    const invitations = await Query<HiringInvitationType[]>(
      HiringInvitationsQuery,
      HiringInvitationsTypeList,
      [],
    )
    return { invitations }
  },
  component: HiringInvitations,
})
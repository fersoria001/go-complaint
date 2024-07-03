import HiringInvitation from "../components/hiring/HiringInvitation"
import { HiringInvitationType } from "../lib/types"
import { Route } from "../routes/_profile/hiring-invitations"

function HiringInvitations() {
    const { invitations }  = Route.useLoaderData()
    return (
        <div className="min-h-[315px] md:min-h-[460px]">
            {
                invitations.map((invitation: HiringInvitationType) => (
                    <div key={invitation.eventID}>
                    <HiringInvitation invitation={invitation} />
                    </div>
                ))
            }
        </div>
    )
}

export default HiringInvitations;
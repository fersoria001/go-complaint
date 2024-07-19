import { graphql } from "@/gql";

const hiringInvitationsByAuthenticatedUserQuery = graphql(`
    query hiringInvitationsByAuthenticatedUserQuery{
        hiringInvitationsByAuthenticatedUser{
            eventId
            enterpriseId
            proposedPosition
            ownerId
            fullName
            enterpriseEmail
            enterprisePhone
            enterpriseLogoImg
            occurredOn
            seen
            status
            reason
        }
    }`)
export default hiringInvitationsByAuthenticatedUserQuery
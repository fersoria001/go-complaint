import { graphql } from "@/gql";


const enterprisesByAuthenticatedUserQuery = graphql(`
query EnterprisesByAuthenticatedUserQuery{
    enterprisesByAuthenticatedUser{
        enterprises{
            authority{
            authority
            enterpriseId
        }
        enterprise{
            name
            logoImg
            bannerImg
            website
            email
            phoneNumber
            address {
                country
                countryState
                city
                }
            industry
            foundationDate
            ownerId
        }
        }
        offices{
            authority{
            authority
            enterpriseId
        }
        enterprise{
            name
            logoImg
            bannerImg
            website
            email
            phoneNumber
            address {
                country
                countryState
                city
                }
            industry
            foundationDate
            ownerId
        }
        }
    }
}
`)

export default enterprisesByAuthenticatedUserQuery

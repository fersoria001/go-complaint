import { graphql } from "@/gql";


const enterprisesByAuthenticatedUserQuery = graphql(`
query EnterprisesByAuthenticatedUserQuery{
    enterprisesByAuthenticatedUser{
        enterprises{
            authority{
            authority
            enterpriseId
            principal
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
            industry {
                id
                name
            }
            foundationDate
            ownerId
        }
        }
        offices{
            authority{
            authority
            enterpriseId
            principal
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
            industry {
                id
                name
            }
            foundationDate
            ownerId
        }
        }
    }
}
`)

export default enterprisesByAuthenticatedUserQuery

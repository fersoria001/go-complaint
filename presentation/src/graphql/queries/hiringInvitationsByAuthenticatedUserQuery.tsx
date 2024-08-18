import { graphql } from "@/gql";

const hiringProcessByAuthenticatedUserQuery = graphql(`
    query hiringProcessByAuthenticatedUserQuery{
        hiringProcessByAuthenticatedUser{
            id
            enterprise {
                id
                subjectName
                subjectThumbnail
                subjectEmail
            }
            user {
                id
                userName
                    person {
                        profileImg
                        email
                        firstName
                        lastName
                        genre
                        pronoun
                        age
                        phoneNumber
                        address { 
                            country
                            countryState
                            city
                            }
                    }
                status
            }
            role
            status
            reason
            emitedBy {
                id
                subjectName
                subjectThumbnail
                subjectEmail
            }
            occurredOn
            lastUpdate
            updatedBy {
                id
                subjectName
                subjectThumbnail
                subjectEmail
            }
            industry {
                id
                name
            }
        }
    }`)
export default hiringProcessByAuthenticatedUserQuery
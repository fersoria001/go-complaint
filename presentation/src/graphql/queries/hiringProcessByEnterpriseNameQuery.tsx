import { graphql } from "@/gql";

const hiringProcessByEnterpriseNameQuery = graphql(`
    query hiringProcessByEnterpriseName($name: String!) {
        hiringProcessByEnterpriseName(name:$name){
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

export default hiringProcessByEnterpriseNameQuery;
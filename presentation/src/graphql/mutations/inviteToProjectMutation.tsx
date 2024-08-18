import { graphql } from "@/gql";

const inviteToProjectMutation = graphql(`
    mutation inviteToProjectMutation($input: InviteToProject!){
        inviteToProject(input: $input){
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
        }
    }`)

export default inviteToProjectMutation
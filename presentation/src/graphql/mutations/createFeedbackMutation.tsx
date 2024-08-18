import { graphql } from "@/gql";

const createFeedbackMutation = graphql(`
    mutation createFeedbackMutation($input: CreateFeedback) {
        createFeedback(input: $input){
            id
            complaintId
            enterpriseId
            replyReview {
                id
                feedbackId
                reviewer {
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
                    }
                    status
                }
                replies {
                    id
                    complaintId
                    sender {
                        id
                        subjectName
                        subjectThumbnail
                        subjectEmail
                    }
                    body
                    createdAt
                    read
                    readAt
                    updatedAt
                    isEnterprise
                    enterpriseId
                }
                review {
                    id
                    comment
                }
                color
                createdAt
            }
            reviewedAt
            updatedAt
            isDone
        }
    }
    `)
export default createFeedbackMutation
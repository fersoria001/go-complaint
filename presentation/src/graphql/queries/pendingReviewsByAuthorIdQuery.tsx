import { graphql } from "@/gql";
const pendingReviewsByAuthorIdQuery = graphql(`
    query pendingReviewsByAuthorId($id:String!, $term: String){
        pendingReviewsByAuthorId(id:$id,term:$term){
            id
            author{
                id
                subjectName
                subjectThumbnail
                isOnline
                isEnterprise
            }
            receiver{
                id
                subjectName
                subjectThumbnail
                isOnline
                isEnterprise
            }
            status
            title
            description
            createdAt
            updatedAt
            rating {
                id
                rate
                comment
                sentToReviewBy{
                    id
                    subjectName
                    subjectThumbnail
                    isOnline
                    isEnterprise
                }
                ratedBy{
                    id
                    subjectName
                    subjectThumbnail
                    isOnline
                    isEnterprise
                }
                createdAt
                lastUpdate
            }
            replies{
                id
                complaintId
                sender{
                    id
                    subjectName
                    subjectThumbnail
                    isEnterprise
                }
                body
                read
                readAt
                createdAt
                updatedAt
            }
        }
    }`)

export default pendingReviewsByAuthorIdQuery;
import { graphql } from "@/gql";

const complaintsByAuthorIdOrReceiverIdQuery = graphql(`
    query complaintsByAuthorIdOrReceiverIdQuery($id: String!){
        complaintsByAuthorOrReceiverId(id:$id){
            id
            author{
                id
                subjectName
                subjectThumbnail
            }
            receiver{
                id
                subjectName
                subjectThumbnail
            }
            status
            title
            description
            createdAt
            updatedAt
            replies{
                id
                complaintId
                sender{
                    id
                    subjectName
                    subjectThumbnail
                }
                body
                readAt
                createdAt
                updatedAt
            }
        }
    }`)

export default complaintsByAuthorIdOrReceiverIdQuery;
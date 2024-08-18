import { graphql } from "@/gql";

const complaintByIdQuery = graphql(`
    query ComplaintByIdQuery($id: String!){
        complaintById(id:$id){
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
                isEnterprise
                enterpriseId
            }
        }
    }`)

export default complaintByIdQuery;
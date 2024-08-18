import { graphql } from "@/gql";

const complaintsOfResolvedFeedbackByEmployeeIdQuery = graphql(
    `query complaintsOfResolvedFeedbackByEmployeeId($id: String!){
        complaintsOfResolvedFeedbackByEmployeeId(id:$id){
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
            }
        }
    }`
)

export default complaintsOfResolvedFeedbackByEmployeeIdQuery;
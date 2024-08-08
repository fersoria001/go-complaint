import { graphql } from "@/gql";


const describeComplaintMutation = graphql(`
    mutation describeComplaintMutation($input:DescribeComplaint!){
        describeComplaint(input:$input){
            id
            receiver{
                id
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
                }
                body
                readAt
                createdAt
                updatedAt
            }
        }
    }
    `)

export default describeComplaintMutation;
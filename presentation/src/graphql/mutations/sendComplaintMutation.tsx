import { graphql } from "@/gql";

const sendComplaintMutation = graphql(`
    mutation sendComplaintMutation($input: SendComplaint!) {
        sendComplaint(input:$input){
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
    }`)

export default sendComplaintMutation;
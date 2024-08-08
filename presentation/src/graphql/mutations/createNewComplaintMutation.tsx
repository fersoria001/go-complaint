import { graphql } from "@/gql";

const createNewComplaintMutation = graphql(`
    mutation createNewComplaintMutation($input: CreateNewComplaint!){
        createNewComplaint(input:$input){
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

export default createNewComplaintMutation
import { graphql } from "@/gql";

const createNewComplaintMutation = graphql(`
    mutation createNewComplaintMutation($input: CreateNewComplaint!){
        createNewComplaint(input:$input){
            id
        }
    }`)

export default createNewComplaintMutation
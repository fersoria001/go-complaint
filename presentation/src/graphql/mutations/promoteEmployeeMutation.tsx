import { graphql } from "@/gql";

const promoteEmployeeMutation = graphql(`
    mutation promoteEmployeeMutation($input: PromoteEmployee!){
        promoteEmployee(input: $input){
            id
            enterprisePosition
        }
    }`)

export default promoteEmployeeMutation;
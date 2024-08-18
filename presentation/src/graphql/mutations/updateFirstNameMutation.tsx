import { graphql } from "@/gql";

const updateFirstNameMutation = graphql(`
    mutation updateFirstNameMutation($input: ChangeUserFirstName!) {
        changeFirstName(input: $input){
            id
        }
    }`)

export default updateFirstNameMutation;
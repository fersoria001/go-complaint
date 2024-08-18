import { graphql } from "@/gql";

const updatePasswordMutation = graphql(`
    mutation updatePasswordMutation($input: ChangePassword!){
        changePassword(input: $input){
            id
        }
    }`)

export default updatePasswordMutation;
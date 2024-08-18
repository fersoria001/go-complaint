import { graphql } from "@/gql";

const updateLastNameMutation = graphql(`
    mutation updateLastNameMutation($input: ChangeUserLastName!){
        changeLastName(input: $input){
            id
        }
    }`)

export default updateLastNameMutation;
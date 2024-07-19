import { graphql } from "@/gql";
const createUserMutation = graphql(`
    mutation CreateUserMutation(
        $input: CreateUser!){
            createUser(
                input: $input
            ){
                userName
            }
    }`)
export default createUserMutation;
import { graphql } from "@/gql";

const recoverPasswordMutation = graphql(`
    mutation recoverPasswordMutation($userName: String!){
        recoverPassword(userName: $userName)
    }
    `)

export default recoverPasswordMutation;
import { graphql } from "@/gql";

const createEnterpriseMutation = graphql(`
    mutation createEnterprise($input: CreateEnterprise!){
        createEnterprise(input: $input){
            name
        }
    }`)

export default createEnterpriseMutation
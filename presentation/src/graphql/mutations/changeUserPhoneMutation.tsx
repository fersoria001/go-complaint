import { graphql } from "@/gql";

const changeUserPhoneMutation = graphql(`
    mutation changeUserPhoneMutation($input: ChangeUserPhone!){
        changeUserPhone(input:$input){
            id
        }
    }`)

export default changeUserPhoneMutation;
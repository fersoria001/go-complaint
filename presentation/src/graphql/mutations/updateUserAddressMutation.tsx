import { graphql } from "@/gql";

const updateUserAddressMutation = graphql(`
    mutation updateUserAddressMutation($input: UpdateUserAddress!){
        updateUserAddress(input: $input){
            id
        }
    }`)
    
export default updateUserAddressMutation;
import { graphql } from "@/gql";

const changeEnterpriseAddressMutation = graphql(`
    mutation changeEnterpriseAddressMutation($input: ChangeEnterpriseAddress!){
        changeEnterpriseAddress(input: $input){
            id
        }
    }`)

export default changeEnterpriseAddressMutation;
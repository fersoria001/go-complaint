import { graphql } from "@/gql";

const changeEnterprisePhoneMutation = graphql(`
    mutation changeEnterprisePhoneMutation($input: ChangeEnterprisePhone!){
        changeEnterprisePhone(input: $input){
            id
        }
    }`)
export default changeEnterprisePhoneMutation;
import { graphql } from "@/gql";

const changeEnterpriseWebsiteMutation = graphql(`
    mutation changeEnterpriseWebsiteMutation($input: ChangeEnterpriseWebsite!){
        changeEnterpriseWebsite(input: $input){
            id
        }
    }`)

export default changeEnterpriseWebsiteMutation;
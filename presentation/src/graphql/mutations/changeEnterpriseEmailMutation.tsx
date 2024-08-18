import { graphql } from "@/gql";

const changeEnterpriseEmailMutation = graphql(`
    mutation changeEnterpriseEmailMutation($input: ChangeEnterpriseEmail!){
        changeEnterpriseEmail(input:$input){
            id
        }
    }`)
export default changeEnterpriseEmailMutation;
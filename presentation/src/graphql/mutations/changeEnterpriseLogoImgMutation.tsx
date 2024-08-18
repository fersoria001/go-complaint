import { graphql } from "@/gql";

const changeEnterpriseLogoImgMutation = graphql(`
    mutation changeEnterpriseLogoMutation($enterpriseId:String!, $file: Upload!){
        changeEnterpriseLogoImg(enterpriseId: $enterpriseId, file: $file){
            id
        }
    }`)

export default changeEnterpriseLogoImgMutation;
import { graphql } from "@/gql";

const changeEnterpriseBannerImgMutation = graphql(`
    mutation changeEnterpriseBannerMutation($enterpriseId:String!, $file: Upload!){
        changeEnterpriseBannerImg(enterpriseId: $enterpriseId, file: $file){
            id
        }
    }`)

export default changeEnterpriseBannerImgMutation;
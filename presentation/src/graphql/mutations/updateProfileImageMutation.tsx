import { graphql } from "@/gql";

const updateProfileImageMutation = graphql(`
    mutation updateProfileImageMutation($id:String!,$file: Upload!) {
        updateProfileImg(id:$id,file: $file) {
            id
        }
    }`)

export default updateProfileImageMutation
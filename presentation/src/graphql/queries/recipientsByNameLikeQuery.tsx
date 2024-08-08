import { graphql } from "@/gql";

const recipientsByNameLikeQuery = graphql(`
    query recipientsByNameLikeQuery($term:String!){
        recipientsByNameLike(term:$term){
            id
            subjectName
            subjectThumbnail
            isEnterprise
        }
    }
`)

export default recipientsByNameLikeQuery;
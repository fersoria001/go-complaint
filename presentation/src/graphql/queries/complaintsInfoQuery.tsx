import { graphql } from "@/gql";

const complaintsInfoQuery = graphql(`
    query complaintsInfoQuery($id:String!){
        complaintsReceivedInfo(id:$id){
            received
            resolved
            reviewed
            pending
            avgRating
            total
        }
    }`)

export default complaintsInfoQuery;
import { graphql } from "@/gql";

const markNotificationAsReadMutation = graphql(`
    mutation markNotificationAsReadMutation($id:String!){
        markNotificationAsRead(id:$id){
            id
        }
    }
    `)

export default markNotificationAsReadMutation
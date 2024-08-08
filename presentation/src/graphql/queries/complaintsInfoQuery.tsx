import { graphql } from "@/gql";

const complaintsInfoQuery = graphql(`
    query complaintsInfoQuery($id: String!) {
        complaintsInfo(id: $id) {
            received {
                id
                ownerId
                complaintId
                occurredOn
                dataType
            }
            resolved {
                id
                ownerId
                complaintId
                occurredOn
                dataType
            }
            reviewed {
                id
                ownerId
                complaintId
                occurredOn
                dataType
            }
            sent {
                id
                ownerId
                complaintId
                occurredOn
                dataType
            }
        }
    }`)

export default complaintsInfoQuery;
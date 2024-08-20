import { graphql } from "@/gql";

const complaintWritingByAuthorIdAndReceiverIdQuery = graphql(`
    query complaintWritingByAuthorIdAndReceiverIdQuery($input:FindComplaintWriting!){
        complaintWritingByAuthorIdAndReceiverId(input:$input){
            id
        }
    }`)
export default complaintWritingByAuthorIdAndReceiverIdQuery
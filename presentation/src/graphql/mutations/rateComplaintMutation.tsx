import { graphql } from "@/gql";

const rateComplaintMutation = graphql(`
mutation rateComplaintMutation($input: RateComplaint!){
    rateComplaint(input: $input){
            id
            rate
            comment
            sentToReviewBy{
                id
                subjectName
                subjectThumbnail
                isOnline
                isEnterprise
            }
            ratedBy{
                id
                subjectName
                subjectThumbnail
                isOnline
                isEnterprise
            }
            createdAt
            lastUpdate  
    }
}`)

export default rateComplaintMutation
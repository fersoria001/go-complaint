import { graphql } from "@/gql";

const removeFeedbackReplyMutation = graphql(`
    mutation removeFeedbackReplyMutation($input:RemoveFeedbackReply!){
        removeFeedbackReply(input: $input) {
            id
        }
    }`)

export default removeFeedbackReplyMutation;
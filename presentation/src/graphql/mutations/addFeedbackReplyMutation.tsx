import { graphql } from "@/gql";

const addFeedbackReplyMutation = graphql(`
    mutation addFeedbackReplyMutation($input: AddFeedbackReply!){
        addFeedbackReply(input: $input){
            id
        }
    }`)
export default addFeedbackReplyMutation
import { graphql } from "@/gql";

const removeFeedbackCommentMutation = graphql(`
    mutation removeFeedbackCommentMutation($input: RemoveFeedbackComment!){
        removeFeedbackCommand(input: $input){
            id
        }
    }`)

export default removeFeedbackCommentMutation
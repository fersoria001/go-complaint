import { graphql } from "@/gql";

const addFeedbackCommentMutation = graphql(`
    mutation addFeedbackCommentMutation($input: AddFeedbackComment!){
        addFeedbackComment(input:$input){
            id
        }
    }`)
export default addFeedbackCommentMutation;
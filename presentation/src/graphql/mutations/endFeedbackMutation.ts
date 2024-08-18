import { graphql } from "@/gql";

const endFeedbackMutation = graphql(`
  mutation endFeedbackMutation($input: EndFeedback!) {
    endFeedback(input: $input) {
      id
    }
  }
`);

export default endFeedbackMutation;
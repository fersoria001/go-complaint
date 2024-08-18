import { graphql } from "@/gql";

const changeUserPronounMutation = graphql(`
    mutation changeUserPronounMutation($input: ChangeUserPronoun!) {
        changeUserPronoun(input: $input) {
            id
        }
    }`)

export default changeUserPronounMutation
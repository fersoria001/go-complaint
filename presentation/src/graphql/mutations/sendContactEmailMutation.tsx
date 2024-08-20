import { graphql } from "@/gql";

const sendContactEmailMutation = graphql(`
    mutation sendContactEmailMutation($input: ContactEmail!){
        contactEmail(input: $input)
    }`)

export default sendContactEmailMutation
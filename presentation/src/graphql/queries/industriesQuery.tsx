import { graphql } from "@/gql";

const industriesQuery = graphql(
    `query industries{
        industries{
            id
            name
        }
    }`
)

export default industriesQuery;
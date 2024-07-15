import { graphql } from "@/gql";

const countriesQuery = graphql(`
    query Countries{
        Countries{
            id
            name
            phoneCode
        }
    }`)
export default countriesQuery;
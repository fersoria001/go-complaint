import { graphql } from "@/gql";

const countriesQuery = graphql(`
    query Countries{
        countries{
            id
            name
            phoneCode
        }
    }`)
export default countriesQuery;
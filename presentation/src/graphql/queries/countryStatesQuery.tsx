import { graphql } from "@/gql";

const countryStatesQuery = graphql(`
    query CountryStatesQuery($id: Int!){
        CountryStates(id: $id) {
            id
            name
        }
    }`)

export default countryStatesQuery;
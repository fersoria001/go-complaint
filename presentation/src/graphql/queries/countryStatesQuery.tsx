import { graphql } from "@/gql";

const countryStatesQuery = graphql(`
    query CountryStatesQuery($id: Int!){
        countryStates(id: $id) {
            id
            name
        }
    }`)

export default countryStatesQuery;
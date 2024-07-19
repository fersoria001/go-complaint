import { graphql } from "@/gql";

const citiesQuery = graphql(`
    query CitiesQuery($id: Int!){
        cities(id:$id){
            id
            name
            countryCode
            latitude
            longitude
        }
    }`)

export default citiesQuery;
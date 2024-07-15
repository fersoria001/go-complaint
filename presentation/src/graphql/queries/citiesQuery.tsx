import { graphql } from "@/gql";

const citiesQuery = graphql(`
    query CitiesQuery($id: Int!){
        Cities(id:$id){
            id
            name
            countryCode
            latitude
            longitude
        }
    }`)

export default citiesQuery;
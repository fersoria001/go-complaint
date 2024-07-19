import { graphql } from "@/gql";

const userByIdQuery = graphql(`
    query User($id:String!){
        userById(id:$id){
                userName
                person {
                    profileImg
                    email
                    firstName
                    lastName
                    genre
                    pronoun
                    age
                    phoneNumber
                    address { 
                        country
                        countryState
                        city
                        }
                }
                status
        }
    }`)

export default userByIdQuery;
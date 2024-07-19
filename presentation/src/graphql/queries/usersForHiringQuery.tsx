import { graphql } from "@/gql";

const usersForHiringQuery = graphql(`
    query usersForHiringQuery($input: SearchWithPagination!) {
        usersForHiring(input:$input){
            users {
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
                    address { country countryState city}
                }
                status
            }
            count
            limit
            offset
            nextCursor
            prevCursor
        }
    }`)
export default usersForHiringQuery
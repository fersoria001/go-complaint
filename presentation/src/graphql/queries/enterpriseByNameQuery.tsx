import { graphql } from "@/gql";

const enterpriseByNameQuery = graphql(`
    query enterpriseByNameQuery($name: String!){
        enterpriseByName(name:$name){
            id
            name
            logoImg
            bannerImg
            website
            email
            phoneNumber
            address {
                country
                countryState
                city
                }
            industry {
                id
                name
            }
            foundationDate
            ownerId
            employees {
                id
                enterpriseId
                userId
                user{
                    id
                    userName
                    person{
                        profileImg
                        email
                        firstName
                        lastName
                        genre
                        pronoun
                        age
                        phoneNumber
                        address{
                            country
                            countryState
                            city
                        }
                    }
                    status
                }
                hiringDate
                approvedHiring
                approvedHiringAt
                enterprisePosition
            }
        }
    }`)

export default enterpriseByNameQuery;
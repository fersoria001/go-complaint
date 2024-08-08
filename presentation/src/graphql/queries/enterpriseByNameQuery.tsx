import { graphql } from "@/gql";

const enterpriseByNameQuery = graphql(`
    query enterpriseByNameQuery($name: String!){
        enterpriseByName(name:$name){
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
            industry
            foundationDate
            ownerId
            employees {
                id
                enterpriseId
                userId
                user{
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
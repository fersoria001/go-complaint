/**
 * 
 * @param id feedbackId
 * @returns 
 */
function enterpriseByIdSubscription(id: string, userId: string) {
    return `
subscription{
    enterpriseById(id: "${id}", userId: "${userId}"){
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
}
`
}

export default enterpriseByIdSubscription;
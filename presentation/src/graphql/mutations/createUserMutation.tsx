import { graphql } from "@/gql";
const createUserMutation = graphql(`
    mutation CreateUserMutation(
        $email:String!,
        $password:String!,
        $firstName:String!,
        $lastName:String!,
        $gender:String!,
        $pronoun:String!,
        $birthDate:String!,
        $phone:String!,
        $countryId:Int!,
        $countryStateId:Int!,
        $cityId:Int!){
        CreateUser(
            email: $email,
            password: $password,
            firstName: $firstName,
            lastName: $lastName,
            gender: $gender,
            pronoun: $pronoun,
            birthDate: $birthDate,
            phone: $phone,
            countryId: $countryId,
            countryStateId: $countryStateId,
            cityId: $cityId
        )
    }`)
export default createUserMutation;
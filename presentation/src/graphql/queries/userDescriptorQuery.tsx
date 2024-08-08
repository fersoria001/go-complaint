import { graphql } from "@/gql";

const userDescriptorQuery = graphql(`
    query UserDescriptor{
        userDescriptor{
            id
            userName
            fullName
            profileImg
            genre
            pronoun
            authorities{
                enterpriseId
                principal
                authority
            }
        }
    }`)

export default userDescriptorQuery
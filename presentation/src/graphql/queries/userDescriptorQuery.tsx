import { graphql } from "@/gql";

const userDescriptorQuery = graphql(`
    query UserDescriptor{
        userDescriptor{
            userName
            fullName
            profileImg
            genre
            pronoun
            authorities{
                enterpriseId
                authority
            }
        }
    }`)

export default userDescriptorQuery
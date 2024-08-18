import { graphql } from "@/gql";

const createEnterpriseChatMutation = graphql(`
    mutation createEnterpriseChatMutation($input: CreateEnterpriseChat!){
        createEnterpriseChat(input: $input){
            id
            enterpriseId
            recipientOne {
                id
                subjectName
                subjectThumbnail
                subjectEmail
            }
            recipientTwo {
                id
                subjectName
                subjectThumbnail
                subjectEmail
            }
            replies {
                id
                chatId
                sender {
                    id
                    subjectName
                    subjectThumbnail
                    subjectEmail
                }
                content
                createdAt
                updatedAt
                seen
            }
        }
    }`)

export default createEnterpriseChatMutation;
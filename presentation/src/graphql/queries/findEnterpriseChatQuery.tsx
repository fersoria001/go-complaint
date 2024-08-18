import { graphql } from "@/gql";

const findEnterpriseChatQuery = graphql(`
    query findEnterpriseChatQuery($input: FindEnterpriseChat!){
        findEnterpriseChat(input: $input){
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

export default findEnterpriseChatQuery;
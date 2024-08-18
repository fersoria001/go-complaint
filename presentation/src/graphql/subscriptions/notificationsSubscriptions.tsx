/**
 * 
 * @param id feedbackId
 * @returns 
 */
function notificationsSubscription(id: string) {
    return `
subscription{
    notifications(id: "${id}"){
        id
        owner {
            id
            subjectName
            subjectThumbnail
            subjectEmail
        }
        sender {
            id
            subjectName
            subjectThumbnail
            subjectEmail
        }
        title
        content
        link
        seen
        occurredOn
    }
}
`
}

export default notificationsSubscription;
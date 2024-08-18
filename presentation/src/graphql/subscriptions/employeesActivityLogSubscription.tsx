/**
 * 
 * @param id enterprise name unescaped
 * @returns 
 */
function employeesActivityLogSubscription(id: string) {
    return `
subscription{
    employeesActivityLog(id: "${id}"){
        id
        user{
            id
            subjectName
            subjectThumbnail
            isOnline
        }
        activityId
        enterpriseId
        enterpriseName
        occurredOn
        activityType
    }
}
`
}

export default employeesActivityLogSubscription;
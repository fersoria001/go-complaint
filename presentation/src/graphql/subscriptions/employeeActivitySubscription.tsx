/**
 * 
 * @param id employeeId
 * @returns 
 */
function employeeActivitySubscription(id: string) {
    return `
subscription{
    employeeActivity(id: "${id}"){
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

export default employeeActivitySubscription;
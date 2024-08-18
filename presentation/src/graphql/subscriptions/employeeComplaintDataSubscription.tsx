/**
 * 
 * @param id employee.user id
 * @returns 
 */
function employeeComplaintDataSubscription(id: string) {
    return `
    subscription{
        employeeComplaintData(id: "${id}"){
                id
                ownerId
                complaintId
                occurredOn
                dataType
        }
    }
    `
}

export default employeeComplaintDataSubscription;
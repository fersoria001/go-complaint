/**
 * 
 * @param id user id or enterprise id
 * @returns 
 */
function complaintDataByOwnershipSubscription(id: string) {
    return `
    subscription{
        complaintDataByOwnership(id: "${id}"){
                id
                ownerId
                complaintId
                occurredOn
                dataType
        }
    }
    `
}

export default complaintDataByOwnershipSubscription;
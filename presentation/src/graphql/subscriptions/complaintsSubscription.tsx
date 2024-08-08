function complaintsSubscription(id: string) {
    return `subscription{
    complaints(id: "${id}"){
              id
              author{
                  id
                  subjectName
                  subjectThumbnail
                  isOnline
              }
              receiver{
                  id
                  subjectName
                  subjectThumbnail
                  isOnline
              }
              status
              title
              description
              createdAt
              updatedAt
              replies{
                  id
                  complaintId
                  sender{
                      id
                      subjectName
                      subjectThumbnail
                  }
                  body
                  readAt
                  createdAt
                  updatedAt
              }
          }
  }`
  }
  

export default complaintsSubscription;
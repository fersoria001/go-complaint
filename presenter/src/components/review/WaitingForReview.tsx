import { ComplaintReviewType, UserDescriptor } from "../../lib/types"
import ComplaintCard from "../complaint/ComplaintCard"


interface Props {
    review: ComplaintReviewType
    user?: UserDescriptor
    enterpriseId?: string
}

function WaitingForReview({ review, user, enterpriseId }: Props) {
    let content = `You asked ${review.complaint.authorFullName} for a review of your attention at this complaint.`

    if (enterpriseId) {
        const lastReply = review.complaint.replies!.shift()?.senderName
        if (user?.email!= review.triggeredBy.email) {
            content = `${review.triggeredBy.firstName} ${review.triggeredBy.lastName} 
        asked ${review.complaint.authorFullName} for a review of ${lastReply == user?.fullName ? 'your' : lastReply } attention at this complaint.`
        } else {
            content = `You
        asked ${review.complaint.authorFullName} for a review of ${lastReply == user?.fullName ? 'your' : lastReply } attention at this complaint.`
        }
    }
    return (
        <div className="mx-auto pt-5">
            <div
                className="flex flex-col p-5 bg-white border border-gray-200 rounded-lg shadow 
            text-gray-700">
                <p className="text-sm md:text-xl text-gray-700 mb-4">
                    {content}
                </p>
                <div className="shadow p-5 ">
                    <ComplaintCard
                        reason={review.complaint.message.title}
                        description={review.complaint.message.description}
                        body={review.complaint.message.body}
                        status={review.complaint.status}
                        createdAt={review.complaint.createdAt}
                    />
                </div>
            </div>
        </div >)
}


export default WaitingForReview;
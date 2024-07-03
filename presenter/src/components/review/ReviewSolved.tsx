import { ComplaintReviewType, UserDescriptor } from "../../lib/types"
import StarIcon from "../icons/StarIcon"


interface Props {
    review: ComplaintReviewType
    user?: UserDescriptor
}

function ReviewSolved({ review, user }: Props) {
    let content
    let comment
    let isEnterprise = false
    let enterpriseName = ''
    const lastReply = review.complaint.replies![review.complaint.replies!.length - 1]
    review.complaint.replies!.forEach((reply) => {
        if (reply.isEnterprise) {
            isEnterprise = true
            enterpriseName = reply.enterpriseID
        }
    })
    if (review.ratedBy.email == user?.email) {
        content = `You reviewed ${lastReply.senderName} attention ${isEnterprise ? `at  this ${enterpriseName} complaint` : 'at this complaint'}`
        comment = `Your comment: ${review.complaint.rating!.comment}`
    } else {
        content = `${review.ratedBy.firstName} ${review.ratedBy.lastName}
         reviewed ${lastReply.senderID == user?.email ? 'your' : lastReply.senderName}
         attention ${isEnterprise ? `at  this ${enterpriseName} complaint` : 'at this complaint'}`
        comment = `Their comment: ${review.complaint.rating!.comment}`
    }

    return <div className="mx-auto pt-5">
        <div
            className="flex flex-col p-5 bg-white border border-gray-200 rounded-lg shadow 
                text-gray-700">
            <p className="text-sm md:text-xl text-gray-700 mb-4">
                {content}
            </p>
            <div className="shadow p-5 ">
                <div className="flex mb-3">
                    <p className="text-sm md:text-xl text-gray-700 mb-4">Complaint title: {` ${review.complaint.message.title} `}</p>
                </div>
                <div className="flex mb-3">
                    <p className="text-sm md:text-xl text-gray-700 mb-4 pr-2">Rating: </p>
                    {
                        [...Array(5)].map((_, index) => {
                            index += 1;
                            return <span
                                key={index}>
                                <StarIcon index={index} rating={review.complaint.rating!.rate} hover={0} fill={""} />
                            </span>
                        })
                    }
                </div>
                <p className="text-sm md:text-xl text-gray-700 mb-4">{comment}</p>
            </div>
        </div>
    </div >
}


export default ReviewSolved;
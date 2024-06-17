
import { timeAgo } from "../../lib/actions";
import { SolvedReview } from "../../lib/types";
import StarIcon from "../icons/StarIcon";
interface Props {
    solvedReview: SolvedReview;
    occurredOn: string;
}

function OldReviewsCard({ solvedReview, occurredOn }: Props) {

    return <div className="mx-auto pt-5">
        <div
            className="flex flex-col p-5
             bg-white border border-gray-200 rounded-lg shadow 
                text-gray-700
             ">
            <p>You reviewed {` ${solvedReview.User.firstName} ${solvedReview.User.lastName} `}
                attention, in regard to the complaint you sent</p>
            <div className="shadow p-5 hover:bg-gray-100">
                <div className="flex mb-3">
                    <p >Complaint title: {` ${solvedReview.Complaint.message.title} `}</p>
                </div>
                <div className="flex mb-3">
                    <p className="pr-2">Rating: </p>
                    {
                        [...Array(5)].map((_, index) => {
                            index += 1;
                            return <span
                                key={index}>
                                <StarIcon index={index} rating={solvedReview.Complaint.rating.rate} hover={0} fill={""} />
                            </span>
                        })
                    }
                </div>
                <p className="mb-3">Your comment: {" " + solvedReview.Complaint.rating.comment}</p>
                <p className="mb-3"> This review was made
                    {` ${timeAgo(occurredOn)} `}</p>
            </div>
        </div>
    </div >
}


export default OldReviewsCard;
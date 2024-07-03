import { Route } from "../../routes/_profile/reviews";
import Review from "../review/Review";
import ReviewSolved from "../review/ReviewSolved";
import WaitingForReview from "../review/WaitingForReview";



function Reviews() {
    const { descriptor, complaintReviews } = Route.useLoaderData();
    return (
        <div className="min-h-[315px] md:min-h-[460px]">
            <ul>
                {complaintReviews.length > 0 ?
                    complaintReviews.map((complaintReview) =>
                        <li key={complaintReview.eventID}>
                            {complaintReview.status === "pending" ?
                                <Review review={complaintReview} user={descriptor} /> :
                                complaintReview.status === "rated" ?
                                    <ReviewSolved
                                        review={complaintReview}
                                        user={descriptor}
                                    /> :
                                    <WaitingForReview review={complaintReview} user={descriptor} />
                            }
                        </li>
                    )
                    :
                    <div className="flex justify-center items-center h-screen">
                        <p className="text-center text-gray-700">
                            Currently you have not reviewed any complaint attention
                        </p>
                    </div>
                }
            </ul>

        </div>
    )
}

export default Reviews;
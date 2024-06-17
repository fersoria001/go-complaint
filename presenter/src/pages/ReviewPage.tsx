import { useLoaderData } from "react-router-dom";
import { InfoForReview, SolvedReview } from "../lib/types";
import ReviewCard from "../components/Review/ReviewCard";
import OldReviewsCard from "../components/Review/OldReviewsCard";

function ReviewPage() {
    const { pendingReviews, solvedReviewss } = useLoaderData() as {
        pendingReviews: { eventID: string; info: InfoForReview }[] | null;
        solvedReviewss:
        | {
            eventID: string;
            solvedReview: SolvedReview;
            occurredOn: string;
        }[]
        | null;
    }
    console.log(solvedReviewss);
    return (
        <div>
            {
                pendingReviews ?
                    <ul>
                        {pendingReviews.length > 0 &&
                            pendingReviews.map((info) =>
                                <li key={info.eventID}>
                                    <ReviewCard info={info.info} notificationID={info.eventID} />
                                </li>
                            )
                        }
                    </ul>
                    :
                    <div className="flex justify-center items-center h-screen">
                        <p className="text-center text-gray-700">
                            Currently you have not reviewed any complaint attention
                        </p>
                    </div>
            }
            {
                solvedReviewss ?
                    <ul>
                        {solvedReviewss.length > 0 &&
                            solvedReviewss.map((solvedReview) =>
                                <li key={solvedReview.eventID}>
                                    <OldReviewsCard solvedReview={solvedReview.solvedReview} occurredOn={solvedReview.occurredOn} />
                                </li>
                            )
                        }
                    </ul>
                    :
                    <div className="flex justify-center items-center h-screen">
                        <p className="text-center text-gray-700">
                            Currently you have not reviewed any complaint attention yet
                        </p>
                    </div>
            }
        </div>
    )

}


export default ReviewPage;
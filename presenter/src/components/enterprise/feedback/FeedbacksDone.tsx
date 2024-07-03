import { Route } from "../../../routes/$enterpriseID/feedback"
import FeedbackDone from "./FeedbackDone"

function FeedbacksDone() {
    const feedbacks = Route.useLoaderData()
    const {enterpriseID } = Route.useParams()
    return (
        <div className="min-h-[360px] md:min-h-[460px]">
            <div className="border rounded-md bg-white">
                {
                    feedbacks.length > 0 ?
                        feedbacks.map((v) => (
                            <FeedbackDone key={v.id} feedback={v} enterpriseID={enterpriseID} />
                        )) :
                        <div>
                            No feedbacks received yet.
                        </div>
                }
            </div>
        </div>
    )
}

export default FeedbacksDone
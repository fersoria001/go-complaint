import { Link } from "@tanstack/react-router"
import { FeedbackType, User } from "../../../lib/types"
import CommunicationIcon from "../../icons/CommunicationIcon"
import { useState } from "react"

interface Props {
    feedback: FeedbackType,
    enterpriseID: string,
}
function FeedbackDone({ enterpriseID, feedback }: Props) {
    const [fill, setFill] = useState('#374151')
    const diffReviewers: User[] = []
    const s = feedback.replyReview
    for (let i = 0; i < s.length; i++) {
        const exists = diffReviewers.find((v) => v.email == s[i].reviewer.email)
        if (!exists) {
            diffReviewers.push(s[i].reviewer)

        }
    }
    const date = new Date(parseInt(feedback.updatedAt)).toLocaleDateString()
    console.error("Date", feedback.updatedAt)
    return (
        <div className="flex flex-col p-3" >
            <div className="flex">
                <div>
                    <p className="text-gray-700 mb-4 text-sm md:text-xl cursor-default"> You have received feedback from </p>
                    <p className="self-baseline mt-10 text-sm md:text-sm text-gray-800 cursor-default">{date}</p>
                </div>
                <div className="flex flex-col mt-8 ml-5" >
                    {
                        diffReviewers.map((v) => {
                            return (
                                <div className="flex items-center align-center" key={v.email}>
                                    <img className="w-5 h-5 md:w-10 md:h-10 rounded-full" src={v.profileIMG} />
                                    <p className="ps-5 text-gray-700 text-sm md:text-xl">{`${v.firstName} ${v.lastName}`}</p>
                                </div>

                            )
                        })
                    }
                </div>
                <div className="ms-auto mt-auto  mr-3">
                    <Link
                        id="feedback-btn"
                        to={`/${enterpriseID}/feedbacks/`}
                        search={
                            {
                                id: feedback.id
                            }
                        }
                        onMouseEnter={() => setFill('#ffffff')}
                        onMouseLeave={() => setFill('#374151')}
                        className={`inline-flex items-center px-1 py-2
                                 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-cyan-200
                        hover:text-white
                        relative z-10  bg-gradient-to-br from-cyan-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`
                        }>
                        <CommunicationIcon fill={fill} />
                        Look at the feedback
                    </Link>
                </div>
            </div>

        </div >
    )
}

export default FeedbackDone;

{/* */ }
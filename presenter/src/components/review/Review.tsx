/* eslint-disable @typescript-eslint/no-unused-vars */
/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useCallback, useState } from "react";

import { useRouter } from "@tanstack/react-router";
import { Mutation, RateComplaintMutation } from "../../lib/mutations";
import { ComplaintReviewType, ErrorType, RateComplaint, UserDescriptor } from "../../lib/types";
import AcceptBtn from "../buttons/AcceptBtn";
import StarIcon from "../icons/StarIcon";


interface Props {
    review: ComplaintReviewType;
    enterpriseId?: string;
    user?: UserDescriptor;
}
const Review: React.FC<Props> = ({ review, enterpriseId }: Props) => {
    const [rating, setRating] = useState(0);
    const [hover, setHover] = useState(0);
    const [comment, setComment] = useState("");
    const [errors, setErrors] = useState<ErrorType>({})
    const router = useRouter();
    const handleSubmit = useCallback(async () => {
        try {
            const ok = await Mutation<RateComplaint>(
                RateComplaintMutation,
                {
                    complaintId: review.complaint.id,
                    eventId: review.eventID,
                    rate: rating,
                    comment: comment,
                }
            )
            return ok
        } catch (e: any) {
            if (e.message.includes("bad request")) {
                setErrors({ review: "complete the review before sending it" })
            } else {
                router.navigate({ to: `/errors` });
            }
            return false
        }
    }, [comment, rating, review, router])

    const cleanUp = useCallback(() => {
        router.invalidate()
        return Promise.resolve(true);
    }, [router])
    let content = `
    ${review.triggeredBy.firstName}  ${review.triggeredBy.lastName} has been asked for you
    to rate ${review.triggeredBy.pronoun == "he" ? "his" : "her"} attention in regard to the ${review.complaint.message.title} complaint
    you sent to ${review.complaint.receiverFullName}
    `
    if (enterpriseId) {
        content = `${review.triggeredBy.firstName}  ${review.triggeredBy.lastName} has been asked for ${enterpriseId}
        to rate ${review.triggeredBy.pronoun == 'he' ? 'his' : 'her'} attention in regard to the ${review.complaint.message.title} complaint`
    }
    return (
        <div className="md:w-4/5 md:mx-0  mx-auto pt-5">
            <div className="flex flex-col p-5 bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
                <p className="text-sm md:text-xl text-gray-700">{content}</p>
                <div className="flex py-5">
                    {[...Array(5)].map((_, index) => {
                        index += 1;
                        return <span
                            key={index}
                            onClick={() => setRating(index)}
                            onMouseEnter={() => setHover(index)}
                            onMouseLeave={() => setHover(0)}>
                            <StarIcon index={index} rating={rating} hover={hover} fill={""} />
                        </span>
                    })}
                </div>
                <textarea
                    className="block p-2.5 w-full text-sm md:text-xl
               text-gray-900 bg-gray-50 rounded-lg border border-gray-300
                focus:ring-blue-500 focus:border-blue-500"
                    rows={4}
                    maxLength={250}
                    placeholder="Write your review here"
                    onChange={(e) => setComment(e.target.value)}
                >
                </textarea>
                {errors?.review && <span className="text-red-500 text-xs italic" >{errors.review}</span>}
                <div className="self-center pt-6 md:self-end">
                    <AcceptBtn
                        variant="primary"
                        text="Rate"
                        callback={handleSubmit}
                        cleanUp={cleanUp}
                    />
                </div>
            </div>
        </div>
    );
}

export default Review;
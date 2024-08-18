'use client'

import { useState } from "react";
import StarIcon from "../icons/StarIcon";
import clsx from "clsx";
import { Complaint, UserDescriptor } from "@/gql/graphql";
import { dateFromMsString } from "@/lib/dateFromMsString";
import { rateComplaint, ReviewFormErrors } from "@/lib/actions/graphqlActions";
import InlineAlert from "../error/InlineAlert";
import { useQueryClient } from "@tanstack/react-query";
interface Props {
    alias: string
    item: Complaint
    currentUser: UserDescriptor
}
const Review: React.FC<Props> = ({ alias, item, currentUser }: Props) => {
    const [rating, setRating] = useState(0);
    const [hover, setHover] = useState(0);
    const [errors, setErrors] = useState<ReviewFormErrors>()
    const queryClient = useQueryClient()

    let assistedBy = item.rating?.sentToReviewBy?.subjectName + " "
    if (item.receiver?.isEnterprise) {
        assistedBy += "from" + " " + item.receiver!.subjectName + " "
    }
    const handleSubmit = async (event: React.SyntheticEvent<HTMLFormElement>) => {
        event.preventDefault()
        const { success: ok, errors: reviewErrors } = await rateComplaint(new FormData(event.currentTarget))
        if (!ok) {
            setErrors(reviewErrors)
            return
        }
        queryClient.refetchQueries({
            queryKey: [
                'pendingReviewsByAuthorIdQuery',
                'complaintsSentForReviewByReceiverIdQuery',
                'complaintsRatedByAuthorIdQuery',
                'complaintsRatedByReceiverIdQuery'
            ]
        })
    
        return
    }
    return (
        <form className="w-full border-t border-b" onSubmit={handleSubmit} >
            <div className="max-w-xl flex flex-col p-5 bg-white">
                <p className="text-gray-500 text-sm md:text-md xl:text-md font-medium text-end mb-2">
                    {dateFromMsString(item.rating?.createdAt!).toDateString()}
                </p>
                <p className="text-sm md:text-md text-gray-700 mb-4">
                    You have been ask to review {assistedBy}
                    attention at this complaint:
                </p>
                <div className="mx-4 border rounded-md p-2">
                    <div className="flex">
                        <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="reason">Reason:</label>
                        <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md">{item.title}</p>
                    </div>
                    <div className="flex">
                        <label className="text-gray-700 text-sm md:text-md xl:text-md font-bold" htmlFor="description">Description:</label>
                        <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md">{item.description}</p>
                    </div>
                    <p className="ms-0.5 text-gray-700 text-sm md:text-md xl:text-md text-end">
                        {dateFromMsString(item.createdAt!).toDateString()}
                    </p>
                </div>
                <input className="hidden" value={currentUser.id} name="userId" />
                <input className="hidden" value={item.id!} name="complaintId" />
                <input className="hidden" value={rating} name="rate" />
                <div className="flex py-5">
                    {[...Array(5)].map((_, index) => {
                        index += 1;
                        return <span
                            key={index}
                            onClick={() => setRating(index)}
                            onMouseEnter={() => setHover(index)}
                            onMouseLeave={() => setHover(0)}>
                            <StarIcon className={clsx('w-6 h-6 md:w-8 md:h-8 fill-gray-200 cursor-pointer', {
                                "fill-yellow-500": index <= (hover || rating)
                            })} />
                        </span>
                    })}
                </div>
                <textarea
                    className="block p-2.5 w-full text-sm md:text-md xl:text-md focus:outline-none appearance-none resize-none text-gray-700 rounded-lg border border-gray-300"
                    name="comment"
                    rows={4}
                    maxLength={250}
                    placeholder="Write your review here">
                </textarea>
                {errors?.fieldErrors?.comment && <InlineAlert
                    className="flex items-center text-sm text-red-800 rounded-lg bg-red-50"
                    errors={errors.fieldErrors.comment} />}
                {errors?.formErrors && <InlineAlert
                    className="flex items-center text-sm text-red-800 rounded-lg bg-red-50"
                    errors={errors.formErrors} />}
                <button
                    type="submit"
                    className="self-center mt-6 md:self-end px-7 py-3 bg-blue-500 font-bold text-white rounded-md hover:bg-blue-600">
                    Rate
                </button>
            </div>
        </form>
    )
}
export default Review;
import { useEffect, useState } from "react";
import StarIcon from "../icons/StarIcon";
import PrimaryButton from "../buttons/PrimaryButton";
import { useNavigate } from "react-router-dom";
import { ErrorType, InfoForReview, RateComplaint, RateValidationSchema } from "../../lib/types";
import { Mutation, RateComplaintMutation } from "../../lib/mutations";
interface Props {
    info: InfoForReview;
    notificationID: string;
}
function ReviewCard({ info, notificationID }: Props) {
    const [rating, setRating] = useState(0);
    const [hover, setHover] = useState(0);
    const [comment, setComment] = useState("");
    const [errors, setErrors] = useState<ErrorType>({});
    const navigate = useNavigate();
    useEffect(() => {
    }, [errors]);
    function handleSubmit() {
        setErrors({});
        const parsed = RateValidationSchema.safeParse({
            rate: rating,
            comment: comment,
        });
        let errorPath: string;
        if (!parsed.success) {
            parsed.error.errors.forEach((error) => {
                errorPath = error.path.join("");
                errors[errorPath] = error.message;
            });
            return setErrors(errors);
        } else {
            const rateAComplaint: RateComplaint = {
                notificationID: notificationID,
                complaintID: info.Complaint.id,
                rate: parsed.data.rate,
                comment: parsed.data.comment,
            };
            Mutation<RateComplaint>(RateComplaintMutation, rateAComplaint).then(() =>
                navigate("/success/rate%20a%20complaint")
            )
        }
    }
    return (
        <div className="md:w-1/2 mx-auto pt-5">
            <div className="flex flex-col p-5 bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
                <p>{info.User.firstName + " " + info.User.lastName + " "}
                    has been asked for you
                    to rate his/her attention in regard to your
                    {" " + info.Complaint.message.title + " "} complaint
                    you sent to {info.Complaint.receiverFullName}</p>
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
                    className="block p-2.5 w-full text-sm
               text-gray-900 bg-gray-50 rounded-lg border border-gray-300
                focus:ring-blue-500 focus:border-blue-500"
                    rows={4}
                    maxLength={250}
                    placeholder="Write your review here"
                    onChange={(e) => setComment(e.target.value)}
                >
                </textarea>
                <span className="self-center py-2.5" onMouseUp={handleSubmit}>
                    <PrimaryButton text="Review" />
                </span>
            </div>
        </div>
    );
}

export default ReviewCard;
'use client'

import { useState } from "react";
import StarIcon from "../icons/StarIcon";
import clsx from "clsx";

const Review: React.FC = () => {
    const [rating, setRating] = useState(0);
    const [hover, setHover] = useState(0);
    return (
        <form className="w-full border-t border-b">
            <div className="max-w-xl flex flex-col p-5 bg-white">
                <p className="text-sm md:text-xl text-gray-700">You have been ask to review someone attention at this complaint</p>
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
                    className="block p-2.5 w-full text-sm md:text-xl focus:outline-none appearance-none resize-none text-gray-700 rounded-lg border border-gray-300"
                    rows={4}
                    maxLength={250}
                    placeholder="Write your review here">
                </textarea>
                {/* {errors?.review && <span className="text-red-500 text-xs italic" >{errors.review}</span>} */}
                {/* <div className="self-center pt-6 md:self-end">
                    <AcceptBtn
                        variant="primary"
                        text="Rate"
                        callback={handleSubmit}
                        cleanUp={cleanUp}
                    />
                </div> */}
                <button
                    type="button"
                    className="self-center mt-6 md:self-end px-7 py-3 bg-blue-500 font-bold text-white rounded-md hover:bg-blue-600">
                    Rate
                </button>
            </div>
        </form>
    )
}
export default Review;
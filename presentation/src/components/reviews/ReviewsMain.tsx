'use client'

import FilterBy from "../search/FilterBy"
import SearchInput from "../search/SearchInput"
import Review from "./Review"

const reviewFilterOptions = [
    {
        id: "0",
        name: "pending",
        value: "pending"
    },
    {
        id: "1",
        name: "solved",
        value: "solved"
    },
    {
        id: "2",
        name: "waiting",
        value: "waiting"
    },
]

const ReviewsMain: React.FC = () => {
    return (
        <div className="py-2 px-2">
            <SearchInput placeholder="Search..." onChange={() => { }} />
            <FilterBy options={reviewFilterOptions} />
            <div className="flex flex-col mt-2">
                <Review />
            </div>
        </div>
    )
}
export default ReviewsMain
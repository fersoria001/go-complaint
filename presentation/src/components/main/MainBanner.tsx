import ComplaintsReview from "../stickers/ComplaintsReview";
import TeamCommunication from "../stickers/TeamCommunication";
import VisualizeComplaints from "./../stickers/VisualizeComplaints";

const MainBanner: React.FC = () => {
    return (
        <div className="w-full bg-white md:grid md:grid-cols-3 md:px-2 py-3 lg:py-5 xl:py-7 2xl:py-12">
            <div className="flex px-2 md:flex-col">
                <VisualizeComplaints className="h-36 w-36 shrink-0 md:self-center lg:h-48 lg:w-48 xl:h-52 xl:w-52 2xl:h-72 2xl:w-72" />
                <p className="pt-6 ps-3 text-sm text-gray-700 md:text-center  md:text-md lg:text-lg xl:text-xl 2xl:text-3xl">
                    At Go Complaint you can visualize your complaints with ease,
                    we have a native integrated search with 
                    full-text search capabilities.
                </p>
            </div>
            <div className="flex px-2 md:flex-col">
                <ComplaintsReview className="h-36 w-36 shrink-0 md:self-center lg:h-48 lg:w-48 xl:h-52 xl:w-52 2xl:h-72 2xl:w-72" />
                <p className="pt-6 ps-3 text-sm text-gray-700 md:text-center md:text-md lg:text-lg xl:text-xl 2xl:text-3xl">
                    Information about the assistance is important,
                    we follow the process of rating the attention for each
                    individual case.
                </p>
            </div>
            <div className="flex px-2 md:flex-col">
                <TeamCommunication className="h-36 w-36 shrink-0 md:self-center lg:h-48 lg:w-48 xl:h-52 xl:w-52 2xl:h-72 2xl:w-72" />
                <p className="pt-6 ps-3 text-sm text-gray-700 md:text-center md:text-md lg:text-lg xl:text-xl 2xl:text-3xl">
                    Complaint assistance communication
                    occurs in real time to quickly answer the customer
                    needs.
                </p>
            </div>
        </div>
    )
}

export default MainBanner;
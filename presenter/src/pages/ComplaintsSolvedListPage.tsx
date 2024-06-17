import { Link, useLoaderData, useParams } from "react-router-dom";
import { Complaint } from "../lib/types";
import StarIcon from "../components/icons/StarIcon";
import ComplaintCard from "../components/complaint/ComplaintCard";
import PrimaryButton from "../components/buttons/PrimaryButton";

function ComplaintsSolvedListPage() {
    const data = useLoaderData() as Complaint[];
    const { id } = useParams();
    return (
        <div className=" bg-white w-full shadow rounded-md p-5 border-gray-100">
            <ul>
                {
                    data.length > 0 ? data.map((complaint) => (
                        <li className="flex flex-col" key={complaint.id}>
                            <ComplaintCard
                                reason={complaint.message.title}
                                description={complaint.message.description}
                                body={complaint.message.body}
                                status={complaint.status}
                                createdAt={complaint.createdAt}
                            />
                            <div className="p-5 border border-gray-100 shadow">
                                <p>{complaint.authorFullName} rated the attention received:</p>
                                <div className="flex mb-3">
                                    <p className="pr-2">Rating: </p>
                                    {
                                        [...Array(5)].map((_, index) => {
                                            index += 1;
                                            return <span
                                                key={index}>
                                                <StarIcon index={index} rating={complaint.rating!.rate} hover={0} fill={""} />
                                            </span>
                                        })
                                    }
                                </div>
                                <p className="mb-3">{" " + complaint.rating!.comment}</p>
                            </div>
                            <div className="pt-3"> 
                                <p className="mb-3 text-gray-700"> This complaint has solved after {complaint.replies!.length} replies </p>
                            </div>
                            <Link to={`/enterprises/${id}/feedback/${complaint.id}`} className="self-end p-3">
                                <PrimaryButton text="Review" />
                            </Link>
                        </li>
                    )) : <p>No complaints solved</p>
                }
            </ul>
        </div>
    )
}

export default ComplaintsSolvedListPage;

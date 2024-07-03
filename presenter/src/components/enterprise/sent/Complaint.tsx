import ComplaintCard from "../../complaint/ComplaintCard";
import ProfileCard from "../../complaint/ProfileCard";
import { Route } from "../../../routes/$enterpriseID/sent/$complaintId";
import ComplaintTabs from "../../complaint/ComplaintTabs";


function Complaint() {
    const { complaint } = Route.useLoaderData();
    const { enterpriseID, complaintId } = Route.useParams()
    return (
        <div className="flex flex-col relative min-h-[315px] md:min-h-[460px]">
            <div className="flex flex-col w-full items-center pb-2 md:pb-6">
                <ProfileCard
                    userFullName={complaint.receiverFullName}
                    profileIMG={complaint.receiverProfileIMG}
                    sub={complaint.industry || ''}
                />
            </div>
            <div className="w-full bg-white ">
                <ComplaintTabs
                    selected="complaint"
                    complaintLink={`/${enterpriseID}/sent/${complaintId}`}
                    chatLink={`/${enterpriseID}/sent/${complaintId}/chat`} />
                <div className="border-t border-gray-200">
                    <div className=" bg-white rounded-lg">
                        <ComplaintCard
                            reason={complaint.message.title}
                            description={complaint.message.description}
                            body={complaint.message.body}
                            status={complaint.status}
                            createdAt={complaint.createdAt} />
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Complaint;
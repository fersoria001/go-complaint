import { useCallback, useState } from "react";
import ComplaintCard from "../../complaint/ComplaintCard";
import ProfileCard from "../../complaint/ProfileCard";
import RateReviewIcon from "../../icons/RateReviewIcon";
import { Route } from "../../../routes/$enterpriseID/inbox/$complaintId";
import ComplaintTabs from "../../complaint/ComplaintTabs";
import SendToReview from "../../complaint/SendToReview";
import { useRouter } from "@tanstack/react-router";
import { Mutation, MarkAsReviewableMutation } from "../../../lib/mutations";


function Complaint() {
    const { complaint } = Route.useLoaderData();
    const { enterpriseID, complaintId } = Route.useParams()
    const [showConfirm, setShowConfirm] = useState(false);
    const router = useRouter();
    const sendToReview = useCallback(async () => {
        try {
            const data = await Mutation<string>(
                MarkAsReviewableMutation,
                complaint.id);
            return data;
        } catch (error) {
            router.navigate({ to: `/errors` });
            return false;
        }
    }, [complaint.id, router])

    const succesCleanUp = useCallback(() => {
        setShowConfirm(false);
        router.navigate({ to: `/${enterpriseID}/inbox/${complaintId}` });
        return Promise.resolve(true);
    }, [enterpriseID, complaintId, router])
    const simpleCleanUp = useCallback(() => {
        setShowConfirm(false);
        return Promise.resolve(true);
    }, [])
    return (
        <div className="flex flex-col relative min-h-[315px] md:min-h-[460px]">
            <div className="flex flex-col w-full items-center pb-2 md:pb-6">
                <ProfileCard
                    userFullName={complaint.authorFullName}
                    profileIMG={complaint.authorProfileIMG}
                    sub={complaint.industry || ''}
                />
                {
                    complaint.status === "IN_DISCUSSION" &&
                    <span
                        onMouseUp={() => setShowConfirm(true)}
                        className="flex rounded-md shadow mt-2 p-2
                         hover:bg-gray-100 cursor-pointer">
                        <RateReviewIcon  />
                        <p>Ask for review</p>
                    </span>
                }
                {showConfirm && <SendToReview confirm={sendToReview}
                    successCleanUp={succesCleanUp}
                    cleanUp={simpleCleanUp}
                />}
            </div>
            <div className="w-full bg-white ">
                <ComplaintTabs
                    selected="complaint"
                    complaintLink={`/${enterpriseID}/inbox/${complaintId}`}
                    chatLink={`/${enterpriseID}/inbox/${complaintId}/chat`} />
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
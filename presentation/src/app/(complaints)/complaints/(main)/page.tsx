import ComplaintsMain from "@/components/complaints/ComplaintsMain";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const Complaints: React.FC = async () => {
    const alias = cookies().get("alias")
    if (!alias) {
        redirect("/")
    }
    return (
        <ComplaintsMain />
    )
}
export default Complaints;
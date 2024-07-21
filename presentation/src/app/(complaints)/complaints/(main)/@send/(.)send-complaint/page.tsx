import PageProps from "@/app/pageProps"
import Complain from "@/components/complaints/Complain"
import DescribeComplaint from "@/components/complaints/DescribeComplaint"
import FindReceiver from "@/components/complaints/FindReceiver"
import Modal from "@/components/modal/Modal"
import { redirect } from "next/navigation"

const SendComplaintPage: React.FC<PageProps> = ({ searchParams }: PageProps) => {
    const modalClassName = "absolute flex flex-col p-2 bg-white border-t shadow-md rounded-md w-full h-screen top-20 mt-[1px] md:h-[390px] md:w-[320px] md:right-0 md:mr-12 md:bottom-0 md:inset-y-52"
    if (!searchParams?.step) {
        redirect('/complaints')
    }
    if (searchParams.step === "1") {
        return <Modal className={modalClassName}>
            <FindReceiver />
        </Modal>
    }
    if (searchParams.step === "2") {
        return <Modal className={modalClassName}>
            <DescribeComplaint />
        </Modal>
    }
    if (searchParams.step === "3") {
        return <Modal className={modalClassName}>
            <Complain />
        </Modal>
    }
}
export default SendComplaintPage
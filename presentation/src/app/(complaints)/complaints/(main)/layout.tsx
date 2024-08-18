import AddIcon from "@/components/icons/AddIcon";
import Link from "next/link";
interface Props  {
    send: React.ReactNode
    children: React.ReactNode
}
const Layout: React.FC<Props> = async ({ send, children }: Props) => {
    return (
        <>
            {children}
            <Link href={"/complaints/send-complaint?step=1"}>
                <AddIcon className="fill-gray-700 fixed bottom-11 right-8" />
            </Link>
            {send}
        </>
    )
}
export default Layout;
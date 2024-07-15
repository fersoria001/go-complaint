import Link from "next/link";
import NavbarDropdown from "./NavbarDropdown";
import Notifications from "../notifications/Notifications";
import UserDescriptorType from "@/lib/types/userDescriptorType";
import NotificationType from "@/lib/types/notificationType";
interface Props {
    user: UserDescriptorType | null
    notifications: NotificationType[]
}
const Navbar: React.FC<Props> = ({ user, notifications }: Props) => {
    return (
        <header
            className="flex absolute top-0 z-2 min-h-[82px] w-full bg-white border-b border-gray-200">
            <div className="flex justify-between w-full self-center">
                <Link href="/" className="self-center ps-5 whitespace-nowrap font-bold text-xl sm:text-2xl md:text-3xl">
                    Go Complaint
                </Link>
                <>
                    {
                        user ?
                            <div className="flex">
                                <Notifications notifications={notifications} />
                                <NavbarDropdown user={user} />
                            </div>
                            :
                            <ul className="flex pe-3 gap-3">
                                <li className="mt-auto text-sm text-gray-700 font-medium hover:text-blue-500">
                                    <Link
                                        href="/sign-in">
                                        Sign in
                                    </Link>
                                </li>
                                <li className="mt-auto text-sm  text-white
                                font-medium hover:bg-blue-600 bg-blue-500
                                 rounded-lg  px-2">
                                    <Link
                                        href="/sign-up">
                                        Sign up
                                    </Link>
                                </li>
                            </ul>
                    }
                </>
            </div>

        </header>
    )
}


export default Navbar;
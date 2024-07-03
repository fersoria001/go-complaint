import { Link } from "@tanstack/react-router";
import { UserDescriptor } from "../../lib/types";
import { useRef, useState } from "react";
import useOutsideDenier from "../../lib/hooks/useOutsideDenier";
import { Logout } from "../../lib/logout";

interface Props {
    descriptor: UserDescriptor
}
function NavbarDropdown({ descriptor }: Props) {
    const [isOpen, setIsOpen] = useState(false)
    const wrapperRef = useRef<HTMLDivElement>(null);
    useOutsideDenier(wrapperRef, (value: boolean) => { setIsOpen(value) });
    return <div className="flex items-center ms-3" ref={wrapperRef}>
        <div>
            <button
                onClick={() => { setIsOpen(!isOpen) }}
                type="button"
                className="flex text-sm bg-gray-800 rounded-full 
                focus:ring-4 focus:ring-gray-300"
                aria-expanded="false"
                data-dropdown-toggle="dropdown-user">
                <span className="sr-only">Open user menu</span>
                <img
                    src={descriptor.profileIMG}
                    className="w-8 h-8 rounded-full"
                    alt="user photo" />
            </button>
            {isOpen &&
                <div
                    className="absolute right-0 top-[65px] z-50  my-4 text-base list-none bg-white divide-y divide-gray-100 rounded shadow"

                    id="dropdown-user">
                    <div className="px-4 py-3" role="none">
                        <p className="text-sm text-gray-900 " role="none">
                            {descriptor.fullName}
                        </p>
                        <p className="text-sm font-medium text-gray-900 truncate " role="none">
                            {descriptor.email}
                        </p>
                    </div>
                    <ul className="py-1" role="none">
                        <li>
                            <Link
                                onClick={() => { setIsOpen(false) }}
                                to="/profile"
                                className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 ">Dashboard</Link>
                        </li>
                        <li>
                            <Link 
                            onClick={() => { setIsOpen(false) }}
                            to="/settings" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 " role="menuitem">Settings</Link>
                        </li>
                        <li>
                            <a href="/" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 "
                                onClick={() => Logout()}
                                role="menuitem">Sign out</a>
                        </li>
                    </ul>
                </div>
            }
        </div>
    </div>
}

export default NavbarDropdown;
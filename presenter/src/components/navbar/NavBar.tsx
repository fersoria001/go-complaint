import { useRef, useState } from "react";
import { Link, useLoaderData } from "react-router-dom";
import { UserDescriptor, UserNotifications } from "../../lib/types";
import useOutsideDenier from "../../lib/hooks/useOutsideDenier";
import { Logout } from "../../lib/actions";

import UserNotificationsLog from "../profile/UserNotificationsLog";


function NavBar() {
    const { user, notifications } = useLoaderData() as {
        user: UserDescriptor | null;
        notifications: UserNotifications | null;
    };
    const signedIn = user;
    const [navBarOpen, setNavBarOpen] = useState(false)
    const wrapperRef = useRef(null);
    useOutsideDenier(wrapperRef, setNavBarOpen);
    return (
        <nav className="absolute top-0 z-50 w-full bg-white border-b border-gray-200">
            <div className="px-3 py-3 lg:px-5 lg:pl-3">
                <div className="flex items-center justify-between">
                    <div className="flex items-center justify-start rtl:justify-end">

                        <Link to="/" className="flex ms-2 md:me-24">
                            <img src="https://flowbite.com/docs/images/logo.svg" className="h-8 me-3" alt="FlowBite Logo" />
                            <span className="self-center text-xl font-semibold sm:text-2xl whitespace-nowrap dark:text-white">Flowbite</span>
                        </Link>
                    </div>
                    <div className="flex items-center">
                        {
                            notifications && <UserNotificationsLog notifications={notifications} />
                        }
                        {signedIn && (<div className="flex items-center ms-3">
                            <div>
                                <button
                                    onClick={() => setNavBarOpen(!navBarOpen)}
                                    type="button"
                                    className="flex text-sm bg-gray-800 rounded-full focus:ring-4 focus:ring-gray-300 "
                                    aria-expanded="false"
                                    data-dropdown-toggle="dropdown-user">
                                    <span className="sr-only">Open user menu</span>
                                    <img className="w-8 h-8 rounded-full"
                                        src={
                                            signedIn.profileIMG === "default.jpg" ?
                                                "/default.jpg" : signedIn.profileIMG
                                        } alt="user photo" />
                                </button>
                            </div>
                            <div className={
                                navBarOpen ?
                                    `absolute right-0 top-10 z-50  my-4 text-base list-none bg-white divide-y divide-gray-100 rounded shadow`
                                    :
                                    `z-50 hidden my-4 text-base list-none bg-white divide-y divide-gray-100 rounded shadow`
                            }
                                ref={wrapperRef}
                                id="dropdown-user">
                                <div className="px-4 py-3" role="none">
                                    <p className="text-sm text-gray-900 " role="none">
                                        {signedIn.fullName}
                                    </p>
                                    <p className="text-sm font-medium text-gray-900 truncate " role="none">
                                        {signedIn.email}
                                    </p>
                                </div>
                                <ul className="py-1" role="none">
                                    <li>
                                        <Link
                                            to={"/profile"}
                                            className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 " role="menuitem">Dashboard</Link>
                                    </li>
                                    <li>
                                        <a href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 " role="menuitem">Settings</a>
                                    </li>
                                    <li>
                                        <a href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 " role="menuitem">Earnings</a>
                                    </li>
                                    <li>
                                        <a href="/" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 "
                                            onClick={() => Logout()}
                                            role="menuitem">Sign out</a>
                                    </li>
                                </ul>
                            </div>
                        </div>) || (
                                <div className="flex items-center ms-3">
                                    <Link to="/sign-in" className="text-sm text-gray-800 hover:text-gray-900 ms-3">Sign in</Link>
                                    <Link to="/sign-up" className="text-sm text-gray-800 hover:text-gray-900 ms-3">Sign up</Link>
                                </div>
                            )}
                    </div>
                </div>
            </div>
        </nav>
    );
}

export default NavBar;
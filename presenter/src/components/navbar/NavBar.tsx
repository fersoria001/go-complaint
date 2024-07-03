import { useContext, useEffect, useState } from "react";
import { Link, useMatches } from "@tanstack/react-router";
import { Route } from "../../routes/__root";
import BurguerMenuIcon from "../icons/BurguerMenuIcon";
import { SideBarContext } from "../../react-context/SideBarContext";
import NavbarDropdown from "./NavbarDropdown";
import NotificationList from "../profile/NotificationList";



function NavBar() {
    const { sideBarOpen, setSideBarOpen } = useContext(SideBarContext)
    const [exists, setExists] = useState(true)
    const { descriptor } = Route.useLoaderData()
    const matches = useMatches()
    const handleSideBar = () => {
        setSideBarOpen(!sideBarOpen)
    }

    useEffect(() => {
        if (matches.find((m) =>
            m.id === '/' ||
            m.id === '/sign-up' ||
            m.id === '/sign-in' ||
            m.id === '/about' ||
            m.id === '/privacy' ||
            m.id === '/licensing' ||
            m.id === '/contact'
        )) {
            setExists(false)
        } else {
            setExists(true)
        }
    }, [matches])


    return (
        <nav>
            <div className="px-3 py-3 lg:px-5 lg:pl-3">
                <div className="flex items-center justify-between">
                    <div className="flex items-center justify-start rtl:justify-end">
                        {exists && <button
                            onClick={handleSideBar}
                            data-drawer-target="logo-sidebar"
                            data-drawer-toggle="logo-sidebar"
                            aria-controls="logo-sidebar"
                            type="button"
                            className="inline-flex items-center p-2 text-sm text-gray-500
                        rounded-lg sm:hidden hover:bg-gray-100 focus:outline-none
                              focus:ring-2 focus:ring-gray-200">
                            <span className="sr-only">Open sidebar</span>
                            <BurguerMenuIcon />
                        </button>}
                        <Link to="/" className="flex ms-2 md:me-24">
                            <span
                                className="self-center text-xl font-semibold sm:text-2xl whitespace-nowrap">
                                Go-Complaint
                            </span>
                        </Link>
                    </div>
                    <div className="flex items-center">
                        {
                            descriptor ?
                                <div className="flex">
                                    <NotificationList />
                                    <NavbarDropdown descriptor={descriptor} />
                                </div>
                                :
                                <div className="flex items-center ms-3">
                                    <Link to="/sign-up" className="text-xs md:text-sm text-gray-500 font-medium hover:text-gray-900 ms-3">Sign up</Link>
                                    <Link to="/sign-in" className="text-xs md:text-sm text-gray-500 font-medium hover:text-gray-900 ms-3">Sign in</Link>
                                </div>
                        }
                    </div>
                </div>
            </div>
        </nav >
    );
}

export default NavBar;
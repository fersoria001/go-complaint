import { Outlet } from "@tanstack/react-router";

import { useContext } from "react";
import { SideBarContext } from "./react-context/SideBarContext";
import SideBar from "./components/sidebar/SideBar";
import { Route } from "./routes/_profile";

function ProfileLayout() {
    const { descriptor } = Route.useLoaderData()
    const { sideBarOpen } = useContext(SideBarContext)
    return (
        <>
            <SideBar descriptor={descriptor} />
            <div className={sideBarOpen ? `p-4 sm:ml-64 blur-md` : `p-4 sm:ml-64`}>
                <div className="p-4 rounded-lg">
                    <Outlet />
                </div>
            </div>
        </>
    )
}

export default ProfileLayout;
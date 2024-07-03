import { Outlet } from "@tanstack/react-router";

import { useContext, useEffect } from "react";
import { SideBarContext } from "./react-context/SideBarContext";
import SideBar from "./components/sidebar/SideBar";
import { Route } from "./routes/$enterpriseID";
import SideChat from "./components/enterprise/chat/SideChat";
import useWindowDimensions from "./lib/hooks/useWindowsDimensions";

function EnterpriseLayout() {
    const { enterprise, role, descriptor } = Route.useLoaderData()
    const { sideBarOpen, rightBarOpen } = useContext(SideBarContext)
    const { width } = useWindowDimensions()
    useEffect(() => {
        if (rightBarOpen && width < 768) {
            document.body.classList.add("overflow-y-hidden")
        } else {
            document.body.classList.remove("overflow-y-hidden")
        }
    }, [rightBarOpen, width])
    return (
        <div id="enterprise-layout" className="relative">
            <SideBar descriptor={descriptor} isEnterprise={true} enterpriseName={enterprise?.name} role={role} />
            <div className={`${sideBarOpen ? 'p-4 sm:ml-64 blur-md' : 'p-4 sm:ml-64'} `}>
                <div className="p-4 rounded-lg">
                    <Outlet />
                </div>
            </div>
            <SideChat />
        </div>
    )
}

export default EnterpriseLayout;
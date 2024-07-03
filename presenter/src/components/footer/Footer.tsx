import { useContext, useEffect, useState } from "react";
import { SideBarContext } from "../../react-context/SideBarContext";
import { Link, useMatches } from "@tanstack/react-router";

function Footer() {
    const { sideBarOpen } = useContext(SideBarContext)
    const matches = useMatches()
    const [haveSidebar, setHaveSidebar] = useState(false)

    useEffect(() => {
        if (matches.find((m) => m.id === '/_profile' || m.routeId.includes("/$enterpriseID"))) {
            setHaveSidebar(true)
        } else {
            setHaveSidebar(false)
        }
    }
        , [matches])
    const attrelemnt = {
        "xmlns:cc": "http://creativecommons.org/ns#",
        "xmlns:dct": "http://purl.org/dc/terms/"
    }
    return (

        < footer className={
            haveSidebar ?
                sideBarOpen ?
                    ` mt-full  sm:ml-64 bg-white rounded-lg   blur-md` :
                    ` mt-full  sm:ml-64  bg-white rounded-lg  ` :
                `   bg-white   rounded-lg `
        }>
            <div className="mx-auto p-4 md:flex md:items-center md:justify-between">

                <div {...attrelemnt} className="flex flex-col md:flex-row text-xs md:text-sm pl-2 text-gray-600">
                    <div className="flex">
                        <p>© 2024</p>
                        <p>
                            <a className="px-1" property="dct:title" rel="cc:attributionURL" href="https://www.go-complaint.com">
                                Go-Complaint</a>by
                            <a className="pl-1 pr-1"
                                rel="cc:attributionURL dct:creator" property="cc:attributionName" href="https://www.go-complaint.com">
                                Fernando Agustín Soria</a>
                        </p>
                    </div>
                    <div className="flex">
                        is licensed under
                        <a className="pl-1 pr-1" href="https://creativecommons.org/licenses/by-nc-nd/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer">
                            CC BY-NC-ND 4.0
                        </a>
                        <div className="flex">
                            <img
                                className="pl-1 w-5 h-5"
                                src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1" alt="" />
                            <img className="pl-1 w-5 h-5"
                                src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1" alt="" />
                            <img className="pl-1 w-5 h-5"
                                src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1" alt="" />
                            <img className="pl-1 w-5 h-5"
                                src="https://mirrors.creativecommons.org/presskit/icons/nd.svg?ref=chooser-v1" alt="" />
                        </div>
                    </div>

                </div>

                <ul className="flex flex-wrap items-center mt-3 text-sm font-medium text-gray-500 sm:mt-0">
                    <li>
                        <Link to={'/about'} className="hover:underline me-4 md:me-6">About</Link>
                    </li>
                    <li>
                        <Link to={'/privacy'} href="#" className="hover:underline me-4 md:me-6">Privacy Policy</Link>
                    </li>
                    <li>
                        <Link to={'/licensing'} className="hover:underline me-4 md:me-6">Licensing</Link>
                    </li>
                    <li>
                        <Link  to={'/contact'} className="hover:underline">Contact</Link>
                    </li>
                </ul>
            </div>
        </footer >

    );
}

export default Footer;


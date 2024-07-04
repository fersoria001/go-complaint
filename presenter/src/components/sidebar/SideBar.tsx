/* eslint-disable react-hooks/exhaustive-deps */
import { useCallback, useContext, useEffect, useState } from "react";
import { SideBarContext } from "../../react-context/SideBarContext";
import useWindowDimensions from "../../lib/hooks/useWindowsDimensions";
import SideBarOptions from "./SideBarOptions";
import { enterpriseOptions, profileOptions } from "./side_bar_options";
import { SideBarOptionsType, UserDescriptor } from "../../lib/types";

interface Props {
    descriptor: UserDescriptor;
    role?: string;
    isEnterprise?: boolean;
    enterpriseName?: string;
}


const SideBar: React.FC<Props> = ({ descriptor, role = "", isEnterprise = false, enterpriseName = "" }: Props) => {
    const { setSideBarOpen, sideBarOpen, reload, setReload } = useContext(SideBarContext)
    const [options, setOptions] = useState<SideBarOptionsType[]>([]);
    const { width } = useWindowDimensions();

    useEffect(() => {
        if (width >= 768) {
            setSideBarOpen(false);
        }
    });
    const closeSideBar = useCallback(() => {
        setSideBarOpen(false);
    }, [setSideBarOpen]);
    useEffect(() => {
        const load = async () => {
            const options = isEnterprise ? await enterpriseOptions(descriptor, enterpriseName) : await profileOptions(descriptor.email, descriptor.email);
            setOptions(options);
        }
        load()
    }, [])
    useEffect(() => {
        if (reload) {
            const load = async () => {
                const options = isEnterprise ? await enterpriseOptions(descriptor, enterpriseName) : await profileOptions(descriptor.email, descriptor.email);
                setOptions(options);
            }
            load()
            setReload(false);
        }
    }, [reload])
    const handleReload = () => {
        setReload(true);
    }
    return (
        <aside id="logo-sidebar"
            className={
                sideBarOpen ?
                    `fixed top-0 left-0 z-40 w-64 h-full pt-20 transition-transform -translate-x-0 duration-200 ease-in-out bg-white border-r border-gray-200 ` :
                    `fixed top-0 left-0 z-40 w-64 h-full pt-20 transition-transform -translate-x-full bg-white border-r border-gray-200 sm:translate-x-0`
            }
            aria-label="Sidebar">
            <div className="flex flex-col h-full px-3 pb-4 overflow-y-auto bg-white ">
                {
                    isEnterprise &&
                    <div className="self-center">
                        <div className="flex-flex-col text-gray-700">
                            <p className="text-xl underline underline-offset-8 mb-2">{enterpriseName}</p>
                            <p className="text-sm text-center">{role}</p>
                        </div>
                    </div>
                }
                <ul className="space-y-2 font-medium">
                    <SideBarOptions callback={handleReload} closeSideBar={closeSideBar} options={options} />
                </ul>
            </div>
        </aside >
    )
}

export default SideBar;
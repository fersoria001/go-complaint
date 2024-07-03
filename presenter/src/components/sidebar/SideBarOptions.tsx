import { Link } from "@tanstack/react-router";
import { SideBarOptionsType } from "../../lib/types";

interface Props {
    options: SideBarOptionsType[]
    closeSideBar: () => void;
    callback: () => void;
}

function SideBarOptions({ options,callback, closeSideBar }: Props) {
    const handleClick = () => {
        callback();
        closeSideBar();
    }
    return (
        <>
            {
                options.map((option, index) => {
                    if (option.unread) {
                        return (<li key={index}>
                            <Link
                                onClick={handleClick}
                                to={option.link} className="flex items-center p-2 text-gray-900 rounded-lg  hover:bg-gray-100 group">
                                {option.icon}
                                <span className="flex-1 ms-3 whitespace-nowrap">{option.title}</span>
                                <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800
                             bg-blue-100 rounded-full">
                                    {option.unread}
                                </span>
                            </Link>
                        </li>)
                    } else {
                        return (<li key={index}>
                            <Link
                                onClick={handleClick}
                                to={option.link} className="flex items-center p-2 text-gray-900 rounded-lg  hover:bg-gray-100 group">
                                {option.icon}
                                <span className="flex-1 ms-3 whitespace-nowrap">{option.title}</span>
                            </Link>
                        </li>)
                    }
                })
            }
        </>
    )
}

export default SideBarOptions;
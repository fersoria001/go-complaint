'use client'

import { useState } from "react";
import KeyboardArrowDownIcon from "../icons/KeyboardArrowDownIcon";
import KeyboardArrowUpIcon from "../icons/KeyboardArrowUpIcon";

const OfficesList: React.FC = () => {
    const [show, setShow] = useState<boolean>(true)
    return (
        <div className="min-h-20">
            <div className="flex items-center">
                <div className="flex w-full bg-gray-200 h-0.5"></div>
                <h3 className="text-gray-700 text-md lg:text-xl whitespace-nowrap px-2.5 font-bold cursor-default">Enterprises you have join to work with</h3>
                <div className="flex w-1/6 bg-gray-200 h-0.5"></div>
                {
                    show ?
                        <span className="py-3">
                            <KeyboardArrowUpIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700" />
                        </span>
                        :
                        <span className="py-3">
                            <KeyboardArrowDownIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700" />
                        </span>
                }
            </div>
            {
                show &&
                <div className="px-8">
                    <p className="text-gray-700 text-sm lg:text-md text-end">You have not joined any enterprise to work with yet.</p>
                </div>
            }
        </div>
    )
}
export default OfficesList;
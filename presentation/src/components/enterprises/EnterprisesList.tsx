'use client'

import { useState } from "react";
import KeyboardArrowDownIcon from "../icons/KeyboardArrowDownIcon";
import { EnterpriseByAuthenticatedUser } from "@/gql/graphql";
import EnterpriseItem from "./EnterpriseItem";
import KeyboardArrowRightIcon from "../icons/KeyboardArrowRightIcon";

interface Props {
    enterprises: EnterpriseByAuthenticatedUser[]
}

const EnterprisesList: React.FC<Props> = ({ enterprises }: Props) => {
    const [show, setShow] = useState<boolean>(true)
    return (
        <div className="min-h-20">
            <div className="flex items-center">
                <div className="flex w-full bg-gray-200 h-0.5"></div>
                <h3 className="text-gray-700 text-md lg:text-xl whitespace-nowrap px-2.5 font-bold cursor-default">
                    Enterprises you have registered
                </h3>
                <div className="flex w-1/6 shrink bg-gray-200 h-0.5 ps-10 md:ps-[5.5rem] lg:ps-[7rem] xl:ps-[9.1rem]"></div>
                {
                    show ?
                        <span className="py-3">
                            <KeyboardArrowDownIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700 cursor-pointer" />
                        </span>
                        :
                        <span className="py-3">
                            <KeyboardArrowRightIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700 cursor-pointer" />
                        </span>
                }
            </div>
            {
                show && enterprises.length > 0 && enterprises.map((enterprise) => {
                    return (
                        <div
                            key={enterprise.enterprise?.name} className="px-2.5" >
                            <EnterpriseItem data={enterprise} />
                        </div>
                    )
                })
                || show && enterprises.length <= 0 &&
                <div className="px-8">
                    <p className="text-gray-700 text-sm lg:text-md text-end">You have not registered any enterprise yet.</p>
                </div>
            }
        </div>
    )
}
export default EnterprisesList;
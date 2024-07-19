'use client'

import { EnterpriseByAuthenticatedUser } from "@/gql/graphql";
import ContactMailIcon from "../icons/ContactMailIcon";
import ContactPhoneIcon from "../icons/ContactPhoneIcon";
import ContactWebsiteIcon from "../icons/ContactWebsiteIcon";
import { useState } from "react";
import Link from "next/dist/client/link";
import ArrowLeftIcon from "../icons/KeyboardArrowLeftIcon";
import KeyboardArrowLeftIcon from "../icons/KeyboardArrowLeftIcon";

interface Props {
    data: EnterpriseByAuthenticatedUser
}

const EnterpriseItem: React.FC<Props> = ({ data }: Props) => {
    const [clicked, setClicked] = useState<string | undefined>(undefined)
    if (clicked) {
        return (
            <div>
                <h3
                    className="text-gray-700 text-md lg:text-xl whitespace-nowrap px-2.5 font-bold cursor-default mb-4">
                    {clicked}
                </h3>
                <div className="flex justify-around">
                    <Link
                        className="text-gray-700 text-md underline mb-4"
                        href={`/enterprises/${data.enterprise?.name}/employees`}>
                        Manage the enterprise employees.
                    </Link>
                    <Link
                        className="text-gray-700 text-md underline mb-4"
                        href={`/settings?enterpriseId=${data.enterprise?.name}`}>
                        Change your enterprise settings.
                    </Link>
                </div>
                <KeyboardArrowLeftIcon
                    onClick={() => setClicked(undefined)}
                    className="fill-white bg-blue-300 rounded-full cursor-pointer" />
            </div>
        )
    }
    return (
        <div
            onClick={() => setClicked(data.enterprise?.name)}
            className={"bg-white hover:bg-gray-50 border border-gray-200 rounded-lg shadow"}>
            <div className="flex flex-col py-3 md:p-5">
                <div className="flex self-start flex-col md:flex-row  md:justify-between w-full px-5">
                    <h5 className="pb-2 mb-2 text-2xl  text-center  font-bold tracking-tight text-gray-900">
                        {data.enterprise?.name}
                    </h5>
                    <div className="">
                        <div className="flex mb-3">
                            <ContactMailIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700">{data.enterprise?.email}</p>
                        </div>
                        <div className="flex mb-3">
                            <ContactPhoneIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700">
                                {data.enterprise?.phoneNumber}
                            </p>
                        </div>
                        <div className="flex mb-3">
                            <ContactWebsiteIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700">
                                {data.enterprise?.website}
                            </p>
                        </div>
                    </div>
                </div>
                <div className="flex flex-col self-center">
                    <p className="self-center mb-3 font-normal text-gray-700 underline underline-offset-8">
                        {data.enterprise?.industry}
                    </p>
                    <p className="mb-3 font-normal text-gray-700 text-center">
                        {data.enterprise?.address.country}, {data.enterprise?.address.countryState}, {data.enterprise?.address.city}.
                    </p>
                </div>
            </div>
        </div>
    )
}
export default EnterpriseItem;
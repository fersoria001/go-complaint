'use client'
import Link from "next/link";
import NavbarDropdown from "./NavbarDropdown";
import Notifications from "../notifications/Notifications";
import { useSuspenseQuery } from "@tanstack/react-query";
import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import { profileOptions } from "@/lib/profileOptions";
import { useState } from "react";
import BurguerMenuIcon from "../icons/BurguerMenuIcon";

const Navbar: React.FC = () => {
    const { data: user } = useSuspenseQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => {
            try {
                return await getGraphQLClient().request(userDescriptorQuery)
            } catch (e: any) {
                return null
            }
        },
        staleTime: Infinity,
        gcTime: Infinity
    })
    const notifications: any = []
    const [show, setShow] = useState<boolean>(false)
    return (
        <>
            <header
                className="flex absolute top-0 z-2 min-h-[82px] w-full bg-white border-b border-gray-200 ">
                <div className="flex w-full self-center">
                    <Link href="/" className="self-center ps-5 xl:px-5 whitespace-nowrap font-bold text-xl sm:text-2xl md:text-3xl">
                        Go Complaint
                    </Link>
                    <>
                        {
                            user ?
                                <div className="flex ms-auto">
                                    <ul className="self-end items-center h-full hidden lg:flex">
                                        {
                                            profileOptions(user?.userDescriptor.id, "").map((option) => {
                                                return (
                                                    <li
                                                        className="mt-2 text-gray-700 text-md font-bold px-2 hover:text-blue-300"
                                                        key={option.title}>
                                                        <Link href={option.link}>
                                                            {option.title}
                                                        </Link>
                                                    </li>
                                                )
                                            })
                                        }
                                    </ul>
                                    <div className="flex">
                                        <span className="p-3 lg:hidden">
                                            <BurguerMenuIcon
                                                className="block"
                                                height={32}
                                                width={32}
                                                fill="#5f6368"
                                                onClick={() => { setShow(!show) }} />
                                        </span>
                                        <Notifications notifications={notifications} />
                                        <NavbarDropdown user={user?.userDescriptor} />
                                    </div>
                                </div>
                                :
                                <ul className="flex pe-3 gap-3 ms-auto">
                                    <li className="mt-auto text-sm text-gray-700 font-medium hover:text-blue-500">
                                        <Link
                                            href="/sign-in">
                                            Sign in
                                        </Link>
                                    </li>
                                    <li className="mt-auto text-sm  text-white font-medium hover:bg-blue-600 bg-blue-500 rounded-lg px-2">
                                        <Link
                                            href="/sign-up">
                                            Sign up
                                        </Link>
                                    </li>
                                </ul>
                        }
                    </>
                </div>
            </header>
            {
                user && show &&
                <ul className="absolute w-full top-20 shadow-md z-30 bg-white divide-y divide-gray-300 py-4 rounded-md">
                    {
                        profileOptions(user?.userDescriptor.id, "").map((option) => {
                            return (
                                <li key={option.title} className="py-4 rounded-b-md text-center font-bold text-gray-500 text-md md:text-xl">
                                    <Link href={option.link}>
                                        {option.title}
                                    </Link>
                                </li>
                            )
                        })
                    }
                </ul>
            }
        </>
    )
}


export default Navbar;
"use client";
import Link from 'next/dist/client/link';
import Image from 'next/image';
import { useRef, useState } from 'react';
import UserDescriptor from '../../lib/types/userDescriptorType';
import useClickOutside from '../../lib/hooks/useClickOutside';
import { logout } from '@/lib/actions/authentication';
interface NavbarDropdownProps {
    user: UserDescriptor;
}
const NavbarDropdown: React.FC<NavbarDropdownProps> = ({ user }: NavbarDropdownProps) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const dropdownRef = useRef<HTMLDivElement>(null)
    useClickOutside(dropdownRef, () => { setIsOpen(false) })
    return (
        <div  className="flex items-center px-3 z-20" ref={dropdownRef}>
            <div>
                <button
                    onClick={() => { setIsOpen(!isOpen) }}
                    type="button"
                    className="flex border rounded-full 
                    focus:ring-4 focus:ring-gray-300"
                    aria-expanded="false"
                    data-dropdown-toggle="dropdown-user">
                    <span className="sr-only">Open user menu</span>
                    <div className='relative w-8 h-8'>
                        <Image
                            src={user.profileImg}
                            className="rounded-full"
                            sizes='32px'
                            fill
                            alt="user photo" />
                    </div>
                </button>
                {isOpen &&
                    <div
                        className="absolute right-0 top-[82px] z-2  bg-white divide-y
                                 divide-gray-100 border-b border-r border-l rounded-md"
                        id="dropdown-user">
                        <div className="px-4 py-3 cursor-default">
                            <p className="text-sm sm:text-md  text-gray-900">
                                {user.fullName}
                            </p>
                            <p className="text-sm sm:text-md font-medium text-gray-900">
                                {user.email}
                            </p>
                        </div>
                        <ul className="py-1">
                            <li>
                                <Link
                                    onClick={() => { setIsOpen(false) }}
                                    href="/profile"
                                    className="block px-4 py-2 text-sm sm:text-md text-gray-700 hover:bg-gray-100">
                                    Dashboard
                                </Link>
                            </li>
                            <li>
                                <Link
                                    onClick={() => { setIsOpen(false) }}
                                    href="/settings"
                                    className="block px-4 py-2 text-sm sm:text-md text-gray-700 hover:bg-gray-100">
                                    Settings
                                </Link>
                            </li>
                            <li>
                                <button
                                    type="button"
                                    className="w-full block px-4 py-2 text-sm sm:text-md text-gray-700"
                                    onClick={() => logout()}>
                                    <p className='hover:text-blue-500'>Sign out</p>
                                </button>
                            </li>
                        </ul>
                    </div>
                }
            </div>
        </div>
    )
}

export default NavbarDropdown;
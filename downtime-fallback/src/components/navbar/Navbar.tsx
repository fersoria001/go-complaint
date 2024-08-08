import Link from "next/link";

const Navbar: React.FC = () => {
    return (
            <header className="flex absolute top-0 z-2 min-h-[82px] w-full bg-white border-b border-gray-200 ">
                <div className="flex w-full self-center">
                    <Link href="/" className="self-center ps-5 xl:px-5 whitespace-nowrap font-bold text-md sm:text-lg md:text-xl">
                        Go Complaint
                    </Link>
                </div>
            </header>
    )
}


export default Navbar;
'use client'
import CloseIcon from "../icons/CloseIcon";
import SearchIcon from "../icons/SearchIcon";
import Link from "next/link";

const FindReceiver: React.FC = () => {
    return (
        <div className="h-full relative flex flex-col py-0.5">
            <form role="search" className="flex items-center max-w-sm mx-auto md:mx-0 gap-2" >
                <label htmlFor="search-receivers" className="sr-only">Search</label>
                <div className="relative w-full">
                    <div
                        className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                        <SearchIcon />
                    </div>
                    <input
                        type="text"
                        id="search-receivers"
                        name="term"
                        //value={ }
                        // onChange={(event) => { }}
                        className="appearance-none focus:outline-none bg-gray-50 border border-gray-300
                         text-gray-700 text-sm rounded-lg block w-full ps-10 p-2.5"
                        placeholder="Send to..."
                        required />
                </div>
            </form>
            <div className="flex flex-col h-80 px-2.5 overflow-y-auto">
                List
                {/* {searching ? <span className="self-center pt-8"><LoadingSpinner /></span> :
                                <ul className="flex flex-col w-full h-auto cursor-pointer">
                                    {results.length > 0 && results.map((receiver) => (
                                        <li
                                            onClick={() => handleSelection(receiver)}
                                            className="flex items-center pb-2 hover:bg-gray-100"
                                            key={receiver.fullName}
                                        >
                                            <img src={receiver.thumbnail} alt={receiver.fullName} className="w-10 h-10 rounded-full" />
                                            <span className="ms-2">{receiver.fullName}</span>
                                        </li>
                                    ))}
                                </ul>
                            } */
                }
            </div>
            <Link
                href="/complaints/send-complaint?step=2"
                className="px-7 py-3 bg-blue-500 hover:bg-blue-600 rounded-md text-white font-bold self-center absolute bottom-28 md:bottom-[1.275rem]">
                Next
            </Link>
        </div>
    )
}
export default FindReceiver;
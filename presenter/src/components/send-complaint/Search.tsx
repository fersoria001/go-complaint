import { useEffect, useState } from "react";
import { Receiver } from "../../lib/types";
import LoadingSpinner from "../icons/LoadingSpinner";
import { Route } from "../../routes/_profile/_send-complaint/send-complaint";
import { Query, FindComplaintReceiversQuery, FindComplaintReceiversTypeList } from "../../lib/queries";
interface Props {
    callback: (receiver: Receiver) => void;
}
function Search({ callback }: Props) {
    const { descriptor, receivers } = Route.useLoaderData();
    const [query, setQuery] = useState<string>("");
    const [searching, setSearching] = useState<boolean>(false);
    const [results, setResults] = useState<Receiver[]>(receivers);
    useEffect(() => {
        if (query.length >= 1) {
            setSearching(true);
        }
        async function search() {
            const results = await Query<Receiver[]>(
                FindComplaintReceiversQuery,
                FindComplaintReceiversTypeList,
                [descriptor.email, query]
            );
            setResults(results);
        }
        search();
        setSearching(false);
    }, [descriptor.email, query])
    const handleSelection = (receiver: Receiver) => {
        document.getElementById("search-receivers")?.setAttribute("value", receiver.fullName);
        setQuery(receiver.fullName);
        callback(receiver);
    }
    const handleInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        document.getElementById("search-receivers")?.setAttribute("value", event.currentTarget.value);
        setQuery(event.currentTarget.value);
        callback({} as Receiver);
    }

    return (
        <>
            <form role="search" className="flex items-center max-w-sm mx-auto" >
                <label htmlFor="search-receivers" className="sr-only">Search</label>
                <div className="relative w-full">
                    <div
                        className="absolute inset-y-0 start-0 flex items-center
                     ps-3 pointer-events-none">
                        <svg
                            className="w-4 h-4"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 20 20">
                            <path
                                stroke="currentColor"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth="2"
                                d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
                        </svg>
                    </div>
                    <input
                        type="text"
                        id="search-receivers"
                        name="term"
                        value={query}
                        onChange={(event) => handleInput(event)}
                        className="bg-gray-50 border
                         border-gray-300 text-gray-900 text-sm 
             rounded-lg focus:ring-blue-500 focus:border-blue-500 
             block w-full ps-10 p-2.5 "
                        placeholder="Send to..."
                        required />
                </div>
            </form>
            <div className="flex flex-col h-36 md:h-56 lg:h-96 px-2 md:px-6 overflow-y-auto  ">
                {searching ? <span className="self-center pt-8"><LoadingSpinner /></span> :
                    <ul className="flex flex-col w-full h-auto cursor-pointer">
                        {results.map((receiver) => (
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
                }
            </div>
        </>
    )
}

export default Search
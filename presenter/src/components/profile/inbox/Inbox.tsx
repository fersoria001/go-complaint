import { useRef, useState } from "react";
import { Link, useRouter } from "@tanstack/react-router";
import { ComplaintType } from "../../../lib/types";
import { Route } from "../../../routes/_profile/inbox";
import useOutsideDenier from "../../../lib/hooks/useOutsideDenier";


function Inbox() {
    const { sent, search } = Route.useLoaderData();
    const [showDropdown, setShowDropdown] = useState(false)
    const [page, setPage] = useState<number>(search.page);
    const [date, setDate] = useState<string>(search.filter.date);
    const [query, setQuery] = useState<string>(search.filter.query);
    const ref = useRef(null)
    const router = useRouter();
    const numberOfPages = Math.floor(sent.count / sent.currentLimit) + 1;
    const handlePrevious = () => {
        if (page === 1) return
        setPage(page - 1)
        Pagination(page - 1)

    }
    const handleNext = () => {
        if (page === numberOfPages) return
        setPage(page + 1)
        Pagination(page + 1)
    }
    const Pagination = (nextPage: number) => {
        router.navigate({
            to: `/inbox`,
            search: {
                page: nextPage,
                filter: {
                    query: query,
                    date: date,
                }
            }
        })
    }
    const Search = () => {
        setPage(1)
        router.navigate({
            to: `/inbox`,
            search: {
                page: 1,
                filter: {
                    query: query,
                    date: date,
                }
            }
        })
    }
    const handleDaysAgo = (nextDate: string) => {
        setDate(nextDate)
        setPage(1)
        setShowDropdown(false)
        router.navigate({
            to: `/inbox`,
            search: {
                page: 1,
                filter: {
                    query: query,
                    date: nextDate,
                }
            }
        })
    }
    useOutsideDenier(ref, () => {
        setShowDropdown(false)
    })
    return (
        <>
            <div className="flex flex-column sm:flex-row flex-wrap space-y-4 sm:space-y-0 items-center justify-between pb-4">
                <div className="relative" ref={ref}>
                    <button
                        id="dropdownRadioButton"
                        onClick={() => setShowDropdown(!showDropdown)}
                        className="inline-flex items-center text-gray-500 bg-white border border-gray-300 focus:outline-none
       hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-3 py-1.5"
                        type="button">
                        <svg className="w-3 h-3 text-gray-500 dark:text-gray-400 me-3" aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M10 0a10 10 0 1 0 10 10A10.011 10.011 0 0 0 10 0Zm3.982 13.982a1 1 0 0 1-1.414 0l-3.274-3.274A1.012 1.012 0 0 1 9 10V6a1 1 0 0 1 2 0v3.586l2.982 2.982a1 1 0 0 1 0 1.414Z" />
                        </svg>
                        <p>{date}</p>
                        <svg className="w-2.5 h-2.5 ms-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 10 6">
                            <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m1 1 4 4 4-4" />
                        </svg>
                    </button>
                    {showDropdown &&
                        <div
                            id="dropdownRadio"
                            className="absolute z-10 w-48 bg-white divide-y divide-gray-100 rounded-lg shadow">
                            <ul className="p-3 space-y-1 text-sm text-gray-700" aria-labelledby="dropdownRadioButton">
                                <li onMouseUp={() => handleDaysAgo("Last day")}>
                                    <div className="flex items-center p-2 rounded hover:bg-gray-100">
                                        <input id="filter-radio-example-1" type="radio" value=""
                                            name="filter-radio" className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500" />
                                        <label htmlFor="filter-radio-example-1"
                                            className="w-full ms-2 text-sm font-medium text-gray-900 rounded">Last day</label>
                                    </div>
                                </li>
                                <li onMouseUp={() => handleDaysAgo("Last 7 days")}>
                                    <div className="flex items-center p-2 rounded hover:bg-gray-100">
                                        <input id="filter-radio-example-2" type="radio" value="" name="filter-radio"
                                            className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500" />
                                        <label htmlFor="filter-radio-example-2" className="w-full ms-2 text-sm font-medium text-gray-900 rounded">
                                            Last 7 days
                                        </label>
                                    </div>
                                </li>
                                <li onMouseUp={() => handleDaysAgo("Last 30 days")}>
                                    <div className="flex items-center p-2 rounded hover:bg-gray-100">
                                        <input id="filter-radio-example-3" type="radio" value="" name="filter-radio"
                                            className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500" />
                                        <label htmlFor="filter-radio-example-3"
                                            className="w-full ms-2 text-sm font-medium text-gray-900 rounded">
                                            Last 30 days
                                        </label>
                                    </div>
                                </li>
                                <li onMouseUp={() => handleDaysAgo("Last month")}>
                                    <div className="flex items-center p-2 rounded hover:bg-gray-100">
                                        <input id="filter-radio-example-4" type="radio" value="" name="filter-radio"
                                            className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500" />
                                        <label htmlFor="filter-radio-example-4" className="w-full ms-2 text-sm font-medium text-gray-900 rounded">
                                            Last month
                                        </label>
                                    </div>
                                </li>
                                <li onMouseUp={() => handleDaysAgo("Last year")}>
                                    <div className="flex items-center p-2 rounded hover:bg-gray-100">
                                        <input id="filter-radio-example-5" type="radio" value="" name="filter-radio"
                                            className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500" />
                                        <label htmlFor="filter-radio-example-5" className="w-full ms-2 text-sm font-medium text-gray-900 rounded">
                                            Last year
                                        </label>
                                    </div>
                                </li>
                            </ul>
                        </div>}
                </div>
                <label htmlFor="table-search" className="sr-only">Search
                </label>
                <div className="relative flex">
                    <div className="absolute inset-y-0 left-0 rtl:inset-r-0 rtl:right-0 flex items-center ps-3 pointer-events-none">
                        <svg className="w-5 h-5 text-gray-500 dark:text-gray-400" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clipRule="evenodd"></path></svg>
                    </div>
                    <input
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        type="text"
                        id="table-search"
                        className="block p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg w-80
                                bg-gray-50 focus:ring-blue-500 focus:border-blue-500"
                        placeholder="" />
                    <button type="button" onClick={Search} className="text-white font-medium rounded-md ms-2 px-4 py-1 bg-group bg-gradient-to-br from-cyan-500 to-blue-500"> search </button>
                </div>
            </div>
            <div className="relative flex flex-col min-h-screen justify-between  overflow-x-auto shadow-md sm:rounded-lg cursor-default">
                <table
                    className="w-full text-sm text-left rtl:text-right text-gray-500 ">
                    <thead
                        className="text-xs text-gray-700 uppercase bg-gray-50">
                        <tr>
                            <th scope="col" className="p-4">
                                <div className="flex items-center">
                                    <input
                                        id="checkbox-all-search"
                                        type="checkbox"
                                        className="w-4 h-4 text-blue-600 bg-gray-10
                                         border-gray-300 rounded focus:ring-blue-500 " />
                                    <label
                                        htmlFor="checkbox-all-search"
                                        className="sr-only">
                                        checkbox
                                    </label>
                                </div>
                            </th>
                            <th scope="col" className="px-6 py-3">
                                From
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Reason
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Status
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Date
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Unreads
                            </th>
                        </tr>
                    </thead>
                    <tbody >
                        {sent && sent.complaints.map((complaint: ComplaintType) => {
                            return (<tr key={complaint.id} className="bg-white border-b hover:bg-gray-50">
                                <td className="w-4 p-4">
                                    <div className="flex items-center">
                                        <input
                                            id="checkbox-table-search-1"
                                            type="checkbox"
                                            className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded
                                      focus:ring-blue-500 focus:ring-2" />
                                        <label htmlFor="checkbox-table-search-1" className="sr-only">
                                            checkbox
                                        </label>
                                    </div>
                                </td>
                                <th scope="row" >
                                    <Link to={`/inbox/${complaint.id}`} className="flex align-center  whitespace-nowrap px-6 py-4">
                                        <img className="w-10 h-10 rounded-full" src={complaint.authorProfileIMG} alt="avatar" />
                                        <div className="ps-3 text-center my-auto font-medium text-gray-900" >
                                            {complaint.authorFullName}
                                        </div>
                                    </Link>
                                </th>
                                <td className="px-6 py-4">
                                    {complaint.message.title}
                                </td>
                                <td className="px-6 py-4">
                                    {complaint.status}
                                </td>
                                <td className="px-6 py-4">
                                    {new Date(parseInt(complaint.createdAt)).toLocaleDateString()}
                                </td>
                                <td
                                    className="px-6 py-4">
                                    {complaint.replies!.filter(reply => reply.senderID != complaint.authorID && !reply.read).length}
                                </td>
                            </tr>
                            )
                        })}

                    </tbody>
                </table>
                <nav
                    className="flex items-center flex-column flex-wrap md:flex-row justify-between px-2 py-4"
                    aria-label="Table navigation">
                    <span
                        className="text-sm font-normal text-gray-500  mb-4 md:mb-0 block w-full md:inline md:w-auto">
                        {"Showing \t"}
                        <span className="font-semibold text-gray-900">{`${sent.currentOffset}-${sent.currentLimit + sent.currentOffset > sent.count ? sent.currentOffset + sent.count - sent.currentOffset : sent.currentLimit} \t`}</span>
                        {"of \t"}
                        <span className="font-semibold text-gray-900">{`${sent.count} \t`}</span>
                    </span>
                    <form role="search">
                        <ul className="inline-flex -space-x-px rtl:space-x-reverse text-sm h-8">
                            <li>
                                <span
                                    onClick={handlePrevious}
                                    className="flex items-center justify-center px-3 h-8 ms-0 leading-tight
                             text-gray-500 bg-white border border-gray-300 rounded-s-lg
                              hover:bg-gray-100 hover:text-gray-700 ">
                                    Previous
                                </span>
                            </li>
                            {Array.from({ length: numberOfPages }, (_, i) => {
                                return (
                                    <li key={i}>
                                        <input
                                            name="page"
                                            value={(i + 1).toString()}
                                            onClick={(event) => {
                                                const p = event.currentTarget.value
                                                setPage(parseInt(p))
                                                Pagination(parseInt(p))
                                            }}
                                            className={i + 1 === page ?
                                                `cursor-default max-w-10 flex items-center justify-center px-3 h-8 text-blue-600 border border-gray-300
                             bg-blue-50` :
                                                `max-w-10 flex items-center justify-center px-3 h-8 leading-tight
                             text-gray-500 bg-white border border-gray-300 hover:bg-gray-100
                              hover:text-gray-700 cursor-pointer`
                                            }
                                            readOnly />

                                    </li>
                                )
                            })}
                            <li>
                                <span
                                    onClick={handleNext}
                                    className="flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border
                              border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700">
                                    Next
                                </span>
                            </li>
                        </ul>
                    </form>
                </nav>
            </div >
        </>
    );

}

export default Inbox;

import { Link, useLoaderData, useParams } from "react-router-dom";
import UserForHiring from "../components/hiring/UserForHiring";
import { useState } from "react";
import useUsersForHiring from "../lib/hooks/useUsersForHiring";

function Hiring() {
    const pages = useLoaderData() as string;
    const [search, setSearch] = useState("")
    const [page, setPage] = useState("1")
    const { id } = useParams();
    const list = useUsersForHiring(id!, page, search)
    const handlePrevious = () => {
        if (parseInt(page) === 1) return
        setPage((parseInt(page) - 1).toString())
    }
    const handleNext = () => {
        if (parseInt(page) === parseInt(pages)) return
        setPage((parseInt(page) + 1).toString())
    }
    return (
        <div className="pt-4">
            <div className="min-h-80 md:min-h-screen">
                <div className="px-2.5">
                    <input
                        className="bg-gray-50 border border-gray-300 text-gray-900 text-sm 
               rounded-lg focus:ring-blue-500 focus:border-blue-500 
               block w-full  p-2.5 mb-2"
                        type="text" placeholder="Search for an user..." onChange={e => setSearch(e.target.value)} />
                </div>
                <ul>
                    {
                        list?.users.map((user) => {
                            return <li key={user.email}>
                                <Link to={`/enterprises/${id}/hiring/${user.email}`}>
                                    <UserForHiring user={user} />
                                </Link>
                            </li>
                        })
                    }
                </ul>
            </div>
            <nav
                className="flex items-center flex-column flex-wrap md:flex-row justify-between py-4"
                aria-label="Table navigation">
                <span
                    className="text-sm font-normal text-gray-500  mb-4 md:mb-0 block w-full md:inline md:w-auto">
                    {"Showing \t"}
                    <span className="font-semibold text-gray-900">{`1-${list?.users.length} \t`}</span>
                    {"of \t"}
                    <span className="font-semibold text-gray-900">{`${list?.count} \t`}</span>
                </span>
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
                    {Array.from({ length: parseInt(pages) }, (_, i) => {
                        return (
                            <li key={i}>
                                <input
                                    name="page"
                                    value={(i + 1).toString()}
                                    onClick={(event) => {
                                        setPage(event.currentTarget.value)
                                    }}
                                    className={i + 1 === parseInt(page) ?
                                        `max-w-10 flex items-center justify-center px-3 h-8 text-blue-600 border border-gray-300
                             bg-blue-50 hover:bg-blue-100 hover:text-blue-700` :
                                        `max-w-10 flex items-center justify-center px-3 h-8 leading-tight
                             text-gray-500 bg-white border border-gray-300 hover:bg-gray-100
                              hover:text-gray-700`
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

            </nav>
        </div>
    )
}

export default Hiring;
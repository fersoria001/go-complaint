import UserForHiring from "../components/hiring/UserForHiring";
import { useState } from "react";
import { Link } from "@tanstack/react-router";
import { Route } from "../routes/$enterpriseID/hiring";

function Hiring() {
    const { search, usersForHiring } = Route.useLoaderData()
    const { enterpriseID } = Route.useParams()
    const navigate = Route.useNavigate()
    const [query, setQuery] = useState(search.filter.query)
    const [page, setPage] = useState<number>(search.filter.page)
    const pages = Math.ceil(usersForHiring.count / 10)
    const handlePrevious = () => {
        if (page === 1) return
        Pagination(page - 1)
    }
    const handleNext = () => {
        if (page === pages) return
        Pagination(page + 1)
    }
    const Search = () => {
        navigate({ search: { filter: { query, page: 1 } } })
    }
    const Pagination = (nextPage: number) => {
        setPage(nextPage)
        navigate({ search: { filter: { query, page: nextPage } } })
    }
    return (
        <div className="pt-4">
            <div className="min-h-80 md:min-h-screen">
                <div className="flex px-2.5 mb-2">
                    <input className="bg-gray-50 border border-gray-300 text-gray-900 text-sm 
               rounded-lg focus:ring-blue-500 focus:border-blue-500 
               block w-full  p-2.5"
                        type="text" placeholder="Search for an user..." onChange={e => setQuery(e.target.value)} />
                    <button type="button" onClick={Search} className="text-white font-medium rounded-md ms-2 px-4 py-1 bg-group bg-gradient-to-br from-cyan-500 to-blue-500"> search </button>
                </div>
                <ul>
                    {
                        usersForHiring.users.map((user) => {
                            return <li key={user.email}>
                                <Link to={`/${enterpriseID}/hire`} search={{ id: { email: user.email } }}>
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
                    <span className="font-semibold text-gray-900">{`1-${usersForHiring.users.length} \t`}</span>
                    {"of \t"}
                    <span className="font-semibold text-gray-900">{`${usersForHiring.count} \t`}</span>
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
                    {Array.from({ length: pages }, (_, i) => {
                        return (
                            <li key={i}>
                                <input
                                    name="page"
                                    value={(i + 1).toString()}
                                    onClick={(event) => {
                                        Pagination(parseInt(event.currentTarget.value))
                                    }}
                                    className={i + 1 === page ?
                                        `max-w-10 flex items-center justify-center px-3 
                                        h-8 text-blue-600 border border-gray-300
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
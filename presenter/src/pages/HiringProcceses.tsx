import { useState } from "react";
import HiringProccess from "../components/enterprise/HiringProccess";
import { Route } from "../routes/$enterpriseID/hiring-procceses";

function HiringProcceses() {
    const { hiringProcceses, search } = Route.useLoaderData();
    console.log(hiringProcceses)
    const navigate = Route.useNavigate();
    const [page, setPage] = useState<number>(search.filter.page)
    const [query, setQuery] = useState(search.filter.query)
    const pages = Math.ceil(hiringProcceses.count / 5) > 0 ? Math.ceil(hiringProcceses.count / 5) : 1;
    const Search = () => {
        navigate({ search: { filter: { query: query, page: 1 } } });
    }
    const Pagination = (nextPage: number) => {
        setPage(nextPage)
        navigate({ search: { filter: { query, page: nextPage } } })
    }
    return (
        <div className="" >
            <div className="flex px-2.5 mb-2">
                <input
                    className="bg-gray-50 border border-gray-300 text-gray-900 text-sm 
               rounded-lg focus:ring-blue-500 focus:border-blue-500 
               block w-full  p-2.5"
                    type="text" placeholder="Search for an user..." onChange={e => setQuery(e.target.value)} />
                <button type="button" onClick={Search} className="text-white font-medium rounded-md ms-2 px-4 py-1 bg-group bg-gradient-to-br from-cyan-500 to-blue-500"> search </button>
            </div>
            <ul className="min-h-[315px] md:min-h-[340px]">
                {hiringProcceses.hiringProccesses.map((hire) => {
                    return (
                        <li key={hire.eventID}>
                            <HiringProccess hiringProccess={hire} />
                        </li>
                    )
                })}
            </ul>
            <nav className="flex flex-col pt-[67px]">
                <ul className="inline-flex -space-x-px self-end text-base h-10">
                    <li>
                        <a className="flex items-center justify-center px-4 h-10 ms-0 leading-tight
                         text-gray-500 bg-white border border-e-0 border-gray-300 rounded-s-lg
                          hover:bg-gray-100 hover:text-gray-700">Previous</a>
                    </li>
                    {Array.from({ length: pages }, (_, i) => {
                        return (
                            <li key={i}>
                                <input
                                    value={(i + 1).toString()}
                                    onClick={(e) => {
                                        const v = e.currentTarget.value
                                        Pagination(parseInt(v))
                                    }}
                                    className={i + 1 === page ?
                                        `cursor-default max-w-10 flex items-center justify-center px-3 h-10 text-blue-600 border border-gray-300
                     bg-blue-50` :
                                        `max-w-10 flex items-center justify-center px-3 h-10 leading-tight
                     text-gray-500 bg-white border border-gray-300 hover:bg-gray-100
                      hover:text-gray-700 cursor-pointer`
                                    } readOnly />
                            </li>)
                    })}
                    <li>
                        <a className="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border
                         border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700">Next</a>
                    </li>
                </ul>
            </nav>
        </div>
    );
}
export default HiringProcceses;
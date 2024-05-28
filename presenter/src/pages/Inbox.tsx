import { Form, useLoaderData, useLocation, useNavigate, useSubmit } from "react-router-dom";
import { Complaint, ComplaintTypeList } from "../lib/types";

function Inbox() {
    const { complaintList, page } = useLoaderData() as { complaintList: ComplaintTypeList | null, page: string };
    const inbox = complaintList;
    const submit = useSubmit();
    const navigate = useNavigate();
    const location = useLocation();
    const nextURL = location.pathname.split("/").slice(0, -1).join("/");
    if (!inbox) { return <h1>Loading...</h1>; }
    const numberOfPages = Math.floor(inbox.count / inbox.currentLimit) + 1;
    const handlePrevious = () => {
        if (parseInt(page) === 1) return
        submit({ page: (parseInt(page) - 1).toString() });
    }
    const handleNext = () => {
        if (parseInt(page) === numberOfPages) return
        submit({ page: (parseInt(page) + 1).toString() });
    }
    return (
        <>
            <div
                className="relative flex flex-col min-h-screen justify-between  overflow-x-auto shadow-md sm:rounded-lg">
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
                                Last Update
                            </th>
                        </tr>
                    </thead>
                    <tbody >
                        {inbox && inbox.complaints.map((complaint: Complaint) => {
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
                                <th
                                    onMouseUp={() => { navigate(`${nextURL}/${complaint.id}`) }}
                                    scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                                    {complaint.authorFullName}
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
                                    {new Date(parseInt(complaint.updatedAt)).toLocaleDateString()}
                                </td>
                            </tr>)
                        })}

                    </tbody>
                </table>
                <nav
                    className="flex items-center flex-column flex-wrap md:flex-row justify-between py-4"
                    aria-label="Table navigation">
                    <span
                        className="text-sm font-normal text-gray-500  mb-4 md:mb-0 block w-full md:inline md:w-auto">
                        {"Showing \t"}
                        <span className="font-semibold text-gray-900">{`1-${inbox.complaints.length} \t`}</span>
                        {"of \t"}
                        <span className="font-semibold text-gray-900">{`${inbox.count} \t`}</span>
                    </span>
                    <Form role="search">
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
                                                const formData = new FormData();
                                                formData.append("page", event.currentTarget.value);
                                                submit(formData);
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
                    </Form>
                </nav>
            </div >
        </>
    );

}

export default Inbox;

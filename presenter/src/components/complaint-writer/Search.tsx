import { useContext, useEffect, useState } from "react";
import { Form, useLoaderData, useNavigation, useSubmit } from "react-router-dom";
import { Receiver } from "../../lib/types";
import LoadingSpinner from "../icons/LoadingSpinner";
import { ComplaintContext } from "../../react-context/ComplaintContext";
interface Props {
    callbackFn: () => void;
}
function Search({ callbackFn }: Props) {
    const { complaintData, updateState } = useContext(
        ComplaintContext
    );
    const { receivers, term } = useLoaderData() as
        { receivers: Receiver[]; term: string };
    const nonMutableReceivers = [...receivers]; //copy of full list
    const [query, setQuery] = useState<string>(complaintData?.fullName || term);
    const submit = useSubmit();
    useEffect(() => {
        setQuery(term);
    }, [term]);
    const navigation = useNavigation();
    const searching = navigation.location &&
        new URLSearchParams(navigation.location.search)
            .has(
                "term"
            );

    const handleSelectItem = (receiver: Receiver) => {
        setQuery(receiver.fullName);
        const isFirstSearch = term == "";
        const formData = new FormData();
        formData.append("term", receiver.fullName);
        submit(formData, {
            replace: !isFirstSearch,
        });
        console.log("got a receiver", receiver.fullName, receiver.ID)
        updateState({
            complaintData: {
                fullName: receiver.fullName,
                senderID: complaintData?.senderID || "",
                receiverID: receiver.ID || "",
                reason: complaintData?.reason || "",
                description: complaintData?.description || "",
                body: complaintData?.body || ""
            }
        });
        console.log("complaint data after update", complaintData)
        callbackFn();
    }
    return (
        <>
            <Form role="search" className="flex items-center max-w-sm mx-auto" >
                <label htmlFor="search-receivers" className="sr-only">Search</label>
                <div className="relative w-full">
                    <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
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
                        defaultValue={query}
                        onChange={(event) => {
                            callbackFn();
                            setQuery(event.currentTarget.value);
                            const isFirstSearch = term == "";
                            submit(event.currentTarget.form, {
                                replace: !isFirstSearch,
                            });
                            updateState({
                                complaintData: {
                                    fullName: event.currentTarget.value,
                                    senderID: complaintData?.senderID || "",
                                    receiverID: nonMutableReceivers.find(
                                        (v) => v.fullName === event.currentTarget.value
                                    )?.ID || "null",
                                    reason: complaintData?.reason || "",
                                    description: complaintData?.description || "",
                                    body: complaintData?.body || ""
                                }
                            });
                        }}
                        className="bg-gray-50 border border-gray-300 text-gray-900 text-sm 
             rounded-lg focus:ring-blue-500 focus:border-blue-500 
             block w-full ps-10 p-2.5 " placeholder="Send to..."
                        required />
                </div>
            </Form>
            <div className="flex flex-col h-40 md:h-56 lg:h-96 px-2 md:px-6 overflow-y-auto  ">
                {searching ? <span className="self-center pt-8"><LoadingSpinner /></span> :
                    <ul className="flex flex-col w-full h-auto">
                        {receivers.map((receiver) => (
                            <li
                                onClick={() => handleSelectItem(receiver)}
                                className="flex items-center pb-4 first:pt-2 last:mb-0 hover:bg-gray-100"
                                key={receiver.fullName}
                            >
                                <img src={receiver.IMG} alt={receiver.fullName} className="w-10 h-10 rounded-full" />
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
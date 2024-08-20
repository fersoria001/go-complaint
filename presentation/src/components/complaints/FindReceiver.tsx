'use client'
import SearchIcon from "../icons/SearchIcon";
import recipientsByNameLikeQuery from "@/graphql/queries/recipientsByNameLikeQuery";
import { useSuspenseQuery } from "@tanstack/react-query";
import getGraphQLClient from "@/graphql/graphQLClient";
import LoadingSpinnerIcon from "../icons/LoadingSpinner";
import Image from "next/image";
import { useState } from "react";
import { Recipient } from "@/gql/graphql";
import { createNewComplaint } from "@/lib/actions/graphqlActions";
import clsx from "clsx";
import InlineAlert from "../error/InlineAlert";

const FindReceiver: React.FC = () => {
    const [term, setTerm] = useState<string>("")
    const [receiver, setReceiver] = useState<Recipient | null>(null)
    const [errors, setErrors] = useState<{ [key: string]: string }>({})
    const { data: recipients, isLoading } = useSuspenseQuery({
        queryKey: ['recipientsByNameLike', term],
        queryFn: async ({ queryKey }) => {
            try {
                return await getGraphQLClient().request(recipientsByNameLikeQuery, { term: queryKey[1] })
            } catch (e: any) {
                console.log("error: ", e)
                return null
            }
        },
    })
    const onChangeSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
        setTerm(e.currentTarget.value)
    }
    const createComplaint = async () => {
        if (!receiver || !receiver.id) {
            setErrors({ "receiver": "pick a receiver before the next step" })
            return
        }
        const id = await createNewComplaint(receiver.id)
    }
    return (
        <div className="h-full relative flex flex-col py-0.5">
            <form role="search" className="flex items-center max-w-sm mx-auto md:mx-0 gap-2" >
                <label htmlFor="search-recipients" className="sr-only">Search</label>
                <div className="relative w-full">
                    <div
                        className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                        <SearchIcon />
                    </div>
                    <input
                        type="text"
                        id="search-recipients"
                        name="term"
                        value={term}
                        onChange={onChangeSearch}
                        className="appearance-none focus:outline-none bg-gray-50 border border-gray-300
                         text-gray-700 text-sm rounded-lg block w-full ps-10 p-2.5"
                        placeholder="Send to..."
                        required />
                </div>
            </form>
            <div className="flex flex-col h-56 overflow-y-auto mb-1">
                {isLoading ? <span className="self-center pt-8"><LoadingSpinnerIcon /></span> :
                    <ul className="flex flex-col w-full h-auto cursor-pointer">
                        {recipients && recipients.recipientsByNameLike.length > 0 && recipients.recipientsByNameLike.map((recipient) => (
                            <li
                                onClick={() => { setReceiver(recipient as Recipient) }}
                                className={clsx("flex items-center pb-2 first:mt-4", {
                                    "bg-gray-100": receiver?.id === recipient.id,
                                    "hover:bg-gray-100": receiver?.id !== recipient.id
                                })}
                                key={recipient.id}
                            >
                                <div className='relative w-10 h-10 rounded-full'>
                                    <Image
                                        src={recipient.subjectThumbnail!}
                                        alt={recipient.subjectName!}
                                        className="rounded-full"
                                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                                        fill
                                    />
                                </div>
                                <span className="ms-2 text-gray-700 text-sm md:text-md xl:text-xl">{recipient.subjectName}</span>
                            </li>
                        ))}
                    </ul>

                }
            </div>
            {errors && errors.receiver && <InlineAlert 
            className={"flex items-center text-sm text-red-800 rounded-lg bg-red-50"}
            errors={[errors.receiver]} />}
            <button
                type="button"
                onClick={createComplaint}
                className="px-7 py-3 bg-blue-500 hover:bg-blue-600 rounded-md text-white font-bold self-center absolute bottom-36 md:bottom-[0rem]">
                Next
            </button>
        </div >
    )
}
export default FindReceiver;
//href="/complaints/send-complaint?step=2"
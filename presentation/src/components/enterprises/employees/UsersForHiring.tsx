'use client'

import { User } from "@/gql/graphql";
import getGraphQLClient from "@/graphql/graphQLClient";
import usersForHiringQuery from "@/graphql/queries/usersForHiringQuery";
import { useSuspenseInfiniteQuery } from "@tanstack/react-query";
import { useParams, useSearchParams } from "next/navigation";
import { useState } from "react";
import UsersForHiringItem from "./UsersForHiringItem";
import Link from "next/dist/client/link";
import clsx from "clsx";
import { useRouter } from "next/navigation";
import SearchInput from "@/components/search/SearchInput";

const UsersForHiring: React.FC = () => {
    const params = useParams()
    const searchParams = useSearchParams()
    const [page, setPage] = useState<number>(parseInt(searchParams.get("page") || "0", 10))
    const query = searchParams.get("query") || ""
    const enterpriseId = params.enterpriseId
    const gqlClient = getGraphQLClient()
    const {
        data,
        error,
        fetchPreviousPage,
        hasPreviousPage,
        fetchNextPage,
        hasNextPage,
        isFetching,
        isFetchingPreviousPage,
        isFetchingNextPage,
        status,
    } = useSuspenseInfiniteQuery({
        queryKey: ['users-for-hiring', enterpriseId, query],
        queryFn: async ({ pageParam, queryKey }) => gqlClient.request(usersForHiringQuery, {
            input: {
                id: queryKey[1] as string,
                query: queryKey[2] as string,
                limit: 10,
                offset: pageParam * 10,
            }
        }),
        initialPageParam: page,
        getNextPageParam: (lastPage: any, pages: any) => {
            if (lastPage.usersForHiring.nextCursor < 0) {
                return undefined
            }
            return lastPage.usersForHiring.nextCursor
        },
        getPreviousPageParam: (lastPage: any, pages: any) => {
            if (lastPage.usersForHiring.prevCursor < 0) {
                return undefined
            }
            return lastPage.usersForHiring.prevCursor
        },
    })
    const router = useRouter()
    const pages = Math.floor(data.pages[0].usersForHiring.count / 10)
    const onChangeSearch = (e: React.ChangeEvent<HTMLInputElement>)=>{
        if(e.currentTarget.value.length > 3) {
            router.push(`/enterprises/${enterpriseId}/employees/hire-new?query=${e.currentTarget.value}`)
        }
    }
    return (
        <div className="p-4 relative">
            <div className="min-h-screen">
                <SearchInput placeholder="Search for an user..." onChange={onChangeSearch} />
                <ul>
                    {
                        data.pages[0].usersForHiring.users.map((user: User) => {
                            return (
                                <li key={user.userName}>
                                    <Link href={`/enterprises/${enterpriseId}/employees/hire?userId=${user.userName}`}>
                                        <UsersForHiringItem user={user} />
                                    </Link>
                                </li>
                            )
                        })
                    }
                </ul>
            </div>

            <nav
                className="flex items-center justify-between"
                aria-label="Table navigation">
                <span
                    className="text-sm font-normal text-gray-700 flex items-center mb-4 md:mb-0 gap-x-1 w-full md:inline md:w-auto">
                    {"Showing \t"}
                    <span className="font-semibold text-gray-700">{`1-${data.pages[0].usersForHiring.users.length} \t`}</span>
                    {"of \t"}
                    <span className="font-semibold text-gray-700">{`${data.pages[0].usersForHiring.count} \t`}</span>
                </span>

                <ul className="inline-flex -space-x-px rtl:space-x-reverse text-sm h-8">
                    <li>
                        <button
                            type="button"
                            onClick={() => {
                                if (!isFetching) {
                                    setPage(data.pages[0].usersForHiring.prevCursor)
                                    fetchPreviousPage()
                                }
                            }}
                            disabled={!hasPreviousPage || isFetchingPreviousPage}
                            className={clsx("flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-gray-300 rounded-s-lg", {
                                'opacity-50': !hasPreviousPage,
                                'hover:bg-gray-100 hover:text-gray-700': hasPreviousPage
                            })}>
                            Previous
                        </button>
                    </li>

                    {Array.from({ length: pages }, (_, i) => {
                        return (
                            <li key={i}>
                                <input
                                    name="page"
                                    value={(i + 1).toString()}
                                    onClick={(event) => {

                                    }}
                                    className={clsx("max-w-10 flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700", {
                                        'bg-blue-50 text-blue-600 hover:text-blue-700': i + 1 === page
                                    })}
                                    readOnly />
                            </li>
                        )
                    })}
                    <li>
                        <button
                            type="button"
                            onClick={() => {
                                if (!isFetching) {
                                    setPage(data.pages[0].usersForHiring.nextCursor)
                                    fetchNextPage()
                                }
                            }}
                            disabled={!hasNextPage || isFetchingNextPage}
                            className={clsx("flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border border-gray-300 rounded-e-lg", {
                                'opacity-50': !hasNextPage,
                                'hover:bg-gray-100 hover:text-gray-700': hasNextPage
                            })}>
                            Next
                        </button>
                    </li>
                </ul>

            </nav>
        </div>
    )
}

export default UsersForHiring;
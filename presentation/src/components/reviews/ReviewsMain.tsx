'use client'

import getGraphQLClient from "@/graphql/graphQLClient"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { useQueryClient, useSuspenseQueries, useSuspenseQuery } from "@tanstack/react-query"
import FilterBy from "../search/FilterBy"
import SearchInput from "../search/SearchInput"
import Review from "./Review"
import clsx from "clsx"
import { useEffect, useState } from "react"
import { setCookie } from "@/lib/actions/cookies"
import complaintsSentForReviewByReceiverIdQuery from "@/graphql/queries/complaintsSentForReviewByReceiverIdQuery"
import complaintsRatedByReceiverIdQuery from "@/graphql/queries/complaintsRatedByReceiverIdQuery"
import complaintsRatedByAuthorIdQuery from "@/graphql/queries/complaintsRatedByAuthorIdQuery"
import pendingReviewsByAuthorIdQuery from "@/graphql/queries/pendingReviewsByAuthorIdQuery"
import { Complaint } from "@/gql/graphql"
import WaitingReview from "./WaitingReview"
import SolvedReview from "./SolvedReview"

enum ReviewSearchFilter {
    Pending = "pending",
    Solved = "solved",
    Waiting = "waiting"
}

const reviewFilterOptions = [
    {
        id: "0",
        name: "pending",
        value: ReviewSearchFilter.Pending
    },
    {
        id: "1",
        name: "solved",
        value: ReviewSearchFilter.Solved
    },
    {
        id: "2",
        name: "waiting",
        value: ReviewSearchFilter.Waiting
    },
]
const ReviewsMain: React.FC = () => {
    const { data: { userDescriptor } } = useSuspenseQuery({
        queryKey: ['userDescriptor'],
        queryFn: async () => getGraphQLClient().request(userDescriptorQuery),
        staleTime: Infinity,
        gcTime: Infinity
    },)
    const [alias, setAlias] = useState<string>(userDescriptor.id)
    const [term, setTerm] = useState<string[]>(["", "", ""])
    const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
        const v = e.currentTarget.value
        switch (filter) {
            case ReviewSearchFilter.Pending:
                setTerm(p => p.map((t, i) => i === 0 ? v : t))
                break;
            case ReviewSearchFilter.Solved:
                setTerm(p => p.map((t, i) => i === 2 ? v : t))
                break;
            case ReviewSearchFilter.Waiting:
                setTerm(p => p.map((t, i) => i === 1 ? v : t))
                break
            default:
                break;
        }
    }
    const [
        { data: { pendingReviewsByAuthorId } },
        { data: { complaintsSentForReviewByReceiverId } },
        { data: { complaintsRatedByAuthorId } },
        { data: { complaintsRatedByReceiverId } },
    ] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['pendingReviewsByAuthorIdQuery', alias, term[0]],
                queryFn: async () => getGraphQLClient()
                    .request(pendingReviewsByAuthorIdQuery, { id: alias, term: term[0] }),
            },
            {
                queryKey: ['complaintsSentForReviewByReceiverIdQuery', alias, term[1]],
                queryFn: async () => getGraphQLClient()
                    .request(complaintsSentForReviewByReceiverIdQuery, { id: alias, term: term[1] }),
            },
            {
                queryKey: ['complaintsRatedByAuthorIdQuery', alias, term[2]],
                queryFn: async () => getGraphQLClient()
                    .request(complaintsRatedByAuthorIdQuery, { id: alias, term: term[2] }),
            },
            {
                queryKey: ['complaintsRatedByReceiverIdQuery', alias, term[2]],
                queryFn: async () => getGraphQLClient()
                    .request(complaintsRatedByReceiverIdQuery, { id: alias, term: term[2] }),
            },

        ]
    })
    const handleAliasSelect = async (value: string) => {
        const ok = await setCookie("alias", value)
        if (!ok) {
            console.error("couln't set the alias")
            return
        }
        setAlias(ok)
    }
    const [filter, setFilter] = useState<ReviewSearchFilter>(reviewFilterOptions[0].value)
    //not reactive
    const toRender = (): JSX.Element[] => {
        let result: JSX.Element[] = []
        switch (filter) {
            case ReviewSearchFilter.Pending:
                result = pendingReviewsByAuthorId.map((v) =>
                    <Review
                        currentUser={userDescriptor}
                        key={v!.id}
                        item={v as Complaint}
                        alias={alias}
                    />
                )
                break;
            case ReviewSearchFilter.Solved:
                result = [...complaintsRatedByAuthorId, ...complaintsRatedByReceiverId].map((v) =>
                    <SolvedReview key={v?.id} item={v as Complaint} />
                )
                break;
            case ReviewSearchFilter.Waiting:
                result = complaintsSentForReviewByReceiverId.map((v) =>
                    <WaitingReview
                        key={v?.id}
                        item={v as Complaint}
                    />
                )
                break
            default:
                break;
        }
        return result
    }
    return (
        <div className="py-2 px-2">
            <div className="flex gap-5 overflow-x-auto w-full self-center my-2 whitespace-nowrap px-2.5">
                <button
                    type="button"
                    onClick={() => { handleAliasSelect(userDescriptor.id) }}
                    className={clsx("bg-blue-500 rounded-lg px-2.5 py-0.5 text-white font-bold", {
                        "scale 110 bg-blue-700": alias === userDescriptor.id,
                        "hover:bg-blue-600": alias != userDescriptor.id
                    })}>
                    {userDescriptor.fullName.split(" ")[0]}(you)
                </button>
                {
                    userDescriptor.authorities?.map((enterprise) => {
                        return (
                            <button
                                key={enterprise?.enterpriseId}
                                type="button"
                                onClick={() => { handleAliasSelect(enterprise?.enterpriseId!) }}
                                className={clsx("bg-blue-500 rounded-lg px-2.5 py-0.5 text-white font-bold", {
                                    "scale 110 bg-blue-700": alias === enterprise?.enterpriseId,
                                    "hover:bg-blue-600": alias != enterprise?.enterpriseId
                                })}>
                                {enterprise?.principal}
                            </button>
                        )
                    })
                }
            </div>
            <SearchInput placeholder="Search..." onChange={handleSearch} />
            <FilterBy options={reviewFilterOptions} callback={(v: string) => {
                setTerm(["", "", ""])
                setFilter(v as ReviewSearchFilter)
            }} />
            <div className="flex flex-col mt-2">
                {
                    toRender().map((c) => { return c })
                }
            </div>
        </div>
    )
}
export default ReviewsMain
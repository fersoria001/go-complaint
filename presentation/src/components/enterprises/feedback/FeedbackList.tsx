'use client'
import KeyboardArrowDownIcon from "@/components/icons/KeyboardArrowDownIcon"
import KeyboardArrowRightIcon from "@/components/icons/KeyboardArrowRightIcon"
import { Complaint } from "@/gql/graphql"
import { useParams, useSearchParams } from "next/navigation"
import { useState } from "react"
import ComplaintForFeedbackItem from "./ComplaintForFeedbackItem"
import Link from "next/link"

interface Props {
    complaints: Complaint[]
    label: string;
    fallbackLabel: string
}
const FeedbackList: React.FC<Props> = ({ complaints, label, fallbackLabel }: Props) => {
    const [show, setShow] = useState<boolean>(true)
    const { enterpriseId } = useParams()
    const enterpriseName = decodeURIComponent(enterpriseId as string)
    return (
        <div className="min-h-20">
            <div className="flex items-center">
                <div className="flex w-full bg-gray-200 h-0.5"></div>
                <h3 className="text-gray-700 text-md lg:text-xl whitespace-nowrap px-2.5 font-bold cursor-default">
                    {label}
                </h3>
                <div className="flex w-1/6 shrink bg-gray-200 h-0.5 ps-10 md:ps-[5.5rem] lg:ps-[7rem] xl:ps-[9.1rem]"></div>
                {
                    show ?
                        <span className="py-3">
                            <KeyboardArrowDownIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700 cursor-pointer" />
                        </span>
                        :
                        <span className="py-3">
                            <KeyboardArrowRightIcon
                                onClick={() => setShow(!show)}
                                className="shrink-0 fill-gray-700 cursor-pointer" />
                        </span>
                }
            </div>
            {
                show && complaints.length > 0 && complaints.map((complaint) => {
                        return (
                            <Link
                                href={`/enterprises/${enterpriseName}/employees/feedback/${complaint.id}`}
                                key={complaint?.id}
                                className="px-2.5" >
                                <ComplaintForFeedbackItem enterpriseName={enterpriseName} item={complaint} />
                            </Link>
                        )
                })
                || show && complaints.length <= 0 &&
                <div className="px-8">
                    <p className="text-gray-700 text-sm lg:text-md text-end">{fallbackLabel}</p>
                </div>
            }
        </div>
    )
}

export default FeedbackList
'use client'
import { getComplaintsInfo } from "@/lib/actions/graphqlActions";
import { useQuery } from "@tanstack/react-query";

interface Props {
    id: string
}

const ComplaintsInfo: React.FC<Props> = ({ id }: Props) => {
    const { isError, isPending, isFetching, data, error } = useQuery({
        queryKey: ['info', id],
        queryFn: ({ queryKey }) => getComplaintsInfo(queryKey[1]),
    })
    if (isPending) {
        return <div>Loading...</div>
    }
    if (isError) {
        return <div>{error.message}</div>
    }
    return (
        <div>
            <div>Received: {data.received}</div>
            <div>Pending: {data.pending}</div>
            <div>Resolved: {data.resolved}</div>
            <div>Reviewed: {data.reviewed}</div>
            <div>Average Rating: {data.avgRating}</div>
            <div>Total: {data.total}</div>
        </div>
    )
}
export default ComplaintsInfo;
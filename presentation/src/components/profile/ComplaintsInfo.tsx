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

        </div>
    )
}
export default ComplaintsInfo;
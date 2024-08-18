import { HiringProccessStatus } from "@/gql/graphql";

export const mayHaveReason = [
    HiringProccessStatus.Rejected,
    HiringProccessStatus.Canceled,
    HiringProccessStatus.Fired,
    HiringProccessStatus.Leaved
]
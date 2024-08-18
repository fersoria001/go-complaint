import { HiringProccessStatus } from "@/gql/graphql";

function colorSwitch(status: HiringProccessStatus): string {
    let color = "";
    switch (status) {
        case HiringProccessStatus.Pending:
            color = "#99CCFF";
            break;
        case HiringProccessStatus.Accepted:
            color = "#99CCFF";
            break;
        case HiringProccessStatus.Rejected:
            color = "#fde68a";
            break;
        case HiringProccessStatus.Canceled:
            color = "#FFCCCC";
            break;
        case HiringProccessStatus.UserAccepted:
            color = "#99FFCC";
            break;
        case HiringProccessStatus.Hired:
            color = "#dbeafe";
            break;
        case HiringProccessStatus.Leaved:
            color = "#fde68a";
            break;
        case HiringProccessStatus.Fired:
            color = "#FFCCCC";
            break;
    }
    return color;
}
export default colorSwitch;
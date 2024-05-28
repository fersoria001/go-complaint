import { SendComplaint, ComplaintState } from "../../lib/types";

export const defaultComplaintState: ComplaintState = {
    complaintData: {} as SendComplaint,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    updateState: (newState?: Partial<ComplaintState>) => { }
}
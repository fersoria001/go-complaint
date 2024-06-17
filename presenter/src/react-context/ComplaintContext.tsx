import React from "react"
import { ComplaintState, SendComplaint } from "../lib/types"

const defaultComplaintState: ComplaintState = {
    complaintData: {} as SendComplaint,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    updateState: (newState?: Partial<ComplaintState>) => { }
}
export const ComplaintContext = React.createContext<ComplaintState>(
    defaultComplaintState
)



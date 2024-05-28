import React from "react"
import { ComplaintState } from "../../lib/types"
import { defaultComplaintState } from "./DefaultComplaintState"
export const ComplaintContext = React.createContext<ComplaintState>(defaultComplaintState)
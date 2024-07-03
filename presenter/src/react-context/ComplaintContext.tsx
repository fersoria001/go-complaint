/* eslint-disable @typescript-eslint/no-unused-vars */
import React from "react"
import { ErrorType, SendComplaintType, SendComplaintValidationSchema } from "../lib/types"
import {  syncParseSchema } from "../lib/parse_schema";

export const ComplaintContext = React.createContext({
    complaint: {
        authorID: "",
        receiverID: "",
        title: "",
        description: "",
        content: "",
    } as SendComplaintType,
    setKeyValue: (_key: keyof SendComplaintType, _value: string): ErrorType => { return {} },
})


interface Props {
    routeParams?: {
        enterpriseId: string;
    },
    children: React.ReactNode;
}

export const ComplaintContextProvider: React.FunctionComponent<Props> = (
    props: Props
): JSX.Element => {
    const [complaint, setComplaint] = React.useState({} as SendComplaintType);
    const setKeyValue = (key: keyof SendComplaintType, value: string): ErrorType => {
        complaint[key] = value;
        setComplaint({ ...complaint });
        const { errors } = syncParseSchema(complaint, SendComplaintValidationSchema)
        return errors
    }
    return (
        <ComplaintContext.Provider value={{
            complaint,
            setKeyValue,
        }}>
            {props.children}
        </ComplaintContext.Provider>
    );
};
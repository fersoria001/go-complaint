import { useState } from "react";
import { ComplaintState } from "../../lib/types";
import { ComplaintContext } from "./ComplaintContext";

interface Props {
    children: React.ReactNode;
}
export const ComplaintContextProvider: React.FunctionComponent<Props> = (
    props: Props
): JSX.Element => {
    const [state, setState] = useState({});
    const updateState = (newState: Partial<ComplaintState>) => {
        setState({ ...state, ...newState });
    };
    return (
        <ComplaintContext.Provider value={{ ...state, updateState }}>
            {props.children}
        </ComplaintContext.Provider>
    );
};
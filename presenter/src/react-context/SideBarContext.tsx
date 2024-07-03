/* eslint-disable @typescript-eslint/no-unused-vars */
import React from "react"

export const SideBarContext = React.createContext({
    reload: false,
    setReload: (_reload: boolean) => { },
    sideBarOpen: false,
    setSideBarOpen: (_open: boolean) => { },
    rightBarOpen: false,
    setRightBarOpen: (_open: boolean) => { },
})

interface Props {
    children: React.ReactNode;
}

export const SideBarContextProvider: React.FunctionComponent<Props> = (
    props: Props
): JSX.Element => {
    const [sideBarOpen, setSideBarOpen] = React.useState(false);
    const [rightBarOpen, setRightBarOpen] = React.useState(false);
    const [reload, setReload] = React.useState(false);
    return (
        <SideBarContext.Provider value={{ reload, setReload, sideBarOpen, setSideBarOpen, rightBarOpen, setRightBarOpen, }}>
            {props.children}
        </SideBarContext.Provider>
    );
};

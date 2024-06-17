import React from "react";
import { UserState } from "../lib/types";

const defaultUserState: UserState = {
    userSession: null,
    userNotifications: null,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    updateState : (newState?: Partial<UserState>) => {},
}

export const UserContext = React.createContext<UserState>(defaultUserState) 
//This will  be updated after persistence layer update
export const restNotifications = async () => {
    const response = await fetch(import.meta.env.VITE_NOTIFICATIONS_ENDPOINT + "?type=user",
        {
            method: "GET",
            credentials: "include",
        })
    const data = await response.json()
    return data
}
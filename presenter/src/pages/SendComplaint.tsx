import { Outlet } from "@tanstack/react-router";
import { ComplaintContextProvider } from "../react-context/ComplaintContext";

function SendComplaint() {
    return (
        <ComplaintContextProvider >
            <Outlet />
        </ComplaintContextProvider>
    )
}

export default SendComplaint;
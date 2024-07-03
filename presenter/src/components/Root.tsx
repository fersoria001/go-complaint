import { Outlet } from "@tanstack/react-router";
import { SideBarContextProvider } from "../react-context/SideBarContext";
import Footer from "./footer/Footer";
import NavBar from "./navbar/NavBar";


function Root() {
    return (
        <>
            <SideBarContextProvider>
                <header className="fixed top-0 z-50 w-full
                 bg-white border-b border-gray-200">
                    <NavBar />
                </header>
                <main className=" mt-14  ">
                    <Outlet />
                </main>
                <Footer />
            </SideBarContextProvider>
        </>
    );
}
export default Root;


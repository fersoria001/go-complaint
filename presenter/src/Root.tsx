import { Outlet } from "react-router-dom"
import Footer from "./components/footer/Footer"
import NavBar from "./components/navbar/NavBar"
function Root() {

  return (
    <>
      <header className="py-12">
        <NavBar />
      </header>
      <main className="min-h-screen">
        <Outlet />
      </main>
      <Footer />
    </>
  )
}

export default Root

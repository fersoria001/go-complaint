
import ReactDOM from "react-dom/client";
import { routeTree } from './routeTree.gen'
import { createRouter, parseSearchWith, RouterProvider, stringifySearchWith } from "@tanstack/react-router";
import "./index.css";
import { fetchUserDescriptor } from "./lib/fetchUserDescriptor";
import { fetchNotifications } from "./lib/fetchNotifications";
import NotFoundPage from "./pages/NotFoundPage";
import { hasPermission, isLoggedIn } from "./lib/is_logged_in";
const router = createRouter({
  routeTree,
  parseSearch: parseSearchWith((value) => JSON.parse(decodeFromBinary(value))),
  stringifySearch: stringifySearchWith((value) =>
    encodeToBinary(JSON.stringify(value)),
  ),
  context: {
    hasPermission,
    fetchUserDescriptor,
    fetchNotifications,
    isLoggedIn, 
  },
  defaultNotFoundComponent: NotFoundPage,
})
function decodeFromBinary(str: string): string {
  return decodeURIComponent(
    Array.prototype.map
      .call(atob(str), function (c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)
      })
      .join(''),
  )
}

function encodeToBinary(str: string): string {
  return btoa(
    encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, function (_, p1) {
      return String.fromCharCode(parseInt(p1, 16))
    }),
  )
}
// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}
// Render the app
const rootElement = document.getElementById('root')!
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement)
  root.render(
    //<StrictMode>
      <RouterProvider router={router} />
    //</StrictMode>,
  )
}



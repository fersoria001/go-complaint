import type { Metadata } from "next";
import { Raleway } from "next/font/google";
import "./globals.css";
import Footer from "@/components/footer/Footer";
import Navbar from "@/components/navbar/Navbar";
import Providers from "./providers";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { cookies } from "next/headers";
import getGraphQLClient from "@/graphql/graphQLClient";
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery";
import SideChat from "@/components/enterprises/chat/SideChat";
import { getCookie } from "@/lib/actions/cookies";

const raleway = Raleway({ subsets: ["latin"] });
export const metadata: Metadata = {
  title: "Go Complaint",
  description: "A site designed to send complaints to different users and enterprises.",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const jwtCookie = cookies().get("jwt")
  const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
  const gqlClient = getGraphQLClient()
  gqlClient.setHeader("Cookie", strCookie)
  const queryClient = new QueryClient()
  await queryClient.prefetchQuery({
    queryKey: ['userDescriptor'],
    queryFn: async () => {
      try {
        return await gqlClient.request(userDescriptorQuery)
      } catch (e: any) {
        console.log("error: ", e)
        return null
      }
    },
    staleTime: Infinity,
    gcTime: Infinity
  })
  await queryClient.prefetchQuery({
    queryKey: ["server-side-jwt"],
    queryFn: async () => {
      const jwt = await getCookie("jwt")
      if (!jwt) return ""
      return jwt
    }
  })
  const isAuth = async () => {
    try {
      const ok = await gqlClient.request(userDescriptorQuery)
      if (ok.userDescriptor.authorities!.length > 0) {
        return true
      }
    } catch {
      return false
    }
  }
  return (
    <html lang="en">
      <body className={raleway.className} id="body-element">
        <Providers>
          <HydrationBoundary state={dehydrate(queryClient)} >
            <Navbar />
            <div id="scroll-top"></div>
            <main className="mt-20 pt-0.5 min-h-screen">
              {children}
            </main>
            {await isAuth() && <SideChat />}
          </HydrationBoundary>
          <Footer />
        </Providers>
      </body>
    </html>
  );
}

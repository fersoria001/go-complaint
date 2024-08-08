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
        console.log("error: ",e)
        return null
      }
    },
    staleTime: Infinity,
    gcTime: Infinity
  })

  return (
    <html lang="en">
      <body className={raleway.className}>
        <Providers>
          <HydrationBoundary state={dehydrate(queryClient)} >
            <Navbar />
            <div id="scroll-top"></div>
            <main className="mt-20 pt-0.5 min-h-screen">
              {children}
            </main>
          </HydrationBoundary>
          <Footer />
        </Providers>
      </body>
    </html>
  );
}

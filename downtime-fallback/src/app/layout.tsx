import type { Metadata } from "next";
import { Raleway } from "next/font/google";
import "./globals.css";
import Footer from "@/components/footer/Footer";
import Navbar from "@/components/navbar/Navbar";


const raleway = Raleway({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Go Complaint",
  description: "A site designed to send complaints to different users and enterprises.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={raleway.className}>
        <Navbar />
        <main className="min-h-screen pt-0.5 mt-20">
          {children}
        </main>
        <Footer />
      </body>

    </html>
  );
}

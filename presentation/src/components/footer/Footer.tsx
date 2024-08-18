import Link from "next/link";
import Image from "next/image";

const Footer: React.FC = () => {
    const attrelemnt = {
        "xmlns:cc": "http://creativecommons.org/ns#",
        "xmlns:dct": "http://purl.org/dc/terms/"
    }
    return (
        <footer id="footer-element"
         className={`absolute  w-full z-50 mt-full bg-white shadow-md border-t p-4 `}>
            <div className="md:flex md:flex-col w-full lg:flex-row lg:justify-between">
                <div {...attrelemnt} className="flex flex-col md:flex-row text-sm pl-2 text-gray-700 md:mx-auto lg:mx-0">
                    <p className="mb-0.5 lg:flex lg:items-center">
                        © 2024
                        <a className="px-1 " property="dct:title" rel="cc:attributionURL" href="https://www.go-complaint.com">
                            Go Complaint</a>by
                        <a className="pl-1 pr-1"
                            rel="cc:attributionURL dct:creator" property="cc:attributionName" href="https://www.go-complaint.com">
                            Fernando Agustín Soria</a>
                        is licensed under
                        <a
                            className="pl-1 pr-1"
                            href="https://creativecommons.org/licenses/by-nc-nd/4.0/?ref=chooser-v1"
                            target="_blank"
                            rel="license noopener noreferrer">
                            CC BY-NC-ND 4.0
                        </a>
                    </p>
                    <div className="flex mx-auto md:mx-0 align-top">
                        <Image
                            width={28} height={28} style={{ padding: "4px", height: "auto" }}
                            src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1" alt="" />
                        <Image width={28} height={28} style={{ padding: "4px", height: "auto" }}
                            src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1" alt="" />
                        <Image width={28} height={28} style={{ padding: "4px", height: "auto" }}
                            src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1" alt="" />
                        <Image width={28} height={28} style={{ padding: "4px", height: "auto" }}
                            src="https://mirrors.creativecommons.org/presskit/icons/nd.svg?ref=chooser-v1" alt="" />
                    </div>
                </div>

                <ul className="flex flex-wrap justify-center  my-3 text-sm gap-2 font-medium text-gray-700 shrink-0">
                    <li>
                        <Link href={'/about'} className="hover:underline ">About</Link>
                    </li>
                    <li>
                        <Link href={'/privacy'} className="hover:underline ">Privacy Policy</Link>
                    </li>
                    <li>
                        <Link href={'/contact'} className="hover:underline">Contact</Link>
                    </li>
                </ul>
            </div>
        </footer >
    )
}
export default Footer;
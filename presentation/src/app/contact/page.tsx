import ContactForm from "@/components/contact/ContactForm";
import Image from "next/image";
import "@/components/wave.css"
import PageProps from "../pageProps";
import ContactSuccess from "@/components/contact/ContactSuccess";

const Contact: React.FC<PageProps> = ({ searchParams }: PageProps) => {

    return (
        <div className="flex flex-col py-5 mt-2.5 max-w-lg grow shrink-0 mx-auto lg:mx-0">
            <div className="flex flex-col px-3">
                <div>
                    <div className="w-24 h-24 rounded-full mb-2">
                        <Image
                            src={'/contact.jpg'}
                            alt={'Image of Fernando Agustín Soria'}
                            width={720}
                            height={720}
                            style={{
                                width: '100%',
                                height: 'auto',
                                borderRadius: '9999px'
                            }}
                        />
                    </div>
                    <p className="text-gray-700 text-md md:text-xl mb-2 md:mb-4 font-medium">Fernando Agustín Soria</p>
                </div>
                <div>
                    <p className="text-gray-700 text-md md:text-xl mb-2 md:mb-4 font-medium">Web developer</p>
                    <div className="flex">
                        <p className="text-gray-700 text-md md:text-xl mb-2 md:mb-4 font-medium"> Number:</p>
                        <p className="pl-1 text-gray-700 text-md md:text-xl mb-2 md:mb-4 ">+54 2944 7818 23</p>
                    </div>
                    <div className="flex">
                        <p className="text-gray-700 text-md md:text-xl mb-2 md:mb-4 font-medium"> Email:</p>
                        <p className="pl-1 text-gray-700 text-md md:text-xl mb-2 md:mb-4 ">bercho001@gmail.com</p>
                    </div>
                </div>
            </div>
            {
                searchParams?.success == "1" ? <ContactSuccess /> : <ContactForm />
            }
        </div>
    )
}
export default Contact;
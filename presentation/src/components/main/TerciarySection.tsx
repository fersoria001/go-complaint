import Link from "next/link";
import "../inverted-wave.css"
import Image from "next/image";
const TerciarySection: React.FC = () => {
    return (
        <div className="pt-16">
            <div className="w-full ml-full lg:flex ">
                <div>
                    <p className="mb-4  px-5 md:text-center lg:text-start text-lg text-gray-700 md:text-xl lg:text-2xl xl:text-3xl 2xl:text-4xl ">
                        Register your enterprises and
                        invite people to work with you!
                    </p>
                    <div className="md:p-1 md:text-center lg:text-start">
                        <Image
                            src={'/main3.png'}
                            alt={'Image of Go complaint profile'}
                            sizes="100vw"
                            style={{
                                width: '100%',
                                height: 'auto',
                            }}
                            width={600}
                            height={300}
                        />
                    </div>
                </div>
                <div>
                    <p className="mb-4  px-5 md:text-center lg:text-start text-lg text-gray-700 md:text-xl lg:text-2xl xl:text-3xl 2xl:text-4xl">
                        Assign them duties and let them
                        take care of customers support.
                    </p>
                    <div className="md:p-1 md:text-center lg:text-start">
                            <Image
                                src={'/main5.png'}
                                alt={'Image of Go complaint profile'}
                                sizes="100vw"
                                style={{
                                    width: '100%',
                                    height: 'auto',
                                }}
                                width={600}
                                height={300}
                            />
                    </div>
                </div>
                <div>
                    <p className="mb-4 px-5 md:text-center lg:text-start text-lg text-gray-700 md:text-xl lg:text-2xl xl:text-3xl 2xl:text-4xl">
                        Provide them feedback so that everyone can
                        improve its attention.
                    </p>
                    <div className="md:p-1 md:text-center lg:text-start">
                            <Image
                                src={'/main4.png'}
                                alt={'Image of Go complaint profile'}
                                sizes="100vw"
                                style={{
                                    width: '100%',
                                    height: 'auto',
                                }}
                                width={600}
                                height={300}
                            />
                    </div>
                </div>
            </div>
            <div className="bg-blue-500 h-[80vh]  w-full  overflow-hidden">
                <div className="container">
                    <div className="wave">
                        <div className="flex flex-col z-100 absolute left-10 2xl:-top-20 xl:-to">
                            <h1 className="px-2 text-center mb-4 text-gray-700 font-medium text-2xl md:text-3xl lg:text-4xl xl:text-5xl 2xl:text-6xl">
                                At Go Complaint we care about customers
                            </h1>
                            <p className="px-5 text-center mb-4 text-md lg:text-lg xl:text-xl 2xl:text-2xl
                             text-gray-700">Start with us today</p>
                            <Link
                                href={"/sign-up"}
                                type="button"
                                className="mx-auto mb-4 py-3 px-5 2xl:py-5 2xl:px-7
                                 text-white text-lg lg:text-xl xl:text-2xl 2xl:text-3xl rounded-md bg-blue-500 hover:bg-blue-600">
                                Register now
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default TerciarySection;

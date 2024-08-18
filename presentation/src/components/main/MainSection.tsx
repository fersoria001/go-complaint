import Image from "next/image";

const MainSection: React.FC = () => {
    return (
        <div className="lg:flex overflow-hidden relative h-screen">
            <article className="pt-8 lg:pt-16 mb-4 2xl:mb-8 lg:w-1/3 lg:ps-5 flex flex-col">
                <h1 className="px-2 text-center mb-4 2xl:mb-8 text-gray-700 font-bold
                 text-2xl md:text-3xl lg:text-4xl xl:text-5xl 2xl:text-6xl lg:text-start">
                    Go Complaint
                </h1>
                <h2 className="px-2 text-center font-semibold text-gray-700 text-lg mb-4 2xl:mb-8 lg:text-start lg:text-xl xl:text-2xl 2xl:text-3xl ">
                    A site designed to send complaints to different users and enterprises.
                </h2>
                <p className="px-5 lg:text-start mb-4 2xl:mb-8 text-md lg:text-lg xl:text-xl 2xl:text-3xl text-gray-700">
                    Get ready to receive complaints, here it&apos;s the place
                    to keep track of them. From the open complaint to the history,
                    being an individual or an enterprise, guide clients trough the proccess
                    and analyze it later.
                </p>
                <button
                    type="button"
                    className="mx-auto mb-4 2xl:mb-8 py-3 px-5 text-white text-lg lg:text-xl xl:py-5 xl:px-7 2xl:text-3xl
                     rounded-md bg-blue-500 hover:bg-blue-600">
                    Register now
                </button>
            </article>
            <div className="w-full
            skew-y-6 
            lg:absolute 
            lg:top-20
            lg:translate-x-1/2
            xl:top-30
            xl:translate-x-1/3
            2xl:top-40">
                <Image
                    src={'/main1.png'}
                    alt={'Image of Go complaint profile'}
                    width={1360}
                    height={620}
                    style={{
                        width: '100%',
                        height: 'auto',
                        borderColor: '#3b82f6',
                        borderWidth: '1rem',
                    }}
                    priority
                />
            </div>
        </div>
    )
}


export default MainSection;
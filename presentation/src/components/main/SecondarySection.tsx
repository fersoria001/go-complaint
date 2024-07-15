import Image from "next/image";
const SecondarySection: React.FC = () => {
    return (
        <div className="h-screen  lg:flex overflow-hidden xl:pt-8">
            <article className="mb-4 lg:ps-5 flex flex-col pt-5 pb-8 pr-8 lg:my-auto lg:max-w-md 2xl:max-w-2xl">
                <p className="px-5 lg:text-start mb-4 text-md lg:text-lg xl:text-xl 2xl:text-3xl text-gray-700">
                    Our services are made for improving customers attention.
                    How clients feel and react to assistants work is key
                    to achieve this goal.
                </p>
                <p className="px-5 lg:text-start mb-4 text-md lg:text-lg xl:text-xl 2xl:text-3xl text-gray-700">
                    At Go Complaint our features are aware about this and
                    the attention can be corrected while is being served and after.
                    Also every single complaint generates a review.
                </p>
            </article>
            <div className="p-2.5
            lg:ps-5 lg:pt-0 bg-blue-500 skew-y-6 relative md:mx-auto md:w-[90%] md:h-[70%] lg:w-[50%] lg:h-[50%] lg:translate-y-36
            xl:h-[70%] xl:w-[70%] xl:translate-y-28 ">
                <div className="
                w-full lg:absolute
                md:translate-y-2
                lg:-translate-x-2
                xl:-translate-y-12
                origin-center -rotate-12 skew-x-6 -skew-y-6
                 shadow-sm shadow-gray-700">
                    <Image
                        src={'/complain.png'}
                        alt={'Image of Go complaint profile'}
                        width={1077}
                        height={633}
                        style={{
                            width: '100%',
                            height: 'auto',
                        }}
                    />
                </div>
            </div>
        </div>
    )
}
export default SecondarySection;
import { useEffect, useState } from "react";

function MainPage() {
    const [src, setSrc] = useState(["./find-receiver.png", "./describe.png", "./complain.png"])
    useEffect(() => {
        const interval = setInterval(() => {
            console.log("change");
            setSrc(prev => {
                const first = prev.shift() as string
                const newArr = [...prev, first]
                return newArr
            })
            const img = document.getElementById('send-complaint')
            img?.animate({
                opacity: [0, 0.8, 0.9, 1],
                offset: [0, 0.5, 1],
                easing: ["ease-in", "ease-out"],
            }, 500
            )
        }, 3000)
        return () => {
            clearInterval(interval)
        }
    }, [])
    return (
        <div className="p-5">
            <h1 className="p-5 font-bold text-gray-900 text-xl md:text-2xl text-center text-nowrap">
                Welcome to Go-Complaint
            </h1>
            <h1 className="p-5 font-bold text-gray-900 text-md md:text-xl text-center">
                a site designed to send complaints to different users and enterprises.
            </h1>
            <div className="border rounded-md w-full p-5 mb-2">
                <div className="flex flex-col md:flex-row justify-between  rounded-md w-full mb-2">
                    <h2 className="text-center pt-12 mb-4 md:mb-0 md:ps-12 text-gray-800 font-medium text-xl md:text-2xl">
                        Manage your own complaints.
                    </h2>
                    <img
                        id="send-complaint"
                        src={src[0]}
                        className="mr-20 w-[450px] h-auto" />
                </div>
                <div className="flex flex-col md:flex-row justify-between  rounded-md w-full mb-2">
                    <h3 className="text-center pt-12 mb-4 md:mb-0 md:ps-24 text-gray-700 text-md md:text-xl">
                        Ask for a review.
                    </h3>
                    <img src="./reviews.png" className="mr-20 w-[450px] h-auto" />
                </div>
            </div>

            <div className="border rounded-md p-5">
                <div className="flex flex-col md:flex-row  justify-between  rounded-md w-full  mb-4">
                    <h2 className="pt-12  mb-4 md:mb-0 md:ps-12 text-gray-800 font-medium text-xl md:text-2xl">
                        Register an enterprise, invite people to work with you.
                    </h2>
                    <img src="./hiring-proccess.png" className="mr-20 w-[450px] h-auto" />
                </div>
                <div className="flex flex-col md:flex-row  justify-between  rounded-md w-full  mb-4">
                    <h3 className="pt-12 mb-4 md:mb-0 md:ps-24 text-gray-700 text-md md:text-xl">
                        We count with an integrated enterprise chat to communicate
                    </h3>
                    <img src="./chat.png" className="mr-20 w-[450px] h-auto" />
                </div>
                <div className="flex flex-col md:flex-row  justify-between  rounded-md w-full  mb-4">
                    <h3 className="pt-12 whitespace-nowrap mb-4 md:mb-0 md:ps-24 text-gray-700 text-md md:text-xl">
                        Make a feedback on the attention.
                    </h3>
                    <img src="./feedbacks.png" className="mr-20 w-[450px] h-auto" />
                </div>
            </div>
        </div>

    )
}

export default MainPage;
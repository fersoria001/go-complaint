
import { Link, useRouter } from "@tanstack/react-router";
import SettingsIcon from "../components/icons/SettingsIcon";
import { Route } from "../routes/$enterpriseID";
import CloseIcon from "../components/icons/CloseIcon";
import { useState } from "react";
import { leaveEnterprise } from "../lib/leave_enterprise";



const EnterprisePage: React.FC = () => {
    const { descriptor, enterprise, role, complaintsInfo, employee } = Route.useLoaderData()
    const [showModal, setShowModal] = useState(false)
    const router = useRouter()
    let managers = 0
    let assistants = 0
    for (let i = 0; i < enterprise.employees!.length; i++) {
        if (enterprise.employees![i].position.toLocaleLowerCase() === "manager") {
            managers++
        } else {
            assistants++
        }
    }

    const handleLeaveEnterprise = async () => {
        if (!employee) return
        const ok = await leaveEnterprise({ employeeID: employee.id, enterpriseName: enterprise.name })
        if (ok) {
            setShowModal(false)
            router.navigate({ to: '/profile' })
        }
    }

    return (
        <div>
            <div className="md:grid md:grid-cols-3 md:grid-rows-2 pt-5 pb-2 bg-white  rounded-md">

                <div className="md:row-start-1 md:row-end-1 md:col-span-1 md:ml-2 md:mt-2">
                    <div className=" flex flex-col border p-5 rounded-md">
                        <div className="self-start">
                            <h1 className="text-sm md:text-xl text-gray-700 mb-2"> {descriptor?.fullName} </h1>
                            <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {role} </h2>
                        </div>
                        <div className="self-end">
                            <h1 className="underline underline-offset-8 text-sm md:text-xl text-gray-700 mb-2"> {enterprise.name} </h1>
                        </div>
                    </div>
                </div>

                <div className="relative md:row-start-1 md:row-end-1 md:col-span-2 flex flex-col md:flex-row border md:border-none p-5 rounded-md">
                    <img src={enterprise.logoIMG} alt="logo" className="self-center md:self-start md:mt-4 w-20 h-20  rounded-full" />

                    <div className="self-center md:self-start md:mt-4 text-center md:whitespace-nowrap">
                        <h1 className="text-sm md:text-xl text-gray-700 mb-2"> {enterprise.email} </h1>
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {enterprise.phone} </h2>
                    </div>

                    <div className="md:w-full text-end">
                        {role === "OWNER" && <Link to={`/${enterprise.name}/settings`}>
                            <SettingsIcon className="cursor-pointer absolute top-4 h-6 w-6 md:relative md:-top-3  md:ms-auto md:mb-10 md:w-10 md:h-10" fill="#374151" />
                        </Link>}
                        <h1 className="text-sm md:text-xl text-gray-700 mb-2"> {enterprise.address.country} </h1>
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {enterprise.address.county} </h2>
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {enterprise.address.city} </h2>
                    </div>
                </div>

                <div className="md:row-start-2 md:row-end-2 md:col-start-3 md:col-end-3">
                    <div className="flex flex-col p-5 border md:border-none rounded-md">
                        <h1 className="underline underline-offset-8 text-sm md:text-xl text-gray-700 mb-2"> {enterprise.name} </h1>
                        <div className="ml-auto">
                            <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {enterprise.employees!.length} {" "} Employees</h2>
                            <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {managers} {" "} Managers</h2>
                            <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {assistants} {" "} Assistants</h2>
                        </div>
                    </div>
                    {role != "OWNER" &&
                        <div className="relative ms-52">
                            <button
                                onClick={() => setShowModal(true)}
                                type="button"
                                className="border py-3 px-1 rounded-md border-cyan-300">
                                Leave enterprise
                            </button>
                            {showModal && <div className="absolute flex flex-col border min-w-52 -top-28 right-16 bg-white">
                                <span
                                    onClick={() => setShowModal(false)}
                                    className="ms-auto mr-1 mt-1 cursor-pointer">
                                    <CloseIcon fill="#374151" />
                                </span>
                                <p className="px-3 text-gray-700"> Are you sure you want to leave {enterprise.name} ?</p>
                                <div className="flex w-full justify-around p-3 ">
                                    <button
                                        onClick={handleLeaveEnterprise}
                                        className="py-2 px-6 border border-red-300 ease-in duration-300 hover:-translate-y-2 hover:scale-110 ">
                                        Yes
                                    </button>
                                    <button
                                        onClick={() => setShowModal(false)}
                                        className="py-2 px-6 border border-blue-300 ease-in duration-300 hover:-translate-y-2 hover:scale-110">
                                        No
                                    </button>
                                </div>
                            </div>}
                        </div>
                    }
                </div>

                <div className="md:row-start-2 md:row-end-2 md:col-start-1 md:col-end-3 flex flex-col p-5  border md:border-none rounded-md">
                    <h1 className="text-center text-sm md:text-xl text-gray-700 mb-2">Complaints received : {complaintsInfo.complaintsReceived}</h1>
                    <div className="self-start">
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {complaintsInfo.complaintsResolved} {" "} Resolved</h2>
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {complaintsInfo.complaintsPending} {" "} Pending</h2>
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {complaintsInfo.complaintsReviewed} {" "} Reviewed by the users</h2>
                        <h2 className="text-sm md:text-xl text-gray-700 mb-2"> {complaintsInfo.averageRating} {" "} Average rating</h2>
                    </div>

                </div>



            </div>
            <img src={enterprise.bannerIMG} alt="cover" className="md:row-start-3 md:col-span-3 md:w-[1200px] md:h-[400px] h-52 w-full rounded-md" />
        </div>
    )
}

export default EnterprisePage;
import { Link, useRouter } from "@tanstack/react-router";
import { EmployeeType } from "../../../lib/types";
import CheerIcon from "../../icons/CheerIcon";
import CommunicationIcon from "../../icons/CommunicationIcon";
import ContactMailIcon from "../../icons/ContactMailIcon";
import MaleFaceIcon from "../../icons/MaleFaceIcon";
import PersonRemoveIcon from "../../icons/PersonRemoveIcon";
import WorkIcon from "../../icons/WorkIcon";
import { useState } from "react";
import SelectIcon from "../../icons/SelectIcon";
import CloseIcon from "../../icons/CloseIcon";
import { fireEmployee, promoteEmployee } from "./employee_lib";


interface Props {
    employee: EmployeeType;
    enterpriseId: string;
}
const Employee: React.FC<Props> = ({ employee, enterpriseId }: Props) => {
    const [fill, setFill] = useState('#374151')
    const [fill1, setFill1] = useState('#374151')
    const [fill2, setFill2] = useState('#374151')
    const [showPromotion, setShowPromotion] = useState(false)
    const [showFire, setShowFire] = useState(false)
    const [selectedPosition, setSelectedPosition] = useState("")
    const positions = ["MANAGER", "ASSISTANT"]
    const router = useRouter()
    const handlePromotion = async () => {
        await promoteEmployee({
            enterpriseName: enterpriseId,
            employeeID: employee.id,
            position: selectedPosition
        })
        router.invalidate()
    }
    const handleFire = async () => {
        await fireEmployee({
            enterpriseName: enterpriseId,
            employeeID: employee.id
        })
        router.invalidate()
    }
    const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedPosition(e.target.value)
    }
    return (
        <div className="flex flex-col md:flex-row justify-around items-center bg-white border border-gray-200 rounded-lg shadow relative">
            <div className="mt-2">
                <img className="h-24 rounded-t-lg" src={employee.profileIMG} alt="avatar" />
                <h5 className=" text-center font-bold tracking-tight text-gray-900">
                    {employee.firstName} {employee.lastName}
                </h5>
            </div>
            <div className="self-start px-2 md:py-5">
                <div className="flex mb-1">
                    <ContactMailIcon fill="#374151" />
                    <p className=" font-normal text-gray-700 underline underline-offset-8">{employee.email}</p>
                </div>

                <div className="flex mb-1">
                    <MaleFaceIcon fill="#374151" />
                    <p className=" font-normal text-gray-700 underline underline-offset-8">Age: {employee.age}</p>
                </div>
                <div className="flex mb-1">
                    <WorkIcon fill="#374151" />
                    <p className="font-normal text-gray-700 underline underline-offset-8">
                        Position: {employee.position}</p>
                </div>
                <p className="font-normal text-gray-700">
                    Hiring date: {new Date(parseInt(employee.approvedHiringAt)).toLocaleDateString()}
                </p>

            </div>


            <div className="self-start px-2 grid grid-cols-1 grid-rows-1 md:grid-cols-3 md:py-3">
                <div className="col-start-1 cold-end-1 row-start-1 row-end-2 md:col-span-3 md:flex">
                    <div className="pr-2">
                        <p className="md:text-start text-sm text-gray-700">
                            Complaints solved: {employee.complaintsSolved}
                        </p>
                        <p className="md:text-start text-sm  text-gray-700">
                            Complaints reviewed: {employee.complaintsRated}
                        </p>
                    </div>

                    <div className="pr-2">
                        <p className="md:text-start text-sm  text-gray-700">
                            Feedback started: {employee.complaintsFeedbacked}
                        </p>
                        <p className="md:text-start text-sm  text-gray-700">
                            Feedback received: {employee.complaintsFeedbacked}
                        </p>
                    </div>

                    <div className="pr-2">
                        <p className="md:text-start text-sm  text-gray-700">
                            Hiring invitations sent: {employee.hireInvitationsSent}
                        </p>
                        <p className="md:text-start text-sm  text-gray-700">
                            Employees hired: {employee.employeesHired}
                        </p>
                        <p className="md:text-start text-sm  text-gray-700">
                            Employees fired: {employee.employeesFired}
                        </p>
                    </div>
                </div>

                <div className="row-start-2 row-end-2 md:col-span-3 md:mx-auto">
                    <div className="md:mx-auto inline-flex rounded-md shadow-sm py-5 md:py-2" role="group">
                        {employee.complaintsSolved > 0 ?
                            <Link
                                id="feedback-btn"
                                to={`/${enterpriseId}/employees/solved/${employee.id}`}
                                onMouseEnter={() => setFill('#ffffff')}
                                onMouseLeave={() => setFill('#374151')}
                                className={`inline-flex items-center px-1 py-2
                                 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-cyan-200
                        hover:text-white
                        relative z-10  bg-gradient-to-br from-cyan-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`
                                }>
                                <CommunicationIcon fill={fill} />
                                Feedback
                            </Link> :
                            <p
                                onMouseEnter={() => setFill('#374151')}
                                onMouseLeave={() => setFill('#374151')}
                                className={`inline-flex items-center px-1 py-2
                                 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-cyan-200
                        opacity-50 cursor-not-allowed`
                                }>
                                <CommunicationIcon fill={fill} />
                                Feedback
                            </p>
                        }
                        <button
                            onMouseEnter={() => setFill1('#ffffff')}
                            onMouseLeave={() => setFill1('#374151')}
                            onMouseUp={() => setShowPromotion(true)}
                            type="button"
                            className={`inline-flex items-center px-1 py-2 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-cyan-200
                        hover:text-white
                        relative z-10 bg-gradient-to-br from-cyan-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`}>
                            <CheerIcon fill={fill1} />
                            Promote
                        </button>
                        <button
                            onMouseUp={() => setShowFire(true)}
                            onMouseEnter={() => setFill2('#ffffff')}
                            onMouseLeave={() => setFill2('#374151')}
                            type="button"
                            className={
                                `inline-flex items-center px-1 py-2 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-cyan-200 
                        hover:text-white
                        relative z-10  bg-gradient-to-br from-cyan-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`
                            }>
                            <PersonRemoveIcon fill={fill2} />
                            Fire
                        </button>
                    </div>
                </div>
            </div>

            {showPromotion &&
                <div
                    className="absolute bottom-8 md:bottom-0 md:top-20  md:right-12 z-20 bg-white flex flex-col  border  w-full md:h-[240px] md:w-[400px] rounded-md">
                    <div
                        onMouseUp={() => setShowPromotion(false)}
                        className="cursor-pointer self-end mr-1 mt-1">
                        <CloseIcon fill="#374151" />
                    </div>
                    <div className="p-5 flex flex-col">
                        <input className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline cursor-none mb-4" value={employee.position} readOnly />
                        <div className="relative mb-4">
                            <select
                                className="block appearance-none w-full bg-gray-200 border
                         border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                          focus:outline-none focus:bg-white focus:border-gray-500"
                                onChange={handleChange}
                                name="gender">
                                {
                                    positions.map((position, index) =>

                                        <option key={index} value={position}>{position}</option>

                                    )
                                }
                            </select>
                            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                                <SelectIcon />
                            </div>
                        </div>
                        <button
                            onMouseUp={handlePromotion}
                            className="self-center rounded-md py-3 px-1.5 text-white font-medium w-1/2 bg-gradient-to-br from-cyan-500  to-blue-500 " type="button">Promote</button>
                    </div>
                </div>
            }
            {showFire && <div
                className="absolute bottom-8 md:bottom-0 md:top-20  md:right-12 z-20 bg-white flex flex-col  border  w-full md:h-[250px] md:w-[400px] rounded-md">
                <div
                    onMouseUp={() => setShowFire(false)}
                    className="cursor-pointer self-end mr-1 mt-1">
                    <CloseIcon fill="#374151" />
                </div>
                <div className="px-5 pb-5 pt-3 flex flex-col">
                    <p className="md:text-start text-sm md:text-xl text-gray-700"> Are you sure you want to fire {employee.firstName + " " + employee.lastName}?</p>
                    <p className="md:text-start text-sm md:text-xl  text-gray-700">
                        This proccess can not be undone, you will still be able to sent a new hiring invitation.
                    </p>
                </div>
                <div className="flex self-center justify-center gap-2 w-full pb-5">
                    <button
                        onMouseUp={handleFire}
                        className="w-1/3 self-center rounded-md py-3 px-1.5 text-white font-medium bg-gradient-to-r from-red-200 via-red-300 to-yellow-200 " type="button">Fire</button>
                    <button
                        onMouseUp={() => setShowFire(false)}
                        className="w-1/3 self-center rounded-md py-3 px-1.5 text-white font-medium bg-gradient-to-br from-cyan-500  to-blue-500 " type="button">Cancel</button>
                </div>
            </div>}




        </div >
    )
}

export default Employee;


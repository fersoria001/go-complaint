import { useState } from "react";
import SelectIcon from "../../icons/SelectIcon";
import CloseIcon from "../../icons/CloseIcon";
import { Employee, UserDescriptor } from "@/gql/graphql";
import CheerIcon from "@/components/icons/CheerIcon";
import CommunicationIcon from "@/components/icons/CommunicationIcon";
import ContactMailIcon from "@/components/icons/ContactMailIcon";
import MaleFaceIcon from "@/components/icons/MaleFaceIcon";
import PersonRemoveIcon from "@/components/icons/PersonRemoveIcon";
import WorkIcon from "@/components/icons/WorkIcon";
import Link from "next/link";
import FemaleFaceIcon from "@/components/icons/FemaleFaceIcon";
import Image from "next/image";
import ChartIcon from "@/components/icons/ChartIcont";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { fireEmployee, promoteEmployee } from "@/lib/actions/graphqlActions";


interface Props {
    employee: Employee
    enterpriseName: string;
    currentUser: UserDescriptor
}
const positions = ["MANAGER", "ASSISTANT"]
const EmployeeCard: React.FC<Props> = ({ employee, enterpriseName, currentUser }: Props) => {
    const [fill, setFill] = useState<Array<string>>(new Array('#374151', '#374151', '#374151', '#374151'))
    const [showPromotion, setShowPromotion] = useState(false)
    const [showFire, setShowFire] = useState(false)
    const [selectedPosition, setSelectedPosition] = useState<string>(positions[0])
    const [reason, setReason] = useState<string>("")
    const queryClient = useQueryClient()


    const handleBtnHover = (index: number, v: boolean) => {
        if (v) {
            setFill(prev => {
                prev[index] = "#ffffff"
                return [...prev]
            })
        } else {
            setFill(prev => {
                prev[index] = "#374151"
                return [...prev]
            })
        }
    }


    const promotionMutation = useMutation({
        mutationFn: async () => promoteEmployee({
            enterpriseName: enterpriseName,
            employeeId: employee.id,
            promoteTo: selectedPosition,
            promotedById: currentUser.id
        }),
        onSuccess: () => {queryClient.invalidateQueries({ queryKey: ['enterprise'] }); setShowPromotion(false)}
    })

    const fireMutation = useMutation({
        mutationFn: async () => fireEmployee({
            employeeId: employee.id,
            enterpriseName: enterpriseName,
            fireReason: reason,
            triggeredBy: currentUser.id
        }),
        onSuccess: () => {queryClient.invalidateQueries({ queryKey: ['enterprise'] }); setShowFire(false)}
    })

    const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedPosition(e.target.value)
    }
    return (
        <div className="flex flex-col md:flex-row justify-around
          bg-white border border-gray-200 rounded-lg shadow relative p-3">

            <div className="mt-2 flex flex-col items-center shrink-0">
                <div className='relative w-36 h-36'>
                    <Image
                        src={employee.user.person.profileImg}
                        className="rounded-t-lg"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                        fill
                        alt="user photo" />
                </div>
                <h5 className="text-center font-bold tracking-tight text-gray-900">
                    {employee.user.person.firstName} {employee.user.person.lastName}
                </h5>
            </div>
            <div className="px-2 flex flex-col h-auto my-auto">
                <div className="flex mb-1">
                    <ContactMailIcon fill="#374151" className="w-6 h-6" />
                    <p className=" font-normal text-gray-700 underline underline-offset-8">{employee.user.person.email}</p>
                </div>
                <div className="flex mb-1">
                    {
                        employee.user.person.genre == "female" ?
                            <FemaleFaceIcon fill="#374151" className="w-6 h-6" /> :
                            <MaleFaceIcon fill="#374151" className="w-6 h-6" />
                    }
                    <p className=" font-normal text-gray-700 underline underline-offset-8">Age: {employee.user.person.age}</p>
                </div>
                <div className="flex mb-1">
                    <WorkIcon fill="#374151" className="w-6 h-6" />
                    <p className="font-normal text-gray-700 underline underline-offset-8">
                        Position: {employee.enterprisePosition}</p>
                </div>
                <p className="font-normal text-gray-700">
                    Hiring date: {new Date(parseInt(employee.approvedHiringAt)).toDateString()}
                </p>

            </div>


            <div className="grid rows-2 lg:inline-flex md:h-28 lg:h-12 md:mx-auto rounded-md shadow-sm my-auto" role="group">
                <Link
                    id="feedback-btn"
                    href={`/enterprises/${enterpriseName}/employees/feedback?id=${employee.id}`}
                    onMouseEnter={() => handleBtnHover(0, true)}
                    onMouseLeave={() => handleBtnHover(0, false)}
                    className={`row-start-1 inline-flex items-center px-1 py-2
                                 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-blue-500
                        hover:text-white
                        relative z-10  bg-gradient-to-br from-blue-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`
                    }>
                    <CommunicationIcon fill={fill[0]} className="w-6 h-6" />
                    Feedback
                </Link>
                <Link
                    id="activity-btn"
                    href={`/enterprises/${enterpriseName}/employees/activity?id=${employee.user.id}`}
                    onMouseEnter={() => handleBtnHover(1, true)}
                    onMouseLeave={() => handleBtnHover(1, false)}
                    className={`row-start-1 inline-flex items-center px-1 py-2
                                 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-blue-500
                        hover:text-white
                        relative z-10  bg-gradient-to-br from-blue-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`
                    }>
                    <ChartIcon fill={fill[1]} className="w-6 h-6" />
                    Activity
                </Link>
                <button
                    onMouseEnter={() => handleBtnHover(2, true)}
                    onMouseLeave={() => handleBtnHover(2, false)}
                    onMouseUp={() => setShowPromotion(true)}
                    type="button"
                    className={`row-start-2 inline-flex items-center px-1 py-2 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-blue-500
                        hover:text-white
                        relative z-10 bg-gradient-to-br from-blue-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`}>
                    <CheerIcon fill={fill[2]} className="w-6 h-6" />
                    Promote
                </button>
                <button
                    onMouseUp={() => setShowFire(true)}
                    onMouseEnter={() => handleBtnHover(3, true)}
                    onMouseLeave={() => handleBtnHover(3, false)}
                    type="button"
                    className={
                        `row-start-2 inline-flex items-center px-1 py-2 md:px-5 md:py-3 text-sm font-medium
                        text-gray-900 bg-transparent border border-blue-500
                        hover:text-white
                        relative z-10 bg-gradient-to-br from-blue-500  to-blue-500 
                        before:absolute  before:top-0 before:left-0 
                        before:w-full before:h-full 
                        before:bg-white 
                        before:-z-10 hover:before:translate-x-full`
                    }>
                    <PersonRemoveIcon fill={fill[3]} className="w-6 h-6" />
                    Fire
                </button>

            </div>


            {showPromotion &&
                <div
                    className="absolute bottom-4 inset-x-4 md:inset-x-auto md:bottom-0 z-20
                     bg-white flex flex-col  border   max-w-md rounded-md">
                    <div
                        onMouseUp={() => setShowPromotion(false)}
                        className="cursor-pointer self-end mr-1 mt-1">
                        <CloseIcon fill="#374151" className="w-6 h-6" />
                    </div>
                    <div className="p-5 flex flex-col">
                        <input className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline cursor-none mb-4" value={employee.enterprisePosition} readOnly />
                        <div className="relative mb-4">
                            <select
                                className="block appearance-none w-full bg-gray-200 border
                         border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                          focus:outline-none focus:bg-white focus:border-gray-500"
                                onChange={handleChange}
                                defaultValue={selectedPosition}>
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
                            onMouseUp={() => promotionMutation.mutate()}
                            className="self-center rounded-md py-3 px-1.5 text-white font-medium w-1/2 
                             bg-blue-500 hover:bg-blue-600" type="button">Promote</button>
                    </div>
                </div>
            }
            {showFire && <div
                className="absolute md:-top-16 md:inset-x-auto z-20 bg-white flex flex-col  border 
                 max-w-md h-[340px] rounded-md">
                <div
                    onMouseUp={() => setShowFire(false)}
                    className="cursor-pointer self-end mr-1 mt-1">
                    <CloseIcon fill="#374151" className="w-6 h-6"/>
                </div>
                <div className="px-5 pb-5 pt-3 flex flex-col">
                    <p className="md:text-start text-sm md:text-md text-gray-700">
                        Are you sure you want to fire {employee.user.person.firstName + " " + employee.user.person.lastName}?
                    </p>
                    <p className="md:text-start text-sm md:text-md text-gray-700">
                        This proccess can not be undone, you will still be able to sent a new hiring invitation.
                    </p>
                </div>
                <div className="p-5">
                    <textarea
                        rows={4}
                        className="text-sm md:text-md border h-full w-full appearance-none focus:outline-none"
                        maxLength={80}
                        placeholder="You can write a reason to fire this employee (optional)"
                        onChange={(e) => setReason(e.currentTarget.value)}
                    />
                </div>
                <div className="flex self-center justify-center gap-2 w-full pb-5">
                    <button
                        onMouseUp={() => fireMutation.mutate()}
                        className="w-1/3 self-center rounded-md py-3 px-1.5 text-white font-medium 
                       bg-blue-500 hover:bg-blue-600" type="button">Fire</button>
                    <button
                        onMouseUp={() => setShowFire(false)}
                        className="w-1/3 self-center rounded-md py-3 px-1.5 text-white font-medium
                        bg-blue-500 hover:bg-blue-600" type="button">Cancel</button>
                </div>
            </div>}
        </div >
    )
}

export default EmployeeCard;
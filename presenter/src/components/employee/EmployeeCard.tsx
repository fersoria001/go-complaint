import { Link } from "react-router-dom";
import { Employee } from "../../lib/types";
import CheerIcon from "../icons/CheerIcon";
import CommunicationIcon from "../icons/CommunicationIcon";
import ContactMailIcon from "../icons/ContactMailIcon";
import MaleFaceIcon from "../icons/MaleFaceIcon";
import PersonRemoveIcon from "../icons/PersonRemoveIcon";
import WorkIcon from "../icons/WorkIcon";
interface Props {
    enterpriseID: string;
    employee: Employee;
}
function EmployeeCard({ enterpriseID, employee }: Props) {
    return (<div className="flex flex-col md:flex-row justify-around items-center
        bg-white border border-gray-200 rounded-lg shadow ">
        <div className="flex flex-col align-center justify-center" >
            <img className="w-full h-48 object-scale-down rounded-t-lg" src={employee.profileIMG} alt="avatar" />
        </div>
        <div className="flex flex-col">
            <div className="flex self-center md:self-auto flex-col md:flex-row  md:justify-around">
                <h5 className="pb-2 mb-2 text-2xl text-center font-bold tracking-tight text-gray-900">
                    {employee.firstName} {employee.lastName}
                </h5>
            </div>
            <div className="flex flex-col self-center">
                <div className="self-start mb-3 ">
                    <div className="flex mb-3">
                        <MaleFaceIcon fill="#374151" />
                        <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Age: {employee.age}</p>
                    </div>
                    <div className="flex mb-3">
                        <WorkIcon fill="#374151" />
                        <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">
                            Position: {employee.position}</p>
                    </div>

                </div>
                <div className="self-start mb-3 ">
                    <div className="flex mb-3">
                        <ContactMailIcon fill="#374151" />
                        <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">{employee.email}</p>
                    </div>
                </div>
                <div className="flex flex-col align-center justify-between w-full">
                    <p className="mb-3  mr-2 font-normal text-gray-700">
                        Hired since: {new Date(parseInt(employee.approvedHiringAt)).toLocaleDateString()}
                    </p>
                </div>
                <div className="inline-flex rounded-md shadow-sm py-5" role="group">
                    <Link
                        to={`/enterprises/${enterpriseID}/feedback?employee=${employee.ID}`}
                        className={parseInt(employee.complaintsSolved) > 0 ? `inline-flex items-center px-4 py-2 text-sm font-medium
             text-gray-900 bg-transparent border border-gray-900 rounded-s-lg
              hover:bg-gray-900 hover:text-white focus:z-10 focus:ring-2 focus:ring-gray-500
               focus:bg-gray-900 focus:text-white`:
                            `inline-flex items-center px-4 py-2 text-sm font-medium
             text-gray-900 bg-transparent border border-gray-900 rounded-s-lg
             cursor-not-allowed opacity-50`
                        }>
                        <CommunicationIcon fill="#374151" />
                        Feedback
                    </Link>
                    <button
                        type="button"
                        className={`inline-flex items-center px-4 py-2 text-sm font-medium
              text-gray-900 bg-transparent border-t border-b
               border-gray-900 hover:bg-gray-900 hover:text-white
                focus:z-10 focus:ring-2 focus:ring-gray-500 focus:bg-gray-900
                 focus:text-white`}>
                        <CheerIcon fill="#374151" />
                        Promote
                    </button>
                    <button
                        type="button"
                        className="inline-flex items-center px-4 py-2 text-sm font-medium
              text-gray-900 bg-transparent border border-gray-900 rounded-e-lg
               hover:bg-gray-900 hover:text-white focus:z-10 focus:ring-2 focus:ring-gray-500
                focus:bg-gray-900 focus:text-white">
                        <PersonRemoveIcon fill="#374151" />
                        Fire
                    </button>
                </div>
            </div>
        </div>
    </div>
    )
}

export default EmployeeCard;


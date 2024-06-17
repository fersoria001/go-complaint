import { Link } from "react-router-dom";
import UpdateIcon from "../icons/UpdateIcon";
import useOffices from "../../lib/hooks/useOffices";
import OfficeCard from "./OfficeCard";

interface Props {
    title: string;
    description: string;
}
function OfficesList({ title, description }: Props) {
    const offices = useOffices();
    console.log("officeS", offices)
    return (
        <>
            <div
                className="block p-6
        bg-white border
        border-gray-200 rounded-lg
        shadow
        hover:bg-gray-100"
            >
                <h5
                    className="mb-2
        text-2xl font-bold tracking-tight text-gray-900"
                >
                    {title}
                </h5>
                <div className="flex">
                    <UpdateIcon fill="#374151" />
                    <p className="font-normal text-gray-700">
                        {description}
                    </p>
                </div>
            </div>
            <div>
                {offices.length < 1 ?
                    <div>
                        <h5>You are not currently at any enterprises</h5>
                    </div> :
                    offices.map((office) => (
                        <Link to={`/office/${office.employeeID.split("-")[0]}`} key={office.employeeID}>
                            <OfficeCard {...office} />
                        </Link>
                    ))}
            </div>
        </>
    );
}

export default OfficesList;
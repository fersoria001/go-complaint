import { Enterprise, UserDescriptor } from "../../lib/types";
import OfficeCard from "./OfficeCard";
import { Link } from "@tanstack/react-router";

interface Props {
    enterprises: Enterprise[];
    descriptor: UserDescriptor;
}
function OfficesList({ enterprises, descriptor }: Props) {
    return (
        <>
            <div
                className="block p-6
        bg-white border
        border-gray-200 rounded-lg
        shadow"
            >
                <h5
                    className="mb-2
        text-2xl font-bold tracking-tight text-gray-900"
                >
                    Offices
                </h5>
                <div className="flex">
                    <p className="font-normal text-gray-700">
                        {"Here's a list of enterprises you have a position in"}
                    </p>
                </div>
            </div>
            <div>
                {enterprises.length < 1 ?
                    <div className="bg-white border mt-2
                     border-gray-200 rounded-lg shadow">
                        <h5
                            className="text-sm md:text-xl text-gray-700 p-4 mb-4">
                            You are not currently working at any enterprises</h5>
                    </div> :
                    enterprises.map((enterprise) => (
                        <Link to={`/${enterprise.name}`} key={enterprise.name}>
                            <OfficeCard
                                enterprise={enterprise}
                                authority={descriptor.grantedAuthorities.find((ga) => ga.enterpriseID === enterprise.name)!} />
                        </Link>
                    ))}
            </div>
        </>
    );
}

export default OfficesList;
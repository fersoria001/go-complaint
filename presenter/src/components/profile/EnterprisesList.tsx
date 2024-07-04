import { Link } from "@tanstack/react-router";
import EnterpriseCard from "./EnterpriseCard";
import { Enterprise } from "../../lib/types";

interface Props {
  enterprises: Enterprise[];
}
function EnterprisesList({ enterprises }: Props) {
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
          Enterprises
        </h5>
        <div className="flex">
          <p className="font-normal text-gray-700">
            Here's a list of enterprises that you have register
          </p>
        </div>
      </div>
      <div>
        {
          enterprises.length < 0 ?
            <div className="bg-white border mt-2
        border-gray-200 rounded-lg shadow">
              <h5
                className="text-sm md:text-xl text-gray-700 p-4 mb-4">
                No enterprises found</h5>
            </div> :
            enterprises.map((enterprise) => (
              <Link to={`/${enterprise.name}`} key={enterprise.name}>
                <EnterpriseCard {...enterprise} />
              </Link>
            ))
        }
      </div>
    </>
  );
}

export default EnterprisesList;
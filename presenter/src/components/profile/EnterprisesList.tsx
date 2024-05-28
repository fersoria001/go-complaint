import { Link } from "react-router-dom";
import useOwnerEnterprises from "../../lib/hooks/useOwnerEnterprises";
import UpdateIcon from "../icons/UpdateIcon";
import EnterpriseCard from "./EnterpriseCard";

interface Props {
  title: string;
  description: string;
}
function EnterprisesList({ title, description }: Props) {
  const enterprises = useOwnerEnterprises();
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
        {enterprises.length < 1 ?
          <div>
            <h5>No enterprises found</h5>
          </div> :
          enterprises.map((enterprise) => (
            <Link to={`/enterprises/${enterprise.name}`} key={enterprise.name}>
              <EnterpriseCard {...enterprise} />
            </Link>
          ))}
      </div>
    </>
  );
}

export default EnterprisesList;
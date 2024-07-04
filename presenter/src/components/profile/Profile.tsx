

import { Route } from "../../routes/_profile";
import EnterprisesList from "./EnterprisesList";
import OfficesList from "./OfficesList";

function Profile() {
    const { descriptor, ownerEnterprises, employeeEnterprises } = Route.useLoaderData();
    return (
        <div className="pb-[30px] md:min-h-[450px]">
            <div className="mb-4">
                <EnterprisesList enterprises={ownerEnterprises} />
            </div>
            <div className="mb-4">
                <OfficesList
                    enterprises={employeeEnterprises}
                    descriptor={descriptor!} />
            </div>
        </div>
    );
}

export default Profile;
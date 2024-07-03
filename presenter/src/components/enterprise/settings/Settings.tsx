import { Route } from "../../../routes/$enterpriseID/settings";
import UpdateAddress from "./UpdateAddress";
import UpdateBannerImage from "./UpdateBannerImage";
import UpdateEmail from "./UpdateEmail";
import UpdateLogoImage from "./UpdateLogoImage";
import UpdatePhone from "./UpdatePhone";
import UpdateWebsite from "./UpdateWebsite";

function Settings() {
    const { countries, enterprise } = Route.useLoaderData();
    const { enterpriseID } = Route.useParams();
    return (
        <div className="bg-white">
            <UpdateBannerImage enterprise={enterprise} />
            <UpdateLogoImage enterprise={enterprise} />
            <UpdateWebsite enterpriseID={enterpriseID} />
            <UpdateEmail enterpriseID={enterpriseID} />
            <UpdatePhone enterpriseID={enterpriseID} countries={countries} />
            <UpdateAddress enterpriseID={enterpriseID} countries={countries} />
        </div>
    );
}

export default Settings;
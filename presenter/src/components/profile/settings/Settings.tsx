import { Route } from "../../../routes/_profile/settings";
import UpdateAddress from "./UpdateAddress";
import UpdateFirstName from "./UpdateFirstName";
import UpdateGenre from "./UpdateGenre";
import UpdateLastName from "./UpdateLastName";
import UpdatePassword from "./UpdatePassword";
import UpdatePhone from "./UpdatePhone";
import UpdateProfileImage from "./UpdateProfileImage";

export const Settings: React.FC = () => {
    const { descriptor, countries } = Route.useLoaderData()
    return (
        <div className=" bg-white ">
            <div>
                <UpdateProfileImage descriptor={descriptor} />
                <UpdatePassword />
                <UpdateGenre descriptor={descriptor} />
                <UpdateFirstName />
                <UpdateLastName />
                <UpdatePhone countries={countries} />
                <UpdateAddress countries={countries} />
            </div>

        </div>
    )
}

export default Settings;
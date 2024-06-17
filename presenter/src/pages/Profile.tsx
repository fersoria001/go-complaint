import EnterprisesList from "../components/profile/EnterprisesList";
import OfficesList from "../components/profile/OfficesList";

function Profile() {
    return (
        <>
            <EnterprisesList
                title={"Enterprises"}
                description={"Here's a list of enterprises that you have register"} />
            <OfficesList
                title={"Offices"}
                description={"Here's a list of enterprises you have a position in"} />
        </>
    );
}

export default Profile;
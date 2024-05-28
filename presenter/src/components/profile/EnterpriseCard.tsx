import { Enterprise } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import ContactPhoneIcon from "../icons/ContactPhoneIcon";
import ContactWebsiteIcon from "../icons/ContactWebsiteIcon";

function EnterpriseCard({
    name,
    email,
    website,
    phone,
    industry,
    address,
}: Enterprise) {
    return (
        <div className=" bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
            <div className="flex flex-col p-5">
                <div className="flex self-center md:self-auto flex-col md:flex-row  md:justify-around">
                    <h5 className="pb-2 mb-2 text-2xl text-center font-bold tracking-tight text-gray-900">
                        {name}
                    </h5>

                    <div>
                        <div className="flex mb-3">
                            <ContactMailIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700">{email}</p>
                        </div>
                        <div className="flex mb-3">
                            <ContactPhoneIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700">
                                {phone}
                            </p>
                        </div>
                        <div className="flex mb-3">
                            <ContactWebsiteIcon fill="#374151" />
                            <p className="pl-2 font-normal text-gray-700">
                                {website}
                            </p>
                        </div>
                    </div>
                </div>
                <div className="flex flex-col self-center">
                    <p className="self-center mb-3 font-normal text-gray-700 underline underline-offset-8">
                        {industry}
                    </p>
                    <p className="mb-3 font-normal text-gray-700 text-center">
                        {address.country}, {address.county}, {address.city}.
                    </p>
                </div>
                    {/* banner image to add later */}
                {/* <img className="h-auto w-full" src={logoIMG} alt="" /> */}

            </div>
        </div>
    );
}
export default EnterpriseCard;
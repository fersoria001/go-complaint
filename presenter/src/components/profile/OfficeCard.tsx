import { Enterprise, GrantedAuthority } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import ContactPhoneIcon from "../icons/ContactPhoneIcon";
import ContactWebsiteIcon from "../icons/ContactWebsiteIcon";
interface Props {
  enterprise: Enterprise;
  authority: GrantedAuthority;
}
function OfficeCard({ enterprise, authority }: Props) {
  return (

    <div className="flex  flex-col md:flex-row bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
      <img
        className="w-full md:w-1/4 self-center pt-2 h-48 object-scale-down rounded-t-lg"
        src={enterprise.logoIMG}
        alt="avatar" />
      <div className="w-full flex flex-col p-5 ">
        <div className="flex gap-2 self-center md:self-start w-full  flex-col md:flex-row  md:justify-between">
          
          <div className="md:flex gap-1">
            <h5 className="pb-2 mb-2  text-xl text-center font-bold  text-gray-900 break-normal">
              {authority.authority}
            </h5>
            <h5 className="pb-2 mb-2  text-xl text-center font-bold  text-gray-900 break-normal">
              at
            </h5>
            <h5 className="pb-2 mb-2  text-xl text-center font-bold  text-gray-900 break-normal">
                {enterprise.name}
            </h5>
          </div>

          <div>
            <div className="flex mb-3">
              <ContactMailIcon fill="#374151" />
              <p className="pl-2 font-normal text-gray-700">{enterprise.email}</p>
            </div>
            <div className="flex mb-3">
              <ContactPhoneIcon fill="#374151" />
              <p className="pl-2 font-normal text-gray-700">{enterprise.phone}</p>
            </div>
            <div className="flex mb-3">
              <ContactWebsiteIcon fill="#374151" />
              <p className="pl-2 font-normal text-gray-700">{enterprise.website}</p>
            </div>
          </div>
        </div>
        <div className="flex flex-col self-center">
          <p className="self-center mb-3 font-normal text-gray-700 underline underline-offset-8">
            {enterprise.industry}
          </p>
          <p className="mb-3 font-normal text-gray-700 text-center">
            {enterprise.address.country}, {enterprise.address.county}, {enterprise.address.city}.
          </p>
        </div>
        {/* banner image to add later */}
        {/* <img className="h-auto w-full" src={logoIMG} alt="" /> */}
      </div>
    </div>
  );
}
export default OfficeCard;

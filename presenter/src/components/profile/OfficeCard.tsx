import { Office } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import ContactPhoneIcon from "../icons/ContactPhoneIcon";
import ContactWebsiteIcon from "../icons/ContactWebsiteIcon";
function OfficeCard({
  employeePosition,
  enterpriseLogoIMG,
  enterpriseName,
  enterpriseWebsite,
  enterprisePhone,
  enterpriseEmail,
  enterpriseIndustry,
  enterpriseAddress,
}: Office) {
  return (

    <div className="flex flex-col md:flex-row bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
      <img className="w-full md:w-1/3 self-center pt-2 h-48 object-scale-down rounded-t-lg" src={enterpriseLogoIMG} alt="avatar" />
      <div className="flex flex-col p-5 ">
        <div className="flex gap-2 self-center md:self-auto flex-col md:flex-row  md:justify-around">
          <h5 className="pb-2 mb-2 md:w-1/3 text-2xl text-center font-bold tracking-tight text-gray-900">
            You are currently {employeePosition} at {enterpriseName}
          </h5>
          
          <div>
            <div className="flex mb-3">
              <ContactMailIcon fill="#374151" />
              <p className="pl-2 font-normal text-gray-700">{enterpriseEmail}</p>
            </div>
            <div className="flex mb-3">
              <ContactPhoneIcon fill="#374151" />
              <p className="pl-2 font-normal text-gray-700">{enterprisePhone}</p>
            </div>
            <div className="flex mb-3">
              <ContactWebsiteIcon fill="#374151" />
              <p className="pl-2 font-normal text-gray-700">{enterpriseWebsite}</p>
            </div>
          </div>
        </div>
        <div className="flex flex-col self-center">
          <p className="self-center mb-3 font-normal text-gray-700 underline underline-offset-8">
            {enterpriseIndustry}
          </p>
          <p className="mb-3 font-normal text-gray-700 text-center">
            {enterpriseAddress.country}, {enterpriseAddress.county}, {enterpriseAddress.city}.
          </p>
        </div>
        {/* banner image to add later */}
        {/* <img className="h-auto w-full" src={logoIMG} alt="" /> */}
      </div>
    </div>
  );
}
export default OfficeCard;

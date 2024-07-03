import { User } from "../../lib/types";
import ContactMailIcon from "../icons/ContactMailIcon";
import FemaleFaceIcon from "../icons/FemaleFaceIcon";
import MaleFaceIcon from "../icons/MaleFaceIcon";

interface Props {
    user: User;
}
function UserForHiring({ user }: Props) {
    return (
        <div className="flex flex-col md:flex-row justify-around items-center
        bg-white border border-gray-200 rounded-lg shadow  hover:bg-gray-100">
            <div className="flex flex-col align-center justify-center" >
                <img className="w-full h-48 object-scale-down rounded-t-lg" src={user.profileIMG} alt="avatar" />
            </div>
            <div className="flex flex-col">
                <div className="flex self-center md:self-auto flex-col md:flex-row  md:justify-around">
                    <h5 className="pb-2 mb-2 text-2xl text-center font-bold tracking-tight text-gray-900">
                        {user.firstName} {user.lastName}
                    </h5>
                </div>
                <div className="flex flex-col self-center">
                    <div className="self-start mb-3 ">
                        <div className="flex mb-3">
                            {
                                user.gender == "female" ?
                                    <FemaleFaceIcon fill="#5f6368" /> :
                                    <MaleFaceIcon fill="#5f6368" />
                            }

                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">Age: {user.age}</p>
                        </div>
                    </div>
                    <div className="self-start mb-3 ">
                        <div className="flex mb-3">
                            <ContactMailIcon fill="#5f6368" />
                            <p className="pl-2 font-normal text-gray-700 underline underline-offset-8">{user.email}</p>
                        </div>
                    </div>
                    <p className="mb-3 font-normal text-gray-700 text-center">
                        {user.address.country}, {user.address.county}, {user.address.city}.
                    </p>
                </div>
            </div>

        </div>
    )
}

export default UserForHiring;
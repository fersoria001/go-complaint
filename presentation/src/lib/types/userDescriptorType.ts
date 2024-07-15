import GrantedAuthorityType from "./grantedAuthorityType";

type UserDescriptorType = {
    email: string;
    fullName: string;
    profileImg: string;
    gender: string;
    pronoun: string;
    loginDate: string;
    ip: string;
    device: string;
    geolocation: {
      latitude: number;
      longitude: number;
    };
    grantedAuthorities: GrantedAuthorityType[];
  };

export default UserDescriptorType;
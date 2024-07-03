import Cookies from "js-cookie";
import { fetchUserDescriptor } from "./fetchUserDescriptor";
export const isLoggedIn = () => {
  return Cookies.get("Authorization") !== undefined;
};


export const hasPermission = async (
  role: string,
  enterpriseID: string
): Promise<boolean> => {
  const descriptor = await fetchUserDescriptor();
  if (!descriptor) {
    return false;
  }
  let authority;
  switch (role) {
    case "ASSISTANT":
      authority = descriptor.grantedAuthorities.find(
        (v) => v.authority === role && v.enterpriseID == enterpriseID
      );
      if (authority) {
        return true;
      }
      return false;
    case "MANAGER":
      authority = descriptor.grantedAuthorities.find(
        (v) => v.authority === "MANAGER" && v.enterpriseID == enterpriseID
      );
      if (!authority) {
        authority = descriptor.grantedAuthorities.find(
          (v) => v.authority === "OWNER" && v.enterpriseID == enterpriseID
        );
        if (!authority) return false;
        return true;
      }
      return true;
    case "OWNER":
      authority = descriptor.grantedAuthorities.find(
        (v) => v.authority === role && v.enterpriseID == enterpriseID
      );
      if (authority) {
        return true;
      }
      return false;
    default:
      return false;
  }
};

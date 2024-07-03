//import { csrf } from "./csrf";
import Cookies from 'js-cookie';
export enum Folder {
  Profile = "profile_img",
  Enterprise = "logo_img",
  Banner = "banner_img",
}
export const uploadFile = async (
  file: File,
  folder: Folder,
  enterpriseID?: string
): Promise<boolean> => {
  // const token = await csrf();
  // if (token != "") {
    const blob = new Blob([file], { type: file.type });
    const form = new FormData();
    const fileName = file.name
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "")
      .replace(/-/g, "");
    form.append(folder, blob, fileName);
    const request = new XMLHttpRequest();
    const bearer = Cookies.get("authorization") || "";
    let url = "http://3.143.110.143:5555" + `?folder=${folder}`;
    if (enterpriseID) {
      url = url + `&id=${enterpriseID}`;
    }
    request.open("POST", url, true);
//    request.withCredentials = true;
    //request.setRequestHeader("x-csrf-token", token);
    request.setRequestHeader("authorization", `Bearer ${bearer}`)
    request.send(form);
    request.onreadystatechange = function () {
      console.log(request.readyState);
      if (request.readyState === 4 && request.status === 200) {
        return true;
      }
    };
    return true;
  // }
  // throw new Error("No CSRF token");
};

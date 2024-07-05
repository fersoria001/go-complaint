//import { csrf } from "./csrf";
import Cookies from 'js-cookie';
import { csrf } from './csrf';
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
  const csrftoken = await csrf();
  if (csrftoken != "") {
    const blob = new Blob([file], { type: file.type });
    const form = new FormData();
    const fileName = file.name
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "")
      .replace(/-/g, "");
    form.append(folder, blob, fileName);
    const request = new XMLHttpRequest();
    const bearer = Cookies.get("authorization") || "";
    let url = "https://api.go-complaint.com/upload" + `?folder=${folder}`;
    if (enterpriseID) {
      url = url + `&id=${enterpriseID}`;
    }
    request.open("POST", url, true);
    request.withCredentials = true;
    request.setRequestHeader("x-csrf-token", csrftoken);
    request.setRequestHeader("authorization", `${bearer}`)
    request.send(form);
    request.onreadystatechange = function () {
      if (request.readyState === 4 && request.status === 200) {
        return true;
      }
    };
    return true;
  }
  throw new Error("No CSRF token");
};

/* eslint-disable @typescript-eslint/no-explicit-any */
import { csrf } from "./csrf";
import { deleteLinebreaks } from "./delete_line_breaks";
import Cookies from 'js-cookie';
export const Publish = async <T>(
    mutationFn: (data: T) => string,
    arg: T,
    subscriptionID: string
  ): Promise<boolean> => {
    const token = await csrf();
    if (token != "") {
      const authorization = Cookies.get("Authorization");
      const headers: any = {
        "Content-Type": "application/json",
        "x-csrf-token": token,
        "subscription-id": subscriptionID,
      };
      if (authorization && authorization != "") {
        headers["Authorization"] = authorization;
      }
      const response = await fetch("https://api.go-complaint.com/graphql/publish", {
        method: "POST",
        headers: headers,
        credentials: "include",
        body: JSON.stringify({ query: deleteLinebreaks(mutationFn(arg)) }),
      });
      if (response.status != 200) {
        throw new Error(`Failed to publish: ${response}`);
      }
      return true;
    }
    throw new Error("No CSRF token");
  };
  
import { csrf } from "./csrf";
import { deleteLinebreaks } from "./delete_line_breaks";

export const Publish = async <T>(
    mutationFn: (data: T) => string,
    arg: T,
    subscriptionID: string
  ): Promise<boolean> => {
    const token = await csrf();
    if (token != "") {
      const response = await fetch("http://3.143.110.143:5555/graphql/publish", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "x-csrf-token": token,
          "subscription-id": subscriptionID,
        },
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
  
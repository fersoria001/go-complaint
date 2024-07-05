import { syncParseSchema } from "./parse_schema";
import { Query, SignInQuery, SignInType } from "./queries";
import { ErrorType, SignInSchema } from "./types";
import Cookies from "js-cookie";
export const signIn = async (
  email: string,
  password: string,
  rememberMe: boolean
): Promise<ErrorType> => {
  const { data, errors } = syncParseSchema(
    { email, password, rememberMe },
    SignInSchema
  );
  if (Object.keys(errors).length > 0) {
    return errors;
  }
  return await Query<string>(SignInQuery, SignInType, [
    data.email,
    data.password,
    data.rememberMe,
  ])
    .then((res) => {
      const date = new Date();
      date.setTime(date.getTime() + 1 * 24 * 60 * 60 * 1000);
      Cookies.set("Authorization", `Bearer ${res}`, {
        path: "/",
        expires: date,
        sameSite: "Lax",
        secure:true,
        domain: ".go-complaint.com",
        
      });
      return {};
    })
    .catch((error) => {
      console.error("Error signing in", error);
      return { form: error.message };
    });
};

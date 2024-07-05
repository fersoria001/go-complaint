import { syncParseSchema } from "./parse_schema";
import { LoginQuery, LoginQueryType, LoginType, Query } from "./queries";
import { ConfirmationCodeValidationSchema, ErrorType } from "./types";
import Cookies from "js-cookie";
export const login = async (confirmationCode: string): Promise<ErrorType> => {
  const { data, errors } = syncParseSchema(
    { confirmationCode },
    ConfirmationCodeValidationSchema
  );
  if (Object.keys(errors).length > 0){
    return errors;
  }
  return await Query<LoginType>(LoginQuery, LoginQueryType, [data.confirmationCode])
    .then((res) => {
      Cookies.remove("Authorization");
      const date = new Date();
      date.setTime(date.getTime() + 1 * 24 * 60 * 60 * 1000);
      Cookies.set("Authorization", `Bearer ${res.token}`, {
        path: "/",
        expires: date,
      });
      console.warn("Login success",res);
      return {};
    })
    .catch((error) => {
      console.error("Error confirmation code login", error);
      return { form: error };
    });
};

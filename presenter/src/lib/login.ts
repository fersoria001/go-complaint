/* eslint-disable @typescript-eslint/no-explicit-any */
import { syncParseSchema } from "./parse_schema";
import { LoginQuery, LoginQueryType, LoginType, Query } from "./queries";
import { ConfirmationCodeValidationSchema, ErrorType } from "./types";
import Cookies from "js-cookie";
export const login = async (confirmationCode: string): Promise<ErrorType> => {
  const { data, errors } = syncParseSchema(
    { confirmationCode },
    ConfirmationCodeValidationSchema
  );
  if (Object.keys(errors).length > 0) {
    return errors;
  }
  try {
    const token = await Query<LoginType>(LoginQuery, LoginQueryType, [
      data.confirmationCode,
    ]);
    Cookies.remove("Authorization");
    const date = new Date();
    date.setTime(date.getTime() + 1 * 24 * 60 * 60 * 1000);
    Cookies.set("Authorization", `Bearer ${token.token}`, {
      path: "/",
      expires: date,
      // secure:true,
      // domain: ".go-complaint.com",
    });
    return {};
  } catch (error: any) {
    return { form: error };
  }
};

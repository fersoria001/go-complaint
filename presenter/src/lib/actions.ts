/* eslint-disable @typescript-eslint/no-explicit-any */
import {
  CreateEnterprise,
  RegisterEnterpriseSchema,
  SignInSchema,
  SignUpSchema,
  User,
} from "./types";
import {
  Mutation,
  CreateUserMutation,
  CreateEnterpriseMutation,
} from "./mutations";
import { redirect } from "react-router-dom";
import { Query, SignIn, SignInQuery } from "./queries";
import Cookies from "universal-cookie";
export const SignUpAction = async ({ request }: any) => {
  const formData = await request.formData();
  const updates = Object.fromEntries(formData);

  const parsed = SignUpSchema.safeParse(updates);
  const errors: { [key: string]: string } = {};
  let errorPath: string;
  console.log(parsed);
  if (!parsed.success) {
    parsed.error.errors.forEach((error) => {
      errorPath = error.path.join("");
      errors[errorPath] = error.message;
    });
    return errors;
  }
  const newUser: User = {
    email: parsed.data.email,
    password: parsed.data.password,
    firstName: parsed.data.firstName,
    lastName: parsed.data.lastName,
    birthDate: parsed.data.birthDate.getMilliseconds().toString(),
    phone: parsed.data.phoneCode.startsWith("+")
      ? parsed.data.phoneCode + parsed.data.phone
      : "+" + parsed.data.phoneCode + parsed.data.phone,
    country: parsed.data.country,
    county: parsed.data.county,
    city: parsed.data.city,
  };
  await Mutation<User>(CreateUserMutation, newUser);
  return redirect("/success/register");
};

export const SignInAction = async ({ request }: any) => {
  const formData = await request.formData();
  const updates = Object.fromEntries(formData);
  const parsed = SignInSchema.safeParse(updates);
  const errors: { [key: string]: string } = {};
  let errorPath: string;
  if (!parsed.success) {
    parsed.error.errors.forEach((error) => {
      errorPath = error.path.join("");
      errors[errorPath] = error.message;
    });
    return errors;
  }
  const newSignIn = [
    parsed.data.email,
    parsed.data.password,
    parsed.data.rememberMe,
  ];
  const token = await Query<string>(SignInQuery, SignIn, newSignIn);
  const cookies = new Cookies();
  const tokenExpires = expirationTime(new Date());
  cookies.set("Authorization", `Bearer ${token}`, {
    path: "/",
    expires: tokenExpires,
  });
  return redirect("/success/login");
};
function expirationTime(date: Date): Date {
  const newDate = new Date(date.getTime());
  newDate.setTime(newDate.getTime() + 24 * 60 * 60 * 1000);
  return newDate;
}
export const Logout = () => {
  const cookies = new Cookies();
  cookies.remove("Authorization");
  return redirect("/");
};

export const RegisterEnterpriseAction = async ({ request }: any) => {
  const formData = await request.formData();
  const updates = Object.fromEntries(formData);

  const parsed = await RegisterEnterpriseSchema.safeParseAsync(updates);
  const errors: { [key: string]: string } = {};
  let errorPath: string;

  if (!parsed.success) {
    parsed.error.errors.forEach((error) => {
      errorPath = error.path.join("");
      errors[errorPath] = error.message;
    });
    return errors;
  }
  const newEnterprise: CreateEnterprise = {
    name: parsed.data.name,
    email: parsed.data.email,
    website: parsed.data.website,
    phone: parsed.data.phoneCode.startsWith("+")
      ? parsed.data.phoneCode + parsed.data.phone
      : "+" + parsed.data.phoneCode + parsed.data.phone,
    industry: parsed.data.industry,
    country: parsed.data.country,
    county: parsed.data.county,
    city: parsed.data.city,
    foundationDate: parsed.data.foundationDate.getMilliseconds().toString(),
  };
  await Mutation<CreateEnterprise>(CreateEnterpriseMutation, newEnterprise);
  return redirect("/success/register enterprise");
};

import {
  ChangePasswordMutation,
  Mutation,
  UpdateUserMutation,
  UpdateUserMutation2,
} from "../../../lib/mutations";
import { UpdateUserType } from "../../../lib/types";

export async function updateGenre(
  genre: string,
  pronoun: string
): Promise<boolean> {
  const promises = [
    Mutation<UpdateUserType>(UpdateUserMutation, {
      updateType: "gender",
      value: genre,
    }),
    Mutation<UpdateUserType>(UpdateUserMutation, {
      updateType: "pronoun",
      value: pronoun,
    }),
  ];
  const res = await Promise.all(promises);
  return res.every((r) => r);
}

export async function updateFirstName(firstName: string): Promise<boolean> {
  const res = await Mutation<UpdateUserType>(UpdateUserMutation, {
    updateType: "firstName",
    value: firstName,
  });
  return res;
}

export async function updateLastName(lastName: string): Promise<boolean> {
  const res = await Mutation<UpdateUserType>(UpdateUserMutation, {
    updateType: "lastName",
    value: lastName,
  });
  return res;
}

export async function updatePhoneNumber(phoneNumber: string): Promise<boolean> {
  const res = await Mutation<UpdateUserType>(UpdateUserMutation, {
    updateType: "phone",
    value: phoneNumber,
  });
  return res;
}

export async function updateAddress(
  countryID: number,
  countyID: number,
  cityID: number
): Promise<boolean> {
  const promises = [
    Mutation<UpdateUserType>(UpdateUserMutation2, {
      updateType: "country",
      numberValue: countryID,
    }),
    Mutation<UpdateUserType>(UpdateUserMutation2, {
      updateType: "countryState",
      numberValue: countyID,
    }),
    Mutation<UpdateUserType>(UpdateUserMutation2, {
      updateType: "city",
      numberValue: cityID,
    }),
  ];
  const res = await Promise.all(promises);
  return res.every((r) => r);
}

export type ChangePasswordType = {
  oldPassword: string;
  newPassword: string;
};
export async function updatePassword(
  oldPassword: string,
  newPassword: string
): Promise<boolean> {
  return Mutation<ChangePasswordType>(ChangePasswordMutation, {
    oldPassword,
    newPassword,
  });
}

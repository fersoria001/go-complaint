import {
  Mutation,
  UpdateEnterpriseMutation,
  UpdateEnterpriseMutation2,
} from "../../../lib/mutations";
import { UpdateEnterpriseType } from "../../../lib/types";

export const updateWebsite = async (
  enterpriseID: string,
  website: string
): Promise<boolean> => {
  const ok = await Mutation<UpdateEnterpriseType>(UpdateEnterpriseMutation, {
    updateType: "website",
    enterpriseID: enterpriseID,
    value: website,
  });
  return ok;
};

export const updateEmail = async (
  enterpriseID: string,
  email: string
): Promise<boolean> => {
  const ok = await Mutation<UpdateEnterpriseType>(UpdateEnterpriseMutation, {
    updateType: "email",
    enterpriseID: enterpriseID,
    value: email,
  });
  return ok;
};

export const updatePhoneNumber = async (
  enterpriseID: string,
  phone: string
): Promise<boolean> => {
  const ok = await Mutation<UpdateEnterpriseType>(UpdateEnterpriseMutation, {
    updateType: "phone",
    enterpriseID: enterpriseID,
    value: phone,
  });
  return ok;
};

export const updateAddress = async (
  enterpriseID: string,
  country: number,
  county: number,
  city: number
): Promise<boolean> => {
  const promises = [
    Mutation<UpdateEnterpriseType>(UpdateEnterpriseMutation2, {
      updateType: "country",
      enterpriseID: enterpriseID,
      numberValue: country,
    }),
    Mutation<UpdateEnterpriseType>(UpdateEnterpriseMutation2, {
      updateType: "countryState",
      enterpriseID: enterpriseID,
      numberValue: county,
    }),
    Mutation<UpdateEnterpriseType>(UpdateEnterpriseMutation2, {
      updateType: "city",
      enterpriseID: enterpriseID,
      numberValue: city,
    }),
  ];
  const ok = await Promise.all(promises);
  return ok.every((x) => x);
};

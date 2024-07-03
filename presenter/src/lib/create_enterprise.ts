import { CreateEnterpriseMutation, Mutation } from "./mutations";
import { RegisterEnterprise } from "./types";

export const createEnterprise = async (
  registerEnterprise: RegisterEnterprise
) => {
  const ok = await Mutation<RegisterEnterprise>(
    CreateEnterpriseMutation,
    registerEnterprise
  );
  return ok;
};

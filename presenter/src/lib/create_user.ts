import { CreateUserMutation, Mutation } from "./mutations";
import { CreateUser } from "./types";

export async function createUser(createUser: CreateUser) {
  const ok = await Mutation<CreateUser>(CreateUserMutation, createUser);
  return ok;
}

import {
  FireEmployeeMutation,
  Mutation,
  PromoteEmployeeMutation,
} from "../../../lib/mutations";
import { FireEmployeeType, PromoteEmployeeType } from "../../../lib/types";

export const promoteEmployee = async ({
  enterpriseName,
  employeeID,
  position,
}: PromoteEmployeeType): Promise<boolean> => {
  return await Mutation<PromoteEmployeeType>(PromoteEmployeeMutation, {
    enterpriseName,
    employeeID,
    position,
  });
};

export const fireEmployee = async ({
  enterpriseName,
  employeeID,
}: FireEmployeeType): Promise<boolean> => {
  return await Mutation<FireEmployeeType>(FireEmployeeMutation, {
    enterpriseName,
    employeeID,
  });
};

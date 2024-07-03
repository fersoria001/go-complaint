import {
  LeaveEnterpriseMutation,
  LeaveEnterpriseType,
  Mutation,
} from "./mutations";

export const leaveEnterprise = async (
  leaveEnterprise: LeaveEnterpriseType
): Promise<boolean> => {
  return await Mutation<LeaveEnterpriseType>(
    LeaveEnterpriseMutation,
    leaveEnterprise
  );
};

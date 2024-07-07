/* eslint-disable @typescript-eslint/no-explicit-any */
import { Mutation, ReplyComplaintMutation } from "./mutations";
export async function replyComplaint(
  complaintID: string,
  replyAuthorID: string,
  replyBody: string,
  replyEnterpriseID: string
): Promise<boolean> {
  try {
    const ok = await Mutation(ReplyComplaintMutation, {
      complaintID,
      replyAuthorID,
      replyBody,
      replyEnterpriseID,
    });
    return ok;
  } catch (error: any) {
    console.log(error);
    return false
  }
}

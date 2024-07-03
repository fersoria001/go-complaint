import { Mutation, ReplyComplaintMutation } from "./mutations";
export async function replyComplaint(
  complaintID: string,
  replyAuthorID: string,
  replyBody: string,
  replyEnterpriseID: string
): Promise<boolean> {
  return await Mutation(ReplyComplaintMutation, {
    complaintID,
    replyAuthorID,
    replyBody,
    replyEnterpriseID,
  }).then((res) => {
    return res;
  });
}

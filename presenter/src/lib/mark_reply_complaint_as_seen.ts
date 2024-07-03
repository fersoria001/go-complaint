/* eslint-disable @typescript-eslint/no-explicit-any */
import { MarkComplaintRepliesAsSeenMutation, Mutation } from "./mutations";
import { MarkComplaintRepliesAsSeenType, Reply } from "./types";

export const markAsSeen = async (
  complaintId: string,
  reply: Reply[]
): Promise<boolean> => {
  const ids = reply.map((r) => r.id);
  const ok = await Mutation<MarkComplaintRepliesAsSeenType>(
    MarkComplaintRepliesAsSeenMutation,
    {
      complaintID: complaintId,
      repliesID: ids,
    }
  );
  console.error(ok);
  return ok
};

export const isReply = (data: any): boolean => {
  if (Array.isArray(data)) return false;
  const reply = data as Reply;
  return reply.complaintStatus !== undefined;
}
export const isReplies = (data: any): boolean => {
  if (Array.isArray(data)) return true;
  return false;
}
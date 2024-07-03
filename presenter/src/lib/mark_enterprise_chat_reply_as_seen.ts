/* eslint-disable @typescript-eslint/no-explicit-any */
import { MarkReplyChatAsSeenMutation, Mutation } from "./mutations";
import {
  EnterpriseChatReplyType,
  MarkReplyChatAsSeenType,
} from "./types";

export const markAsSeen = async (
  enterpriseName: string,
  replies: EnterpriseChatReplyType[]
): Promise<boolean> => {
  if (replies.length <= 0) return false;
  const ids = replies.map((r) => r.id);
  const chatID = replies[0].chatID
  const ok = await Mutation<MarkReplyChatAsSeenType>(
    MarkReplyChatAsSeenMutation,
    {
      chatID,
      enterpriseName,
      repliesID: ids,
    }
  );
  console.error(ok);
  return ok;
};

export const isReplies = (data: any): boolean => {
  if (Array.isArray(data)) return true;
  return false;
};

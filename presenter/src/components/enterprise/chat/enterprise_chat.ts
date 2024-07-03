/* eslint-disable @typescript-eslint/no-explicit-any */
import { Mutation, ReplyChatMutation } from "../../../lib/mutations";
import {
  UserDescriptor,
  User,
  ReplyEnterpriseChatType,
  EnterpriseChatReplyType,
} from "../../../lib/types";

export async function replyEnterpriseChat(
  enterpriseID: string,
  descriptor: UserDescriptor,
  user: User,
  content: string
): Promise<boolean> {
  const res = await Mutation<ReplyEnterpriseChatType>(ReplyChatMutation, {
    id: `chat:${enterpriseID}=${descriptor.email}#${user.email}`,
    enterpriseName: enterpriseID,
    senderID: descriptor.email,
    content: content,
  });
  return res;
}

export const isPair = (item: any): boolean => {
  if (!item) return false;
  const pair = item as Pair;
  return pair.one !== undefined && pair.two !== undefined;
};
export const isReply = (item: any): boolean => {
  if (!item) return false;
  const reply = item as EnterpriseChatReplyType;
  return reply.content !== undefined && reply.user !== undefined;
};
export type Pair = {
  one: string;
  two: string;
};
export type TabType = {
  user: User;
  enterpriseID: string;
  index: number;
  isActive: boolean;
};

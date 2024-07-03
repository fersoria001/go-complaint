import Cookies from 'js-cookie';
import { deleteLinebreaks } from "./delete_line_breaks";
import { FeedbackType, Notifications, Reply, Subscription } from "./types";
import UnauthorizedError from "../components/error/UnauthorizedError";

/* eslint-disable @typescript-eslint/no-explicit-any */
export function ComplaintLastReplySubscription(id: string) {
  return `
    subscription {
        complaintLastReply(id: "${id}") {
            id
            complaintID
            senderID
            senderIMG
            senderName
            body
            createdAt
            read
            readAt
            updatedAt
            isEnterprise
            enterpriseID
            complaintStatus
        }
}
    `;
}

export function ComplaintLastReplyReturnType(data: any): Reply {
  return data;
}
export const NotificationSubscription = (id: string): string => `
subscription {
  notifications(id:"${id}") {
    id
    title
    content
    thumbnail
    occurredOn
    seen
    link
  }
}
`;

export const NotificationReturnType = (data: any): Notifications => {
  return data;
};

export const FeedbackSubscription = (): string => `
` 
export const FeedbackReturnType = (data: any): FeedbackType => {
  return data;
};

export function createSubscription<T>(
  id: string,
  subscriptionFunc: (...args: any[]) => string,
  args: any[],
  subscriptionReturnType: (data: any) => T
): Subscription<T> {
  const cookie = Cookies.get("Authorization");
  const token = cookie?.slice(7);
  const subscription = deleteLinebreaks(subscriptionFunc(...args));
  if (!token) {
    throw new UnauthorizedError();
  }
  const connection_ack = {
    type: "connection_ack",
    payload: {
      query: subscription,
      subscription_id: id,
      token: token,
    },
  };
  return {
    connection_ack,
    subscription,
    subscriptionID: id,
    subscriptionReturnType,
  };
}

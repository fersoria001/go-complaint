import { NotificationQuery, NotificationTypeList, Query } from "./queries";
import { Notifications } from "./types";

export const fetchNotifications = async (
  id: string
): Promise<Notifications[]> => {
  return await Query<Notifications[]>(NotificationQuery, NotificationTypeList, [
    id,
  ])
    .then((data) => {
      return data;
    })
    .catch((error) => {
      console.error("Error fetching notifications", error);
      return [];
    });
};

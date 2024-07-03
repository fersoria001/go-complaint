import { Query, UserDescriptorQuery, UserDescriptorType } from "./queries";
import { UserDescriptor } from "./types";

export const fetchUserDescriptor = async (): Promise<UserDescriptor> => {
    const descriptor = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType,
      []
    );
    return descriptor;
};

import Cookies from "universal-cookie";
import {
  Query,
  CountriesQuery,
  CountryListType,
  UserDescriptorQuery,
  UserDescriptorType,
  IndustriesQuery,
  IndustryListType,
  FindReceiverQuery,
  ReceiverTypeList,
  DraftQuery,
  DraftTypeList,
  SentQuery,
  SentTypeList,
  ComplaintQuery,
  ComplaintType,
  EnterpriseQuery,
  EnterpriseType,
  UsersForHiringQuery,
  UsersForHiringType,
  UserQuery,
  UserType,
  EmployeeQuery,
  EmployeeType,
  EmployeesQuery,
  EmployeesTypeList,
} from "./queries";
import {
  AuthMsg,
  Complaint,
  ComplaintTypeList,
  Country,
  Employee,
  EmployeeWaitingForApproval,
  Enterprise,
  EnterpriseNotifications,
  Industry,
  newAuthMsg,
  newSender,
  Receiver,
  Sender,
  User,
  UserDescriptor,
  UserNotifications,
  UsersForHiring,
} from "./types";

export async function RootLoader(): Promise<{
  user: UserDescriptor | null;
  notifications: UserNotifications | null;
}> {
  const cookies = new Cookies();
  if (cookies.get("Authorization")) {
    const user = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType
    );
    const notifications = await fetch(
      import.meta.env.VITE_NOTIFICATIONS_ENDPOINT + "?type=user",
      {
        method: "GET",
        credentials: "include",
      }
    );
    const data = await notifications.json();
    if (data) {
      return { user, notifications: data };
    }
  }
  return { user: null, notifications: null };
}
export async function ProfileLoader(): Promise<{
  user: UserDescriptor | null;
  notifications: UserNotifications | null;
}> {
  const cookies = new Cookies();
  if (cookies.get("Authorization")) {
    const user = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType
    );
    const notifications = await fetch(
      import.meta.env.VITE_NOTIFICATIONS_ENDPOINT + "?type=user",
      {
        method: "GET",
        credentials: "include",
      }
    );
    const data = await notifications.json();
    if (data) {
      return { user, notifications: data };
    }
  }
  return { user: null, notifications: null };
}
export async function SignUpLoader(): Promise<Country[]> {
  return await Query<Country[]>(CountriesQuery, CountryListType);
}

export async function RegisterEnterpriseLoader(): Promise<{
  countries: Country[];
  industries: Industry[];
}> {
  const result: { countries: Country[]; industries: Industry[] } = {
    countries: [],
    industries: [],
  };
  result.countries = await Query<Country[]>(CountriesQuery, CountryListType);
  result.industries = await Query<Industry[]>(
    IndustriesQuery,
    IndustryListType
  );
  return result;
}

export async function FindReceiverLoader({
  request,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ receivers: Receiver[]; term: string }> {
  const url = new URL(request.url);
  const term = url.searchParams.get("term") || "";
  const receivers = await Query<Receiver[]>(
    FindReceiverQuery,
    ReceiverTypeList,
    [term]
  );

  return { receivers, term };
}

export async function UserSendComplaintLoader(): Promise<UserDescriptor | null> {
  const cookies = new Cookies();
  if (cookies.get("Authorization")) {
    return await Query<UserDescriptor>(UserDescriptorQuery, UserDescriptorType);
  }
  return null;
}
export async function EnterpriseSentComplaintLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<Enterprise> {
  const enterprise = await Query<Enterprise>(EnterpriseQuery, EnterpriseType, [
    params.id,
  ]);
  return enterprise;
}

export async function UserInboxLoader({
  request,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ complaintList: ComplaintTypeList | null; page: string }> {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const offset = (parseInt(page) - 1) * 10;
  const cookies = new Cookies();
  if (cookies.get("Authorization")) {
    const descriptor = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType
    );
    const complaintList = await Query<ComplaintTypeList>(
      DraftQuery,
      DraftTypeList,
      [descriptor.email, 10, offset]
    );
    return { complaintList, page };
  }
  return { complaintList: null, page };
}
export async function EnterpriseInboxLoader({
  request,
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ complaintList: ComplaintTypeList | null; page: string }> {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const offset = (parseInt(page) - 1) * 10;
  if (params.id) {
    const complaintList = await Query<ComplaintTypeList>(
      DraftQuery,
      DraftTypeList,
      [params.id, 10, offset]
    );
    return { complaintList, page };
  }
  return { complaintList: null, page };
}

export async function UserSentLoader({
  request,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ complaintList: ComplaintTypeList | null; page: string }> {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const offset = (parseInt(page) - 1) * 10;
  const cookies = new Cookies();
  if (cookies.get("Authorization")) {
    const descriptor = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType
    );
    const complaintList = await Query<ComplaintTypeList>(
      SentQuery,
      SentTypeList,
      [descriptor.email, 10, offset]
    );
    return { complaintList, page };
  }
  return { complaintList: null, page };
}
export async function EnterpriseSentLoader({
  request,
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ complaintList: ComplaintTypeList | null; page: string }> {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const offset = (parseInt(page) - 1) * 10;
  if (params.id) {
    const complaintList = await Query<ComplaintTypeList>(
      SentQuery,
      SentTypeList,
      [params.id, 10, offset]
    );
    return { complaintList, page };
  }
  return { complaintList: null, page };
}

export async function ComplaintLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{
  complaints: Complaint[] | null;
  id: string;
  sender: Sender | null;
  authMsg: AuthMsg | null;
}> {
  const cookies = new Cookies();
  const bearer = cookies.get("Authorization");
  if (bearer) {
    const descriptor = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType
    );
    
    const complaints = await Query<Complaint[]>(ComplaintQuery, ComplaintType, [
      params.complaintID,
    ]);
    let authMsg = null;
    let sender = null;
    if (params.id) {
      authMsg = newAuthMsg(bearer, params.id);
      const enterprise = await Query<Enterprise>(
        EnterpriseQuery,
        EnterpriseType,
        [params.id]
      );
      sender = newSender(enterprise.logoIMG, enterprise.name);
    } else {
      authMsg = newAuthMsg(bearer);
      sender = newSender(descriptor.profileIMG, descriptor.fullName);
    }
    
    return { complaints, id: params.complaintID, sender, authMsg };
  }
  return { complaints: null, id: params.complaintID, sender: null, authMsg: null };
}

export async function EnterpriseLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{
  enterprise: Enterprise | null;
  notifications: EnterpriseNotifications | null;
}> {
  const enterprise = await Query<Enterprise>(EnterpriseQuery, EnterpriseType, [
    params.id,
  ]);
  const notifications = await fetch(
    import.meta.env.VITE_NOTIFICATIONS_ENDPOINT +
      `?type=enterprise&id=${params.id}`,
    {
      method: "GET",
      credentials: "include",
    }
  );
  const data = await notifications.json();
  if (data) {
    return { enterprise, notifications: data };
  }
  return { enterprise: null, notifications: null };
}

export async function HiringLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<string> {
  const usersForHiring = await Query<UsersForHiring>(
    UsersForHiringQuery,
    UsersForHiringType,
    [params.id, 10, 0]
  );
  const pages =
    Math.floor(usersForHiring.count / usersForHiring.currentLimit) + 1;
  return pages.toString();
}

export async function HireLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ enterprise: Enterprise | null; user: User | null }> {
  const enterprise = await Query<Enterprise>(EnterpriseQuery, EnterpriseType, [
    params.id,
  ]);
  const user = await Query<User>(UserQuery, UserType, [params.userID]);
  if (enterprise && user) {
    return { enterprise, user };
  }
  return { enterprise: null, user: null };
}

export async function AcceptInvitationLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<UserDescriptor | null> {
  const cookies = new Cookies();
  if (cookies.get("Authorization")) {
    const notification = await fetch(
      import.meta.env.VITE_NOTIFICATION_ENDPOINT + `${params.id}`,
      {
        method: "GET",
        credentials: "include",
      }
    );
    const data = await notification.json();
    console.log(data);
    if (data) {
      return data;
    }
  }
  return null;
}
export async function PendingHiresLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ eventID: string; employee: Employee }[] | null> {
  const notificationsUrl = `${
    import.meta.env.VITE_NOTIFICATIONS_ENDPOINT
  }?type=enterprise&id=${params.id}`;

  try {
    const notificationsResponse = await fetch(notificationsUrl, {
      method: "GET",
      credentials: "include",
    });

    if (!notificationsResponse.ok) {
      throw new Error(
        `Notifications API request failed with status: ${notificationsResponse.status}`
      );
    }

    const notificationsData = await notificationsResponse.json();

    if (!notificationsData) {
      return null; // No notifications data, return null as before
    }

    const employeeWaitingForApproval =
      notificationsData.employee_waiting_for_approval;

    if (
      !employeeWaitingForApproval ||
      employeeWaitingForApproval.length === 0
    ) {
      return null; // No pending hires, return null
    }
    const notSeen = employeeWaitingForApproval.filter(
      (hire: EmployeeWaitingForApproval) => !hire.seen
    );
    // Use Promise.all to efficiently fetch employee details concurrently
    const employeePromises = notSeen.map((hire: EmployeeWaitingForApproval) =>
      Query<Employee>(EmployeeQuery, EmployeeType, [hire.employee_id])
    );

    const resolvedEmployees = await Promise.all(employeePromises);
    const results = resolvedEmployees.map((employee, index) => ({
      eventID: notSeen[index].event_id,
      employee,
    }));
    return results as { eventID: string; employee: Employee }[];
  } catch (error) {
    console.error("Error fetching pending hires:", error);
    return null; // Handle errors gracefully, return null
  }
}

export async function EmployeesLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any) {
  const employees = await Query<Employee[]>(EmployeesQuery, EmployeesTypeList, [
    params.id,
  ]);
  return employees;
}

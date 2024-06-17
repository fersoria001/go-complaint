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
  EmployeesQuery,
  EmployeesTypeList,
  InfoForReviewQuery,
  InfoForReviewType,
  OfficeTypeList,
  OfficesQuery,
  SolvedReviewQuery,
  SolvedReviewType,
  EmployeeQuery,
  EmployeeType,
} from "./queries";
import {
  AuthMsg,
  Complaint,
  ComplaintRated,
  ComplaintTypeList,
  Country,
  Employee,
  Enterprise,
  EnterpriseNotifications,
  Industry,
  InfoForReview,
  newAuthMsg,
  Office,
  Receiver,
  Sender,
  SolvedReview,
  User,
  UserDescriptor,
  UserLog,
  UserNotifications,
  UsersForHiring,
  WaitingForReview,
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
      return { user, notifications: data as UserNotifications };
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
  request,
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{
  complaints: Complaint | null;
  complaintID: string;
  sender: Sender | null;
  authMsg: AuthMsg | null;
  office: Office | null;
}> {
  const url = new URL(request.url);
  const cookies = new Cookies();
  const bearer = cookies.get("Authorization");
  if (bearer) {
    const descriptor = await Query<UserDescriptor>(
      UserDescriptorQuery,
      UserDescriptorType
    );
    const complaints = await Query<Complaint>(ComplaintQuery, ComplaintType, [
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
      sender = {
        thumbnail: enterprise.logoIMG,
        fullName: descriptor.fullName,
        isEnterprise: true,
        enterpriseID: enterprise.name,
      };
    } else {
      authMsg = newAuthMsg(bearer);
      sender = {
        thumbnail: descriptor.profileIMG,
        fullName: descriptor.fullName,
        isEnterprise: false,
        enterpriseID: "",
      };
    }

    if (url.pathname.includes("office")) {
      const offices = await Query<Office[]>(OfficesQuery, OfficeTypeList);
      const currentOffice = offices.find(
        (office) => office.employeeID.split("-")[0] === params.id
      );
      if (currentOffice) {
        return {
          complaints,
          complaintID: params.complaintID,
          sender,
          authMsg,
          office: currentOffice,
        };
      }
    }
    return {
      complaints,
      complaintID: params.complaintID,
      sender,
      authMsg,
      office: null,
    };
  }
  return {
    complaints: null,
    complaintID: params.complaintID,
    sender: null,
    authMsg: null,
    office: null,
  };
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
any): Promise<
  | { eventID: string; employee: User; position: string; pendingDate: string }[]
  | null
> {
  const notifications = await fetchEnterpriseNotifications(params.id);
  if (notifications.employee_waiting_for_approval.length < 1) {
    return null;
  }
  const pendingHiresPromises = [];
  for (let i = 0; i < notifications.employee_waiting_for_approval.length; i++) {
    const employee = Query<User>(UserQuery, UserType, [
      notifications.employee_waiting_for_approval[i].invited_user_id,
    ]);
    pendingHiresPromises.push(employee);
  }
  const resolvedPendingHires = await Promise.all(pendingHiresPromises);
  const pendingHires = [];
  for (let i = 0; i < resolvedPendingHires.length; i++) {
    pendingHires.push({
      eventID: notifications.employee_waiting_for_approval[i].id,
      employee: resolvedPendingHires[i],
      position:
        notifications.employee_waiting_for_approval[i].proposed_position,
      pendingDate: new Date(
        parseInt(notifications.employee_waiting_for_approval[i].occurred_on)
      ).toLocaleDateString(),
    });
  }
  return pendingHires;
}

async function fetchEnterpriseNotifications(
  id: string
): Promise<EnterpriseNotifications> {
  const notificationsUrl = `${
    import.meta.env.VITE_NOTIFICATIONS_ENDPOINT
  }?type=enterprise&id=${id}`;
  const notificationsResponse = await fetch(notificationsUrl, {
    method: "GET",
    credentials: "include",
  });
  const notificationsData = await notificationsResponse.json();
  return notificationsData;
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

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export async function ComplaintsSolvedListPageLoader({ request }: any) {
  const url = new URL(request.url);
  const employeeID = url.searchParams.get("employee");
  const employee = await Query<Employee>(EmployeeQuery, EmployeeType, [
    employeeID,
  ]);
  let resolved: Complaint[] = [];
  if (employee && employee.complaintsSolvedIDs.length > 0) {
    const promises = [];
    for (let i = 0; i < employee.complaintsSolvedIDs.length; i++) {
      promises.push(
        Query<Complaint>(ComplaintQuery, ComplaintType, [
          employee.complaintsSolvedIDs[i],
        ])
      );
    }
    resolved = await Promise.all(promises);
  }
  return resolved;
}

export async function FeedbackPageLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ user: UserDescriptor; complaint: Complaint }> {
  const user = await Query<UserDescriptor>(
    UserDescriptorQuery,
    UserDescriptorType
  );
  const complaint = await Query<Complaint>(ComplaintQuery, ComplaintType, [
    params.complaintID,
  ]);
  return { user, complaint };
}

export async function ReviewComplaintLoader(): Promise<{
  pendingReviews: { eventID: string; info: InfoForReview }[] | null;
  solvedReviewss:
    | {
        eventID: string;
        solvedReview: SolvedReview;
        occurredOn: string;
      }[]
    | null;
}> {
  const bearer = new Cookies().get("Authorization");
  if (bearer) {
    let pendingReviews: { eventID: string; info: InfoForReview }[] | null =
      null;
    let solvedReviewss:
      | {
          eventID: string;
          solvedReview: SolvedReview;
          occurredOn: string;
        }[]
      | null = null;
    const userNotificationss = await userNotifications();
    const userLogg = await userLog();
    if (userNotificationss) {
      pendingReviews = await infoForReview(
        userNotificationss.waiting_for_review
      );
    }
    if (userLogg) {
      solvedReviewss = await solvedReviews(userLogg.complaint_rated);
    }
    return { pendingReviews, solvedReviewss };
  } else {
    return { pendingReviews: null, solvedReviewss: null };
  }
}
async function userNotifications(): Promise<UserNotifications | null> {
  const notificationsUrl = `${
    import.meta.env.VITE_NOTIFICATIONS_ENDPOINT
  }?type=user`;
  const response = await fetch(notificationsUrl, {
    method: "GET",
    credentials: "include",
  });
  const notificationsData = await response.json();
  if (!notificationsData) {
    return null;
  }
  return notificationsData;
}
async function userLog(): Promise<UserLog | null> {
  const eventsLogURL = `${import.meta.env.VITE_EVENT_LOG_ENDPOINT}?type=user`;
  const res = await fetch(eventsLogURL, {
    method: "GET",
    credentials: "include",
  });
  const eventLogData = await res.json();
  if (!eventLogData) {
    return null;
  }
  return eventLogData;
}
async function infoForReview(
  waitingForReview: WaitingForReview[]
): Promise<{ eventID: string; info: InfoForReview }[] | null> {
  const pendingReviews: { eventID: string; info: InfoForReview }[] = [];
  const infoPromisesArray = [];
  for (let i = 0; i < waitingForReview.length; i++) {
    const info = Query<InfoForReview>(InfoForReviewQuery, InfoForReviewType, [
      waitingForReview[i].triggered_by,
      waitingForReview[i].complaint_id,
    ]);
    infoPromisesArray.push(info);
  }
  const resolvedInfo = await Promise.all(infoPromisesArray);
  for (let i = 0; i < resolvedInfo.length; i++) {
    pendingReviews.push({
      eventID: waitingForReview[i].event_id,
      info: resolvedInfo[i],
    });
  }
  return pendingReviews;
}

async function solvedReviews(
  complaintRateds: ComplaintRated[]
): Promise<
  { eventID: string; solvedReview: SolvedReview; occurredOn: string }[] | null
> {
  const solvedReviewsPromises = [];
  for (let i = 0; i < complaintRateds.length; i++) {
    const solvedReview = Query<SolvedReview>(
      SolvedReviewQuery,
      SolvedReviewType,
      [complaintRateds[i].complaint_id, complaintRateds[i].rated_by]
    );
    solvedReviewsPromises.push(solvedReview);
  }
  const resolvedSolvedReviews = await Promise.all(solvedReviewsPromises);
  const solvedReviews = [];
  for (let i = 0; i < resolvedSolvedReviews.length; i++) {
    solvedReviews.push({
      eventID: complaintRateds[i].event_id,
      solvedReview: resolvedSolvedReviews[i],
      occurredOn: complaintRateds[i].occurred_on,
    });
  }
  return solvedReviews;
}
export async function ReviewEnterpriseComplaintLoader({
  params,
}: // eslint-disable-next-line @typescript-eslint/no-explicit-any
any): Promise<{ notificationID: string; info: InfoForReview }[] | null> {
  const bearer = new Cookies().get("Authorization");
  if (bearer) {
    const notificationsUrl = `${
      import.meta.env.VITE_NOTIFICATIONS_ENDPOINT
    }?type=enterprise&id=${params.id}`;
    const response = await fetch(notificationsUrl, {
      method: "GET",
      credentials: "include",
    });
    const notificationsData = await response.json();
    if (!notificationsData) {
      return null;
    }
    const waitingForReview =
      notificationsData.waiting_for_review as WaitingForReview[];
    if (!waitingForReview || waitingForReview.length === 0) {
      return null;
    }
    const infoPromisesArray = [];
    for (let i = 0; i < waitingForReview.length; i++) {
      const info = Query<InfoForReview>(InfoForReviewQuery, InfoForReviewType, [
        waitingForReview[i].triggered_by,
        waitingForReview[i].complaint_id,
      ]);
      infoPromisesArray.push(info);
    }
    const resolvedInfo = await Promise.all(infoPromisesArray);
    const result: { notificationID: string; info: InfoForReview }[] = [];
    for (let i = 0; i < resolvedInfo.length; i++) {
      result.push({
        notificationID: waitingForReview[i].event_id,
        info: resolvedInfo[i],
      });
    }
    return result;
  }
  return null;
}

/* eslint-disable @typescript-eslint/no-explicit-any */
import UnauthorizedError from "../components/error/UnauthorizedError";
import PasswordNotMatchError from "../components/sign-in/PasswordNotMatchError";
import UserNotFoundError from "../components/sign-in/UserNotFoundError";
import { csrf } from "./csrf";
import { deleteLinebreaks } from "./delete_line_breaks";
import {
  City,
  ComplaintInfo,
  ComplaintReviewType,
  ComplaintType,
  ComplaintTypeList,
  Country,
  CountryState,
  EmployeeType,
  Enterprise,
  FeedbackType,
  HiringInvitationType,
  HiringProccessList,
  Industry,
  Notifications,
  Office,
  Receiver,
  SolvedReview,
  User,
  UserDescriptor,
  UsersForHiring,
} from "./types";

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export const CountriesQuery = (): string => "{Countries {id name phoneCode}}";
export const CountryListType = (data: any): Country[] => {
  return data.data.Countries as Country[];
};
export const CountryStatesQuery = (id: number): string => `
  {
    CountryStates(ID: ${id}) {
      id
      name
    }
  }    
`;
export const CountryStateListType = (data: any): CountryState[] => {
  return data.data.CountryStates as CountryState[];
};
export const CitiesQuery = (id: number): string => `
  {
    Cities(ID: ${id}) {
      id
      name
      countryCode
      latitude
      longitude
    }
  }
`;
export const CityListType = (data: any): City[] => {
  return data.data.Cities as City[];
};

export const SignInQuery = (
  email: string,
  password: string,
  rememberMe: boolean
): string => `
   {
    SignIn(email: "${email}", password: "${password}", rememberMe: ${rememberMe}) {
      token
    }
  }
`;
export const SignInType = (data: any): string => {
  if (data.errors) {
    switch (data.errors[0].message) {
      case "crypto/bcrypt: hashedPassword is not the hash of the given password":
        throw new PasswordNotMatchError();
      case "no rows in result set":
        throw new UserNotFoundError();
    }
  }
  return data.data.SignIn.token;
};
export const LoginQuery = (confirmationCode: string): string => `
{
  Login(confirmationCode: ${confirmationCode}) {
  token
  }
}
`;
export const LoginType = (data: any): string => {
  return data.data.Login.token;
};

export const UserDescriptorQuery = (): string => `
  {
    UserDescriptor {
      email
      fullName
      profileIMG
      gender
      pronoun
      loginDate
      ip
      device
      geolocation { latitude longitude }
      grantedAuthorities { enterpriseID authority }
    }
  }
`;
export const UserDescriptorType = (data: any): UserDescriptor => {
  if (data.errors) {
    switch (data.errors[0].message) {
      case "Unauthorized: User not found in context":
        throw new UnauthorizedError();
      default:
        throw new Error(data.errors[0].message);
    }
  }
  return data.data.UserDescriptor as UserDescriptor;
};

export const NotificationQuery = (id: string): string => `
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

export const NotificationTypeList = (data: any): Notifications[] => {
  if (data.errors) {
    console.error("NotificationTypeList err ", data.errors[0].message);
  }
  return data.data.notifications;
};

export const IndustriesQuery = (): string => `
  {
    Industries {
      id
      name
    }
  }
`;
export const IndustryListType = (data: any): Industry[] => {
  if (data.errors) {
    console.error("IndustryListType err ", data.errors[0].message);
  }
  return data.data.Industries as Industry[];
};
export const IsEnterpriseNameAvailableQuery = (name: string): string => `
  {
    IsEnterpriseNameAvailable(id: "${name}")
  }
`;
export const IsEnterpriseNameAvailable = (data: any): boolean => {
  if (data.errors) {
    console.error("IsEnterpriseNameAvailable err ", data.errors[0].message);
  }
  return data.data.IsEnterpriseNameAvailable;
};

export const FindComplaintReceiversQuery = (
  id: string,
  name: string
): string => `
  {
    FindComplaintReceivers(id:"${id}", term: "${name}") {
      id
      fullName
      thumbnail
    }
  }
`;

export const FindComplaintReceiversTypeList = (data: any): Receiver[] => {
  return data.data.FindComplaintReceivers;
};

export const IsValidReceiverQuery = (name: string): string => `
  {
    IsValidReceiver(ID: "${name}")
  }
`;
export const IsValidReceiver = (data: any): boolean => {
  return data.IsValidReceiver;
};

export const FindAuthorByIDQuery = (id: string): string => `
  {
    FindAuthorByID(id: "${id}") {
      fullName
      thumbnail
    }
  }
`;
export const FindAuthorByIDType = (data: any): Receiver => {
  if (data.errors) {
    console.error("FindAuthorByIDType err ", data.errors[0].message);
  }
  return data.data.FindAuthorByID;
};
export const DraftQuery = (
  id: string,
  limit: number,
  offset: number
): string => `
  {
    ComplaintInbox(id: "${id}", limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt

      }
      count
      currentLimit
      currentOffset
    }
  }
`;

export const DraftTypeList = (data: any): ComplaintTypeList => {
  if (data.errors) {
    console.error("DraftTypeList err ", data.errors[0].message);
  }
  return data.data.ComplaintInbox as ComplaintTypeList;
};

export const SearchInDraftQuery = (
  id: string,
  term: string,
  after: string,
  before: string,
  limit: number,
  offset: number
): string => `
  {
    ComplaintInboxSearch(id: "${id}", afterDate: "${after}",
     beforeDate: "${before}", term: "${term}",
     limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      status
      message { title description body }
      replies {
        id
        complaintID
        senderID
        createdAt
        read
        readAt
        isEnterprise
        enterpriseID
        }
      createdAt
      updatedAt
      }
      count
      currentLimit
      currentOffset
    }
  }
`;

export const SearchInDraftTypeList = (data: any): ComplaintTypeList => {
  if (data.errors) {
    console.error("DraftTypeList err ", data.errors[0].message);
  }
  return data.data.ComplaintInboxSearch as ComplaintTypeList;
};

export const ComplaintHistoryQuery = (
  id: string,
  term: string,
  after: string,
  before: string,
  limit: number,
  offset: number
): string => `
  {
    ComplaintHistory(id: "${id}", afterDate: "${after}",
     beforeDate: "${before}", term: "${term}",
     limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      }
      count
      currentLimit
      currentOffset
    }
  }
`;

export const ComplaintHistoryTypeList = (data: any): ComplaintTypeList => {
  if (data.errors) {
    console.error("DraftTypeList err ", data.errors[0].message);
  }
  return data.data.ComplaintHistory as ComplaintTypeList;
};

export const SentQuery = (
  id: string,
  limit: number,
  offset: number
): string => `
  {
    ComplaintsSent(id: "${id}", limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      receiverFullName
      receiverProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      }
      count
      currentLimit
      currentOffset
    }
  }
`;
export const SentTypeList = (data: any): ComplaintTypeList => {
  if (data.errors) {
    console.error("SentTypeList err ", data.errors[0].message);
  }
  return data.data.ComplaintsSent as ComplaintTypeList;
};
export const SentSearchQuery = (
  id: string,
  term: string,
  after: string,
  before: string,
  limit: number,
  offset: number
): string => `
  {
    ComplaintsSentSearch(id: "${id}", afterDate: "${after}",
     beforeDate: "${before}", term: "${term}",
     limit:${limit}, offset:${offset}) {
      complaints {
      id
      authorFullName
      authorProfileIMG
      receiverFullName
      receiverProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      replies {
        id
        complaintID
        senderID
        createdAt
        read
        readAt
        isEnterprise
        enterpriseID
        }
      }
      count
      currentLimit
      currentOffset
    }
  }
`;
export const SentSearchTypeList = (data: any): ComplaintTypeList => {
  if (data.errors) {
    console.error("SentTypeList err ", data.errors[0].message);
  }
  console.error("daata", data);
  return data.data.ComplaintsSentSearch as ComplaintTypeList;
};
export const ComplaintQuery = (id: string): string => `
  {
    Complaint(id: "${id}") {
      id
      authorID
      authorFullName
      authorProfileIMG
      receiverID
      receiverFullName
      receiverProfileIMG
      status
      message { title description body }
      createdAt
      updatedAt
      rating { rate comment }
      replies {
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
      }
    }
  }
`;

export const ComplaintQueryType = (data: any): ComplaintType => {
  if (data.errors) {
    console.error("complainttype err ", data.errors[0].message);
  }
  return data.data.Complaint as ComplaintType;
};

export const ComplaintsReceivedInfoQuery = (id: string) => `
{
  ComplaintsReceivedInfo(id: "${id}") {
    complaintsReceived
    complaintsResolved
    complaintsReviewed
    complaintsPending
    averageRating
  }
}
`;

export const ComplaintReceivedInfoType = (data: any): ComplaintInfo => {
  if (data.errors) {
    console.error("complaintReceivedInfoType err ", data.errors[0].message);
  }
  return data.data.ComplaintsReceivedInfo as ComplaintInfo;
};

export const CompleteEnterpriseQuery = (id: string): string => `
{
  Enterprise(id: "${id}") {
    name
    bannerIMG
    logoIMG
    email
    website
    phone
    industry
    address { country county city }
    foundationDate
    employees {
    id
    profileIMG
    firstName
    lastName
    age
    email
    phone
    hiringDate
    approvedHiring
    approvedHiringAt
    position
    }
  }
}
`;

export const EnterpriseQuery = (id: string): string => `
  {
    Enterprise(id: "${id}") {
      name
      bannerIMG
      logoIMG
      email
      website
      phone
      industry
      address { country county city }
      foundationDate
    }
  }
`;

export const EnterpriseType = (data: any): Enterprise => {
  if (data.errors) {
    console.error("enterprisetype err ", data.errors[0].message);
  }
  return data.data.Enterprise as Enterprise;
};

export const UsersForHiringQuery = (
  id: string,
  limit: number,
  offset: number,
  query: string
): string => `
  {
    UsersForHiring(id: "${id}", limit:${limit},
     offset:${offset}, query: "${query}") {
      users{
      profileIMG
      email
      firstName
      lastName
      gender
      pronoun
      age
      phone
      address { country county city } 
      }
      count
      currentLimit
      currentOffset
    }
  }
`;

export const UsersForHiringType = (data: any): UsersForHiring => {
  if (data.errors) {
    console.error("UsersForHiringType err ", data.errors[0].message);
  }
  return data.data.UsersForHiring as UsersForHiring;
};

export const HiringProccessesQuery = (
  id: string,
  term: string,
  offset: string,
  limit: string
): string => `
{
  HiringProcceses(id: "${id}", query: "${term}", offset: ${offset}, limit: ${limit}) {
  hiringProccesses{
    eventID
    user {
      profileIMG
      email
      firstName
      lastName
      gender
      pronoun
      age
      phone
      address { country county city } }
    position
    status
    occurredOn
    reason
    lastUpdate
    emitedBy {
      profileIMG
      email
      firstName
      lastName
        }
      }
    count
    currentLimit
    currentOffset
  }
}
`;

export const HiringProccessesTypeList = (data: any): HiringProccessList => {
  if (data.errors) {
    console.error("HiringProccessesType err ", data.errors[0].message);
  }
  return data.data.HiringProcceses;
};

export const OnlineUsersQuery = (id: string): string => ` 
{
  OnlineUsers(id: "${id}") {
      profileIMG
      email
      firstName
      lastName
      gender
      pronoun
      age
      status
  }
}
`;

export const OnlineUsersType = (data: any): User[] => {
  if (data.errors) {
    console.error("OnlineUsersType err ", data.errors[0].message);
  }
  return data.data.OnlineUsers as User[];
};

export const EnterpriseChatQuery = (
  enterpriseID: string,
  chatID: string
): string => `
{
  EnterpriseChat(enterpriseID: "${enterpriseID}", chatID: "${chatID}") {
    id
    replies {
      id
      chatID
      user {
        email
        firstName
        lastName
      }
      content
      seen
      createdAt
      updatedAt
    }
  }
}
`;

export const EnterpriseChatTypeCast = (data: any): any => {
  if (data.errors) {
    console.error("EnterpriseChatTypeCast err ", data.errors[0].message);
  }
  return data.data.EnterpriseChat;
};
export const UserQuery = (id: string): string => `
  {
    User(id: "${id}") {
      profileIMG
      email
      firstName
      lastName
      gender
      pronoun
      age
      phone
      address { country county city }
    }
  }
`;

export const UserType = (data: any): User => {
  if (data.errors) {
    console.error("UserType err ", data.errors[0].message);
  }
  return data.data.User as User;
};
export const HiringInvitationsQuery = (): string =>
  `
{
  HiringInvitations {
  eventID
  ownerID
  enterpriseID
  proposedPosition
  fullName
  enterpriseEmail
  enterprisePhone
  enterpriseLogoIMG
  occurredOn
  seen
  status
  }
  }
`;

export const HiringInvitationsTypeList = (
  data: any
): HiringInvitationType[] => {
  if (data.errors) {
    console.error("HiringInvitationsTypeList err ", data.errors);
  }
  return data.data.HiringInvitations;
};
export const EmployeeQuery = (
  enterpriseName: string,
  employeeId: string
): string => `
  {
    Employee(enterpriseName: "${enterpriseName}", employeeID: "${employeeId}") {
      id
      profileIMG
      firstName
      lastName
      age
      email
      phone
      hiringDate
      approvedHiring
      approvedHiringAt
      position
      complaintsSolved
      complaintsSolvedIds
      complaintsRated
      complaintsRatedIDs
      complaintsFeedbacked
      complaintsFeedbackedIDs
      feedbackReceived
      feedbackReceivedIDs
      hireInvitationsSent
      employeesHired
      employeesFired
    }
  }
`;

export const EmployeeQueryType = (data: any): EmployeeType => {
  if (data.errors) {
    console.error("EmployeeType err ", data.errors[0].message);
  }
  return data.data.Employee as EmployeeType;
};

export const EmployeesQuery = (id: string): string => `
  {
    Employees(enterpriseName: "${id}") {
      id
      profileIMG
      firstName
      lastName
      age
      email
      phone
      hiringDate
      approvedHiring
      approvedHiringAt
      position
      complaintsSolved
      complaintsSolvedIds
      complaintsRated
      complaintsRatedIDs
      complaintsFeedbacked
      complaintsFeedbackedIDs
      feedbackReceived
      feedbackReceivedIDs
      hireInvitationsSent
      employeesHired
      employeesFired
    }
  }
`;

export const EmployeesTypeList = (data: any): EmployeeType[] => {
  if (data.errors) {
    console.error("EmployeesTypeList err ", data.errors[0].message);
  }
  return data.data.Employees as EmployeeType[];
};

export const OfficesQuery = (): string => `
  {
    Offices {
      employeeID
      employeeFirstName
      employeePosition
      enterpriseLogoIMG
      enterpriseName
      enterpriseWebsite
      enterprisePhone
      enterpriseEmail
      enterpriseIndustry
      enterpriseAddress { country county city }
      ownerFullName
    }
  }
`;

export const OfficeTypeList = (data: any): Office[] => {
  return data.Offices as Office[];
};

export const ComplaintReviews = (id: string): string => `
{
  PendingComplaintReviews(id: "${id}") {
    eventID
    triggeredBy {
      email
      firstName
      lastName
      pronoun
    }
    complaint {
      id
      message { title description body  }
      receiverFullName
      receiverID
      authorFullName
      authorID
      status 
      createdAt
      rating { rate comment }
      replies { senderID senderName isEnterprise enterpriseID}
    }
    ratedBy {
    email
    firstName
    lastName
    pronoun
    }
    status
  }
}
`;
export const ComplaintReviewsTypeList = (data: any): ComplaintReviewType[] => {
  if (data.errors) {
    console.error("ComplaintReviewsTypeList err ", data.errors[0]);
  }
  return data.data.PendingComplaintReviews as ComplaintReviewType[];
};
export const SolvedReviewQuery = (
  complaintID: string,
  assistantUserID: string
): string => `
{
  User(ID: "${assistantUserID}") {
    firstName
    lastName
  }
  Complaint(ID: "${complaintID}") {
    id
    message { title  }
    rating { rate comment }
  }
}
`;

export const SolvedReviewType = (data: any): SolvedReview => {
  return data as SolvedReview;
};

export const FeedbackByComplaintIDQuery = (id: string): string => `
{
  FeedbackByComplaintID(id: "${id}"){
    id
    complaintID
    enterpriseID
    replyReview {
      id
      feedbackID
        replies {
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
        reviewer {
          profileIMG
          email
          firstName
          lastName
          gender
          pronoun
          }
        review {
          replyReviewID 
          comment
        }
        color
        createdAt
      }
      feedbackAnswer {
        id
        feedbackID
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
      }
      isDone
  }
}
`;
export const FeedbackByComplaintIDType = (data: any): FeedbackType => {
  if (data.errors) {
    console.error("FeedbackByComplaintIDType err ", data.errors[0].message);
  }
  return data.data.FeedbackByComplaintID;
};

export const FeedbackByRevieweeIDQuery = (id: string): string => `
  {
  FeedbackByRevieweeID(id: "${id}"){
    id
    complaintID
    enterpriseID
    updatedAt
    replyReview {
      id
      feedbackID
        replies {
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
        reviewer {
          profileIMG
          email
          firstName
          lastName
          gender
          pronoun
          }
        review {
          replyReviewID 
          comment
        }
        color
        createdAt
      }
      feedbackAnswer {
        id
        feedbackID
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
      }
      isDone
  }
}
`;
export const FeedbackByRevieweeIDType = (data: any): FeedbackType[] => {
  if (data.errors) {
    console.error("FeedbackByRevieweeIDType err ", data.errors[0].message);
  }
  return data.data.FeedbackByRevieweeID;
};
export const FeedbackByIDQuery = (id: string): string => `
{
  FeedbackByID(id: "${id}"){
    id
    complaintID
    enterpriseID
    replyReview {
      id
      feedbackID
        replies {
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
        reviewer {
          profileIMG
          email
          firstName
          lastName
          gender
          pronoun
          }
        review {
          replyReviewID 
          comment
        }
        color
        createdAt
      }
      feedbackAnswer {
        id
        feedbackID
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
      }
      isDone
  }
}
`;
export const FeedbackByIDType = (data: any): FeedbackType => {
  if (data.errors) {
    console.error("FeedbackByIDType err ", data.errors[0].message);
  }
  return data.data.FeedbackByID;
};
export async function Query<T>(
  queryFn: (...args: any[]) => string,
  castToFn: (data: any) => T,
  args: any[] = []
): Promise<T> {
  const token = await csrf();
  if (token != "") {
    const strBody = JSON.stringify({
      query: deleteLinebreaks(queryFn(...args)),
    });
    const response = await fetch("http://3.143.110.143:5555", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-csrf-token": token,
      },
      credentials: "include",
      body: strBody,
    });
    const data = await response.json();
    return castToFn(data);
  }
  throw new Error("No CSRF token");
}

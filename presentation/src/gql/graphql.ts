/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Address = {
  __typename?: 'Address';
  city?: Maybe<Scalars['String']['output']>;
  country?: Maybe<Scalars['String']['output']>;
  county?: Maybe<Scalars['String']['output']>;
};

export type ChatReply = {
  __typename?: 'ChatReply';
  chatID?: Maybe<Scalars['String']['output']>;
  content?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  seen?: Maybe<Scalars['Boolean']['output']>;
  updatedAt?: Maybe<Scalars['String']['output']>;
  user?: Maybe<User>;
};

export type City = {
  __typename?: 'City';
  countryCode?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['Int']['output']>;
  latitude?: Maybe<Scalars['Float']['output']>;
  longitude?: Maybe<Scalars['Float']['output']>;
  name?: Maybe<Scalars['String']['output']>;
};

export type Complaint = {
  __typename?: 'Complaint';
  authorFullName?: Maybe<Scalars['String']['output']>;
  authorID?: Maybe<Scalars['String']['output']>;
  authorProfileIMG?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['ID']['output']>;
  message?: Maybe<ComplaintMessage>;
  rating?: Maybe<ComplaintRating>;
  receiverFullName?: Maybe<Scalars['String']['output']>;
  receiverID?: Maybe<Scalars['String']['output']>;
  receiverProfileIMG?: Maybe<Scalars['String']['output']>;
  replies?: Maybe<Array<Maybe<ComplaintReply>>>;
  status?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['String']['output']>;
};

export type ComplaintInfo = {
  __typename?: 'ComplaintInfo';
  averageRating?: Maybe<Scalars['Float']['output']>;
  complaintsPending?: Maybe<Scalars['Int']['output']>;
  complaintsReceived?: Maybe<Scalars['Int']['output']>;
  complaintsResolved?: Maybe<Scalars['Int']['output']>;
  complaintsReviewed?: Maybe<Scalars['Int']['output']>;
};

export type ComplaintList = {
  __typename?: 'ComplaintList';
  complaints?: Maybe<Array<Maybe<Complaint>>>;
  count?: Maybe<Scalars['Int']['output']>;
  currentLimit?: Maybe<Scalars['Int']['output']>;
  currentOffset?: Maybe<Scalars['Int']['output']>;
};

export type ComplaintMessage = {
  __typename?: 'ComplaintMessage';
  body?: Maybe<Scalars['String']['output']>;
  description?: Maybe<Scalars['String']['output']>;
  title?: Maybe<Scalars['String']['output']>;
};

export type ComplaintRating = {
  __typename?: 'ComplaintRating';
  comment?: Maybe<Scalars['String']['output']>;
  rate?: Maybe<Scalars['Int']['output']>;
};

export type ComplaintReceiver = {
  __typename?: 'ComplaintReceiver';
  fullName?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  thumbnail?: Maybe<Scalars['String']['output']>;
};

export type ComplaintReply = {
  __typename?: 'ComplaintReply';
  body?: Maybe<Scalars['String']['output']>;
  complaintID?: Maybe<Scalars['ID']['output']>;
  complaintStatus?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  enterpriseID?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['ID']['output']>;
  isEnterprise?: Maybe<Scalars['Boolean']['output']>;
  read?: Maybe<Scalars['Boolean']['output']>;
  readAt?: Maybe<Scalars['String']['output']>;
  senderID?: Maybe<Scalars['String']['output']>;
  senderIMG?: Maybe<Scalars['String']['output']>;
  senderName?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['String']['output']>;
};

export type Country = {
  __typename?: 'Country';
  id?: Maybe<Scalars['Int']['output']>;
  name?: Maybe<Scalars['String']['output']>;
  phoneCode?: Maybe<Scalars['String']['output']>;
};

export type County = {
  __typename?: 'County';
  id?: Maybe<Scalars['Int']['output']>;
  name?: Maybe<Scalars['String']['output']>;
};

export type Employee = {
  __typename?: 'Employee';
  age?: Maybe<Scalars['Int']['output']>;
  approvedHiring?: Maybe<Scalars['Boolean']['output']>;
  approvedHiringAt?: Maybe<Scalars['String']['output']>;
  complaintsFeedbacked?: Maybe<Scalars['Int']['output']>;
  complaintsFeedbackedIDs?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  complaintsRated?: Maybe<Scalars['Int']['output']>;
  complaintsRatedIDs?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  complaintsSolved?: Maybe<Scalars['Int']['output']>;
  complaintsSolvedIds?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  email?: Maybe<Scalars['String']['output']>;
  employeesFired?: Maybe<Scalars['Int']['output']>;
  employeesHired?: Maybe<Scalars['Int']['output']>;
  feedbackReceived?: Maybe<Scalars['Int']['output']>;
  feedbackReceivedIDs?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  firstName?: Maybe<Scalars['String']['output']>;
  hireInvitationsSent?: Maybe<Scalars['Int']['output']>;
  hiringDate?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  lastName?: Maybe<Scalars['String']['output']>;
  phone?: Maybe<Scalars['String']['output']>;
  position?: Maybe<Scalars['String']['output']>;
  profileIMG?: Maybe<Scalars['String']['output']>;
};

export type Enterprise = {
  __typename?: 'Enterprise';
  address?: Maybe<Address>;
  bannerIMG?: Maybe<Scalars['String']['output']>;
  email?: Maybe<Scalars['String']['output']>;
  employees?: Maybe<Array<Maybe<Employee>>>;
  foundationDate?: Maybe<Scalars['String']['output']>;
  industry?: Maybe<Scalars['String']['output']>;
  logoIMG?: Maybe<Scalars['String']['output']>;
  name?: Maybe<Scalars['String']['output']>;
  ownerID?: Maybe<Scalars['String']['output']>;
  phone?: Maybe<Scalars['String']['output']>;
  website?: Maybe<Scalars['String']['output']>;
};

export type EnterpriseChat = {
  __typename?: 'EnterpriseChat';
  id?: Maybe<Scalars['String']['output']>;
  replies?: Maybe<Array<Maybe<ChatReply>>>;
};

export type Feedback = {
  __typename?: 'Feedback';
  complaintID?: Maybe<Scalars['String']['output']>;
  enterpriseID?: Maybe<Scalars['String']['output']>;
  feedbackAnswer?: Maybe<Array<Maybe<FeedbackAnswer>>>;
  id?: Maybe<Scalars['String']['output']>;
  isDone?: Maybe<Scalars['Boolean']['output']>;
  replyReview?: Maybe<Array<Maybe<ReplyReview>>>;
  reviewedAt?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['String']['output']>;
};

export type FeedbackAnswer = {
  __typename?: 'FeedbackAnswer';
  body?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  enterpriseID?: Maybe<Scalars['String']['output']>;
  feedbackID?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  isEnterprise?: Maybe<Scalars['Boolean']['output']>;
  read?: Maybe<Scalars['Boolean']['output']>;
  readAt?: Maybe<Scalars['String']['output']>;
  senderID?: Maybe<Scalars['String']['output']>;
  senderIMG?: Maybe<Scalars['String']['output']>;
  senderName?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['String']['output']>;
};

export type Geolocation = {
  __typename?: 'Geolocation';
  latitude?: Maybe<Scalars['Float']['output']>;
  longitude?: Maybe<Scalars['Float']['output']>;
};

export type GrantedAuthority = {
  __typename?: 'GrantedAuthority';
  authority?: Maybe<Scalars['String']['output']>;
  enterpriseID?: Maybe<Scalars['String']['output']>;
};

export type HiringInvitation = {
  __typename?: 'HiringInvitation';
  enterpriseEmail?: Maybe<Scalars['String']['output']>;
  enterpriseID?: Maybe<Scalars['String']['output']>;
  enterpriseLogoIMG?: Maybe<Scalars['String']['output']>;
  enterprisePhone?: Maybe<Scalars['String']['output']>;
  eventID?: Maybe<Scalars['String']['output']>;
  fullName?: Maybe<Scalars['String']['output']>;
  occurredOn?: Maybe<Scalars['String']['output']>;
  ownerID?: Maybe<Scalars['String']['output']>;
  proposedPosition?: Maybe<Scalars['String']['output']>;
  reason?: Maybe<Scalars['String']['output']>;
  seen?: Maybe<Scalars['String']['output']>;
  status?: Maybe<Scalars['String']['output']>;
};

export type HiringProccess = {
  __typename?: 'HiringProccess';
  emitedBy?: Maybe<User>;
  eventID?: Maybe<Scalars['String']['output']>;
  lastUpdate?: Maybe<Scalars['String']['output']>;
  occurredOn?: Maybe<Scalars['String']['output']>;
  position?: Maybe<Scalars['String']['output']>;
  reason?: Maybe<Scalars['String']['output']>;
  status?: Maybe<Scalars['String']['output']>;
  user?: Maybe<User>;
};

export type HiringProccessList = {
  __typename?: 'HiringProccessList';
  count?: Maybe<Scalars['Int']['output']>;
  currentLimit?: Maybe<Scalars['Int']['output']>;
  currentOffset?: Maybe<Scalars['Int']['output']>;
  hiringProccesses?: Maybe<Array<Maybe<HiringProccess>>>;
};

export type Industry = {
  __typename?: 'Industry';
  id?: Maybe<Scalars['Int']['output']>;
  name?: Maybe<Scalars['String']['output']>;
};

export type JwtToken = {
  __typename?: 'JwtToken';
  token: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  /** Accept the invitation to join the enterprise */
  AcceptEnterpriseInvitation?: Maybe<Scalars['Boolean']['output']>;
  /** Add a comment to a feedback */
  AddComment?: Maybe<Scalars['Boolean']['output']>;
  /** Add a review to a reply */
  AddReply?: Maybe<Scalars['Boolean']['output']>;
  /** Answer a feedback */
  AnswerFeedback?: Maybe<Scalars['Boolean']['output']>;
  /** Cancel the hiring proccess */
  CancelHiringProccess?: Maybe<Scalars['Boolean']['output']>;
  /** Set a new random password for the user and send it by email */
  ChangePassword?: Maybe<Scalars['Boolean']['output']>;
  /** Send a contact email */
  Contact?: Maybe<Scalars['Boolean']['output']>;
  /** Create a new enterprise */
  CreateEnterprise?: Maybe<Scalars['Boolean']['output']>;
  /** Create a feedback */
  CreateFeedback?: Maybe<Scalars['Boolean']['output']>;
  /** Create new user */
  CreateUser?: Maybe<Scalars['Boolean']['output']>;
  /** Delete a comment from a feedback */
  DeleteComment?: Maybe<Scalars['Boolean']['output']>;
  /** End the feedback */
  EndFeedback?: Maybe<Scalars['Boolean']['output']>;
  /** Fire an employee from the enterprise */
  FireEmployee?: Maybe<Scalars['Boolean']['output']>;
  /** Hire an employee to the enterprise */
  HireEmployee?: Maybe<Scalars['Boolean']['output']>;
  /** Invite a user to join the enterprise */
  InviteToEnterprise?: Maybe<Scalars['Boolean']['output']>;
  /** Leave the enterprise */
  LeaveEnterprise?: Maybe<Scalars['Boolean']['output']>;
  /** Mark a complaint reply as seen */
  MarkAsSeen?: Maybe<Scalars['Boolean']['output']>;
  /** Mark a notification as read */
  MarkNotificationAsRead?: Maybe<Scalars['Boolean']['output']>;
  /** Mark an enterprise chat reply as seen */
  MarkReplyChatAsSeen?: Maybe<Scalars['Boolean']['output']>;
  /** Promote an employee */
  PromoteEmployee?: Maybe<Scalars['Boolean']['output']>;
  /** Rate a complaint */
  RateComplaint?: Maybe<Scalars['Boolean']['output']>;
  /** Set a new random password for the user and send it by email */
  RecoverPassword?: Maybe<Scalars['Boolean']['output']>;
  /** Reject the invitation to join the enterprise */
  RejectEnterpriseInvitation?: Maybe<Scalars['Boolean']['output']>;
  /** Remove a review from a reply */
  RemoveReply?: Maybe<Scalars['Boolean']['output']>;
  /** Reply a chat */
  ReplyChat?: Maybe<Scalars['Boolean']['output']>;
  /** Reply a complaint */
  ReplyComplaint?: Maybe<Scalars['Boolean']['output']>;
  /** Send a new complaint */
  SendComplaint?: Maybe<Scalars['Boolean']['output']>;
  /** Send a complaint for reviewing */
  SendForReviewing?: Maybe<Scalars['Boolean']['output']>;
  /** Update the enterprise */
  UpdateEnterprise?: Maybe<Scalars['Boolean']['output']>;
  /** Update the user personal information */
  UpdateUser?: Maybe<Scalars['Boolean']['output']>;
  /** Verify the email from the link sent by email */
  VerifyEmail?: Maybe<Scalars['Boolean']['output']>;
};


export type MutationAcceptEnterpriseInvitationArgs = {
  id: Scalars['String']['input'];
};


export type MutationAddCommentArgs = {
  color?: InputMaybe<Scalars['String']['input']>;
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintID?: InputMaybe<Scalars['String']['input']>;
  enterpriseID: Scalars['String']['input'];
  feedbackID?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  reviewerID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationAddReplyArgs = {
  color?: InputMaybe<Scalars['String']['input']>;
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintID?: InputMaybe<Scalars['String']['input']>;
  enterpriseID: Scalars['String']['input'];
  feedbackID?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  reviewerID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationAnswerFeedbackArgs = {
  answerBody?: InputMaybe<Scalars['String']['input']>;
  feedbackID: Scalars['String']['input'];
};


export type MutationCancelHiringProccessArgs = {
  enterpriseName: Scalars['String']['input'];
  eventID: Scalars['String']['input'];
  reason?: InputMaybe<Scalars['String']['input']>;
};


export type MutationChangePasswordArgs = {
  newPassword: Scalars['String']['input'];
  oldPassword: Scalars['String']['input'];
};


export type MutationContactArgs = {
  email: Scalars['String']['input'];
  text: Scalars['String']['input'];
};


export type MutationCreateEnterpriseArgs = {
  cityID: Scalars['Int']['input'];
  countryID: Scalars['Int']['input'];
  countryStateID: Scalars['Int']['input'];
  email: Scalars['String']['input'];
  foundationDate: Scalars['String']['input'];
  industryID: Scalars['Int']['input'];
  name: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  phoneCode: Scalars['String']['input'];
  website?: InputMaybe<Scalars['String']['input']>;
};


export type MutationCreateFeedbackArgs = {
  color?: InputMaybe<Scalars['String']['input']>;
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintID?: InputMaybe<Scalars['String']['input']>;
  enterpriseID: Scalars['String']['input'];
  feedbackID?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  reviewerID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationCreateUserArgs = {
  birthDate: Scalars['String']['input'];
  cityId: Scalars['Int']['input'];
  countryId: Scalars['Int']['input'];
  countryStateId: Scalars['Int']['input'];
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  gender: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  password: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  pronoun: Scalars['String']['input'];
};


export type MutationDeleteCommentArgs = {
  color?: InputMaybe<Scalars['String']['input']>;
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintID?: InputMaybe<Scalars['String']['input']>;
  enterpriseID: Scalars['String']['input'];
  feedbackID?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  reviewerID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationEndFeedbackArgs = {
  color?: InputMaybe<Scalars['String']['input']>;
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintID?: InputMaybe<Scalars['String']['input']>;
  enterpriseID: Scalars['String']['input'];
  feedbackID?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  reviewerID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationFireEmployeeArgs = {
  employeeID?: InputMaybe<Scalars['String']['input']>;
  enterpriseName?: InputMaybe<Scalars['String']['input']>;
};


export type MutationHireEmployeeArgs = {
  enterpriseName: Scalars['String']['input'];
  eventID: Scalars['String']['input'];
  reason?: InputMaybe<Scalars['String']['input']>;
};


export type MutationInviteToEnterpriseArgs = {
  enterpriseName: Scalars['String']['input'];
  proposeTo: Scalars['String']['input'];
  proposedPosition: Scalars['String']['input'];
};


export type MutationLeaveEnterpriseArgs = {
  employeeID?: InputMaybe<Scalars['String']['input']>;
  enterpriseName?: InputMaybe<Scalars['String']['input']>;
};


export type MutationMarkAsSeenArgs = {
  complaintID: Scalars['String']['input'];
  ids?: InputMaybe<Scalars['String']['input']>;
};


export type MutationMarkNotificationAsReadArgs = {
  id: Scalars['String']['input'];
};


export type MutationMarkReplyChatAsSeenArgs = {
  chatID?: InputMaybe<Scalars['String']['input']>;
  enterpriseName?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationPromoteEmployeeArgs = {
  employeeID: Scalars['String']['input'];
  enterpriseName: Scalars['String']['input'];
  position: Scalars['String']['input'];
};


export type MutationRateComplaintArgs = {
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintId: Scalars['String']['input'];
  eventId: Scalars['String']['input'];
  rate: Scalars['Int']['input'];
};


export type MutationRecoverPasswordArgs = {
  id: Scalars['String']['input'];
};


export type MutationRejectEnterpriseInvitationArgs = {
  id: Scalars['String']['input'];
  reason?: InputMaybe<Scalars['String']['input']>;
};


export type MutationRemoveReplyArgs = {
  color?: InputMaybe<Scalars['String']['input']>;
  comment?: InputMaybe<Scalars['String']['input']>;
  complaintID?: InputMaybe<Scalars['String']['input']>;
  enterpriseID: Scalars['String']['input'];
  feedbackID?: InputMaybe<Scalars['String']['input']>;
  repliesID?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  reviewerID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationReplyChatArgs = {
  content?: InputMaybe<Scalars['String']['input']>;
  enterpriseName?: InputMaybe<Scalars['String']['input']>;
  id?: InputMaybe<Scalars['String']['input']>;
  senderID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationReplyComplaintArgs = {
  complaintID: Scalars['String']['input'];
  replyAuthorID: Scalars['String']['input'];
  replyBody: Scalars['String']['input'];
  replyEnterpriseID?: InputMaybe<Scalars['String']['input']>;
};


export type MutationSendComplaintArgs = {
  authorID: Scalars['String']['input'];
  content: Scalars['String']['input'];
  description: Scalars['String']['input'];
  receiverFullName: Scalars['String']['input'];
  receiverID: Scalars['String']['input'];
  receiverProfileIMG: Scalars['String']['input'];
  title: Scalars['String']['input'];
};


export type MutationSendForReviewingArgs = {
  id: Scalars['String']['input'];
};


export type MutationUpdateEnterpriseArgs = {
  enterpriseID: Scalars['String']['input'];
  numberValue?: InputMaybe<Scalars['Int']['input']>;
  updateType: Scalars['String']['input'];
  value?: InputMaybe<Scalars['String']['input']>;
};


export type MutationUpdateUserArgs = {
  numberValue?: InputMaybe<Scalars['Int']['input']>;
  updateType: Scalars['String']['input'];
  value?: InputMaybe<Scalars['String']['input']>;
};


export type MutationVerifyEmailArgs = {
  id: Scalars['String']['input'];
};

export type Notifications = {
  __typename?: 'Notifications';
  content?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  link?: Maybe<Scalars['String']['output']>;
  occurredOn?: Maybe<Scalars['String']['output']>;
  ownerID?: Maybe<Scalars['String']['output']>;
  seen?: Maybe<Scalars['Boolean']['output']>;
  thumbnail?: Maybe<Scalars['String']['output']>;
  title?: Maybe<Scalars['String']['output']>;
};

export type PendingReview = {
  __typename?: 'PendingReview';
  complaint?: Maybe<Complaint>;
  eventID?: Maybe<Scalars['String']['output']>;
  occurredOn?: Maybe<Scalars['String']['output']>;
  ratedBy?: Maybe<User>;
  status?: Maybe<Scalars['String']['output']>;
  triggeredBy?: Maybe<User>;
};

export type Query = {
  __typename?: 'Query';
  /** find all cities by county ID */
  Cities?: Maybe<Array<Maybe<City>>>;
  /** Get a complaint by it's ID */
  Complaint?: Maybe<Complaint>;
  /** Get the list of inbox complaints for the current user */
  ComplaintHistory?: Maybe<ComplaintList>;
  /** Get the list of inbox complaints for the current user */
  ComplaintInbox?: Maybe<ComplaintList>;
  /** Search the inbox complaints for the current user */
  ComplaintInboxSearch?: Maybe<ComplaintList>;
  /** Get the complaints received info for the current user */
  ComplaintsReceivedInfo?: Maybe<ComplaintInfo>;
  /** Get the list of sent complaints for the current user */
  ComplaintsSent?: Maybe<ComplaintList>;
  /** Get the list of sent complaints for the current user */
  ComplaintsSentSearch?: Maybe<ComplaintList>;
  /** Get the list of countries */
  Countries?: Maybe<Array<Maybe<Country>>>;
  /** find all counties by country ID */
  CountryStates?: Maybe<Array<Maybe<County>>>;
  /** Get the employee by it's ID, enterpriseID required for auth for authorization */
  Employee?: Maybe<Employee>;
  /** Get the list of employees for the enterprise, enterpriseID required for authorization */
  Employees?: Maybe<Array<Maybe<Employee>>>;
  /** Return the enterprise info */
  Enterprise?: Maybe<Enterprise>;
  /** Get the chat for the enterprise */
  EnterpriseChat?: Maybe<EnterpriseChat>;
  /** Get the feedback by complaint ID */
  FeedbackByComplaintID?: Maybe<Feedback>;
  /** Get the reviews by reviewer ID */
  FeedbackByID?: Maybe<Feedback>;
  /** Get the reviews by reviee ID */
  FeedbackByRevieweeID?: Maybe<Array<Maybe<Feedback>>>;
  /** Find the author by ID */
  FindAuthorByID?: Maybe<ComplaintReceiver>;
  /** Find the receivers for a complaint */
  FindComplaintReceivers?: Maybe<Array<Maybe<ComplaintReceiver>>>;
  /** Get the list of hiring invitations */
  HiringInvitations?: Maybe<Array<Maybe<HiringInvitation>>>;
  /** Get the list of hiring invitations accepted */
  HiringProcceses?: Maybe<HiringProccessList>;
  /** Get the list of industries */
  Industries?: Maybe<Array<Maybe<Industry>>>;
  /** Check if the enterprise name is available */
  IsEnterpriseNameAvailable?: Maybe<Scalars['Boolean']['output']>;
  /** Check if the receiver is valid for a complaint */
  IsValidComplaintReceiver?: Maybe<Scalars['Boolean']['output']>;
  /** Authenticate the user with the token and confirmation code it got the token from the request header */
  Login?: Maybe<JwtToken>;
  /** Get the list of online users */
  OnlineUsers?: Maybe<Array<Maybe<User>>>;
  /** Get the list of complaints waiting for review */
  PendingComplaintReviews?: Maybe<Array<Maybe<PendingReview>>>;
  /** Get the token for the authenticated user or error */
  SignIn?: Maybe<JwtToken>;
  /** Get the list of solved complaints for the current user */
  SolvedComplaints?: Maybe<Array<Maybe<Complaint>>>;
  /** Get a user without private information by it's ID */
  User?: Maybe<User>;
  /** Get the user descriptor for the current session */
  UserDescriptor?: Maybe<UserDescriptor>;
  /** Get the list of users for hiring */
  UsersForHiring?: Maybe<UserList>;
};


export type QueryCitiesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryComplaintArgs = {
  id: Scalars['String']['input'];
};


export type QueryComplaintHistoryArgs = {
  afterDate?: InputMaybe<Scalars['String']['input']>;
  beforeDate?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryComplaintInboxArgs = {
  afterDate?: InputMaybe<Scalars['String']['input']>;
  beforeDate?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryComplaintInboxSearchArgs = {
  afterDate?: InputMaybe<Scalars['String']['input']>;
  beforeDate?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryComplaintsReceivedInfoArgs = {
  id: Scalars['String']['input'];
};


export type QueryComplaintsSentArgs = {
  afterDate?: InputMaybe<Scalars['String']['input']>;
  beforeDate?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryComplaintsSentSearchArgs = {
  afterDate?: InputMaybe<Scalars['String']['input']>;
  beforeDate?: InputMaybe<Scalars['String']['input']>;
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  status?: InputMaybe<Scalars['String']['input']>;
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryCountryStatesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryEmployeeArgs = {
  employeeID?: InputMaybe<Scalars['String']['input']>;
  enterpriseName?: InputMaybe<Scalars['String']['input']>;
};


export type QueryEmployeesArgs = {
  employeeID?: InputMaybe<Scalars['String']['input']>;
  enterpriseName?: InputMaybe<Scalars['String']['input']>;
};


export type QueryEnterpriseArgs = {
  id: Scalars['String']['input'];
};


export type QueryEnterpriseChatArgs = {
  chatID: Scalars['String']['input'];
  enterpriseID: Scalars['String']['input'];
};


export type QueryFeedbackByComplaintIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryFeedbackByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryFeedbackByRevieweeIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryFindAuthorByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryFindComplaintReceiversArgs = {
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryHiringProccesesArgs = {
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  query?: InputMaybe<Scalars['String']['input']>;
};


export type QueryIsEnterpriseNameAvailableArgs = {
  id: Scalars['String']['input'];
};


export type QueryIsValidComplaintReceiverArgs = {
  id: Scalars['String']['input'];
};


export type QueryLoginArgs = {
  confirmationCode: Scalars['Int']['input'];
};


export type QueryOnlineUsersArgs = {
  id: Scalars['String']['input'];
};


export type QueryPendingComplaintReviewsArgs = {
  id: Scalars['String']['input'];
};


export type QuerySignInArgs = {
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
  rememberMe?: InputMaybe<Scalars['Boolean']['input']>;
};


export type QuerySolvedComplaintsArgs = {
  id: Scalars['String']['input'];
};


export type QueryUserArgs = {
  id: Scalars['String']['input'];
};


export type QueryUsersForHiringArgs = {
  id: Scalars['String']['input'];
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
  query?: InputMaybe<Scalars['String']['input']>;
};

export type ReplyReview = {
  __typename?: 'ReplyReview';
  color?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  feedbackID?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  replies?: Maybe<Array<Maybe<ComplaintReply>>>;
  review?: Maybe<Review>;
  reviewer?: Maybe<User>;
};

export type Review = {
  __typename?: 'Review';
  comment?: Maybe<Scalars['String']['output']>;
  replyReviewID?: Maybe<Scalars['String']['output']>;
};

export type Subscription = {
  __typename?: 'Subscription';
  notifications?: Maybe<Array<Maybe<Notifications>>>;
};


export type SubscriptionNotificationsArgs = {
  id: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  address?: Maybe<Address>;
  age?: Maybe<Scalars['Int']['output']>;
  email?: Maybe<Scalars['String']['output']>;
  firstName?: Maybe<Scalars['String']['output']>;
  gender?: Maybe<Scalars['String']['output']>;
  lastName?: Maybe<Scalars['String']['output']>;
  phone?: Maybe<Scalars['String']['output']>;
  profileIMG?: Maybe<Scalars['String']['output']>;
  pronoun?: Maybe<Scalars['String']['output']>;
  status?: Maybe<Scalars['String']['output']>;
};

export type UserDescriptor = {
  __typename?: 'UserDescriptor';
  device?: Maybe<Scalars['String']['output']>;
  email?: Maybe<Scalars['String']['output']>;
  fullName?: Maybe<Scalars['String']['output']>;
  gender?: Maybe<Scalars['String']['output']>;
  geolocation?: Maybe<Geolocation>;
  grantedAuthorities?: Maybe<Array<Maybe<GrantedAuthority>>>;
  ip?: Maybe<Scalars['String']['output']>;
  loginDate?: Maybe<Scalars['String']['output']>;
  profileIMG?: Maybe<Scalars['String']['output']>;
  pronoun?: Maybe<Scalars['String']['output']>;
};

export type UserList = {
  __typename?: 'UserList';
  count?: Maybe<Scalars['Int']['output']>;
  currentLimit?: Maybe<Scalars['Int']['output']>;
  currentOffset?: Maybe<Scalars['Int']['output']>;
  users?: Maybe<Array<Maybe<User>>>;
};

export type CreateUserMutationMutationVariables = Exact<{
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  gender: Scalars['String']['input'];
  pronoun: Scalars['String']['input'];
  birthDate: Scalars['String']['input'];
  phone: Scalars['String']['input'];
  countryId: Scalars['Int']['input'];
  countryStateId: Scalars['Int']['input'];
  cityId: Scalars['Int']['input'];
}>;


export type CreateUserMutationMutation = { __typename?: 'Mutation', CreateUser?: boolean | null };

export type CitiesQueryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CitiesQueryQuery = { __typename?: 'Query', Cities?: Array<{ __typename?: 'City', id?: number | null, name?: string | null, countryCode?: string | null, latitude?: number | null, longitude?: number | null } | null> | null };

export type CountriesQueryVariables = Exact<{ [key: string]: never; }>;


export type CountriesQuery = { __typename?: 'Query', Countries?: Array<{ __typename?: 'Country', id?: number | null, name?: string | null, phoneCode?: string | null } | null> | null };

export type CountryStatesQueryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CountryStatesQueryQuery = { __typename?: 'Query', CountryStates?: Array<{ __typename?: 'County', id?: number | null, name?: string | null } | null> | null };


export const CreateUserMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUserMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"email"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"password"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"firstName"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"lastName"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"gender"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pronoun"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"birthDate"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"phone"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"countryId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"countryStateId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"cityId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"CreateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"email"},"value":{"kind":"Variable","name":{"kind":"Name","value":"email"}}},{"kind":"Argument","name":{"kind":"Name","value":"password"},"value":{"kind":"Variable","name":{"kind":"Name","value":"password"}}},{"kind":"Argument","name":{"kind":"Name","value":"firstName"},"value":{"kind":"Variable","name":{"kind":"Name","value":"firstName"}}},{"kind":"Argument","name":{"kind":"Name","value":"lastName"},"value":{"kind":"Variable","name":{"kind":"Name","value":"lastName"}}},{"kind":"Argument","name":{"kind":"Name","value":"gender"},"value":{"kind":"Variable","name":{"kind":"Name","value":"gender"}}},{"kind":"Argument","name":{"kind":"Name","value":"pronoun"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pronoun"}}},{"kind":"Argument","name":{"kind":"Name","value":"birthDate"},"value":{"kind":"Variable","name":{"kind":"Name","value":"birthDate"}}},{"kind":"Argument","name":{"kind":"Name","value":"phone"},"value":{"kind":"Variable","name":{"kind":"Name","value":"phone"}}},{"kind":"Argument","name":{"kind":"Name","value":"countryId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"countryId"}}},{"kind":"Argument","name":{"kind":"Name","value":"countryStateId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"countryStateId"}}},{"kind":"Argument","name":{"kind":"Name","value":"cityId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"cityId"}}}]}]}}]} as unknown as DocumentNode<CreateUserMutationMutation, CreateUserMutationMutationVariables>;
export const CitiesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CitiesQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Cities"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"countryCode"}},{"kind":"Field","name":{"kind":"Name","value":"latitude"}},{"kind":"Field","name":{"kind":"Name","value":"longitude"}}]}}]}}]} as unknown as DocumentNode<CitiesQueryQuery, CitiesQueryQueryVariables>;
export const CountriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Countries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Countries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"phoneCode"}}]}}]}}]} as unknown as DocumentNode<CountriesQuery, CountriesQueryVariables>;
export const CountryStatesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CountryStatesQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"CountryStates"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CountryStatesQueryQuery, CountryStatesQueryQueryVariables>;
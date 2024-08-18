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
  Upload: { input: any; output: any; }
};

export type AcceptHiringInvitation = {
  hiringProcessId: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type AddFeedbackComment = {
  color: Scalars['String']['input'];
  comment: Scalars['String']['input'];
  feedbackId: Scalars['String']['input'];
};

export type AddFeedbackReply = {
  color: Scalars['String']['input'];
  feedbackId: Scalars['String']['input'];
  repliesIds: Array<Scalars['String']['input']>;
  reviewerId: Scalars['String']['input'];
};

export type Address = {
  __typename?: 'Address';
  city: Scalars['String']['output'];
  country: Scalars['String']['output'];
  countryState: Scalars['String']['output'];
};

export type CancelHiringProcess = {
  cancelationReason: Scalars['String']['input'];
  canceledBy: Scalars['String']['input'];
  enterpriseId: Scalars['String']['input'];
  hiringProcessId: Scalars['String']['input'];
};

export type ChangeEnterpriseAddress = {
  enterpriseId: Scalars['String']['input'];
  newCityId: Scalars['Int']['input'];
  newCountryId: Scalars['Int']['input'];
  newCountyId: Scalars['Int']['input'];
};

export type ChangeEnterpriseEmail = {
  enterpriseId: Scalars['String']['input'];
  newEmail: Scalars['String']['input'];
};

export type ChangeEnterprisePhone = {
  enterpriseId: Scalars['String']['input'];
  newPhone: Scalars['String']['input'];
};

export type ChangeEnterpriseWebsite = {
  enterpriseId: Scalars['String']['input'];
  newWebsite: Scalars['String']['input'];
};

export type ChangePassword = {
  newPassword: Scalars['String']['input'];
  oldPassword: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type ChangeUserFirstName = {
  newFirstName: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type ChangeUserGenre = {
  newGenre: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type ChangeUserLastName = {
  newLastName: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type ChangeUserPhone = {
  newPhone: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type ChangeUserPronoun = {
  newPronoun: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type Chat = {
  __typename?: 'Chat';
  enterpriseId: Scalars['String']['output'];
  id: Scalars['String']['output'];
  recipientOne: Recipient;
  recipientTwo: Recipient;
  replies: Array<Maybe<ChatReply>>;
};

export type ChatReply = {
  __typename?: 'ChatReply';
  chatId: Scalars['String']['output'];
  content: Scalars['String']['output'];
  createdAt: Scalars['String']['output'];
  id: Scalars['String']['output'];
  seen: Scalars['Boolean']['output'];
  sender: Recipient;
  updatedAt: Scalars['String']['output'];
};

export type City = {
  __typename?: 'City';
  countryCode: Scalars['String']['output'];
  id: Scalars['Int']['output'];
  latitude: Scalars['Float']['output'];
  longitude: Scalars['Float']['output'];
  name: Scalars['String']['output'];
};

export type Complaint = {
  __typename?: 'Complaint';
  author?: Maybe<Recipient>;
  createdAt?: Maybe<Scalars['String']['output']>;
  description?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  rating?: Maybe<Rating>;
  receiver?: Maybe<Recipient>;
  replies?: Maybe<Array<Maybe<ComplaintReply>>>;
  status?: Maybe<ComplaintStatus>;
  title?: Maybe<Scalars['String']['output']>;
  updatedAt?: Maybe<Scalars['String']['output']>;
};

export type ComplaintData = {
  __typename?: 'ComplaintData';
  complaintId: Scalars['String']['output'];
  dataType: ComplaintDataType;
  id: Scalars['String']['output'];
  occurredOn: Scalars['String']['output'];
  ownerId: Scalars['String']['output'];
};

export enum ComplaintDataType {
  Received = 'RECEIVED',
  Resolved = 'RESOLVED',
  Reviewed = 'REVIEWED',
  Sent = 'SENT'
}

export type ComplaintReply = {
  __typename?: 'ComplaintReply';
  body?: Maybe<Scalars['String']['output']>;
  complaintId?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  enterpriseId?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  isEnterprise?: Maybe<Scalars['Boolean']['output']>;
  read?: Maybe<Scalars['Boolean']['output']>;
  readAt?: Maybe<Scalars['String']['output']>;
  sender?: Maybe<Recipient>;
  updatedAt?: Maybe<Scalars['String']['output']>;
};

export enum ComplaintStatus {
  Closed = 'CLOSED',
  InDiscussion = 'IN_DISCUSSION',
  InHistory = 'IN_HISTORY',
  InReview = 'IN_REVIEW',
  Open = 'OPEN',
  Started = 'STARTED',
  Writing = 'WRITING'
}

export type Country = {
  __typename?: 'Country';
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
  phoneCode: Scalars['String']['output'];
};

export type CountryState = {
  __typename?: 'CountryState';
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type CreateEnterprise = {
  cityId: Scalars['Int']['input'];
  countryId: Scalars['Int']['input'];
  countryStateId: Scalars['Int']['input'];
  email: Scalars['String']['input'];
  foundationDate: Scalars['String']['input'];
  industryId: Scalars['Int']['input'];
  name: Scalars['String']['input'];
  phoneNumber: Scalars['String']['input'];
  website: Scalars['String']['input'];
};

export type CreateEnterpriseChat = {
  enterpriseId: Scalars['String']['input'];
  receiverId: Scalars['String']['input'];
  senderId: Scalars['String']['input'];
};

export type CreateFeedback = {
  complaintId: Scalars['String']['input'];
  enterpriseId: Scalars['String']['input'];
};

export type CreateHiringInvitation = {
  enterpriseId: Scalars['String']['input'];
  proposeTo: Scalars['String']['input'];
  proposedPosition: Scalars['String']['input'];
};

export type CreateNewComplaint = {
  authorId: Scalars['String']['input'];
  receiverId: Scalars['String']['input'];
};

export type CreateUser = {
  birthDate: Scalars['String']['input'];
  cityId: Scalars['Int']['input'];
  countryId: Scalars['Int']['input'];
  countryStateId: Scalars['Int']['input'];
  firstName: Scalars['String']['input'];
  genre: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  password: Scalars['String']['input'];
  phoneNumber: Scalars['String']['input'];
  pronoun: Scalars['String']['input'];
  userName: Scalars['String']['input'];
};

export type DescribeComplaint = {
  complaintId: Scalars['String']['input'];
  description: Scalars['String']['input'];
  title: Scalars['String']['input'];
};

export type Employee = {
  __typename?: 'Employee';
  approvedHiring: Scalars['Boolean']['output'];
  approvedHiringAt: Scalars['String']['output'];
  enterpriseId: Scalars['String']['output'];
  enterprisePosition: Scalars['String']['output'];
  hiringDate: Scalars['String']['output'];
  id: Scalars['String']['output'];
  user: User;
  userId: Scalars['String']['output'];
};

export type EndFeedback = {
  feedbackId: Scalars['String']['input'];
  reviewerId: Scalars['String']['input'];
};

export type Enterprise = {
  __typename?: 'Enterprise';
  address: Address;
  bannerImg: Scalars['String']['output'];
  email: Scalars['String']['output'];
  employees: Array<Maybe<Employee>>;
  foundationDate: Scalars['String']['output'];
  id: Scalars['String']['output'];
  industry: Industry;
  logoImg: Scalars['String']['output'];
  name: Scalars['String']['output'];
  ownerId: Scalars['String']['output'];
  phoneNumber: Scalars['String']['output'];
  website: Scalars['String']['output'];
};

export type EnterpriseActivity = {
  __typename?: 'EnterpriseActivity';
  activityId: Scalars['String']['output'];
  activityType: EnterpriseActivityType;
  enterpriseId: Scalars['String']['output'];
  enterpriseName: Scalars['String']['output'];
  id: Scalars['String']['output'];
  occurredOn: Scalars['String']['output'];
  user: Recipient;
};

export enum EnterpriseActivityType {
  ComplaintResolved = 'COMPLAINT_RESOLVED',
  ComplaintReviewed = 'COMPLAINT_REVIEWED',
  ComplaintSent = 'COMPLAINT_SENT',
  EmployeesFired = 'EMPLOYEES_FIRED',
  EmployeesHired = 'EMPLOYEES_HIRED',
  FeedbacksReceived = 'FEEDBACKS_RECEIVED',
  FeedbacksStarted = 'FEEDBACKS_STARTED',
  JobProposalsSent = 'JOB_PROPOSALS_SENT'
}

export type EnterpriseByAuthenticatedUser = {
  __typename?: 'EnterpriseByAuthenticatedUser';
  authority: GrantedAuthority;
  enterprise?: Maybe<Enterprise>;
};

export type EnterprisesByAuthenticatedUserResult = {
  __typename?: 'EnterprisesByAuthenticatedUserResult';
  enterprises: Array<EnterpriseByAuthenticatedUser>;
  offices: Array<EnterpriseByAuthenticatedUser>;
};

export type Feedback = {
  __typename?: 'Feedback';
  complaintId: Scalars['String']['output'];
  enterpriseId: Scalars['String']['output'];
  id: Scalars['String']['output'];
  isDone: Scalars['Boolean']['output'];
  replyReview: Array<Maybe<ReplyReview>>;
  reviewedAt: Scalars['String']['output'];
  updatedAt: Scalars['String']['output'];
};

export type FindEnterpriseChat = {
  enterpriseId: Scalars['String']['input'];
  recipientOneId: Scalars['String']['input'];
  recipientTwoId: Scalars['String']['input'];
};

export type FireEmployee = {
  employeeId: Scalars['String']['input'];
  enterpriseName: Scalars['String']['input'];
  fireReason: Scalars['String']['input'];
  triggeredBy: Scalars['String']['input'];
};

export type GrantedAuthority = {
  __typename?: 'GrantedAuthority';
  authority: Roles;
  enterpriseId: Scalars['String']['output'];
  principal: Scalars['String']['output'];
};

export type HireEmployee = {
  enterpriseId: Scalars['String']['input'];
  hiredById: Scalars['String']['input'];
  hiringProcessId: Scalars['String']['input'];
};

export enum HiringProccessStatus {
  Accepted = 'ACCEPTED',
  Canceled = 'CANCELED',
  Fired = 'FIRED',
  Hired = 'HIRED',
  Leaved = 'LEAVED',
  Pending = 'PENDING',
  Rated = 'RATED',
  Rejected = 'REJECTED',
  UserAccepted = 'USER_ACCEPTED',
  Waiting = 'WAITING'
}

export type HiringProcess = {
  __typename?: 'HiringProcess';
  emitedBy: Recipient;
  enterprise?: Maybe<Recipient>;
  id: Scalars['String']['output'];
  industry?: Maybe<Industry>;
  lastUpdate: Scalars['String']['output'];
  occurredOn: Scalars['String']['output'];
  reason?: Maybe<Scalars['String']['output']>;
  role: Scalars['String']['output'];
  status?: Maybe<HiringProccessStatus>;
  updatedBy?: Maybe<Recipient>;
  user: User;
};

export type Industry = {
  __typename?: 'Industry';
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type InviteToProject = {
  enterpriseId: Scalars['String']['input'];
  proposeTo: Scalars['String']['input'];
  proposedBy: Scalars['String']['input'];
  role: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  acceptHiringInvitation: HiringProcess;
  addFeedbackComment: Feedback;
  addFeedbackReply: Feedback;
  cancelHiringProcess: HiringProcess;
  changeEnterpriseAddress: Enterprise;
  changeEnterpriseBannerImg: Enterprise;
  changeEnterpriseEmail: Enterprise;
  changeEnterpriseLogoImg: Enterprise;
  changeEnterprisePhone: Enterprise;
  changeEnterpriseWebsite: Enterprise;
  changeFirstName: User;
  changeLastName: User;
  changePassword: User;
  changeUserGenre: User;
  changeUserPhone: User;
  changeUserPronoun: User;
  createEnterprise: Enterprise;
  createEnterpriseChat: Chat;
  createFeedback: Feedback;
  createNewComplaint?: Maybe<Complaint>;
  createUser: User;
  describeComplaint?: Maybe<Complaint>;
  endFeedback: Feedback;
  fireEmployee: Enterprise;
  hireEmployee: HiringProcess;
  inviteToProject: HiringProcess;
  markNotificationAsRead: NotificationLink;
  promoteEmployee: Employee;
  rateComplaint: Rating;
  rejectHiringInvitation: HiringProcess;
  removeFeedbackCommand: Feedback;
  removeFeedbackReply: Feedback;
  sendComplaint?: Maybe<Complaint>;
  updateProfileImg: User;
  updateUserAddress: User;
};


export type MutationAcceptHiringInvitationArgs = {
  input: AcceptHiringInvitation;
};


export type MutationAddFeedbackCommentArgs = {
  input: AddFeedbackComment;
};


export type MutationAddFeedbackReplyArgs = {
  input: AddFeedbackReply;
};


export type MutationCancelHiringProcessArgs = {
  input: CancelHiringProcess;
};


export type MutationChangeEnterpriseAddressArgs = {
  input: ChangeEnterpriseAddress;
};


export type MutationChangeEnterpriseBannerImgArgs = {
  enterpriseId: Scalars['String']['input'];
  file: Scalars['Upload']['input'];
};


export type MutationChangeEnterpriseEmailArgs = {
  input: ChangeEnterpriseEmail;
};


export type MutationChangeEnterpriseLogoImgArgs = {
  enterpriseId: Scalars['String']['input'];
  file: Scalars['Upload']['input'];
};


export type MutationChangeEnterprisePhoneArgs = {
  input: ChangeEnterprisePhone;
};


export type MutationChangeEnterpriseWebsiteArgs = {
  input: ChangeEnterpriseWebsite;
};


export type MutationChangeFirstNameArgs = {
  input: ChangeUserFirstName;
};


export type MutationChangeLastNameArgs = {
  input: ChangeUserLastName;
};


export type MutationChangePasswordArgs = {
  input: ChangePassword;
};


export type MutationChangeUserGenreArgs = {
  input: ChangeUserGenre;
};


export type MutationChangeUserPhoneArgs = {
  input: ChangeUserPhone;
};


export type MutationChangeUserPronounArgs = {
  input: ChangeUserPronoun;
};


export type MutationCreateEnterpriseArgs = {
  input: CreateEnterprise;
};


export type MutationCreateEnterpriseChatArgs = {
  input?: InputMaybe<CreateEnterpriseChat>;
};


export type MutationCreateFeedbackArgs = {
  input?: InputMaybe<CreateFeedback>;
};


export type MutationCreateNewComplaintArgs = {
  input: CreateNewComplaint;
};


export type MutationCreateUserArgs = {
  input: CreateUser;
};


export type MutationDescribeComplaintArgs = {
  input: DescribeComplaint;
};


export type MutationEndFeedbackArgs = {
  input: EndFeedback;
};


export type MutationFireEmployeeArgs = {
  input: FireEmployee;
};


export type MutationHireEmployeeArgs = {
  input: HireEmployee;
};


export type MutationInviteToProjectArgs = {
  input: InviteToProject;
};


export type MutationMarkNotificationAsReadArgs = {
  id: Scalars['String']['input'];
};


export type MutationPromoteEmployeeArgs = {
  input: PromoteEmployee;
};


export type MutationRateComplaintArgs = {
  input: RateComplaint;
};


export type MutationRejectHiringInvitationArgs = {
  input: RejectHiringInvitation;
};


export type MutationRemoveFeedbackCommandArgs = {
  input: RemoveFeedbackComment;
};


export type MutationRemoveFeedbackReplyArgs = {
  input: RemoveFeedbackReply;
};


export type MutationSendComplaintArgs = {
  input: SendComplaint;
};


export type MutationUpdateProfileImgArgs = {
  file: Scalars['Upload']['input'];
  id: Scalars['String']['input'];
};


export type MutationUpdateUserAddressArgs = {
  input: UpdateUserAddress;
};

export type NotificationLink = {
  __typename?: 'NotificationLink';
  content: Scalars['String']['output'];
  id: Scalars['String']['output'];
  link: Scalars['String']['output'];
  occurredOn: Scalars['String']['output'];
  owner: Recipient;
  seen: Scalars['Boolean']['output'];
  sender: Recipient;
  title: Scalars['String']['output'];
};

export type Person = {
  __typename?: 'Person';
  address: Address;
  age: Scalars['Int']['output'];
  email: Scalars['String']['output'];
  firstName: Scalars['String']['output'];
  genre: Scalars['String']['output'];
  lastName: Scalars['String']['output'];
  phoneNumber: Scalars['String']['output'];
  profileImg: Scalars['String']['output'];
  pronoun: Scalars['String']['output'];
};

export type PromoteEmployee = {
  employeeId: Scalars['String']['input'];
  enterpriseName: Scalars['String']['input'];
  promoteTo: Scalars['String']['input'];
  promotedById: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  cities: Array<City>;
  complaintById: Complaint;
  complaintsByAuthorOrReceiverId: Array<Maybe<Complaint>>;
  complaintsForFeedbackByEmployeeId: Array<Maybe<Complaint>>;
  complaintsOfResolvedFeedbackByEmployeeId: Array<Maybe<Complaint>>;
  complaintsRatedByAuthorId: Array<Maybe<Complaint>>;
  complaintsRatedByReceiverId: Array<Maybe<Complaint>>;
  complaintsSentForReviewByReceiverId: Array<Maybe<Complaint>>;
  countries: Array<Country>;
  countryStates: Array<CountryState>;
  enterpriseByName: Enterprise;
  enterprisesByAuthenticatedUser: EnterprisesByAuthenticatedUserResult;
  findEnterpriseChat: Chat;
  hiringProcessByAuthenticatedUser: Array<Maybe<HiringProcess>>;
  hiringProcessByEnterpriseName: Array<Maybe<HiringProcess>>;
  industries: Array<Industry>;
  pendingReviewsByAuthorId: Array<Maybe<Complaint>>;
  recipientsByNameLike: Array<Recipient>;
  userById: User;
  userDescriptor: UserDescriptor;
  usersForHiring: UsersForHiringResult;
};


export type QueryCitiesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryComplaintByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryComplaintsByAuthorOrReceiverIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryComplaintsForFeedbackByEmployeeIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryComplaintsOfResolvedFeedbackByEmployeeIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryComplaintsRatedByAuthorIdArgs = {
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryComplaintsRatedByReceiverIdArgs = {
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryComplaintsSentForReviewByReceiverIdArgs = {
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryCountryStatesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryEnterpriseByNameArgs = {
  name: Scalars['String']['input'];
};


export type QueryFindEnterpriseChatArgs = {
  input: FindEnterpriseChat;
};


export type QueryHiringProcessByEnterpriseNameArgs = {
  name: Scalars['String']['input'];
};


export type QueryPendingReviewsByAuthorIdArgs = {
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
};


export type QueryRecipientsByNameLikeArgs = {
  term: Scalars['String']['input'];
};


export type QueryUserByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryUsersForHiringArgs = {
  input: SearchWithPagination;
};

export type RateComplaint = {
  comment: Scalars['String']['input'];
  complaintId: Scalars['String']['input'];
  rate: Scalars['Int']['input'];
  userId: Scalars['String']['input'];
};

export type Rating = {
  __typename?: 'Rating';
  comment?: Maybe<Scalars['String']['output']>;
  createdAt?: Maybe<Scalars['String']['output']>;
  id?: Maybe<Scalars['String']['output']>;
  lastUpdate?: Maybe<Scalars['String']['output']>;
  rate?: Maybe<Scalars['Int']['output']>;
  ratedBy?: Maybe<Recipient>;
  sentToReviewBy?: Maybe<Recipient>;
};

export type Recipient = {
  __typename?: 'Recipient';
  id?: Maybe<Scalars['String']['output']>;
  isEnterprise?: Maybe<Scalars['Boolean']['output']>;
  isOnline?: Maybe<Scalars['Boolean']['output']>;
  subjectEmail?: Maybe<Scalars['String']['output']>;
  subjectName?: Maybe<Scalars['String']['output']>;
  subjectThumbnail?: Maybe<Scalars['String']['output']>;
};

export type RejectHiringInvitation = {
  hiringProcessId: Scalars['String']['input'];
  rejectionReason?: InputMaybe<Scalars['String']['input']>;
  userId: Scalars['String']['input'];
};

export type RemoveFeedbackComment = {
  color: Scalars['String']['input'];
  feedbackId: Scalars['String']['input'];
};

export type RemoveFeedbackReply = {
  color: Scalars['String']['input'];
  feedbackId: Scalars['String']['input'];
  repliesIds: Array<Scalars['String']['input']>;
};

export type ReplyReview = {
  __typename?: 'ReplyReview';
  color: Scalars['String']['output'];
  createdAt: Scalars['String']['output'];
  feedbackId: Scalars['String']['output'];
  id: Scalars['String']['output'];
  replies: Array<Maybe<ComplaintReply>>;
  review?: Maybe<Review>;
  reviewer: User;
};

export type Review = {
  __typename?: 'Review';
  comment?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
};

export enum Roles {
  Assistant = 'ASSISTANT',
  Manager = 'MANAGER',
  Owner = 'OWNER'
}

export type SearchWithPagination = {
  id: Scalars['String']['input'];
  limit: Scalars['Int']['input'];
  offset: Scalars['Int']['input'];
  query: Scalars['String']['input'];
};

export type SendComplaint = {
  body: Scalars['String']['input'];
  complaintId: Scalars['String']['input'];
};

export type Subscription = {
  __typename?: 'Subscription';
  complaintDataByOwnership: ComplaintData;
  complaints: Complaint;
  employeeActivity: EnterpriseActivity;
  employeeComplaintData: ComplaintData;
  employeesActivityLog: EnterpriseActivity;
  enterpriseById: Enterprise;
  feedback: Feedback;
  notifications: NotificationLink;
};


export type SubscriptionComplaintDataByOwnershipArgs = {
  id: Scalars['String']['input'];
};


export type SubscriptionComplaintsArgs = {
  id: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};


export type SubscriptionEmployeeActivityArgs = {
  id: Scalars['String']['input'];
};


export type SubscriptionEmployeeComplaintDataArgs = {
  id: Scalars['String']['input'];
};


export type SubscriptionEmployeesActivityLogArgs = {
  id: Scalars['String']['input'];
};


export type SubscriptionEnterpriseByIdArgs = {
  id: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};


export type SubscriptionFeedbackArgs = {
  feedbackId: Scalars['String']['input'];
};


export type SubscriptionNotificationsArgs = {
  id: Scalars['String']['input'];
};

export type UpdateUserAddress = {
  newCityId: Scalars['Int']['input'];
  newCountryId: Scalars['Int']['input'];
  newCountyId: Scalars['Int']['input'];
  userId: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['String']['output'];
  person: Person;
  status: UserStatus;
  userName: Scalars['String']['output'];
};

export type UserDescriptor = {
  __typename?: 'UserDescriptor';
  authorities?: Maybe<Array<Maybe<GrantedAuthority>>>;
  device?: Maybe<Scalars['String']['output']>;
  fullName: Scalars['String']['output'];
  genre: Scalars['String']['output'];
  geolocation?: Maybe<Scalars['String']['output']>;
  id: Scalars['String']['output'];
  ip?: Maybe<Scalars['String']['output']>;
  loginDate?: Maybe<Scalars['String']['output']>;
  profileImg: Scalars['String']['output'];
  pronoun: Scalars['String']['output'];
  userName: Scalars['String']['output'];
};

export enum UserStatus {
  Offline = 'OFFLINE',
  Online = 'ONLINE'
}

export type UsersForHiringResult = {
  __typename?: 'UsersForHiringResult';
  count: Scalars['Int']['output'];
  limit: Scalars['Int']['output'];
  nextCursor: Scalars['Int']['output'];
  offset: Scalars['Int']['output'];
  prevCursor: Scalars['Int']['output'];
  users: Array<User>;
};

export type CountryFragmentFragment = { __typename?: 'Country', id: number, name: string } & { ' $fragmentName'?: 'CountryFragmentFragment' };

export type AcceptHiringInvitationMutationMutationVariables = Exact<{
  input: AcceptHiringInvitation;
}>;


export type AcceptHiringInvitationMutationMutation = { __typename?: 'Mutation', acceptHiringInvitation: { __typename?: 'HiringProcess', id: string, role: string, status?: HiringProccessStatus | null, reason?: string | null, occurredOn: string, lastUpdate: string, enterprise?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, user: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }, emitedBy: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, updatedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null } };

export type AddFeedbackCommentMutationMutationVariables = Exact<{
  input: AddFeedbackComment;
}>;


export type AddFeedbackCommentMutationMutation = { __typename?: 'Mutation', addFeedbackComment: { __typename?: 'Feedback', id: string } };

export type AddFeedbackReplyMutationMutationVariables = Exact<{
  input: AddFeedbackReply;
}>;


export type AddFeedbackReplyMutationMutation = { __typename?: 'Mutation', addFeedbackReply: { __typename?: 'Feedback', id: string } };

export type CancelHiringProcessMutationMutationVariables = Exact<{
  input: CancelHiringProcess;
}>;


export type CancelHiringProcessMutationMutation = { __typename?: 'Mutation', cancelHiringProcess: { __typename?: 'HiringProcess', id: string } };

export type ChangeEnterpriseAddressMutationMutationVariables = Exact<{
  input: ChangeEnterpriseAddress;
}>;


export type ChangeEnterpriseAddressMutationMutation = { __typename?: 'Mutation', changeEnterpriseAddress: { __typename?: 'Enterprise', id: string } };

export type ChangeEnterpriseBannerMutationMutationVariables = Exact<{
  enterpriseId: Scalars['String']['input'];
  file: Scalars['Upload']['input'];
}>;


export type ChangeEnterpriseBannerMutationMutation = { __typename?: 'Mutation', changeEnterpriseBannerImg: { __typename?: 'Enterprise', id: string } };

export type ChangeEnterpriseEmailMutationMutationVariables = Exact<{
  input: ChangeEnterpriseEmail;
}>;


export type ChangeEnterpriseEmailMutationMutation = { __typename?: 'Mutation', changeEnterpriseEmail: { __typename?: 'Enterprise', id: string } };

export type ChangeEnterpriseLogoMutationMutationVariables = Exact<{
  enterpriseId: Scalars['String']['input'];
  file: Scalars['Upload']['input'];
}>;


export type ChangeEnterpriseLogoMutationMutation = { __typename?: 'Mutation', changeEnterpriseLogoImg: { __typename?: 'Enterprise', id: string } };

export type ChangeEnterprisePhoneMutationMutationVariables = Exact<{
  input: ChangeEnterprisePhone;
}>;


export type ChangeEnterprisePhoneMutationMutation = { __typename?: 'Mutation', changeEnterprisePhone: { __typename?: 'Enterprise', id: string } };

export type ChangeEnterpriseWebsiteMutationMutationVariables = Exact<{
  input: ChangeEnterpriseWebsite;
}>;


export type ChangeEnterpriseWebsiteMutationMutation = { __typename?: 'Mutation', changeEnterpriseWebsite: { __typename?: 'Enterprise', id: string } };

export type ChangeUserGenreMutationMutationVariables = Exact<{
  input: ChangeUserGenre;
}>;


export type ChangeUserGenreMutationMutation = { __typename?: 'Mutation', changeUserGenre: { __typename?: 'User', id: string } };

export type ChangeUserPhoneMutationMutationVariables = Exact<{
  input: ChangeUserPhone;
}>;


export type ChangeUserPhoneMutationMutation = { __typename?: 'Mutation', changeUserPhone: { __typename?: 'User', id: string } };

export type ChangeUserPronounMutationMutationVariables = Exact<{
  input: ChangeUserPronoun;
}>;


export type ChangeUserPronounMutationMutation = { __typename?: 'Mutation', changeUserPronoun: { __typename?: 'User', id: string } };

export type CreateEnterpriseChatMutationMutationVariables = Exact<{
  input: CreateEnterpriseChat;
}>;


export type CreateEnterpriseChatMutationMutation = { __typename?: 'Mutation', createEnterpriseChat: { __typename?: 'Chat', id: string, enterpriseId: string, recipientOne: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, recipientTwo: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, replies: Array<{ __typename?: 'ChatReply', id: string, chatId: string, content: string, createdAt: string, updatedAt: string, seen: boolean, sender: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } } | null> } };

export type CreateEnterpriseMutationVariables = Exact<{
  input: CreateEnterprise;
}>;


export type CreateEnterpriseMutation = { __typename?: 'Mutation', createEnterprise: { __typename?: 'Enterprise', name: string } };

export type CreateFeedbackMutationMutationVariables = Exact<{
  input?: InputMaybe<CreateFeedback>;
}>;


export type CreateFeedbackMutationMutation = { __typename?: 'Mutation', createFeedback: { __typename?: 'Feedback', id: string, complaintId: string, enterpriseId: string, reviewedAt: string, updatedAt: string, isDone: boolean, replyReview: Array<{ __typename?: 'ReplyReview', id: string, feedbackId: string, color: string, createdAt: string, reviewer: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string } }, replies: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, createdAt?: string | null, read?: boolean | null, readAt?: string | null, updatedAt?: string | null, isEnterprise?: boolean | null, enterpriseId?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null } | null>, review?: { __typename?: 'Review', id: string, comment?: string | null } | null } | null> } };

export type CreateNewComplaintMutationMutationVariables = Exact<{
  input: CreateNewComplaint;
}>;


export type CreateNewComplaintMutationMutation = { __typename?: 'Mutation', createNewComplaint?: { __typename?: 'Complaint', id?: string | null } | null };

export type CreateUserMutationMutationVariables = Exact<{
  input: CreateUser;
}>;


export type CreateUserMutationMutation = { __typename?: 'Mutation', createUser: { __typename?: 'User', userName: string } };

export type DescribeComplaintMutationMutationVariables = Exact<{
  input: DescribeComplaint;
}>;


export type DescribeComplaintMutationMutation = { __typename?: 'Mutation', describeComplaint?: { __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, receiver?: { __typename?: 'Recipient', id?: string | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null } | null } | null> | null } | null };

export type EndFeedbackMutationMutationVariables = Exact<{
  input: EndFeedback;
}>;


export type EndFeedbackMutationMutation = { __typename?: 'Mutation', endFeedback: { __typename?: 'Feedback', id: string } };

export type FireEmployeeMutationMutationVariables = Exact<{
  input: FireEmployee;
}>;


export type FireEmployeeMutationMutation = { __typename?: 'Mutation', fireEmployee: { __typename?: 'Enterprise', id: string } };

export type HireEmployeeMutationMutationVariables = Exact<{
  input: HireEmployee;
}>;


export type HireEmployeeMutationMutation = { __typename?: 'Mutation', hireEmployee: { __typename?: 'HiringProcess', id: string } };

export type InviteToProjectMutationMutationVariables = Exact<{
  input: InviteToProject;
}>;


export type InviteToProjectMutationMutation = { __typename?: 'Mutation', inviteToProject: { __typename?: 'HiringProcess', id: string, role: string, status?: HiringProccessStatus | null, reason?: string | null, occurredOn: string, lastUpdate: string, enterprise?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, user: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }, emitedBy: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, updatedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null } };

export type MarkNotificationAsReadMutationMutationVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type MarkNotificationAsReadMutationMutation = { __typename?: 'Mutation', markNotificationAsRead: { __typename?: 'NotificationLink', id: string } };

export type PromoteEmployeeMutationMutationVariables = Exact<{
  input: PromoteEmployee;
}>;


export type PromoteEmployeeMutationMutation = { __typename?: 'Mutation', promoteEmployee: { __typename?: 'Employee', id: string, enterprisePosition: string } };

export type RateComplaintMutationMutationVariables = Exact<{
  input: RateComplaint;
}>;


export type RateComplaintMutationMutation = { __typename?: 'Mutation', rateComplaint: { __typename?: 'Rating', id?: string | null, rate?: number | null, comment?: string | null, createdAt?: string | null, lastUpdate?: string | null, sentToReviewBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, ratedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null } };

export type RejectHiringInvitationMutationMutationVariables = Exact<{
  input: RejectHiringInvitation;
}>;


export type RejectHiringInvitationMutationMutation = { __typename?: 'Mutation', rejectHiringInvitation: { __typename?: 'HiringProcess', id: string, role: string, status?: HiringProccessStatus | null, reason?: string | null, occurredOn: string, lastUpdate: string, enterprise?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, user: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }, emitedBy: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, updatedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null } };

export type RemoveFeedbackCommentMutationMutationVariables = Exact<{
  input: RemoveFeedbackComment;
}>;


export type RemoveFeedbackCommentMutationMutation = { __typename?: 'Mutation', removeFeedbackCommand: { __typename?: 'Feedback', id: string } };

export type RemoveFeedbackReplyMutationMutationVariables = Exact<{
  input: RemoveFeedbackReply;
}>;


export type RemoveFeedbackReplyMutationMutation = { __typename?: 'Mutation', removeFeedbackReply: { __typename?: 'Feedback', id: string } };

export type SendComplaintMutationMutationVariables = Exact<{
  input: SendComplaint;
}>;


export type SendComplaintMutationMutation = { __typename?: 'Mutation', sendComplaint?: { __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, receiver?: { __typename?: 'Recipient', id?: string | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null } | null } | null> | null } | null };

export type UpdateFirstNameMutationMutationVariables = Exact<{
  input: ChangeUserFirstName;
}>;


export type UpdateFirstNameMutationMutation = { __typename?: 'Mutation', changeFirstName: { __typename?: 'User', id: string } };

export type UpdateLastNameMutationMutationVariables = Exact<{
  input: ChangeUserLastName;
}>;


export type UpdateLastNameMutationMutation = { __typename?: 'Mutation', changeLastName: { __typename?: 'User', id: string } };

export type UpdatePasswordMutationMutationVariables = Exact<{
  input: ChangePassword;
}>;


export type UpdatePasswordMutationMutation = { __typename?: 'Mutation', changePassword: { __typename?: 'User', id: string } };

export type UpdateProfileImageMutationMutationVariables = Exact<{
  id: Scalars['String']['input'];
  file: Scalars['Upload']['input'];
}>;


export type UpdateProfileImageMutationMutation = { __typename?: 'Mutation', updateProfileImg: { __typename?: 'User', id: string } };

export type UpdateUserAddressMutationMutationVariables = Exact<{
  input: UpdateUserAddress;
}>;


export type UpdateUserAddressMutationMutation = { __typename?: 'Mutation', updateUserAddress: { __typename?: 'User', id: string } };

export type CitiesQueryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CitiesQueryQuery = { __typename?: 'Query', cities: Array<{ __typename?: 'City', id: number, name: string, countryCode: string, latitude: number, longitude: number }> };

export type ComplaintByIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ComplaintByIdQueryQuery = { __typename?: 'Query', complaintById: { __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, isEnterprise?: boolean | null, enterpriseId?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null, isOnline?: boolean | null } | null } | null> | null } };

export type ComplaintsByAuthorIdOrReceiverIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ComplaintsByAuthorIdOrReceiverIdQueryQuery = { __typename?: 'Query', complaintsByAuthorOrReceiverId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null } | null } | null> | null } | null> };

export type ComplaintsForFeedbackByEmployeeIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ComplaintsForFeedbackByEmployeeIdQueryQuery = { __typename?: 'Query', complaintsForFeedbackByEmployeeId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null } | null } | null> | null } | null> };

export type ComplaintsOfResolvedFeedbackByEmployeeIdQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ComplaintsOfResolvedFeedbackByEmployeeIdQuery = { __typename?: 'Query', complaintsOfResolvedFeedbackByEmployeeId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null } | null } | null> | null } | null> };

export type ComplaintsRatedByAuthorIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
}>;


export type ComplaintsRatedByAuthorIdQueryQuery = { __typename?: 'Query', complaintsRatedByAuthorId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, rating?: { __typename?: 'Rating', id?: string | null, rate?: number | null, comment?: string | null, createdAt?: string | null, lastUpdate?: string | null, sentToReviewBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, ratedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null } | null } | null> | null } | null> };

export type ComplaintsRatedByReceiverIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
}>;


export type ComplaintsRatedByReceiverIdQueryQuery = { __typename?: 'Query', complaintsRatedByReceiverId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, rating?: { __typename?: 'Rating', id?: string | null, rate?: number | null, comment?: string | null, createdAt?: string | null, lastUpdate?: string | null, sentToReviewBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, ratedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null } | null } | null> | null } | null> };

export type ComplaintsSentForReviewByReceiverIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
}>;


export type ComplaintsSentForReviewByReceiverIdQueryQuery = { __typename?: 'Query', complaintsSentForReviewByReceiverId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, rating?: { __typename?: 'Rating', id?: string | null, rate?: number | null, comment?: string | null, createdAt?: string | null, lastUpdate?: string | null, sentToReviewBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, ratedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null } | null } | null> | null } | null> };

export type CountriesQueryVariables = Exact<{ [key: string]: never; }>;


export type CountriesQuery = { __typename?: 'Query', countries: Array<{ __typename?: 'Country', id: number, name: string, phoneCode: string }> };

export type CountryStatesQueryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CountryStatesQueryQuery = { __typename?: 'Query', countryStates: Array<{ __typename?: 'CountryState', id: number, name: string }> };

export type EnterpriseByNameQueryQueryVariables = Exact<{
  name: Scalars['String']['input'];
}>;


export type EnterpriseByNameQueryQuery = { __typename?: 'Query', enterpriseByName: { __typename?: 'Enterprise', id: string, name: string, logoImg: string, bannerImg: string, website: string, email: string, phoneNumber: string, foundationDate: string, ownerId: string, address: { __typename?: 'Address', country: string, countryState: string, city: string }, industry: { __typename?: 'Industry', id: number, name: string }, employees: Array<{ __typename?: 'Employee', id: string, enterpriseId: string, userId: string, hiringDate: string, approvedHiring: boolean, approvedHiringAt: string, enterprisePosition: string, user: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } } } | null> } };

export type EnterprisesByAuthenticatedUserQueryQueryVariables = Exact<{ [key: string]: never; }>;


export type EnterprisesByAuthenticatedUserQueryQuery = { __typename?: 'Query', enterprisesByAuthenticatedUser: { __typename?: 'EnterprisesByAuthenticatedUserResult', enterprises: Array<{ __typename?: 'EnterpriseByAuthenticatedUser', authority: { __typename?: 'GrantedAuthority', authority: Roles, enterpriseId: string, principal: string }, enterprise?: { __typename?: 'Enterprise', name: string, logoImg: string, bannerImg: string, website: string, email: string, phoneNumber: string, foundationDate: string, ownerId: string, address: { __typename?: 'Address', country: string, countryState: string, city: string }, industry: { __typename?: 'Industry', id: number, name: string } } | null }>, offices: Array<{ __typename?: 'EnterpriseByAuthenticatedUser', authority: { __typename?: 'GrantedAuthority', authority: Roles, enterpriseId: string, principal: string }, enterprise?: { __typename?: 'Enterprise', name: string, logoImg: string, bannerImg: string, website: string, email: string, phoneNumber: string, foundationDate: string, ownerId: string, address: { __typename?: 'Address', country: string, countryState: string, city: string }, industry: { __typename?: 'Industry', id: number, name: string } } | null }> } };

export type FindEnterpriseChatQueryQueryVariables = Exact<{
  input: FindEnterpriseChat;
}>;


export type FindEnterpriseChatQueryQuery = { __typename?: 'Query', findEnterpriseChat: { __typename?: 'Chat', id: string, enterpriseId: string, recipientOne: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, recipientTwo: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, replies: Array<{ __typename?: 'ChatReply', id: string, chatId: string, content: string, createdAt: string, updatedAt: string, seen: boolean, sender: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } } | null> } };

export type HiringProcessByAuthenticatedUserQueryQueryVariables = Exact<{ [key: string]: never; }>;


export type HiringProcessByAuthenticatedUserQueryQuery = { __typename?: 'Query', hiringProcessByAuthenticatedUser: Array<{ __typename?: 'HiringProcess', id: string, role: string, status?: HiringProccessStatus | null, reason?: string | null, occurredOn: string, lastUpdate: string, enterprise?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, user: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }, emitedBy: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, updatedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, industry?: { __typename?: 'Industry', id: number, name: string } | null } | null> };

export type HiringProcessByEnterpriseNameQueryVariables = Exact<{
  name: Scalars['String']['input'];
}>;


export type HiringProcessByEnterpriseNameQuery = { __typename?: 'Query', hiringProcessByEnterpriseName: Array<{ __typename?: 'HiringProcess', id: string, role: string, status?: HiringProccessStatus | null, reason?: string | null, occurredOn: string, lastUpdate: string, enterprise?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, user: { __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }, emitedBy: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null }, updatedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, subjectEmail?: string | null } | null, industry?: { __typename?: 'Industry', id: number, name: string } | null } | null> };

export type IndustriesQueryVariables = Exact<{ [key: string]: never; }>;


export type IndustriesQuery = { __typename?: 'Query', industries: Array<{ __typename?: 'Industry', id: number, name: string }> };

export type PendingReviewsByAuthorIdQueryVariables = Exact<{
  id: Scalars['String']['input'];
  term?: InputMaybe<Scalars['String']['input']>;
}>;


export type PendingReviewsByAuthorIdQuery = { __typename?: 'Query', pendingReviewsByAuthorId: Array<{ __typename?: 'Complaint', id?: string | null, status?: ComplaintStatus | null, title?: string | null, description?: string | null, createdAt?: string | null, updatedAt?: string | null, author?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, receiver?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, rating?: { __typename?: 'Rating', id?: string | null, rate?: number | null, comment?: string | null, createdAt?: string | null, lastUpdate?: string | null, sentToReviewBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null, ratedBy?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isOnline?: boolean | null, isEnterprise?: boolean | null } | null } | null, replies?: Array<{ __typename?: 'ComplaintReply', id?: string | null, complaintId?: string | null, body?: string | null, read?: boolean | null, readAt?: string | null, createdAt?: string | null, updatedAt?: string | null, sender?: { __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null } | null } | null> | null } | null> };

export type RecipientsByNameLikeQueryQueryVariables = Exact<{
  term: Scalars['String']['input'];
}>;


export type RecipientsByNameLikeQueryQuery = { __typename?: 'Query', recipientsByNameLike: Array<{ __typename?: 'Recipient', id?: string | null, subjectName?: string | null, subjectThumbnail?: string | null, isEnterprise?: boolean | null }> };

export type UserQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type UserQuery = { __typename?: 'Query', userById: { __typename?: 'User', userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } } };

export type UserDescriptorQueryVariables = Exact<{ [key: string]: never; }>;


export type UserDescriptorQuery = { __typename?: 'Query', userDescriptor: { __typename?: 'UserDescriptor', id: string, userName: string, fullName: string, profileImg: string, genre: string, pronoun: string, authorities?: Array<{ __typename?: 'GrantedAuthority', enterpriseId: string, principal: string, authority: Roles } | null> | null } };

export type UsersForHiringQueryQueryVariables = Exact<{
  input: SearchWithPagination;
}>;


export type UsersForHiringQueryQuery = { __typename?: 'Query', usersForHiring: { __typename?: 'UsersForHiringResult', count: number, limit: number, offset: number, nextCursor: number, prevCursor: number, users: Array<{ __typename?: 'User', id: string, userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }> } };

export const CountryFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"CountryFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Country"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]} as unknown as DocumentNode<CountryFragmentFragment, unknown>;
export const AcceptHiringInvitationMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"acceptHiringInvitationMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"AcceptHiringInvitation"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"acceptHiringInvitation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"reason"}},{"kind":"Field","name":{"kind":"Name","value":"emitedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"occurredOn"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}}]}}]}}]} as unknown as DocumentNode<AcceptHiringInvitationMutationMutation, AcceptHiringInvitationMutationMutationVariables>;
export const AddFeedbackCommentMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addFeedbackCommentMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"AddFeedbackComment"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addFeedbackComment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<AddFeedbackCommentMutationMutation, AddFeedbackCommentMutationMutationVariables>;
export const AddFeedbackReplyMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addFeedbackReplyMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"AddFeedbackReply"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addFeedbackReply"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<AddFeedbackReplyMutationMutation, AddFeedbackReplyMutationMutationVariables>;
export const CancelHiringProcessMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"cancelHiringProcessMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CancelHiringProcess"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cancelHiringProcess"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CancelHiringProcessMutationMutation, CancelHiringProcessMutationMutationVariables>;
export const ChangeEnterpriseAddressMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeEnterpriseAddressMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeEnterpriseAddress"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeEnterpriseAddress"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeEnterpriseAddressMutationMutation, ChangeEnterpriseAddressMutationMutationVariables>;
export const ChangeEnterpriseBannerMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeEnterpriseBannerMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"enterpriseId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"file"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Upload"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeEnterpriseBannerImg"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"enterpriseId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"enterpriseId"}}},{"kind":"Argument","name":{"kind":"Name","value":"file"},"value":{"kind":"Variable","name":{"kind":"Name","value":"file"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeEnterpriseBannerMutationMutation, ChangeEnterpriseBannerMutationMutationVariables>;
export const ChangeEnterpriseEmailMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeEnterpriseEmailMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeEnterpriseEmail"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeEnterpriseEmail"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeEnterpriseEmailMutationMutation, ChangeEnterpriseEmailMutationMutationVariables>;
export const ChangeEnterpriseLogoMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeEnterpriseLogoMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"enterpriseId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"file"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Upload"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeEnterpriseLogoImg"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"enterpriseId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"enterpriseId"}}},{"kind":"Argument","name":{"kind":"Name","value":"file"},"value":{"kind":"Variable","name":{"kind":"Name","value":"file"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeEnterpriseLogoMutationMutation, ChangeEnterpriseLogoMutationMutationVariables>;
export const ChangeEnterprisePhoneMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeEnterprisePhoneMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeEnterprisePhone"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeEnterprisePhone"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeEnterprisePhoneMutationMutation, ChangeEnterprisePhoneMutationMutationVariables>;
export const ChangeEnterpriseWebsiteMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeEnterpriseWebsiteMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeEnterpriseWebsite"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeEnterpriseWebsite"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeEnterpriseWebsiteMutationMutation, ChangeEnterpriseWebsiteMutationMutationVariables>;
export const ChangeUserGenreMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeUserGenreMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeUserGenre"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeUserGenre"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeUserGenreMutationMutation, ChangeUserGenreMutationMutationVariables>;
export const ChangeUserPhoneMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeUserPhoneMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeUserPhone"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeUserPhone"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeUserPhoneMutationMutation, ChangeUserPhoneMutationMutationVariables>;
export const ChangeUserPronounMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeUserPronounMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeUserPronoun"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeUserPronoun"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeUserPronounMutationMutation, ChangeUserPronounMutationMutationVariables>;
export const CreateEnterpriseChatMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createEnterpriseChatMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateEnterpriseChat"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createEnterpriseChat"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"recipientOne"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"recipientTwo"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"chatId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"seen"}}]}}]}}]}}]} as unknown as DocumentNode<CreateEnterpriseChatMutationMutation, CreateEnterpriseChatMutationMutationVariables>;
export const CreateEnterpriseDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createEnterprise"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateEnterprise"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createEnterprise"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CreateEnterpriseMutation, CreateEnterpriseMutationVariables>;
export const CreateFeedbackMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createFeedbackMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateFeedback"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createFeedback"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"replyReview"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"feedbackId"}},{"kind":"Field","name":{"kind":"Name","value":"reviewer"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}}]}},{"kind":"Field","name":{"kind":"Name","value":"review"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}}]}},{"kind":"Field","name":{"kind":"Name","value":"color"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"reviewedAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"isDone"}}]}}]}}]} as unknown as DocumentNode<CreateFeedbackMutationMutation, CreateFeedbackMutationMutationVariables>;
export const CreateNewComplaintMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createNewComplaintMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateNewComplaint"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createNewComplaint"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateNewComplaintMutationMutation, CreateNewComplaintMutationMutationVariables>;
export const CreateUserMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUserMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateUser"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}}]}}]}}]} as unknown as DocumentNode<CreateUserMutationMutation, CreateUserMutationMutationVariables>;
export const DescribeComplaintMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"describeComplaintMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"DescribeComplaint"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"describeComplaint"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<DescribeComplaintMutationMutation, DescribeComplaintMutationMutationVariables>;
export const EndFeedbackMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"endFeedbackMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"EndFeedback"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"endFeedback"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<EndFeedbackMutationMutation, EndFeedbackMutationMutationVariables>;
export const FireEmployeeMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"fireEmployeeMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"FireEmployee"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"fireEmployee"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<FireEmployeeMutationMutation, FireEmployeeMutationMutationVariables>;
export const HireEmployeeMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"hireEmployeeMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"HireEmployee"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hireEmployee"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<HireEmployeeMutationMutation, HireEmployeeMutationMutationVariables>;
export const InviteToProjectMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"inviteToProjectMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"InviteToProject"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"inviteToProject"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"reason"}},{"kind":"Field","name":{"kind":"Name","value":"emitedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"occurredOn"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}}]}}]}}]} as unknown as DocumentNode<InviteToProjectMutationMutation, InviteToProjectMutationMutationVariables>;
export const MarkNotificationAsReadMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"markNotificationAsReadMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"markNotificationAsRead"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<MarkNotificationAsReadMutationMutation, MarkNotificationAsReadMutationMutationVariables>;
export const PromoteEmployeeMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"promoteEmployeeMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"PromoteEmployee"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"promoteEmployee"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterprisePosition"}}]}}]}}]} as unknown as DocumentNode<PromoteEmployeeMutationMutation, PromoteEmployeeMutationMutationVariables>;
export const RateComplaintMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"rateComplaintMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RateComplaint"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"rateComplaint"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"sentToReviewBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"ratedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}}]}}]}}]} as unknown as DocumentNode<RateComplaintMutationMutation, RateComplaintMutationMutationVariables>;
export const RejectHiringInvitationMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"rejectHiringInvitationMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RejectHiringInvitation"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"rejectHiringInvitation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"reason"}},{"kind":"Field","name":{"kind":"Name","value":"emitedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"occurredOn"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}}]}}]}}]} as unknown as DocumentNode<RejectHiringInvitationMutationMutation, RejectHiringInvitationMutationMutationVariables>;
export const RemoveFeedbackCommentMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"removeFeedbackCommentMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RemoveFeedbackComment"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"removeFeedbackCommand"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<RemoveFeedbackCommentMutationMutation, RemoveFeedbackCommentMutationMutationVariables>;
export const RemoveFeedbackReplyMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"removeFeedbackReplyMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"RemoveFeedbackReply"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"removeFeedbackReply"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<RemoveFeedbackReplyMutationMutation, RemoveFeedbackReplyMutationMutationVariables>;
export const SendComplaintMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"sendComplaintMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"SendComplaint"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sendComplaint"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<SendComplaintMutationMutation, SendComplaintMutationMutationVariables>;
export const UpdateFirstNameMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateFirstNameMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeUserFirstName"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeFirstName"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateFirstNameMutationMutation, UpdateFirstNameMutationMutationVariables>;
export const UpdateLastNameMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateLastNameMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangeUserLastName"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changeLastName"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateLastNameMutationMutation, UpdateLastNameMutationMutationVariables>;
export const UpdatePasswordMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updatePasswordMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ChangePassword"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"changePassword"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdatePasswordMutationMutation, UpdatePasswordMutationMutationVariables>;
export const UpdateProfileImageMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateProfileImageMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"file"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Upload"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateProfileImg"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"file"},"value":{"kind":"Variable","name":{"kind":"Name","value":"file"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateProfileImageMutationMutation, UpdateProfileImageMutationMutationVariables>;
export const UpdateUserAddressMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateUserAddressMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UpdateUserAddress"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUserAddress"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateUserAddressMutationMutation, UpdateUserAddressMutationMutationVariables>;
export const CitiesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CitiesQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cities"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"countryCode"}},{"kind":"Field","name":{"kind":"Name","value":"latitude"}},{"kind":"Field","name":{"kind":"Name","value":"longitude"}}]}}]}}]} as unknown as DocumentNode<CitiesQueryQuery, CitiesQueryQueryVariables>;
export const ComplaintByIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"ComplaintByIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintById"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintByIdQueryQuery, ComplaintByIdQueryQueryVariables>;
export const ComplaintsByAuthorIdOrReceiverIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsByAuthorIdOrReceiverIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsByAuthorOrReceiverId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintsByAuthorIdOrReceiverIdQueryQuery, ComplaintsByAuthorIdOrReceiverIdQueryQueryVariables>;
export const ComplaintsForFeedbackByEmployeeIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsForFeedbackByEmployeeIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsForFeedbackByEmployeeId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintsForFeedbackByEmployeeIdQueryQuery, ComplaintsForFeedbackByEmployeeIdQueryQueryVariables>;
export const ComplaintsOfResolvedFeedbackByEmployeeIdDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsOfResolvedFeedbackByEmployeeId"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsOfResolvedFeedbackByEmployeeId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintsOfResolvedFeedbackByEmployeeIdQuery, ComplaintsOfResolvedFeedbackByEmployeeIdQueryVariables>;
export const ComplaintsRatedByAuthorIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsRatedByAuthorIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"term"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsRatedByAuthorId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"term"},"value":{"kind":"Variable","name":{"kind":"Name","value":"term"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"rating"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"sentToReviewBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"ratedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintsRatedByAuthorIdQueryQuery, ComplaintsRatedByAuthorIdQueryQueryVariables>;
export const ComplaintsRatedByReceiverIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsRatedByReceiverIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"term"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsRatedByReceiverId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"term"},"value":{"kind":"Variable","name":{"kind":"Name","value":"term"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"rating"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"sentToReviewBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"ratedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintsRatedByReceiverIdQueryQuery, ComplaintsRatedByReceiverIdQueryQueryVariables>;
export const ComplaintsSentForReviewByReceiverIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsSentForReviewByReceiverIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"term"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsSentForReviewByReceiverId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"term"},"value":{"kind":"Variable","name":{"kind":"Name","value":"term"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"rating"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"sentToReviewBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"ratedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<ComplaintsSentForReviewByReceiverIdQueryQuery, ComplaintsSentForReviewByReceiverIdQueryQueryVariables>;
export const CountriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Countries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"countries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"phoneCode"}}]}}]}}]} as unknown as DocumentNode<CountriesQuery, CountriesQueryVariables>;
export const CountryStatesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CountryStatesQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"countryStates"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CountryStatesQueryQuery, CountryStatesQueryQueryVariables>;
export const EnterpriseByNameQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"enterpriseByNameQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterpriseByName"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoImg"}},{"kind":"Field","name":{"kind":"Name","value":"bannerImg"}},{"kind":"Field","name":{"kind":"Name","value":"website"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"foundationDate"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}},{"kind":"Field","name":{"kind":"Name","value":"employees"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"hiringDate"}},{"kind":"Field","name":{"kind":"Name","value":"approvedHiring"}},{"kind":"Field","name":{"kind":"Name","value":"approvedHiringAt"}},{"kind":"Field","name":{"kind":"Name","value":"enterprisePosition"}}]}}]}}]}}]} as unknown as DocumentNode<EnterpriseByNameQueryQuery, EnterpriseByNameQueryQueryVariables>;
export const EnterprisesByAuthenticatedUserQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"EnterprisesByAuthenticatedUserQuery"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterprisesByAuthenticatedUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterprises"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"principal"}}]}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoImg"}},{"kind":"Field","name":{"kind":"Name","value":"bannerImg"}},{"kind":"Field","name":{"kind":"Name","value":"website"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"foundationDate"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"offices"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"principal"}}]}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoImg"}},{"kind":"Field","name":{"kind":"Name","value":"bannerImg"}},{"kind":"Field","name":{"kind":"Name","value":"website"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"foundationDate"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}}]}}]}}]}}]}}]} as unknown as DocumentNode<EnterprisesByAuthenticatedUserQueryQuery, EnterprisesByAuthenticatedUserQueryQueryVariables>;
export const FindEnterpriseChatQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"findEnterpriseChatQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"FindEnterpriseChat"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"findEnterpriseChat"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"recipientOne"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"recipientTwo"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"chatId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"seen"}}]}}]}}]}}]} as unknown as DocumentNode<FindEnterpriseChatQueryQuery, FindEnterpriseChatQueryQueryVariables>;
export const HiringProcessByAuthenticatedUserQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"hiringProcessByAuthenticatedUserQuery"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hiringProcessByAuthenticatedUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"reason"}},{"kind":"Field","name":{"kind":"Name","value":"emitedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"occurredOn"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<HiringProcessByAuthenticatedUserQueryQuery, HiringProcessByAuthenticatedUserQueryQueryVariables>;
export const HiringProcessByEnterpriseNameDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"hiringProcessByEnterpriseName"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hiringProcessByEnterpriseName"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"reason"}},{"kind":"Field","name":{"kind":"Name","value":"emitedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"occurredOn"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}},{"kind":"Field","name":{"kind":"Name","value":"updatedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"subjectEmail"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<HiringProcessByEnterpriseNameQuery, HiringProcessByEnterpriseNameQueryVariables>;
export const IndustriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"industries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"industries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<IndustriesQuery, IndustriesQueryVariables>;
export const PendingReviewsByAuthorIdDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"pendingReviewsByAuthorId"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"term"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pendingReviewsByAuthorId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"term"},"value":{"kind":"Variable","name":{"kind":"Name","value":"term"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"author"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}},{"kind":"Field","name":{"kind":"Name","value":"rating"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"rate"}},{"kind":"Field","name":{"kind":"Name","value":"comment"}},{"kind":"Field","name":{"kind":"Name","value":"sentToReviewBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"ratedBy"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isOnline"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"lastUpdate"}}]}},{"kind":"Field","name":{"kind":"Name","value":"replies"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"complaintId"}},{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"read"}},{"kind":"Field","name":{"kind":"Name","value":"readAt"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<PendingReviewsByAuthorIdQuery, PendingReviewsByAuthorIdQueryVariables>;
export const RecipientsByNameLikeQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"recipientsByNameLikeQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"term"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"recipientsByNameLike"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"term"},"value":{"kind":"Variable","name":{"kind":"Name","value":"term"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"subjectName"}},{"kind":"Field","name":{"kind":"Name","value":"subjectThumbnail"}},{"kind":"Field","name":{"kind":"Name","value":"isEnterprise"}}]}}]}}]} as unknown as DocumentNode<RecipientsByNameLikeQueryQuery, RecipientsByNameLikeQueryQueryVariables>;
export const UserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"User"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userById"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<UserQuery, UserQueryVariables>;
export const UserDescriptorDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"UserDescriptor"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userDescriptor"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"fullName"}},{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"authorities"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"principal"}},{"kind":"Field","name":{"kind":"Name","value":"authority"}}]}}]}}]}}]} as unknown as DocumentNode<UserDescriptorQuery, UserDescriptorQueryVariables>;
export const UsersForHiringQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"usersForHiringQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"SearchWithPagination"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"usersForHiring"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"users"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"count"}},{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"nextCursor"}},{"kind":"Field","name":{"kind":"Name","value":"prevCursor"}}]}}]}}]} as unknown as DocumentNode<UsersForHiringQueryQuery, UsersForHiringQueryQueryVariables>;
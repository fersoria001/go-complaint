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
  city: Scalars['String']['output'];
  country: Scalars['String']['output'];
  countryState: Scalars['String']['output'];
};

export type City = {
  __typename?: 'City';
  countryCode: Scalars['String']['output'];
  id: Scalars['Int']['output'];
  latitude: Scalars['Float']['output'];
  longitude: Scalars['Float']['output'];
  name: Scalars['String']['output'];
};

export type ComplaintInfo = {
  __typename?: 'ComplaintInfo';
  avgRating: Scalars['Float']['output'];
  pending: Scalars['Int']['output'];
  received: Scalars['Int']['output'];
  resolved: Scalars['Int']['output'];
  reviewed: Scalars['Int']['output'];
  total: Scalars['Int']['output'];
};

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

export type Employee = {
  __typename?: 'Employee';
  approvedHiring: Scalars['Boolean']['output'];
  approvedHiringAt: Scalars['String']['output'];
  enterpriseId: Scalars['String']['output'];
  enterprisePosition: Scalars['String']['output'];
  hiringDate: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  user: User;
  userId: Scalars['String']['output'];
};

export type Enterprise = {
  __typename?: 'Enterprise';
  address: Address;
  bannerImg: Scalars['String']['output'];
  email: Scalars['String']['output'];
  employees: Array<Maybe<Employee>>;
  foundationDate: Scalars['String']['output'];
  industry: Scalars['String']['output'];
  logoImg: Scalars['String']['output'];
  name: Scalars['String']['output'];
  ownerId: Scalars['String']['output'];
  phoneNumber: Scalars['String']['output'];
  website: Scalars['String']['output'];
};

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

export type GrantedAuthority = {
  __typename?: 'GrantedAuthority';
  authority: Scalars['String']['output'];
  enterpriseId: Scalars['String']['output'];
};

export type HiringInvitation = {
  __typename?: 'HiringInvitation';
  enterpriseEmail: Scalars['String']['output'];
  enterpriseId: Scalars['String']['output'];
  enterpriseLogoImg: Scalars['String']['output'];
  enterprisePhone: Scalars['String']['output'];
  eventId: Scalars['String']['output'];
  fullName: Scalars['String']['output'];
  occurredOn: Scalars['String']['output'];
  ownerId: Scalars['String']['output'];
  proposedPosition: Scalars['String']['output'];
  reason: Scalars['String']['output'];
  seen: Scalars['Boolean']['output'];
  status: HiringProccessState;
};

export enum HiringProccessState {
  Accepted = 'accepted',
  Canceled = 'canceled',
  Fired = 'fired',
  Hired = 'hired',
  Leaved = 'leaved',
  Pending = 'pending',
  Rated = 'rated',
  Rejected = 'rejected',
  UserAccepted = 'user_accepted',
  Waiting = 'waiting'
}

export type Industry = {
  __typename?: 'Industry';
  id: Scalars['Int']['output'];
  name: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createEnterprise: Enterprise;
  createUser: User;
};


export type MutationCreateEnterpriseArgs = {
  input: CreateEnterprise;
};


export type MutationCreateUserArgs = {
  input: CreateUser;
};

export type Notification = {
  __typename?: 'Notification';
  content: Scalars['String']['output'];
  id: Scalars['String']['output'];
  link: Scalars['String']['output'];
  occurredOn: Scalars['String']['output'];
  ownerId: Scalars['String']['output'];
  seen: Scalars['Boolean']['output'];
  thumbnail: Scalars['String']['output'];
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

export type Query = {
  __typename?: 'Query';
  cities: Array<City>;
  complaintsReceivedInfo: ComplaintInfo;
  countries: Array<Country>;
  countryStates: Array<CountryState>;
  enterpriseById: Enterprise;
  enterprisesByAuthenticatedUser: EnterprisesByAuthenticatedUserResult;
  hiringInvitationsByAuthenticatedUser: Array<HiringInvitation>;
  industries: Array<Industry>;
  userById: User;
  userDescriptor: UserDescriptor;
  usersForHiring: UsersForHiringResult;
};


export type QueryCitiesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryComplaintsReceivedInfoArgs = {
  id: Scalars['String']['input'];
};


export type QueryCountryStatesArgs = {
  id: Scalars['Int']['input'];
};


export type QueryEnterpriseByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryUserByIdArgs = {
  id: Scalars['String']['input'];
};


export type QueryUsersForHiringArgs = {
  input: SearchWithPagination;
};

export type SearchWithPagination = {
  id: Scalars['String']['input'];
  limit: Scalars['Int']['input'];
  offset: Scalars['Int']['input'];
  query: Scalars['String']['input'];
};

export type Subscription = {
  __typename?: 'Subscription';
  notifications: Notification;
};


export type SubscriptionNotificationsArgs = {
  id: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
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

export type CreateEnterpriseMutationVariables = Exact<{
  input: CreateEnterprise;
}>;


export type CreateEnterpriseMutation = { __typename?: 'Mutation', createEnterprise: { __typename?: 'Enterprise', name: string } };

export type CreateUserMutationMutationVariables = Exact<{
  input: CreateUser;
}>;


export type CreateUserMutationMutation = { __typename?: 'Mutation', createUser: { __typename?: 'User', userName: string } };

export type CitiesQueryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CitiesQueryQuery = { __typename?: 'Query', cities: Array<{ __typename?: 'City', id: number, name: string, countryCode: string, latitude: number, longitude: number }> };

export type ComplaintsInfoQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ComplaintsInfoQueryQuery = { __typename?: 'Query', complaintsReceivedInfo: { __typename?: 'ComplaintInfo', received: number, resolved: number, reviewed: number, pending: number, avgRating: number, total: number } };

export type CountriesQueryVariables = Exact<{ [key: string]: never; }>;


export type CountriesQuery = { __typename?: 'Query', countries: Array<{ __typename?: 'Country', id: number, name: string, phoneCode: string }> };

export type CountryStatesQueryQueryVariables = Exact<{
  id: Scalars['Int']['input'];
}>;


export type CountryStatesQueryQuery = { __typename?: 'Query', countryStates: Array<{ __typename?: 'CountryState', id: number, name: string }> };

export type EnterpriseByIdQueryQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type EnterpriseByIdQueryQuery = { __typename?: 'Query', enterpriseById: { __typename?: 'Enterprise', name: string, logoImg: string, bannerImg: string, website: string, email: string, phoneNumber: string, industry: string, foundationDate: string, ownerId: string, address: { __typename?: 'Address', country: string, countryState: string, city: string }, employees: Array<{ __typename?: 'Employee', id: string, enterpriseId: string, userId: string, hiringDate: string, approvedHiring: boolean, approvedHiringAt: string, enterprisePosition: string, user: { __typename?: 'User', userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } } } | null> } };

export type EnterprisesByAuthenticatedUserQueryQueryVariables = Exact<{ [key: string]: never; }>;


export type EnterprisesByAuthenticatedUserQueryQuery = { __typename?: 'Query', enterprisesByAuthenticatedUser: { __typename?: 'EnterprisesByAuthenticatedUserResult', enterprises: Array<{ __typename?: 'EnterpriseByAuthenticatedUser', authority: { __typename?: 'GrantedAuthority', authority: string, enterpriseId: string }, enterprise?: { __typename?: 'Enterprise', name: string, logoImg: string, bannerImg: string, website: string, email: string, phoneNumber: string, industry: string, foundationDate: string, ownerId: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } | null }>, offices: Array<{ __typename?: 'EnterpriseByAuthenticatedUser', authority: { __typename?: 'GrantedAuthority', authority: string, enterpriseId: string }, enterprise?: { __typename?: 'Enterprise', name: string, logoImg: string, bannerImg: string, website: string, email: string, phoneNumber: string, industry: string, foundationDate: string, ownerId: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } | null }> } };

export type HiringInvitationsByAuthenticatedUserQueryQueryVariables = Exact<{ [key: string]: never; }>;


export type HiringInvitationsByAuthenticatedUserQueryQuery = { __typename?: 'Query', hiringInvitationsByAuthenticatedUser: Array<{ __typename?: 'HiringInvitation', eventId: string, enterpriseId: string, proposedPosition: string, ownerId: string, fullName: string, enterpriseEmail: string, enterprisePhone: string, enterpriseLogoImg: string, occurredOn: string, seen: boolean, status: HiringProccessState, reason: string }> };

export type IndustriesQueryVariables = Exact<{ [key: string]: never; }>;


export type IndustriesQuery = { __typename?: 'Query', industries: Array<{ __typename?: 'Industry', id: number, name: string }> };

export type UserQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type UserQuery = { __typename?: 'Query', userById: { __typename?: 'User', userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } } };

export type UserDescriptorQueryVariables = Exact<{ [key: string]: never; }>;


export type UserDescriptorQuery = { __typename?: 'Query', userDescriptor: { __typename?: 'UserDescriptor', userName: string, fullName: string, profileImg: string, genre: string, pronoun: string, authorities?: Array<{ __typename?: 'GrantedAuthority', enterpriseId: string, authority: string } | null> | null } };

export type UsersForHiringQueryQueryVariables = Exact<{
  input: SearchWithPagination;
}>;


export type UsersForHiringQueryQuery = { __typename?: 'Query', usersForHiring: { __typename?: 'UsersForHiringResult', count: number, limit: number, offset: number, nextCursor: number, prevCursor: number, users: Array<{ __typename?: 'User', userName: string, status: UserStatus, person: { __typename?: 'Person', profileImg: string, email: string, firstName: string, lastName: string, genre: string, pronoun: string, age: number, phoneNumber: string, address: { __typename?: 'Address', country: string, countryState: string, city: string } } }> } };

export const CountryFragmentFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"CountryFragment"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Country"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]} as unknown as DocumentNode<CountryFragmentFragment, unknown>;
export const CreateEnterpriseDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createEnterprise"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateEnterprise"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createEnterprise"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CreateEnterpriseMutation, CreateEnterpriseMutationVariables>;
export const CreateUserMutationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUserMutation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateUser"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}}]}}]}}]} as unknown as DocumentNode<CreateUserMutationMutation, CreateUserMutationMutationVariables>;
export const CitiesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CitiesQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cities"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"countryCode"}},{"kind":"Field","name":{"kind":"Name","value":"latitude"}},{"kind":"Field","name":{"kind":"Name","value":"longitude"}}]}}]}}]} as unknown as DocumentNode<CitiesQueryQuery, CitiesQueryQueryVariables>;
export const ComplaintsInfoQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"complaintsInfoQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"complaintsReceivedInfo"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"received"}},{"kind":"Field","name":{"kind":"Name","value":"resolved"}},{"kind":"Field","name":{"kind":"Name","value":"reviewed"}},{"kind":"Field","name":{"kind":"Name","value":"pending"}},{"kind":"Field","name":{"kind":"Name","value":"avgRating"}},{"kind":"Field","name":{"kind":"Name","value":"total"}}]}}]}}]} as unknown as DocumentNode<ComplaintsInfoQueryQuery, ComplaintsInfoQueryQueryVariables>;
export const CountriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Countries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"countries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"phoneCode"}}]}}]}}]} as unknown as DocumentNode<CountriesQuery, CountriesQueryVariables>;
export const CountryStatesQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CountryStatesQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"countryStates"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CountryStatesQueryQuery, CountryStatesQueryQueryVariables>;
export const EnterpriseByIdQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"enterpriseByIdQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterpriseById"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoImg"}},{"kind":"Field","name":{"kind":"Name","value":"bannerImg"}},{"kind":"Field","name":{"kind":"Name","value":"website"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"}},{"kind":"Field","name":{"kind":"Name","value":"foundationDate"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}},{"kind":"Field","name":{"kind":"Name","value":"employees"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"userId"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"hiringDate"}},{"kind":"Field","name":{"kind":"Name","value":"approvedHiring"}},{"kind":"Field","name":{"kind":"Name","value":"approvedHiringAt"}},{"kind":"Field","name":{"kind":"Name","value":"enterprisePosition"}}]}}]}}]}}]} as unknown as DocumentNode<EnterpriseByIdQueryQuery, EnterpriseByIdQueryQueryVariables>;
export const EnterprisesByAuthenticatedUserQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"EnterprisesByAuthenticatedUserQuery"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterprisesByAuthenticatedUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterprises"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}}]}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoImg"}},{"kind":"Field","name":{"kind":"Name","value":"bannerImg"}},{"kind":"Field","name":{"kind":"Name","value":"website"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"}},{"kind":"Field","name":{"kind":"Name","value":"foundationDate"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"offices"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authority"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}}]}},{"kind":"Field","name":{"kind":"Name","value":"enterprise"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"logoImg"}},{"kind":"Field","name":{"kind":"Name","value":"bannerImg"}},{"kind":"Field","name":{"kind":"Name","value":"website"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}},{"kind":"Field","name":{"kind":"Name","value":"industry"}},{"kind":"Field","name":{"kind":"Name","value":"foundationDate"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}}]}}]}}]}}]}}]} as unknown as DocumentNode<EnterprisesByAuthenticatedUserQueryQuery, EnterprisesByAuthenticatedUserQueryQueryVariables>;
export const HiringInvitationsByAuthenticatedUserQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"hiringInvitationsByAuthenticatedUserQuery"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hiringInvitationsByAuthenticatedUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"eventId"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"proposedPosition"}},{"kind":"Field","name":{"kind":"Name","value":"ownerId"}},{"kind":"Field","name":{"kind":"Name","value":"fullName"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseEmail"}},{"kind":"Field","name":{"kind":"Name","value":"enterprisePhone"}},{"kind":"Field","name":{"kind":"Name","value":"enterpriseLogoImg"}},{"kind":"Field","name":{"kind":"Name","value":"occurredOn"}},{"kind":"Field","name":{"kind":"Name","value":"seen"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"reason"}}]}}]}}]} as unknown as DocumentNode<HiringInvitationsByAuthenticatedUserQueryQuery, HiringInvitationsByAuthenticatedUserQueryQueryVariables>;
export const IndustriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"industries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"industries"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<IndustriesQuery, IndustriesQueryVariables>;
export const UserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"User"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userById"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<UserQuery, UserQueryVariables>;
export const UserDescriptorDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"UserDescriptor"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userDescriptor"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"fullName"}},{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"authorities"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"enterpriseId"}},{"kind":"Field","name":{"kind":"Name","value":"authority"}}]}}]}}]}}]} as unknown as DocumentNode<UserDescriptorQuery, UserDescriptorQueryVariables>;
export const UsersForHiringQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"usersForHiringQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"SearchWithPagination"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"usersForHiring"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"users"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userName"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"profileImg"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"genre"}},{"kind":"Field","name":{"kind":"Name","value":"pronoun"}},{"kind":"Field","name":{"kind":"Name","value":"age"}},{"kind":"Field","name":{"kind":"Name","value":"phoneNumber"}},{"kind":"Field","name":{"kind":"Name","value":"address"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"countryState"}},{"kind":"Field","name":{"kind":"Name","value":"city"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"count"}},{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"nextCursor"}},{"kind":"Field","name":{"kind":"Name","value":"prevCursor"}}]}}]}}]} as unknown as DocumentNode<UsersForHiringQueryQuery, UsersForHiringQueryQueryVariables>;
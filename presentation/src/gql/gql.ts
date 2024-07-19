/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  fragment CountryFragment on Country {\n    id\n    name\n  }\n": types.CountryFragmentFragmentDoc,
    "\n    mutation createEnterprise($input: CreateEnterprise!){\n        createEnterprise(input: $input){\n            name\n        }\n    }": types.CreateEnterpriseDocument,
    "\n    mutation CreateUserMutation(\n        $input: CreateUser!){\n            createUser(\n                input: $input\n            ){\n                userName\n            }\n    }": types.CreateUserMutationDocument,
    "\n    query CitiesQuery($id: Int!){\n        cities(id:$id){\n            id\n            name\n            countryCode\n            latitude\n            longitude\n        }\n    }": types.CitiesQueryDocument,
    "\n    query complaintsInfoQuery($id:String!){\n        complaintsReceivedInfo(id:$id){\n            received\n            resolved\n            reviewed\n            pending\n            avgRating\n            total\n        }\n    }": types.ComplaintsInfoQueryDocument,
    "\n    query Countries{\n        countries{\n            id\n            name\n            phoneCode\n        }\n    }": types.CountriesDocument,
    "\n    query CountryStatesQuery($id: Int!){\n        countryStates(id: $id) {\n            id\n            name\n        }\n    }": types.CountryStatesQueryDocument,
    "\n    query enterpriseByIdQuery($id: String!){\n        enterpriseById(id:$id){\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n            employees {\n                id\n                enterpriseId\n                userId\n                user{\n                    userName\n                    person{\n                        profileImg\n                        email\n                        firstName\n                        lastName\n                        genre\n                        pronoun\n                        age\n                        phoneNumber\n                        address{\n                            country\n                            countryState\n                            city\n                        }\n                    }\n                    status\n                }\n                hiringDate\n                approvedHiring\n                approvedHiringAt\n                enterprisePosition\n            }\n        }\n    }": types.EnterpriseByIdQueryDocument,
    "\nquery EnterprisesByAuthenticatedUserQuery{\n    enterprisesByAuthenticatedUser{\n        enterprises{\n            authority{\n            authority\n            enterpriseId\n        }\n        enterprise{\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n        }\n        }\n        offices{\n            authority{\n            authority\n            enterpriseId\n        }\n        enterprise{\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n        }\n        }\n    }\n}\n": types.EnterprisesByAuthenticatedUserQueryDocument,
    "\n    query hiringInvitationsByAuthenticatedUserQuery{\n        hiringInvitationsByAuthenticatedUser{\n            eventId\n            enterpriseId\n            proposedPosition\n            ownerId\n            fullName\n            enterpriseEmail\n            enterprisePhone\n            enterpriseLogoImg\n            occurredOn\n            seen\n            status\n            reason\n        }\n    }": types.HiringInvitationsByAuthenticatedUserQueryDocument,
    "query industries{\n        industries{\n            id\n            name\n        }\n    }": types.IndustriesDocument,
    "\n    query User($id:String!){\n        userById(id:$id){\n                userName\n                person {\n                    profileImg\n                    email\n                    firstName\n                    lastName\n                    genre\n                    pronoun\n                    age\n                    phoneNumber\n                    address { \n                        country\n                        countryState\n                        city\n                        }\n                }\n                status\n        }\n    }": types.UserDocument,
    "\n    query UserDescriptor{\n        userDescriptor{\n            userName\n            fullName\n            profileImg\n            genre\n            pronoun\n            authorities{\n                enterpriseId\n                authority\n            }\n        }\n    }": types.UserDescriptorDocument,
    "\n    query usersForHiringQuery($input: SearchWithPagination!) {\n        usersForHiring(input:$input){\n            users {\n                userName\n                person {\n                    profileImg\n                    email\n                    firstName\n                    lastName\n                    genre\n                    pronoun\n                    age\n                    phoneNumber\n                    address { country countryState city}\n                }\n                status\n            }\n            count\n            limit\n            offset\n            nextCursor\n            prevCursor\n        }\n    }": types.UsersForHiringQueryDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment CountryFragment on Country {\n    id\n    name\n  }\n"): (typeof documents)["\n  fragment CountryFragment on Country {\n    id\n    name\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createEnterprise($input: CreateEnterprise!){\n        createEnterprise(input: $input){\n            name\n        }\n    }"): (typeof documents)["\n    mutation createEnterprise($input: CreateEnterprise!){\n        createEnterprise(input: $input){\n            name\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation CreateUserMutation(\n        $input: CreateUser!){\n            createUser(\n                input: $input\n            ){\n                userName\n            }\n    }"): (typeof documents)["\n    mutation CreateUserMutation(\n        $input: CreateUser!){\n            createUser(\n                input: $input\n            ){\n                userName\n            }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query CitiesQuery($id: Int!){\n        cities(id:$id){\n            id\n            name\n            countryCode\n            latitude\n            longitude\n        }\n    }"): (typeof documents)["\n    query CitiesQuery($id: Int!){\n        cities(id:$id){\n            id\n            name\n            countryCode\n            latitude\n            longitude\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query complaintsInfoQuery($id:String!){\n        complaintsReceivedInfo(id:$id){\n            received\n            resolved\n            reviewed\n            pending\n            avgRating\n            total\n        }\n    }"): (typeof documents)["\n    query complaintsInfoQuery($id:String!){\n        complaintsReceivedInfo(id:$id){\n            received\n            resolved\n            reviewed\n            pending\n            avgRating\n            total\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query Countries{\n        countries{\n            id\n            name\n            phoneCode\n        }\n    }"): (typeof documents)["\n    query Countries{\n        countries{\n            id\n            name\n            phoneCode\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query CountryStatesQuery($id: Int!){\n        countryStates(id: $id) {\n            id\n            name\n        }\n    }"): (typeof documents)["\n    query CountryStatesQuery($id: Int!){\n        countryStates(id: $id) {\n            id\n            name\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query enterpriseByIdQuery($id: String!){\n        enterpriseById(id:$id){\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n            employees {\n                id\n                enterpriseId\n                userId\n                user{\n                    userName\n                    person{\n                        profileImg\n                        email\n                        firstName\n                        lastName\n                        genre\n                        pronoun\n                        age\n                        phoneNumber\n                        address{\n                            country\n                            countryState\n                            city\n                        }\n                    }\n                    status\n                }\n                hiringDate\n                approvedHiring\n                approvedHiringAt\n                enterprisePosition\n            }\n        }\n    }"): (typeof documents)["\n    query enterpriseByIdQuery($id: String!){\n        enterpriseById(id:$id){\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n            employees {\n                id\n                enterpriseId\n                userId\n                user{\n                    userName\n                    person{\n                        profileImg\n                        email\n                        firstName\n                        lastName\n                        genre\n                        pronoun\n                        age\n                        phoneNumber\n                        address{\n                            country\n                            countryState\n                            city\n                        }\n                    }\n                    status\n                }\n                hiringDate\n                approvedHiring\n                approvedHiringAt\n                enterprisePosition\n            }\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\nquery EnterprisesByAuthenticatedUserQuery{\n    enterprisesByAuthenticatedUser{\n        enterprises{\n            authority{\n            authority\n            enterpriseId\n        }\n        enterprise{\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n        }\n        }\n        offices{\n            authority{\n            authority\n            enterpriseId\n        }\n        enterprise{\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n        }\n        }\n    }\n}\n"): (typeof documents)["\nquery EnterprisesByAuthenticatedUserQuery{\n    enterprisesByAuthenticatedUser{\n        enterprises{\n            authority{\n            authority\n            enterpriseId\n        }\n        enterprise{\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n        }\n        }\n        offices{\n            authority{\n            authority\n            enterpriseId\n        }\n        enterprise{\n            name\n            logoImg\n            bannerImg\n            website\n            email\n            phoneNumber\n            address {\n                country\n                countryState\n                city\n                }\n            industry\n            foundationDate\n            ownerId\n        }\n        }\n    }\n}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query hiringInvitationsByAuthenticatedUserQuery{\n        hiringInvitationsByAuthenticatedUser{\n            eventId\n            enterpriseId\n            proposedPosition\n            ownerId\n            fullName\n            enterpriseEmail\n            enterprisePhone\n            enterpriseLogoImg\n            occurredOn\n            seen\n            status\n            reason\n        }\n    }"): (typeof documents)["\n    query hiringInvitationsByAuthenticatedUserQuery{\n        hiringInvitationsByAuthenticatedUser{\n            eventId\n            enterpriseId\n            proposedPosition\n            ownerId\n            fullName\n            enterpriseEmail\n            enterprisePhone\n            enterpriseLogoImg\n            occurredOn\n            seen\n            status\n            reason\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "query industries{\n        industries{\n            id\n            name\n        }\n    }"): (typeof documents)["query industries{\n        industries{\n            id\n            name\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query User($id:String!){\n        userById(id:$id){\n                userName\n                person {\n                    profileImg\n                    email\n                    firstName\n                    lastName\n                    genre\n                    pronoun\n                    age\n                    phoneNumber\n                    address { \n                        country\n                        countryState\n                        city\n                        }\n                }\n                status\n        }\n    }"): (typeof documents)["\n    query User($id:String!){\n        userById(id:$id){\n                userName\n                person {\n                    profileImg\n                    email\n                    firstName\n                    lastName\n                    genre\n                    pronoun\n                    age\n                    phoneNumber\n                    address { \n                        country\n                        countryState\n                        city\n                        }\n                }\n                status\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query UserDescriptor{\n        userDescriptor{\n            userName\n            fullName\n            profileImg\n            genre\n            pronoun\n            authorities{\n                enterpriseId\n                authority\n            }\n        }\n    }"): (typeof documents)["\n    query UserDescriptor{\n        userDescriptor{\n            userName\n            fullName\n            profileImg\n            genre\n            pronoun\n            authorities{\n                enterpriseId\n                authority\n            }\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query usersForHiringQuery($input: SearchWithPagination!) {\n        usersForHiring(input:$input){\n            users {\n                userName\n                person {\n                    profileImg\n                    email\n                    firstName\n                    lastName\n                    genre\n                    pronoun\n                    age\n                    phoneNumber\n                    address { country countryState city}\n                }\n                status\n            }\n            count\n            limit\n            offset\n            nextCursor\n            prevCursor\n        }\n    }"): (typeof documents)["\n    query usersForHiringQuery($input: SearchWithPagination!) {\n        usersForHiring(input:$input){\n            users {\n                userName\n                person {\n                    profileImg\n                    email\n                    firstName\n                    lastName\n                    genre\n                    pronoun\n                    age\n                    phoneNumber\n                    address { country countryState city}\n                }\n                status\n            }\n            count\n            limit\n            offset\n            nextCursor\n            prevCursor\n        }\n    }"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
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
    "\n    mutation CreateUserMutation(\n        $email:String!,\n        $password:String!,\n        $firstName:String!,\n        $lastName:String!,\n        $gender:String!,\n        $pronoun:String!,\n        $birthDate:String!,\n        $phone:String!,\n        $countryId:Int!,\n        $countryStateId:Int!,\n        $cityId:Int!){\n        CreateUser(\n            email: $email,\n            password: $password,\n            firstName: $firstName,\n            lastName: $lastName,\n            gender: $gender,\n            pronoun: $pronoun,\n            birthDate: $birthDate,\n            phone: $phone,\n            countryId: $countryId,\n            countryStateId: $countryStateId,\n            cityId: $cityId\n        )\n    }": types.CreateUserMutationDocument,
    "\n    query CitiesQuery($id: Int!){\n        Cities(id:$id){\n            id\n            name\n            countryCode\n            latitude\n            longitude\n        }\n    }": types.CitiesQueryDocument,
    "\n    query Countries{\n        Countries{\n            id\n            name\n            phoneCode\n        }\n    }": types.CountriesDocument,
    "\n    query CountryStatesQuery($id: Int!){\n        CountryStates(id: $id) {\n            id\n            name\n        }\n    }": types.CountryStatesQueryDocument,
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
export function graphql(source: "\n    mutation CreateUserMutation(\n        $email:String!,\n        $password:String!,\n        $firstName:String!,\n        $lastName:String!,\n        $gender:String!,\n        $pronoun:String!,\n        $birthDate:String!,\n        $phone:String!,\n        $countryId:Int!,\n        $countryStateId:Int!,\n        $cityId:Int!){\n        CreateUser(\n            email: $email,\n            password: $password,\n            firstName: $firstName,\n            lastName: $lastName,\n            gender: $gender,\n            pronoun: $pronoun,\n            birthDate: $birthDate,\n            phone: $phone,\n            countryId: $countryId,\n            countryStateId: $countryStateId,\n            cityId: $cityId\n        )\n    }"): (typeof documents)["\n    mutation CreateUserMutation(\n        $email:String!,\n        $password:String!,\n        $firstName:String!,\n        $lastName:String!,\n        $gender:String!,\n        $pronoun:String!,\n        $birthDate:String!,\n        $phone:String!,\n        $countryId:Int!,\n        $countryStateId:Int!,\n        $cityId:Int!){\n        CreateUser(\n            email: $email,\n            password: $password,\n            firstName: $firstName,\n            lastName: $lastName,\n            gender: $gender,\n            pronoun: $pronoun,\n            birthDate: $birthDate,\n            phone: $phone,\n            countryId: $countryId,\n            countryStateId: $countryStateId,\n            cityId: $cityId\n        )\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query CitiesQuery($id: Int!){\n        Cities(id:$id){\n            id\n            name\n            countryCode\n            latitude\n            longitude\n        }\n    }"): (typeof documents)["\n    query CitiesQuery($id: Int!){\n        Cities(id:$id){\n            id\n            name\n            countryCode\n            latitude\n            longitude\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query Countries{\n        Countries{\n            id\n            name\n            phoneCode\n        }\n    }"): (typeof documents)["\n    query Countries{\n        Countries{\n            id\n            name\n            phoneCode\n        }\n    }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query CountryStatesQuery($id: Int!){\n        CountryStates(id: $id) {\n            id\n            name\n        }\n    }"): (typeof documents)["\n    query CountryStatesQuery($id: Int!){\n        CountryStates(id: $id) {\n            id\n            name\n        }\n    }"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
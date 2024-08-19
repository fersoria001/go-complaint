import { GraphQLClient } from "graphql-request";

const getGraphQLClient = () => {
  const endpoint = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
  if (!endpoint) {
    throw new Error("endpoint for getGraphQLClient is undefined");
  }
  return new GraphQLClient(endpoint, {
    method: "POST",
    credentials: "include",
    errorPolicy: "none",
  });
};

export default getGraphQLClient;

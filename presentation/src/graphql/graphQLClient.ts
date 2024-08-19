import { GraphQLClient } from "graphql-request";

const getGraphQLClient = () => {
  const endpoint = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
  if (!endpoint) {
    throw new Error("endpoint for getGraphQLClient is undefined");
  }
  
  const apiKey = process.env.NEXT_PUBLIC_API_KEY;
  if (!apiKey) {
    throw new Error("api key axios instance not defined in process env");
  }
  const headers = new Headers();
  headers.append("api_key", apiKey);
  return new GraphQLClient(endpoint, {
    method: "POST",
    credentials: "include",
    errorPolicy: "none",
    headers: headers,
  });
};

export default getGraphQLClient;

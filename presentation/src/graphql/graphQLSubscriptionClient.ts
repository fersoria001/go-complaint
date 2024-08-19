import { Client, createClient } from "graphql-ws";
let url = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
if (!url) {
  throw new Error("url is not defined in useSubscription hook");
}
const apiKey = process.env.NEXT_PUBLIC_API_KEY;
if (!apiKey) {
  throw new Error("api key axios instance not defined in process env");
}
url = url.replace("http://", "ws://");
url += `?api_key=${apiKey}`;
const graphQLSubscriptionClient: Client | undefined = undefined;
const getGraphQLSubscriptionClient = () =>
  graphQLSubscriptionClient
    ? graphQLSubscriptionClient
    : createClient({
        url,
      });
export default getGraphQLSubscriptionClient;

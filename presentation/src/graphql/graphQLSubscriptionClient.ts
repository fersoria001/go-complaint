import { Client, createClient } from "graphql-ws";
let url = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
if (!url) {
  throw new Error("url is not defined in useSubscription hook");
}
url = url.replace("http://", "ws://");
const graphQLSubscriptionClient: Client | undefined = undefined;
const getGraphQLSubscriptionClient = () =>
  graphQLSubscriptionClient
    ? graphQLSubscriptionClient
    : createClient({
        url,
      });
export default getGraphQLSubscriptionClient;

import { createClient } from "graphql-ws";
let url = process.env.NEXT_PUBLIC_GRAPHQL_ENDPOINT;
if (!url) {
  throw new Error("url is not defined in useSubscription hook");
}
url = url.replace("http://", "ws://");
const graphQLSubscriptionClient = createClient({
  url,
});

export default graphQLSubscriptionClient;
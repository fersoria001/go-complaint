import { useEffect, useState } from "react";
import { ConnectionACK, Subscription } from "../types";

function useSubscription<T>(subscription: Subscription<T> | null) {
  const [incomingMessage, setIncomingMessage] = useState<T | null>(null);

  useEffect(() => {
    if (!subscription) {
      return;
    }
    const ws = open(subscription.connection_ack);
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === "data") {
        setIncomingMessage(subscription.subscriptionReturnType(data.payload));
      }
    };
    return () => close(ws);
  }, [subscription]);

  return { incomingMessage, open, close };
}
const open = (connection_ack: ConnectionACK) => {
  const ws = new WebSocket("ws://docker-go-complaint-server-latest.onrender.com/subscriptions");
  ws.onopen = () => {
    ws.send(JSON.stringify(connection_ack));
  };
  return ws;
};
const close = (ws: WebSocket) => {
  if (ws) {
    ws.close();
  }
};
export default useSubscription;

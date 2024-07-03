/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import Cookies from "js-cookie";
function useSubscriber(id: string | null) {
  const [incomingMessage, setIncomingMessage] = useState<any | null>(null);
  useEffect(() => {
    if (!id || id === "") {
      return;
    }
    const authorization = Cookies.get("Authorization");
    if (!authorization) {
      return;
    }
    const connection_ack = {
      type: "connection_ack",
      payload: {
        query: `any,not used`,
        subscription_id: id,
        token: authorization.slice(7),
      },
    };
    const ws = new WebSocket(
      "ws://go-complaint-server-latest.onrender.com/graphql"
    );
    ws.onopen = () => {
      console.log("id", id)
      ws.send(JSON.stringify(connection_ack));
    };
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      console.log("wsdata",data)
      if (data.type === "data") {
        setIncomingMessage(data.payload);
      }
    };
    return () => {
      return ws.close();
    };
  }, [id]);

  return { incomingMessage };
}

export default useSubscriber;

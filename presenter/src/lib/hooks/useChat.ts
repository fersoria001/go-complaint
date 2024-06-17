import { useState } from "react";
import { AuthMsg, newReply, Reply, Sender } from "../types";
import { useRevalidator } from "react-router-dom";

const useChat = (id: string, authMsg: AuthMsg) => {
  const [webSocket, setWebSocket] = useState<WebSocket | null>(null);
  const [webSocketReady, setWebSocketReady] = useState(false); // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [incomingMessages, setIncomingMessages] = useState<Reply[]>([]);
  const [closed, setClosed] = useState<boolean>(false);
  const revalidator = useRevalidator();
  const open = () => {
    if (webSocket) {
      console.error("trying to open an already open WebSocket");
      return;
    }
    const ws = new WebSocket(`ws://localhost:8080/chat?room=${id}`);
    setWebSocket(ws);
    ws.onopen = () => {
      ws.send(JSON.stringify(authMsg));
    };
    ws.onmessage = (event) => {
      const parsed = JSON.parse(event.data);
      switch (parsed.content) {
        case "auth":
          if (!parsed.success) {
            console.error("Failed to authenticate WebSocket");
            ws.close();
            return;
          }
          setWebSocketReady(true);
          break;
        case "reply":
          console.log("received message", parsed.reply.body);
          if (parsed.reply.body === "status_change:review_default_message") {
            console.log("the room has been closed remotely");
            revalidator.revalidate();
            setClosed(true);
          }
          setIncomingMessages((prevMessages) => [
            ...prevMessages,
            parsed.reply,
          ]);
          break;
        default:
          break;
      }
    };
  };

  const close = () => {
    if (webSocket) {
      webSocket.close();
      setWebSocketReady(false);
    }
  };

  const send = (input: string, sender: Sender) => {
    if (!webSocketReady || !webSocket) {
      console.error(
        "trying to send to a closed WebSocket",
        webSocketReady,
        webSocket
      );

      return;
    }
    if (incomingMessages.length === 2) {
      revalidator.revalidate();
    }
    const message = newReply(sender.thumbnail, sender.fullName, input, sender.isEnterprise, sender.enterpriseID);
    try {
      webSocket.send(JSON.stringify({ content: "reply", reply: message }));
    } catch (error) {
      console.error("Failed to send message:", error);
    }
  };

  const closeRoom = (sender: Sender) => {
    if (!webSocketReady || !webSocket) {
      console.error("trying to send to a closed WebSocket");
      return;
    }
    const messageText = "status_change:review_default_message";
    const message = newReply(sender.thumbnail, sender.fullName, messageText, sender.isEnterprise, sender.enterpriseID);
    try {
      webSocket.send(
        JSON.stringify({
          content: "reply",
          reply: message,
          status_changed: true,
        })
      );
    } catch (error) {
      console.error("Failed to send message:", error);
    }
  };

  return {
    closed,
    webSocketReady,
    incomingMessages,
    send,
    open,
    close,
    closeRoom,
  };
};

export default useChat;

import { useEffect, useState } from "react";

export enum ChatSubProtocols {
  COMPLAINT = "complaint",
  ENTERPRISE_CHAT = "enterpriseChat",
}
export enum ChatMessageType {
  ConnectionInit = "connection_init",
  ConnectionAcknowledged = "connection_ack",
  Data = "data",
  Complete = "complete",
  UserOnline = "user_online",
  UserOffline = "user_offline",
}
export type ChatMessage = {
  type: string;
  payload: any;
};

function useChat(id: string, subProtocol: ChatSubProtocols, jwt: string) {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [isReady, setIsReady] = useState<boolean>(false);
  const [incomingMsg, setIncomingMsg] = useState<any | null>(null);
  useEffect(() => {
    const url = process.env.NEXT_PUBLIC_CHAT_ENDPOINT;
    if (!url) {
      console.error("ws url is undefined");
      return;
    }
    const apiKey = process.env.NEXT_PUBLIC_API_KEY;
    if (!apiKey) {
      throw new Error("api key axios instance not defined in process env");
    }
    const websocket = new WebSocket(url + `?id=${id}&api_key=${apiKey}`, subProtocol);

    websocket.onopen = () => {
      const msg: ChatMessage = {
        type: "connection_init",
        payload: encodeToBinary(jwt),
      };
      console.log("onOpenSend", msg);
      websocket.send(JSON.stringify(msg));
    };
    websocket.onerror = (e: any) => {
      console.log("ws error", e);
    };
    websocket.onmessage = (event: any) => {
      const jsonMsg = JSON.parse(event.data) as ChatMessage;
      console.log("onMessage", jsonMsg);
      switch (jsonMsg.type) {
        case ChatMessageType.ConnectionAcknowledged: {
          const decodedPayload = decodeFromBinary(jsonMsg.payload);
          console.log("authResponse", decodedPayload);
          if (decodedPayload === "true") {
            setIsReady(true);
            setWs(websocket);
          }
          break;
        }
        case ChatMessageType.Data: {
          const p = JSON.parse(decodeFromBinary(jsonMsg.payload));
          setIncomingMsg(p);
          break;
        }
        case ChatMessageType.Complete: {
          console.log("type complete not implemented");
          break;
        }
        case ChatMessageType.UserOffline: {
          const newMsg = {
            subProtocolDataType: ChatMessageType.UserOffline,
            result: decodeFromBinary(jsonMsg.payload),
          };
          setIncomingMsg(newMsg);
          break;
        }
        case ChatMessageType.UserOnline: {
          const newMsg = {
            subProtocolDataType: ChatMessageType.UserOnline,
            result: decodeFromBinary(jsonMsg.payload),
          };
          setIncomingMsg(newMsg);
          break;
        }
        default: {
          console.log(
            "type",
            jsonMsg.type,
            "default case not implemented",
            decodeFromBinary(jsonMsg.payload)
          );
          break;
        }
      }
    };
    return () => {
      websocket.close(1000);
    };
  }, [id, jwt, subProtocol]);
  function send(payload: Object) {
    const msg: ChatMessage = {
      type: ChatMessageType.Data,
      payload: encodeToBinary(JSON.stringify(payload)),
    };
    if (ws && isReady) {
      ws.send(JSON.stringify(msg));
    }
  }
  return { isReady, incomingMsg, send };
}

export default useChat;

export function decodeFromBinary(str: string): string {
  return decodeURIComponent(
    Array.prototype.map
      .call(atob(str), function (c) {
        return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
      })
      .join("")
  );
}
export function encodeToBinary(str: string): string {
  return btoa(
    encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, function (match, p1) {
      return String.fromCharCode(parseInt(p1, 16));
    })
  );
}

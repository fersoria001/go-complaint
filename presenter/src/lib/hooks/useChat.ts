import { useState, useEffect } from 'react';
import { AuthMsg, Reply } from '../types';

const useChat = (id: string, authMsg : AuthMsg, tab : string) => {
    const [webSocket, setWebSocket] = useState<WebSocket | null>(null);
    const [webSocketReady, setWebSocketReady] = useState(false);
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [incomingMessages, setIncomingMessages] = useState<Reply[]>([]);
    const [sortedMessages, setSortedMessages] = useState<Reply[]>([]);

    useEffect(() => {
        if (tab !== "chat") {
            return;
        }
        const ws = new WebSocket(`ws://localhost:8080/chat?room=${id}`);
        setWebSocket(ws);

        ws.onopen = () => {
            console.log(JSON.stringify(authMsg));
            ws.send(JSON.stringify(authMsg));
        };

        ws.onmessage = (event) => {
            const parsed = JSON.parse(event.data);
            switch (parsed.content) {
                case "auth": {
                    if (!parsed.success) {
                        throw new Error("Auth failed");
                    }
                    setWebSocketReady(true);
                    break;
                }
                case "reply": {
                    setIncomingMessages(prevMessages  => {
                        const newMessages = [...prevMessages, parsed.reply];
                        const sortedNewMessages = newMessages.sort((a, b) => parseInt(a.created_at) - parseInt(b.created_at));
                        setSortedMessages(sortedNewMessages);
                        return newMessages;
                    });
                    break;
                }
                default:
                    break;
            }
        };
        return () => {
            ws.close();
        };
    }, [id, authMsg, tab]);

    return { webSocket, webSocketReady, sortedMessages };
};

export default useChat;
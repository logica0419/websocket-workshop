import ReconnectingWebSocket from "reconnecting-websocket";

export const useWebSocket = (
  url: string,
  onMessage: (e: MessageEvent<string>) => void
) => {
  const RWS = new ReconnectingWebSocket(url);

  RWS.onopen = () => {
    alert("WebSocketに接続しました");
  };
  RWS.onmessage = onMessage;

  return RWS;
};

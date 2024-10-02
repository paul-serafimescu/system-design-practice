import { useEffect, useState } from "react";
import LoginForm from "./LoginForm";

function App() {
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    if (!ws) return;

    ws.onopen = () => console.log("connected to websocket");

    ws.onerror = console.error;

    ws.onclose = () => console.log("disconnected from websocket");

    ws.onmessage = event => console.log(event.data);

    return () => {
      ws.onopen = null;
      ws.onerror = null;
      ws.onclose = null;
      ws.onmessage = null;
    };
  }, [ws]);  

  return (
    <LoginForm setWs={setWs} />
  );
}

export default App;

import { useEffect, useState } from "react";
import LoginForm from "./LoginForm";
import { Route, Routes } from "react-router-dom";
import { LandingPage } from "./pages";
import Navbar from "./components/Navbar";
import { createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import DirectMessagesPage from "./pages/DirectMessagesPage";

function App() {
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    if (!ws) return;

    ws.onopen = () => console.log("connected to websocket");

    ws.onerror = console.error;

    ws.onclose = () => console.log("disconnected from websocket");

    ws.onmessage = (event) => console.log(event.data);

    return () => {
      ws.onopen = null;
      ws.onerror = null;
      ws.onclose = null;
      ws.onmessage = null;

      ws.close();
    };
  }, [ws]);

  const [darkMode, setDarkMode] = useState(false);

  const theme = createTheme({
    palette: {
      mode: darkMode ? "dark" : "light",
    },
  });

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      {/* Render Navbar only for specific routes */}
      {location.pathname !== "/login" &&
        location.pathname !== "/channels/@me" && (
          <Navbar darkMode={darkMode} setDarkMode={setDarkMode} />
        )}
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/login" element={<LoginForm setWs={setWs} />} />
        <Route path="/channels/@me" element={<DirectMessagesPage />} />
      </Routes>
    </ThemeProvider>
  );
}

export default App;

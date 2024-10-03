import React, { useState } from "react";
import axios from "axios";
import {
  Box,
  Button,
  Container,
  TextField,
  Typography,
  Alert,
  CssBaseline,
  ThemeProvider,
  createTheme,
} from "@mui/material";
import SHA256 from "crypto-js/sha256";
import { useNavigate } from "react-router-dom";

const theme = createTheme();

interface IProps {
  setWs: (ws: WebSocket) => void;
}

const LoginForm: React.FC<IProps> = ({ setWs }) => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [username, setUsername] = useState<string>(""); // New state for username
  const [firstname, setFirstname] = useState<string>(""); // New state for firstname
  const [lastname, setLastname] = useState<string>(""); // New state for lastname
  const [error, setError] = useState<string | null>(null);
  const [isSignup, setIsSignup] = useState<boolean>(false);

  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const url = isSignup
        ? "http://localhost:8080/auth/signup"
        : "http://localhost/api/auth/login";

      // Prepare data based on the form mode
      const data = isSignup
        ? {
            username,
            firstname,
            lastname,
            email,
            password: SHA256(password).toString(),
          }
        : {
            email,
            password: SHA256(password).toString(),
          };

      const response = await axios.post<{
        token: string;
        chat_hostname: string;
        chat_port: number;
      }>(url, data, {
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!isSignup) {
        const {
          token,
          chat_hostname: hostname,
          chat_port: port,
        } = response.data;
        console.log(token, hostname, port);

        const wsUrl = `ws://localhost:${port}/connect`;
        setWs(new WebSocket(wsUrl));

        navigate("/channels/@me");
      }

      // Reset fields after successful sign-up
      if (isSignup) {
        setUsername("");
        setFirstname("");
        setLastname("");
        setEmail("");
        setPassword("");

        navigate("/login");
      }
    } catch (err) {
      console.error(`${isSignup ? "Sign up" : "Login"} failed:`, err);
      setError(
        `${isSignup ? "Sign up" : "Login"} failed. Please check your credentials and try again.`,
      );
    }
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Container component="main" maxWidth="xs">
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography component="h1" variant="h5">
            {isSignup ? "Sign Up" : "Login"}
          </Typography>
          <Box
            component="form"
            onSubmit={handleSubmit}
            noValidate
            sx={{ mt: 1 }}
          >
            {isSignup && (
              <>
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="username"
                  label="Username"
                  name="username"
                  autoComplete="username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="firstname"
                  label="First Name"
                  name="firstname"
                  value={firstname}
                  onChange={(e) => setFirstname(e.target.value)}
                />
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="lastname"
                  label="Last Name"
                  name="lastname"
                  value={lastname}
                  onChange={(e) => setLastname(e.target.value)}
                />
              </>
            )}
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              name="email"
              autoComplete="email"
              autoFocus
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            {error && <Alert severity="error">{error}</Alert>}
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              sx={{ mt: 3, mb: 2 }}
            >
              {isSignup ? "Sign Up" : "Login"}
            </Button>
            <Button
              fullWidth
              variant="text"
              onClick={() => setIsSignup((prev) => !prev)}
              sx={{ mt: 2 }}
            >
              {isSignup
                ? "Already have an account? Login"
                : "Don't have an account? Sign Up"}
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
};

export default LoginForm;

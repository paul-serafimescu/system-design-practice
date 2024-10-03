import React from "react";
import { useNavigate } from "react-router-dom";
import { Button, Container, Typography } from "@mui/material";

export const LandingPage: React.FC = () => {
  const navigate = useNavigate();

  const handleLoginRedirect = () => {
    navigate("/login");
  };

  return (
    <Container style={{ textAlign: "center", marginTop: "50px" }}>
      <Typography variant="h3" gutterBottom>
        Welcome to ChatApp
      </Typography>
      <Typography variant="h5" gutterBottom>
        Connect and chat with your friends instantly.
      </Typography>
      <Button
        variant="contained"
        color="primary"
        onClick={handleLoginRedirect}
        style={{ marginTop: "20px" }}
      >
        Login
      </Button>
    </Container>
  );
};

export default LandingPage;

import React from "react";
import { Grid, Container, Typography } from "@mui/material";
import ConversationList from "./ConversationList";
import ChatWindow from "./ChatWindow";

const DirectMessages: React.FC = () => {
  return (
    <Container
      disableGutters
      maxWidth={false}
      sx={{ height: "100vh", display: "flex", flexDirection: "row" }}
    >
      <Grid container sx={{ height: "100%", width: "100%" }}>
        <Grid
          xs={3} // Changed to 3 for a narrower sidebar
          md={2}
          sx={{
            backgroundColor: "#2f3136",
            color: "white",
            padding: 1, // Reduced padding for a tighter look
            borderRight: "1px solid #424549",
            overflowY: "auto",
          }}
        >
          <Typography variant="h6" gutterBottom>
            Conversations
          </Typography>
          <ConversationList />
        </Grid>
        <Grid
          xs={9} // Increased to take more space
          md={10}
          sx={{
            backgroundColor: "#36393f",
            color: "white",
            padding: 2,
            display: "flex",
            flexDirection: "column",
            height: "100%",
          }}
        >
          <Typography variant="h6" gutterBottom>
            Chat
          </Typography>
          <ChatWindow />
        </Grid>
      </Grid>
    </Container>
  );
};

export default DirectMessages;

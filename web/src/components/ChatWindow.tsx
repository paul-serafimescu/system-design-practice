import React from "react";
import {
  Box,
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
} from "@mui/material";

const ChatWindow: React.FC = () => {
  const messages = ["Hello", "How are you?", "I am fine, thank you!"];

  const handleSend = () => {
    // Logic to send message
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
        borderRadius: 1,
        padding: 2,
      }}
    >
      <List sx={{ flexGrow: 1, overflowY: "auto", marginBottom: 2 }}>
        {messages.map((text, index) => (
          <ListItem key={index}>
            <ListItemText
              primary={text}
              sx={{ color: "white", fontSize: "14px" }}
            />
          </ListItem>
        ))}
      </List>
      <Box sx={{ display: "flex", marginTop: 2 }}>
        <TextField
          variant="outlined"
          fullWidth
          placeholder="Type a message"
          sx={{ backgroundColor: "white" }}
        />
        <Button variant="contained" color="primary" onClick={handleSend}>
          Send
        </Button>
      </Box>
    </Box>
  );
};

export default ChatWindow;

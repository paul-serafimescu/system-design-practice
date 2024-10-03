import React from "react";
import { List, ListItem, ListItemText, Divider } from "@mui/material";

const ConversationList: React.FC = () => {
  const conversations = ["Alice", "Bob", "Charlie"];

  return (
    <List sx={{ padding: 0 }}>
      {conversations.map((name, index) => (
        <React.Fragment key={index}>
          <ListItem component="button">
            <ListItemText
              primary={name}
              sx={{ color: "white", fontSize: "14px" }}
            />
          </ListItem>
          <Divider sx={{ backgroundColor: "#424549" }} />
        </React.Fragment>
      ))}
    </List>
  );
};

export default ConversationList;

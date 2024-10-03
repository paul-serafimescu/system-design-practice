import React from "react";
import {
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  Switch,
  Container,
} from "@mui/material";
import { Brightness4, Brightness7 } from "@mui/icons-material";

interface NavbarProps {
  darkMode: boolean;
  setDarkMode: (darkmode: boolean) => void;
  children?: React.ReactNode;
}

const Navbar: React.FC<NavbarProps> = ({ darkMode, setDarkMode, children }) => {
  const handleThemeChange = () => {
    setDarkMode(!darkMode);
  };

  return (
    <React.Fragment>
      <AppBar position="static">
        <Container>
          <Toolbar>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              ChatApp
            </Typography>
            <IconButton color="inherit" onClick={handleThemeChange}>
              {darkMode ? <Brightness7 /> : <Brightness4 />}
            </IconButton>
            <Switch
              checked={darkMode}
              onChange={handleThemeChange}
              color="default"
            />
          </Toolbar>
        </Container>
      </AppBar>
      {children}
    </React.Fragment>
  );
};

export default Navbar;

import React from "react";
import LogoImg from "../images/logo_pentagol.svg";
import DarkModeOutlinedIcon from "@mui/icons-material/DarkModeOutlined";
import LightModeOutlinedIcon from "@mui/icons-material/LightModeOutlined";
import { IconButton } from "@mui/material";
import { useStore } from "../context/store";
import { useStatics } from "../context/statics";
import { Link } from "react-router-dom";

const Header = () => {
  const { DARK, LIGHT } = useStatics();
  const { theme, setTheme } = useStore();

  return (
    <div className="header">
      <div className="header__content container">
        <Link to="/">
          <img src={LogoImg} alt="Pentagol" />
        </Link>
        <IconButton
          className={`header__btn ${theme}`}
          sx={{ backgroundColor: "#fff" }}
          onClick={() => setTheme(theme === LIGHT ? DARK : LIGHT)}
        >
          {theme === LIGHT ? (
            <LightModeOutlinedIcon sx={{ fontSize: "20px" }} />
          ) : (
            <DarkModeOutlinedIcon sx={{ fontSize: "20px" }} />
          )}
        </IconButton>
      </div>
    </div>
  );
};

export default Header;

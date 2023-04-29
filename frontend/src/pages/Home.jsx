import React from "react";
import { Header, Footer, SelectBar } from "../components";
import SeriaALogo from "../images/seria_logo.svg";
const User = () => {
  return (
    <>
      <Header />

      <div className="container">
        <SelectBar leagues={leagues} />
      </div>

      <Footer />
    </>
  );
};

export default User;

const leagues = [
  {
    id: 1,
    name: "Italiya. Seriya A",
    img: SeriaALogo,
  },
  {
    id: 2,
    name: "Premier Liga",
    img: SeriaALogo,
  },
  {
    id: 3,
    name: "LaLiga",
    img: SeriaALogo,
  },
  {
    id: 4,
    name: "Bundesliga",
    img: SeriaALogo,
  },
  {
    id: 5,
    name: "Ligue 1",
    img: SeriaALogo,
  },
];

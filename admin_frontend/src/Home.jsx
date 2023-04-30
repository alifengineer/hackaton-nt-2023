import React, { useState, useEffect } from "react";
import Sidebar from "./components/Sidebar";
import NewsForm from "./components/NewsForm";
import { CreateMatch } from "./features/match";
import { CreateTeam } from "./features/team";
import Cookies from "js-cookie";
import axios from "./utils/axios";
import { authAxios } from "./utils/axios";

const menus = ["Match", "News", "Team"];

const Home = () => {
  const [activeMenu, setActiveMenu] = useState(menus[0]);
  useEffect(() => {
    const token = Cookies.get("token");
    if (token) {
      axios.defaults.headers.Authorization = `Bearer ${token}`;
      authAxios.defaults.headers.Authorization = `Bearer ${token}`;
    }
  }, []);

  return (
    <div className="d-flex">
      <Sidebar handleChange={setActiveMenu} menus={menus} active={activeMenu} />
      <div style={{ padding: "40px" }}>
        {activeMenu === "News" && <NewsForm />}
        {activeMenu === "Match" && <CreateMatch />}
        {activeMenu === "Team" && <CreateTeam />}
      </div>
    </div>
  );
};

export default Home;

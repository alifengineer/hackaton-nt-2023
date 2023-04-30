import React, { useState } from "react";
import Sidebar from "./components/Sidebar";
import NewsForm from "./components/NewsForm";
import { CreateMatch } from "./features/match";

const menus = ["Match", "News"];

const Home = () => {
  const [activeMenu, setActiveMenu] = useState(menus[0]);

  return (
    <div className="d-flex">
      <Sidebar handleChange={setActiveMenu} menus={menus} active={activeMenu} />
      <div style={{ padding: "40px" }}>
        {activeMenu === "News" && <NewsForm />}
        {activeMenu === "Match" && <CreateMatch />}
      </div>
    </div>
  );
};

export default Home;

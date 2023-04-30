import React from "react";
import Sidebar from "./components/Sidebar";
import NewsForm from "./components/NewsForm";
import { CreateMatch } from "./features/match";

const Home = () => {
  return (
    <div className="d-flex">
      <Sidebar />
      <NewsForm />
    </div>
  );
};

export default Home;

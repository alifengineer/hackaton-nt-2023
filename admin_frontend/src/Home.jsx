import React from "react";
import Sidebar from "./components/Sidebar";
import NewsForm from "./components/NewsForm";

const Home = () => {
  return (
    <div className="d-flex">
      <Sidebar />
      <NewsForm />
    </div>
  );
};

export default Home;

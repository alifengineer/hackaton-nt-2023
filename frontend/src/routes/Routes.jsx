import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from "../pages/Home";
import NewsPage from "../pages/NewsPage";

export default function Routers() {
  return (
    <Routes>
      <Route index element={<Home />} />
      <Route element={<NewsPage />} path="/news" />
    </Routes>
  );
}

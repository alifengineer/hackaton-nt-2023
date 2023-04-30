import { BrowserRouter } from "react-router-dom";
import { Routes, Route } from "react-router-dom";
import Login from "./components/Login";
import Home from "./Home";
// style
import "./style.scss";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Home />} />
        <Route element={<Login />} path="/login" />
      </Routes>
    </BrowserRouter>
  );
}
export default App;

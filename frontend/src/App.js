
import { BrowserRouter } from "react-router-dom";
import Routers from "./routes/Routes.jsx";
import "./styles/style.scss"
function App() {
  return (
    <BrowserRouter>
      <Routers />
    </BrowserRouter>
  );
}
export default App;

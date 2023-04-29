import { BrowserRouter } from "react-router-dom";
import Routers from "./routes/Routes.jsx";
import { StoreProvider } from "./context/store.js";
import { StaticsProvider } from "./context/statics.js";

// style
import "./styles/style.scss";

function App() {
  return (
    <BrowserRouter>
      <StaticsProvider>
        <StoreProvider>
          <Routers />
        </StoreProvider>
      </StaticsProvider>
    </BrowserRouter>
  );
}
export default App;

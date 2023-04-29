import { useState, createContext, useContext } from "react";
import { useStatics } from "./statics";

const StoreContext = createContext({});

export const StoreProvider = ({ children }) => {
  const { DARK } = useStatics();
  const [theme, setTheme] = useState(DARK);
  const [activeNextTour, setActiveNextTour] = useState(1);

  return (
    <StoreContext.Provider
      value={{
        theme,
        setTheme,
        activeNextTour,
        setActiveNextTour,
      }}
    >
      {children}
    </StoreContext.Provider>
  );
};

export const useStore = () => useContext(StoreContext);

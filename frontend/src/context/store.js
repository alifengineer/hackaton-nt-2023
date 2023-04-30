import { useState, createContext, useContext } from "react";
import { useStatics } from "./statics";

const StoreContext = createContext({});

export const StoreProvider = ({ children }) => {
  const { DARK } = useStatics();
  const [theme, setTheme] = useState(DARK);
  const [activeNextTour, setActiveNextTour] = useState({ tur: null, id: null });
  const [activeTour, setActiveTour] = useState({ tur: null, id: null });
  const [lastNewsPag, setLastNewsPag] = useState(0);
  const [allNewsPag, setAllNewsPag] = useState(0);

  const getTime = (time) => {
    if (time) {
      const date = new Date(time);
      return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(
        2,
        "0"
      )}-${String(date.getDate()).padStart(2, "0")}`;
    } else {
      return time;
    }
  };

  return (
    <StoreContext.Provider
      value={{
        theme,
        setTheme,
        activeNextTour,
        setActiveNextTour,
        activeTour,
        setActiveTour,
        getTime,
        setLastNewsPag,
        lastNewsPag,
        setAllNewsPag,
        allNewsPag,
      }}
    >
      {children}
    </StoreContext.Provider>
  );
};

export const useStore = () => useContext(StoreContext);

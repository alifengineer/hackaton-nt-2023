import { createContext, useContext } from "react";

const StaticsContext = createContext({});

export const StaticsProvider = ({ children }) => {
  const DARK = "dark";
  const LIGHT = "light";

  return (
    <StaticsContext.Provider
      value={{
        DARK,
        LIGHT,
      }}
    >
      {children}
    </StaticsContext.Provider>
  );
};

export const useStatics = () => useContext(StaticsContext);

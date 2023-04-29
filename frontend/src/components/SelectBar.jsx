import React from "react";
import { useStore } from "../context/store";

const SelectBar = ({ leagues }) => {
  const { activeNextTour, setActiveNextTour } = useStore();

  return (
    <div className="selectbar mt_30">
      <div className="leagues">
        {leagues?.map((league) => {
          return (
            <div
              className={`league ${
                activeNextTour === league.id && "active__league"
              }`}
              key={league.id}
              onClick={() => setActiveNextTour(league.id)}
            >
              <img src={league.img} alt={leagues.name} />
              <p>{league.name}</p>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default SelectBar;

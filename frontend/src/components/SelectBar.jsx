import React from "react";
import imgDef from "../images/default_user.png";

const SelectBar = ({
  leagues,
  activeId,
  handleActiveId,
  check,
  handleClick,
}) => {
  return (
    <div className="selectbar mt_30">
      <div className="leagues">
        {leagues?.map((league, index) => {
          return (
            <div
              className={`league ${
                activeId.id === league.id && "active__league"
              }`}
              key={index}
              onClick={() => {
                handleActiveId({
                  id: league.id,
                  tur: check ? league.next_tur : league.prev_tur,
                });
                handleClick(league);
              }}
            >
              <img src={league.image || imgDef} alt={leagues.name} />
              <p>{league.name}</p>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default SelectBar;

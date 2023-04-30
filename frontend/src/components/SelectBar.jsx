import React from "react";

const SelectBar = ({ leagues, activeId, handleActiveId }) => {
  return (
    <div className="selectbar mt_30">
      <div className="leagues">
        {leagues?.map((league, index) => {
          return (
            <div
              className={`league ${activeId === index && "active__league"}`}
              key={index}
              onClick={() => handleActiveId(index)}
            >
              <img src={league.image} alt={leagues.name} />
              <p>{league.name}</p>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default SelectBar;

import React from "react";
import { OneMatch } from "../components";

const MatchBox = ({ matches }) => {
  return (
    <div className="matchbox mt_20">
      {matches?.map((match, index) => {
        return <OneMatch key={index + 1} match={match} />;
      })}
    </div>
  );
};

export default MatchBox;

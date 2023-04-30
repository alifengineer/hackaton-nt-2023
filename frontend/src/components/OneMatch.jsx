import React from "react";

const OneMatch = ({ match }) => {
  return (
    <div className="match">
      <div className="match__head">
        <div className="team">
          <img src={match.home_team.image} alt={match.home_team.name} />
          <p className={`home--${match.winner}`}>{match.home_team.name}</p>
        </div>
        <div className="score">
          {match.winner ? (
            <div className="score__box">
              <p className={`home--${match.winner}`}>{match.home_team.score}</p>
              <b>:</b>
              <p className={`away--${match.winner}`}>{match.away_team.score}</p>
            </div>
          ) : (
            <p>vs</p>
          )}
        </div>
        <div className="team">
          <p className={`away--${match.winner}`}>{match.away_team.name}</p>
          <img src={match.away_team.image} alt={match.away_team.name} />
        </div>
      </div>
      <div className="match__bottom">
        <p>{match.m_date}</p>
      </div>
    </div>
  );
};

export default OneMatch;

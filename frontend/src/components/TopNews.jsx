import React from "react";
import { useStore } from "../context/store";
import { Link } from "react-router-dom";
const TopNews = ({ topnews }) => {
  const { getTime } = useStore();
  return (
    <div className="topnews">
      {topnews?.map((obj, index) => {
        return (
          <div
            className={`topnews__box topnews__box--${index + 1}`}
            key={index + 1}
          >
            <Link to={`/news?id=${obj.id}`}>
              <img src={obj.image} alt={obj.title} />
              <div className="topnews__content">
                <h2>{obj.title}</h2>
                <p className="mt_10">{getTime(obj.created_at)}</p>
              </div>
            </Link>
          </div>
        );
      })}
    </div>
  );
};

export default TopNews;

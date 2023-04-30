import React from "react";

const TopNews = ({ topnews }) => {
  return (
    <div className="topnews">
      {topnews?.map((obj, index) => {
        return (
          <div
            className={`topnews__box topnews__box--${index + 1}`}
            key={index + 1}
          >
            <img src={obj.image} alt={obj.title} />
            <div className="topnews__content">
              <h2>{obj.title}</h2>
              <p className="mt_10">{obj.date}</p>
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default TopNews;

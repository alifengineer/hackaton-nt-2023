import React from "react";

const News = ({ news }) => {
  return (
    <div className="news mt_20">
      {news?.map((obj, index) => {
        return (
          <div className="news__box" key={index + 1}>
            <img className="news__img mt_7" src={obj.image} alt={obj.title} />
            <h4 className="news__title mt_7">{obj.title}</h4>
            <p className="news__text mt_7">{obj.content}</p>
            <p className="news__date mt_7">{obj.created_at}</p>
          </div>
        );
      })}
    </div>
  );
};

export default News;

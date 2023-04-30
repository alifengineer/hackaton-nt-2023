import React from "react";
import { Link } from "react-router-dom";
import { useStore } from "../context/store";
import defImg from "../images/defimg.jpg";

const News = ({ news }) => {
  const { getTime } = useStore();
  return (
    <div className="news mt_20">
      {news?.map((obj, index) => {
        return (
          <Link to={`/news?id=${obj.id}`} className="news__box" key={index + 1}>
            <img
              className="news__img mt_7"
              src={obj.image || defImg}
              alt={obj.title}
            />
            <h4 className="news__title mt_7">{obj.title}</h4>
            <p className="news__text mt_7">{obj.content}</p>
            <p className="news__date mt_7">{getTime(obj.created_at)}</p>
          </Link>
        );
      })}
    </div>
  );
};

export default News;

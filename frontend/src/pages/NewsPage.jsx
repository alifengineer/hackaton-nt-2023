import React, { useEffect, useState } from "react";
import { Header, Footer, News, Heading } from "../components";
import { useLocation } from "react-router-dom";
import axios from "../utils/axios";
import { useStore } from "../context/store";
import { Link } from "react-router-dom";

const NewsPage = () => {
  const { getTime } = useStore();
  const [lastNews, setLastNews] = useState([]);
  const [info, setInfo] = useState(null);
  const [news, setNews] = useState([]);
  const [mainNews, setMainNews] = useState([]);
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const id = searchParams.get("id");

  useEffect(() => {
    const getData = async () => {
      const res = await axios.get(`/api/v1/news/${id}`);
      setInfo(res.data.data.news);
    };
    if (id) {
      getData();
    }
  }, [id]);

  useEffect(() => {
    const getData = async () => {
      const data = await axios.get(`/api/v1/news/?limit=8&offset=0`);
      setNews(data.data.data.news_list);
    };
    getData();
  }, []);

  useEffect(() => {
    const getData = async () => {
      const data = await axios.get(`/api/v1/news/?limit=4&offset=0`, {
        is_importtant: true,
      });
      setMainNews(data.data.data.news_list);
    };
    getData();
  }, []);

  useEffect(() => {
    const getData = async () => {
      const data = await axios.get(`/api/v1/news/latest?limit=4&offset=0`);
      setLastNews(data?.data.data.news_list);
    };
    getData();
  }, []);

  return (
    <>
      <Header />

      <div className="container mt_30">
        <div className="newspage__head">
          <p className="newspage__date">{getTime(info?.created_at)}</p>
          <h3 className="newspage__title mt_7">{info?.title}</h3>
        </div>
        {id && (
          <div className="newspage__main mt_20">
            <div className="newspage__content">
              <img src={info?.image} alt="news" />
              <p className="newspage__text mt_20">{info?.content}</p>
            </div>
            <div className="newspage__aside">
              <div className="main__news">
                <h2 className="main__news--title">Asosiy</h2>
                {mainNews?.map((obj, index) => {
                  return (
                    <div className="main__news--box mt_20" key={index + 1}>
                      <Link to={`/news?id=${obj.id}`}>
                        <h3 className="main__news--heading">{obj?.title}</h3>
                        <p className="main__news--date mt_7">
                          {getTime(obj?.created_at)}
                        </p>
                      </Link>
                    </div>
                  );
                })}
              </div>
              <div className="main__news mt_20">
                <h2 className="main__news--title">So'ngi</h2>
                {lastNews?.map((obj, index) => {
                  return (
                    <div className="main__news--box mt_20" key={index + 1}>
                      <h3 className="main__news--heading">{obj.title}</h3>
                      <p className="main__news--date mt_7">
                        {getTime(obj.created_at)}
                      </p>
                    </div>
                  );
                })}
              </div>
            </div>
          </div>
        )}

        <div className="mt_50">
          <Heading title="So'ngi yangiliklar" />
        </div>

        <News news={news} />
      </div>

      <Footer />
    </>
  );
};

export default NewsPage;

import React, { useEffect, useState } from "react";
import {
  Header,
  Footer,
  SelectBar,
  MatchBox,
  TopNews,
  Heading,
  LeagueTable,
  News,
} from "../components";
import { useStore } from "../context/store";
import SeriaALogo from "../images/seria_logo.svg";
import barsaImg from "../images/barsa.png";
import liverImg from "../images/liverpul.png";
import bigImg from "../images/big_img.png";
import smallImg from "../images/small_img.png";
import newsImg from "../images/news_img.png";
import axios from "../utils/axios";

const User = () => {
  const {
    activeNextTour,
    setActiveNextTour,
    activeTour,
    setActiveTour,
    setLastNewsPag,
    lastNewsPag,
    setAllNewsPag,
    allNewsPag,
  } = useStore();
  const [lastNews, setLastNews] = useState([]);
  const [allNews, setAllNews] = useState([]);
  const [leagues, setLeagues] = useState([]);
  const [nextMatches, setNextMatches] = useState([]);
  const [prevMatches, setPrevMatches] = useState([]);
  const [leagueTable, setLeagueTable] = useState([]);
  const [topNews, setTopNews] = useState();

  useEffect(() => {
    const getLeagues = async () => {
      const data = await axios.get(`/user-api/v1/league/`);
      setLeagues(data?.data.data);
      const obj = data?.data?.data[0];

      if (!obj) {
        return null;
      }
      setActiveNextTour({ id: obj.id, tur: obj.next_tur });
      setActiveTour({ id: obj.id, tur: obj.prev_tur });

      await axios
        .post(`/user-api/v1/match/tur`, {
          league_id: obj.id,
          tur: obj.next_tur,
        })
        .then((res) => {
          setNextMatches(res?.data.data.matches);
        })
        .catch((err) => {
          console.log(err);
        });

      await axios
        .post(`/user-api/v1/match/tur`, {
          league_id: obj.id,
          tur: obj.prev_tur,
        })
        .then((res) => {
          setPrevMatches(res?.data.data.matches);
        })
        .catch((err) => {
          console.log(err);
        });

      await axios
        .post(`/user-api/v1/league/top_teams`, {
          league_id: obj.id,
        })
        .then((res) => {
          setLeagueTable(res?.data.data.teams);
        })
        .catch((err) => {
          console.log(err);
        });
    };
    getLeagues();
  }, []);

  useEffect(() => {
    const getData = async () => {
      const data = await axios.get(
        `/user-api/v1/news/latest?limit=8&offset=${lastNewsPag}`
      );
      setLastNews([...lastNews, ...data?.data.data.news_list]);
    };
    getData();
  }, [lastNewsPag]);

  useEffect(() => {
    const getData = async () => {
      const data = await axios.get(
        `/user-api/v1/news/?limit=8&offset=${allNewsPag}`
      );
      setAllNews([...allNews, ...data.data.data.news_list]);
    };
    getData();
  }, [allNewsPag]);

  useEffect(() => {
    const getData = async () => {
      const res = await axios.get(`/user-api/v1/news/latest?limit=7&offset=0`);
      setTopNews(res.data.data.news_list);
    };
    getData();
  }, []);

  const getTable = async (iid) => {
    await axios
      .post(`/user-api/v1/league/top_teams`, {
        league_id: iid,
      })
      .then((res) => {
        setLeagueTable(res?.data.data.teams);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const getNextMatches = async (obj) => {
    await axios
      .post(`/user-api/v1/match/tur`, { league_id: obj.id, tur: obj.next_tur })
      .then((res) => {
        setNextMatches(res?.data.data.matches);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const getPrevMatches = async (obj) => {
    await axios
      .post(`/user-api/v1/match/tur`, { league_id: obj.id, tur: obj.prev_tur })
      .then((res) => {
        setPrevMatches(res?.data.data.matches);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <>
      <Header />

      <div className="container">
        <SelectBar
          check={true}
          leagues={leagues}
          activeId={activeNextTour}
          handleActiveId={setActiveNextTour}
          handleClick={getNextMatches}
        />

        <MatchBox matches={nextMatches} />

        <SelectBar
          check={false}
          leagues={leagues}
          activeId={activeTour}
          handleActiveId={setActiveTour}
          handleClick={getPrevMatches}
        />

        <MatchBox matches={prevMatches} />

        <div className="main mt_50">
          <TopNews topnews={topNews} />
          <LeagueTable
            handleTable={getTable}
            leagues={leagues}
            teams={[...leagueTable, ...leagueTable]}
          />
        </div>

        <div className="mt_50">
          <Heading title="So'ngi yangiliklar" />
        </div>

        <News news={lastNews} />

        <div
          className="view__more mt_30"
          onClick={() => {
            setLastNewsPag(lastNewsPag + 8);
          }}
        >
          <p>Barchasi ko’rish</p>
        </div>

        <div className="mt_50">
          <Heading title="Barcha yangiliklar" />
        </div>

        <News news={allNews} />

        <div
          className="view__more mt_30"
          onClick={() => {
            setAllNewsPag(allNewsPag + 8);
          }}
        >
          <p>Barchasi ko’rish</p>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default User;

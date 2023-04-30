import React from "react";
import { Header, Footer, News, Heading } from "../components";
import newsImg from "../images/news_img.png";
import bigImg from "../images/big_img.png";

const NewsPage = () => {
  return (
    <>
      <Header />

      <div className="container mt_30">
        <div className="newspage__head">
          <p className="newspage__date">{data.date}</p>
          <h3 className="newspage__title mt_7">{data.title}</h3>
        </div>

        <div className="newspage__main mt_20">
          <div className="newspage__content">
            <img src={data.image} alt="news" />
            <p className="newspage__text mt_20">{data.content}</p>
          </div>
          <div className="newspage__aside">
            <div className="main__news">
              <h2 className="main__news--title">Asosiy</h2>
              {main_news.map((obj, index) => {
                return (
                  <div className="main__news--box mt_20" key={index + 1}>
                    <h3 className="main__news--heading">{obj.title}</h3>
                    <p className="main__news--date mt_7">{obj.date}</p>
                  </div>
                );
              })}
            </div>
            <div className="main__news mt_20">
              <h2 className="main__news--title">Asosiy</h2>
              {main_news.map((obj, index) => {
                return (
                  <div className="main__news--box mt_20" key={index + 1}>
                    <h3 className="main__news--heading">{obj.title}</h3>
                    <p className="main__news--date mt_7">{obj.date}</p>
                  </div>
                );
              })}
            </div>
          </div>
        </div>

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

const news = [
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
];
const main_news = [
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
  {
    id: "uuid",
    title: "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi",
    content:
      "Zidan Manchester Yunayted ga Premer-ligaga tayyor emasligini aytdi ",
    date: "2020-10-10",
    image: newsImg,
  },
];

const data = {
  id: "uuid",
  title:
    "Faqatgina muvaffaqiyatli qur'a Rossiyani Jahon chempionatiga olib boradi: bo'g'inlar oldidagi barcha maketlar",
  content:
    "Ilon Mask (Elon Mask, Tesla va boshqa loyihalar rahbari) bir yarim hafta ichida ikkinchi marta Xitoyning hukmron doiralari va fuqarolariga murojaat qildi. Videoda u ularni ushbu mamlakatda biznesni rivojlantirish va kengaytirishga sarmoya kiritishga tayyor ekanligiga ishontirdi. Qayta ishlangan ma'lumotlarning xavfsizligiga alohida e'tibor qaratiladi Ular faqat Xitoyning o'zida saqlanadi va shuning uchun istalgan vaqtda Orta Qirollik hukumati foydalanishi mumkin boladi.Ilon Mask Xitoyda hafta oxirida bo‘lib o‘tgan Butunjahon internet konferensiyasining tomoshabinlariga videomurojaati chog‘ida yangi bayonotlar berdi. Kompaniya Cisco Systems (Chak Robbins), Intel (Pat Gelsinger) va Qualcomm (Kristiano Amon) rahbarlaridan tashkil topgan bo'lib, Xitoy Xalq Respublikasi biznesi manfaatlarini Alibaba va Xiaomi rahbarlari himoya qilishgan.Tadbirni Xitoy Xalq Respublikasi Davlat kengashi bosh vaziri o‘rinbosari Lyu Xe ochib berdi va Si Szinpinning Osmon imperiyasining shaffof raqamli iqtisodiyotni yaratish uchun barcha kuchlar bilan ishlash istagi haqidagi so‘zlarini keltirdi.Yaqin kelajakda Tesla nafaqat Shanxay filialida elektromobillar ishlab chiqarishni kengaytiribgina qolmay, balki mahalliy studiya yordamida elektromobilning arzon modelini (narxi 25 000 dollardan kam) ishlab chiqmoqchi.Xitoy Xalq Respublikasida yig‘ilgan Tesla Model Y va Model 3 elektromobillari allaqachon Yevropaga eksport qilinmoqda.",
  image: bigImg,
  date: "2020-10-10",
};

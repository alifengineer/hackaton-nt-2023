import React from "react";
import Logo from "../images/logo_pentagol.svg";
import Twitter from "../images/Twitter.png";
import Telegram from "../images/Telegram.png";
import YouTube from "../images/YouTube.png";
import Facebook from "../images/Facebook.png";
import Instagram from "../images/Instagram.png";
import Skype from "../images/Skype.png";
import VKontakte from "../images/Vkontakte (VK).png";

const Footer = () => {
  return (
    <div className="footer mt_50">
      <div className="container footer__content">
        <div className="footer__left">
          <div className="footer__col">
            <img src={Logo} alt="pentagol" />
          </div>
          <div className="footer__menu">
            <p>Main</p>
            <p>Mach center</p>
            <p>Football</p>
            <p>Tennis</p>
            <p>Cybersport</p>
          </div>
          <div className="footer__menu">
            <p>Main</p>
            <p>Mach center</p>
            <p>Football</p>
            <p>Tennis</p>
            <p>Cybersport</p>
          </div>
          <div className="footer__menu">
            <p>Main</p>
            <p>Mach center</p>
            <p>Football</p>
            <p>Tennis</p>
            <p>Cybersport</p>
          </div>
        </div>
        <div className="footer__right">
          <div className="footer__social">
            <div className="social__row">
              <img src={Twitter} alt="Twitter" />
              <img src={YouTube} alt="YouTube" />
              <img src={VKontakte} alt="VKontakte" />
            </div>
            <div className="social__row">
              <img src={Skype} alt="Skype" />
              <img src={Instagram} alt="Instagram" />
              <img src={Facebook} alt="Facebook" />
              <img src={Telegram} alt="Telegram" />
            </div>
          </div>
        </div>
      </div>
      <div className="footer__copyright mt_30">
        © 2023 Footboll. All rights reserved
      </div>
    </div>
  );
};

export default Footer;

import React from "react";
import Google from "../Assets/google.png";
import Slack from "../Assets/slack.png";
import Atlassian from "../Assets/atlassian.png";
import Dropbox from "../Assets/dropbox.png";
import Shopify from "../Assets/shopify.png";
import "../styles/FirstSection.css";

const Company = () => {
   return (
      <section className="company-container">
         <img src={Google} alt="Google" />
         <img src={Slack} alt="Slack" />
         <img src={Atlassian} alt="Atlassain" />
         <img src={Dropbox} alt="Dropbox" />
         <img src={Shopify} alt="Shopify" />
      </section>
   );
};

export default Company;

import React from "react";
import MyApps from "./Assets/possibility.png";
import "./styles/Fourth.css";

const FourthSection = () => {
   return (
      <section className="fourth-section">
         <article className="image-container">
            <img src={MyApps} alt="MyApps" />
         </article>
         <article className="text-container">
            <p>Request Early Access to Get Started</p>
            <h3>The possibilities are beyond your imagination</h3>
            <p>
               Yet bed any for travelling assistance indulgence unpleasing. Not
               thoughts all exercise blessing. Indulgence way everything joy
               alteration boisterous the attachment. Party we years to order
               allow asked of.
            </p>
            <p>Request Early Access to Get Started</p>
         </article>
      </section>
   );
};

export default FourthSection;

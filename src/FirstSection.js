import React from "react";
import NavBar from "./componets/NavBar.js";
import "./styles/FirstSection.css";
import AI from "./Assets/ai.png";
import Company from "./componets/Company.js";
import People from "./Assets/people.png";

const FirstSection = () => {
   return (
      <>
         <NavBar />
         <section className="first-section">
            <article className="text-section">
               <h2>Let’s Build Something amazing with GPT-3 OpenAI</h2>
               <p>
                  Yet bed any for travelling assistance indulgence unpleasing.
                  Not thoughts all exercise blessing. Indulgence way everything
                  joy alteration boisterous the attachment. Party we years to
                  order allow asked of.
               </p>
               <div className="input-people-container">
                  <div className="input-container">
                     <input type="email" placeholder="Your Email Address" />
                     <button className="btn3"> Get Started </button>
                  </div>
                  <div className="people-container">
                     <img src={People} alt="People" />
                     <p>
                        1,600 people requested access a visit in last 24 hours
                     </p>
                  </div>
               </div>
            </article>
            <article className="image-section">
               <img src={AI} alt="AI-Pic" />
            </article>
         </section>
         <Company />
      </>
   );
};

export default FirstSection;

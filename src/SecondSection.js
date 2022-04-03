import React from "react";
import "./styles/Second.css";

const SecondSection = () => {
   return (
      <section
         data-aos="fade-right"
         data-aos-offset="500"
         data-aos-duration="500"
         className="second-section"
      >
         <article className="main">
            {/* First Article */}
            <aside className="first">
               <div className="first-title">
                  <h2>What is GPT-3</h2>
               </div>
               <div className="first-text">
                  <p>
                     We so opinion friends me message as delight. Whole front do
                     of plate heard oh ought. His defective nor convinced
                     residence own. Connection has put impossible own apartments
                     boisterous. At jointure ladyship an insisted so humanity
                     he. Friendly bachelor entrance to on by.
                  </p>
               </div>
            </aside>
            {/* Second Article */}
            <aside className="second">
               <h2>
                  The possibilities are beyond <br /> your imagination
               </h2>
               <p>Explore The Library</p>
            </aside>
            {/* Third Article */}
            <aside className="third">
               <div className="item">
                  <h2>ChatBots</h2>
                  <p>
                     We so opinion friends me message as delight. Whole front do
                     of plate heard oh ought.
                  </p>
               </div>
               <div className="item">
                  <h2>Knowledgebase</h2>
                  <p>
                     At jointure ladyship an insisted so humanity he. Friendly
                     bachelor entrance to on by. As put impossible own
                     apartment.
                  </p>
               </div>
               <div className="item">
                  <h2>Education</h2>
                  <p>
                     At jointure ladyship an insisted so humanity he. Friendly
                     bachelor entrance to on by. As put impossible own
                     apartments.
                  </p>
               </div>
            </aside>
         </article>
      </section>
   );
};

export default SecondSection;

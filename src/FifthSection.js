import React from "react";
import "./styles/Fifth.css";
import Blog01 from "./Assets/blog01.png";

import Blog02 from "./Assets/blog02.png";

import Blog03 from "./Assets/blog03.png";

import Blog04 from "./Assets/blog04.png";

import Blog05 from "./Assets/blog05.png";
import Article from "./componets/Article.js";

const FifthSection = () => {
   return (
      <section
         data-aos="fade-left"
         data-aos-duration="500"
         className="fifth-section"
      >
         <article className="title-container">
            <h2>
               A lot is happening, <br /> We are blogging about it.
            </h2>
         </article>
         <article class="container">
            <div className="group-a">
               <Article imageURL={Blog01} />
            </div>
            <div className="group-b">
               <Article imageURL={Blog02} />
               <Article imageURL={Blog03} />
               <Article imageURL={Blog04} />
               <Article imageURL={Blog05} />
            </div>
         </article>
      </section>
   );
};

export default FifthSection;

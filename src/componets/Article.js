import React from "react";

const Article = ({ imageURL }) => {
   return (
      <div className="one-box">
         <div className="image-container-1">
            <img src={imageURL} alt="Blog 1" />
         </div>
         <div className="text-container-1">
            <div>
               <p>Sep 26, 2021</p>
               <h3>
                  GPT-3 and Open AI is the future. Let us exlore how it is?
               </h3>
            </div>
            <p>Read Full Article</p>
         </div>
      </div>
   );
};

export default Article;

import React, { useState } from "react";
import "../styles/NavBar.css";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars } from "@fortawesome/free-solid-svg-icons";
import { faTimes } from "@fortawesome/free-solid-svg-icons";

const NavBar = () => {
   const [show, setShow] = useState(false);

   const navButton = () => {
      setShow(true);
   };

   const closeButton = () => {
      setShow(false);
   };

   return (
      <header className="header">
         <h1>GPT-3</h1>
         <nav className={!show ? "nav-bar" : "side-bar"}>
            <button className="close-button" onClick={closeButton}>
               <FontAwesomeIcon icon={faTimes} />
            </button>
            <ul>
               <li>Home</li>

               <li>What is GPT?</li>

               <li>Open AI</li>

               <li>Case Studies</li>

               <li>Library</li>
            </ul>

            <div className="btn-container">
               <button className="btn-1">Sign In</button>
               <button className="btn-2">Sign Up</button>
            </div>
         </nav>
         {!show && (
            <button className="nav-toggle" onClick={navButton}>
               <FontAwesomeIcon icon={faBars} />
            </button>
         )}
      </header>
   );
};

export default NavBar;

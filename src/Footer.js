import React from "react";
import "./styles/Footer.css";

const Footer = () => {
   return (
      <footer className="footer">
         <article className="footer-first">
            <h2>GPt-3</h2>
            <p>
               Crechterwoord K12 182 DK <br /> Alknjkcb, All Rights Reserved
            </p>
         </article>
         <article className="footer-second">
            {/* First List */}
            <div>
               <h3>Links</h3>
               <ul>
                  <li>Overons</li>
                  <li>Social Media</li>
                  <li>Counters</li>
                  <li>Contact</li>
               </ul>
            </div>
            {/* Second List */}
            <div>
               <h3>Company</h3>
               <ul>
                  <li>Terms & Conditions</li>
                  <li>Privacy Policy</li>
                  <li>Contact</li>
               </ul>
            </div>
            {/* Third List */}
            <div>
               <h3>Get in touch</h3>
               <ul>
                  <li>
                     Crechterwoord K12 182 DK <br /> Alknjkcb
                  </li>
                  <li>085-132567</li>
                  <li>info@payme.net</li>
               </ul>
            </div>
         </article>
      </footer>
   );
};

export default Footer;

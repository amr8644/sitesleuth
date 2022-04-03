import FirstSection from "./FirstSection.js";

import SecondSection from "./SecondSection.js";
import ThirdSection from "./ThirdSection.js";
import FourthSection from "./FourthSection.js";
import SubSection from "./componets/SubSection.js";
import FifthSection from "./FifthSection.js";
import LastSection from "./LastSection.js";
import Footer from "./Footer.js";
import NavBar from "./componets/NavBar.js";

function App() {
   return (
      <>
         <NavBar />
         {/* First Section */}

         <FirstSection />

         <SecondSection />
         <ThirdSection />
         <FourthSection />
         <SubSection />
         <FifthSection />
         <LastSection />
         <Footer />
      </>
   );
}

export default App;

import FirstSection from "./FirstSection.js";
import {
   Animator,
   ScrollContainer,
   ScrollPage,
   batch,
   Fade,
   FadeIn,
   Move,
   MoveIn,
   MoveOut,
   Sticky,
   StickyIn,
   ZoomIn,
} from "react-scroll-motion";
import SecondSection from "./SecondSection.js";
import ThirdSection from "./ThirdSection.js";
import FourthSection from "./FourthSection.js";
import SubSection from "./componets/SubSection.js";
import FifthSection from "./FifthSection.js";
import LastSection from "./LastSection.js";
import Footer from "./Footer.js";

const ZoomInScrollOut = batch(StickyIn(), FadeIn(), ZoomIn());
const FadeUp = batch(Fade(), Move(), Sticky());

function App() {
   return (
      <>
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

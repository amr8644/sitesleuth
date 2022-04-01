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
      </>
   );
}

export default App;

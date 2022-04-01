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

const ZoomInScrollOut = batch(StickyIn(), FadeIn(), ZoomIn());
const FadeUp = batch(Fade(), Move(), Sticky());

function App() {
   return (
      <>
         <FirstSection />
         <SecondSection />
      </>
   );
}

export default App;

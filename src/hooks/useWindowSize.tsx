import { useEffect, useState } from "react";

let height: number;
let width: number;

export default function useWindowSize() {
  const [windowSize, setWindowSize] = useState({ width, height });

  useEffect(() => {
    function handleResize() {
      setWindowSize({
        width: window.innerWidth,
        height: window.innerHeight,
      });
    }

    window.addEventListener("resize", handleResize);
    handleResize();
    return () => window.removeEventListener("resize", handleResize);
  }, []);
  return windowSize;
}
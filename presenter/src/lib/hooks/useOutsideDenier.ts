/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect } from "react";

/**
 * Hook that set to false a setter when  clicks outside of the passed ref
 */
function useOutsideDenier(
  ref: any,
  setter: (value : boolean) => void
) {
  useEffect(() => {
    /**
     * Use the setter to set the state to false when the user clicks outside of the passed ref
     * This modify an external state
     */
    function handleClickOutside(event: any) {
      if (ref.current && !ref.current.contains(event.target)) {
        setter(false);
      }
    }
    // Bind the event listener
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      // Unbind the event listener on clean up
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [ref, setter]);
}

export default useOutsideDenier;

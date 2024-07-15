import { RefObject, useEffect } from "react";

/**
 * Hook that set to false a closeFn when  clicks outside of the passed ref
 * 
 */
function useClickOutside(
  ref: RefObject<any>,
  closeFn: () => void
) {
  useEffect(() => {
    function handleClickOutside(event: any) {
      if (ref.current && !ref.current.contains(event.target)) {
        closeFn();
      }
    }
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [ref, closeFn]);
}

export default useClickOutside;
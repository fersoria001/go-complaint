import { useEffect, useState } from "react";
import { Query, CountiesQuery, CountyListType } from "../queries";
import { County } from "../types";

function useCounties(countryID: number): County[] {
  const [counties, setCounties] = useState<County[]>([]);
  useEffect(() => {
    const countyFallback = [{ id: 0, name: "No counties found" }];
    Query<County[]>(CountiesQuery, CountyListType, [countryID]).then((data) => {
      if (data.length > 0) {
        setCounties(data);
        return;
      }
      setCounties(countyFallback);
      return;
    });
  }, [countryID]);

  return counties;
}

export default useCounties;

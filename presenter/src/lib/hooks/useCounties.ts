import { useEffect, useState } from "react";
import { CountryStateListType, CountryStatesQuery, Query } from "../queries";
import { CountryState } from "../types";

function useCounties(countryID: number): CountryState[] {
  const [counties, setCounties] = useState<CountryState[]>([]);
  useEffect(() => {
    async function fetchCountryStates() {
      const data = await Query<CountryState[]>(
        CountryStatesQuery,
        CountryStateListType,
        [countryID]
      );
      setCounties(data);
    }
    fetchCountryStates();
  }, [countryID]);

  return counties;
}

export default useCounties;

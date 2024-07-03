import { useEffect, useState } from "react";
import { Query, CitiesQuery, CityListType } from "../queries";
import { City } from "../types";

function useCities(countID: number): City[] {
  const [cities, setCities] = useState<City[]>([]);
  useEffect(() => {
    async function fetchCities() {
      const data = await Query<City[]>(CitiesQuery, CityListType, [countID]);
      setCities(data);
    }
    fetchCities();
  }, [countID]);
  return cities;
}

export default useCities;

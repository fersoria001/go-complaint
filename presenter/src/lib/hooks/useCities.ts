import { useEffect, useState } from "react";
import { Query, CitiesQuery, CityListType } from "../queries";
import { City } from "../types";

function useCities(countID : number) : City[]{
    const [cities, setCities] = useState<City[]>([]);
    useEffect(() => {
        const cityFallback = [{ id: 0, name: "No cities found" }];
        Query<City[]>(CitiesQuery, CityListType, [countID]).then((data) => {
            if (data.length > 0) {
                setCities(data);
                return
            }
            setCities(cityFallback);
            return
        });
    }, [countID]);
    return cities;
}

export default useCities;
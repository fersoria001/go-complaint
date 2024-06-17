import { useEffect, useState } from "react";
import {
  Query,
  OfficesQuery,
  OfficeTypeList,
} from "../queries";
import { Office } from "../types";

function useOffices(): Office[] {
  const [offices, setOffices] = useState<Office[]>([]);
  useEffect(() => {
    Query<Office[]>(OfficesQuery, OfficeTypeList)
      .then((data) => setOffices(data))
  }, []);
  return offices;
}

export default useOffices;

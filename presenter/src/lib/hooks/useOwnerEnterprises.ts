import { useEffect, useState } from "react";
import {
  Query,
  OwnerEnterprisesQuery,
  OwnerEnterprisesTypeList,
} from "../queries";
import { Enterprise } from "../types";

function useOwnerEnterprises(): Enterprise[] {
  const [enterprises, setEnterprises] = useState<Enterprise[]>([]);
  useEffect(() => {
    Query<Enterprise[]>(OwnerEnterprisesQuery, OwnerEnterprisesTypeList)
      .then((data) => setEnterprises(data))
  }, []);
  return enterprises;
}

export default useOwnerEnterprises;

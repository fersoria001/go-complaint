import { useEffect, useState } from "react";
import { Country } from "../types";

function usePhonecode(countryID: number, countries: Country[]): string {
  const [phoneCode, setPhoneCode] = useState<string>("");
  useEffect(() => {
    function fetchPhoneCode() {
      const country = countries.find((c) => c.id == countryID)
      if (country) {
        setPhoneCode(country.phoneCode);
      }
    }
    fetchPhoneCode();
  }, [countryID, countries]);
  return phoneCode;
}

export default usePhonecode;

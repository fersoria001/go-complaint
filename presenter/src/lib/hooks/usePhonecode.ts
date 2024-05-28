import { useEffect, useState } from "react";
import { Query, PhoneCodeQuery, PhoneCodeType } from "../queries";
import { PhoneCode } from "../types";

function usePhonecode(countryID: number): PhoneCode {
  const [phoneCode, setPhoneCode] = useState<PhoneCode>({ id: 0, code: "" });
  useEffect(() => {
    const phoneCodeFallback = { id: 0, code: "" };
    Query<PhoneCode>(PhoneCodeQuery, PhoneCodeType, [countryID]).then(
      (data) => {
        if (data) {
          setPhoneCode(data);
          return;
        }
        setPhoneCode(phoneCodeFallback);
        return;
      }
    );
  }, [countryID]);
  return phoneCode;
}

export default usePhonecode;

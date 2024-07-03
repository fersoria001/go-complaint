import { z } from "zod";
import { Country, ErrorType } from "../../../lib/types";
import { useState } from "react";
import AcceptBtn from "../../buttons/AcceptBtn";
import { updatePhoneNumber } from "./settings_lib";


interface Props {
    countries: Country[],
    enterpriseID: string
}
const validatePhone = z
    .string({ message: "We could not validate your phone number" })
    .min(6, { message: "We could not validate your phone number" })
    .transform((val, ctx) => {
        const parsed = parseInt(val);
        if (isNaN(parsed)) {
            ctx.addIssue({
                code: z.ZodIssueCode.custom,
                message: "Not a number",
            });
            return z.NEVER;
        }
        return parsed;
    })
export const UpdatePhone: React.FC<Props> = ({ enterpriseID, countries }: Props) => {
    const [phoneCode, setPhoneCode] = useState<string>(countries[0].phoneCode)
    const [phone, setPhone] = useState<string>("")
    const [reset, setReset] = useState<boolean>(false)
    const [errors, setErrors] = useState<ErrorType>({})
    const handleUpdate = async () => {
        const result = validatePhone.safeParse(phone)
        if (!result.success) {
            setErrors({ phone: result.error.errors[0].message })
            return false
        }
        return updatePhoneNumber(enterpriseID, phoneCode + phone)
    }
    const handlePhoneCodeChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setErrors({})
        setPhoneCode(e.target.value)
        setReset(true)
    }
    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setPhone(e.target.value)
        setReset(true)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    return (
        <div className="flex flex-col">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="phone">Phone</label>
                <div className="w-full  flex mb-4">
                    <select
                        onChange={handlePhoneCodeChange}
                        value={phoneCode}
                        className="
                                w-2/4 md:w-1/4
                                bg-gray-100
                                border border-gray-300
                                  text-gray-900 text-sm rounded-lg
                                   focus:ring-blue-500 focus:border-blue-500 
                                   block p-2.5">
                        {
                            countries.map((country, index) => {
                                return <option key={index} value={country.phoneCode}>{country.phoneCode}</option>
                            })
                        }
                    </select>
                    <input
                        onChange={handleInputChange}
                        value={phone}
                        className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        name="phone" type="tel" />
                </div>
                {errors?.phone && <span className="text-red-500 text-xs italic">{errors.phone}</span>}
            </div>
            <div className="self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdatePhone;
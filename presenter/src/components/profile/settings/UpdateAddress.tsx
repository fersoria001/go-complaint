/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";
import useCities from "../../../lib/hooks/useCities";
import useCounties from "../../../lib/hooks/useCounties";
import { Country, ErrorType } from "../../../lib/types";
import SelectIcon from "../../icons/SelectIcon";
import { z } from "zod";

import AcceptBtn from "../../buttons/AcceptBtn";
import { updateAddress } from "./settings_lib";


interface Props {
    countries: Country[]
}
const addressSchema = z.object({
    country: z.number(),
    county: z.number(),
    city: z.number(),
})
const UpdateAddress: React.FC<Props> = ({ countries }: Props) => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [reset, setReset] = useState<boolean>(false)
    const [selectedCountry, setSelectedCountry] = useState<number>(0);
    const [selectedCounty, setSelectedCounty] = useState<number>(0);
    const [selectedCity, setSelectedCity] = useState<number>(0);
    const counties = useCounties(selectedCountry);
    const cities = useCities(selectedCounty);
    const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCountry(parseInt(e.target.value));
        setReset(true)
    }
    const handleCountyChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCounty(parseInt(e.target.value));
        setReset(true)
    }
    const handleCityChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCity(parseInt(e.target.value));
        setReset(true)
    }

    const handleUpdate = async () => {
        const parsed = addressSchema.safeParse({
            country: selectedCountry,
            county: selectedCounty,
            city: selectedCity
        })
        const errors: ErrorType = {};
        let errorPath: string;
        if (!parsed.success) {
            parsed.error.errors.forEach((error: any) => {
                errorPath = error.path.join("");
                errors[errorPath] = error.message;
            });
            setErrors(errors);
            return false;
        }
        return await updateAddress(selectedCountry, selectedCounty, selectedCity)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    return (
        <div>
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="country">Country</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                         border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                          focus:outline-none focus:bg-white focus:border-gray-500"
                        name="country"
                        onChange={handleCountryChange}
                    >
                        {countries && countries.map((country) => (
                            <option
                                key={country.id} value={country.id}>{country.name}</option>
                        )) || <option disabled>No countries found</option>}
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
                {errors?.country && <span className="text-red-500 text-xs italic">{errors.country}</span>}
            </div>

            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="county">County</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="county"
                        onChange={handleCountyChange} defaultValue={"Select a country first"}>
                        {counties && counties.map((county) => (
                            <option key={county.id} value={county.id}>{county.name}</option>
                        )) || <option disabled>Select a country first</option>}
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
                {errors?.county && <span className="text-red-500 text-xs italic">{errors.county}</span>}
            </div>

            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="city">City</label>
                <div className="relative">
                    <select
                        onChange={handleCityChange}
                        name="city"
                        className="block appearance-none w-full bg-gray-200 border border-gray-200
                         text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        defaultValue={"Select a county first"}
                    >
                        {cities && cities.map((city) => (
                            <option key={city.id} value={city.id}>{city.name}</option>
                        )) || <option disabled>Select a county first</option>}
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                    {errors?.city && <span className="text-red-500 text-xs italic">{errors.city}</span>}
                </div>
            </div>
            <div className="self-end mt-2">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdateAddress;
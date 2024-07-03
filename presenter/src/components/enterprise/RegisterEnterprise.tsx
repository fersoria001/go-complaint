import SelectIcon from "../icons/SelectIcon";
import PrimaryButton from "../buttons/PrimaryButton";
import { useState } from "react";
import useCities from "../../lib/hooks/useCities";
import useCounties from "../../lib/hooks/useCounties";

import { Route } from "../../routes/_profile/register-enterprise";
import usePhonecode from "../../lib/hooks/usePhonecode";
import { useRouter } from "@tanstack/react-router";
import { ErrorType, RegisterEnterpriseSchema } from "../../lib/types";
import { parseSchema } from "../../lib/parse_schema";
import { createEnterprise } from "../../lib/create_enterprise";

const RegisterEnterprise: React.FC = () => {
    const { countries, industries } = Route.useLoaderData();
    const [selectedCountry, setSelectedCountry] = useState<number>(0);
    const [selectedCounty, setSelectedCounty] = useState<number>(0);
    const [errors, setErrors] = useState<ErrorType>({});
    const counties = useCounties(selectedCountry);
    const cities = useCities(selectedCounty);
    const phoneCode = usePhonecode(selectedCountry, countries);

    const router = useRouter();
    const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCountry(parseInt(e.target.value));
    }
    const handleCountyChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCounty(parseInt(e.target.value));
    }
    const handleSubmit = async () => {
        const form = document.getElementById("register-enterprise-form") as HTMLFormElement;
        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());
        const { data: parsed, errors } = await parseSchema(data, RegisterEnterpriseSchema, true);
        if (Object.keys(errors).length > 0) {
            setErrors(errors);
            return
        }
        const ok = await createEnterprise(parsed);
        if (ok) {
            router.navigate({ to: `/${parsed.name}` })
        }
    }
    return (
        <form
            id="register-enterprise-form"
            className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4  max-w-lg mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="name">Name</label>
                <input
                    className="text-sm md:text-xl shadow appearance-none border rounded w-full py-2 px-2
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="name" name="name" />
                {errors?.name && <span className="text-red-500 text-xs italic" >{errors.name}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="website">Website</label>
                <input
                    className="text-sm  md:text-xl shadow appearance-none border rounded w-full py-2 px-2
                     text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="website" name="website" placeholder="http://www.mywebsite.com" />
                {errors?.website && <span className="text-red-500 text-xs italic" >{errors.website}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    className="text-sm md:text-xl shadow appearance-none border rounded w-full py-2 px-2
                     text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="email" id="email" name="email" />
                {errors?.email && <span className="text-red-500 text-xs italic" >{errors.email}</span>}
            </div>
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="countryID">Country</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200
                        text-sm md:text-xl
                         text-gray-700 py-3 px-2 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="countryID"
                        onChange={e => handleCountryChange(e)}
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
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="countryStateID">State</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-2
                        text-sm md:text-xl pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="countryStateID"
                        onChange={e => handleCountyChange(e)} defaultValue={"Select a country first"}>
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
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="cityID">City</label>
                <div className="relative">
                    <select
                        name="cityID"
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 
                        text-sm md:text-xl
                        py-3 px-2 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        defaultValue={"Select a county first"}>
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

            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="phone">Phone</label>
                {phoneCode &&
                    <div className="w-full flex mb-4">
                        <input
                            className="
                                w-1/4
                                bg-gray-100
                                border border-gray-300
                                  text-gray-900 text-sm md:text-xl rounded-lg
                                   focus:ring-blue-500 focus:border-blue-500 
                                   block p-2.5 cursor-not-allowed
"
                            name="phoneCode" type="tel" value={phoneCode} readOnly />
                        <input
                            className="text-sm md:text-xl shadow appearance-none border rounded w-full py-2 px-2
                             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            name="phone" type="tel" />
                    </div>
                    ||
                    <div className="mb-4">
                        <input

                            type="tel" disabled />
                        <input

                            type="tel" disabled />
                    </div>
                }
                {errors?.phone && <span className="text-red-500 text-xs italic">{errors.phone}</span>}
            </div>
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="industryID">Industry</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                        text-sm md:text-xl
                         border-gray-200 text-gray-700 py-3 px-2 pr-8 rounded 
                         leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="industryID"
                    >
                        {industries && industries.map((industry) => (
                            <option
                                key={industry.id} value={industry.id}>{industry.name}</option>
                        )) || <option disabled>No industries found</option>}
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
                {errors?.country && <span className="text-red-500 text-xs italic">{errors.country}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="foundationDate">Foundation Date</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-2
                    text-sm md:text-xl
                     text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="date" id="foundationDate" name="foundationDate" />
                {errors?.foundationDate && <span className="text-red-500 text-xs italic" >{errors.foundationDate}</span>}
            </div>
            <div className="mb-4">
                <div className="flex items-start">
                    <div className="flex items-center h-5 md:h-8">
                        <input
                            name="terms"
                            type="checkbox"
                            className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300" />
                    </div>
                    <label
                        className="block text-gray-700 text-sm md:text-xl font-bold mb-2 ms-2"
                        htmlFor="terms">
                        I agree to the terms and conditions
                    </label>
                </div>
                {errors?.terms && <span className="text-red-500 text-xs italic">{errors.terms}</span>}
            </div>
            <div
                onMouseUp={handleSubmit}
                className="flex flex-col justify-center items-center">
                <PrimaryButton text="Sign up" />
            </div>
        </form >
    );
}

export default RegisterEnterprise;
import { Form, useActionData, useLoaderData } from "react-router-dom";
import SelectIcon from "../components/icons/SelectIcon";
import PrimaryButton from "../components/buttons/PrimaryButton";
import { useState } from "react";
import useCities from "../lib/hooks/useCities";
import useCounties from "../lib/hooks/useCounties";
import usePhonecode from "../lib/hooks/usePhonecode";
import { Country, Industry } from "../lib/types";

function RegisterEnterprise() {
    const loaderData = useLoaderData() as { countries: Country[]; industries: Industry[] }
    let countries: Country[] = [{ id: 0, name: "No countries found" }]
    let industries: Industry[] = [{ id: 0, name: "No industries found" }]
    if (loaderData) {
        countries = loaderData.countries as Country[];
        industries = loaderData.industries as Industry[];
    }
    const [selectedCountry, setSelectedCountry] = useState<number>(0);
    const [selectedCounty, setSelectedCounty] = useState<number>(0);
    const counties = useCounties(selectedCountry);
    const cities = useCities(selectedCounty);
    const phoneCode = usePhonecode(selectedCountry);
    const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCountry(parseInt(e.target.value));
    }
    const handleCountyChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCounty(parseInt(e.target.value));
    }
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const errors: any = useActionData();
    return (
        <Form method="post"
            className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4  max-w-lg mx-auto"
        >
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="name">Name</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="name" name="name"/>
                {errors?.name && <span className="text-red-500 text-xs italic" >{errors.name}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="website">Website</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="website" name="website" placeholder="http://www.mywebsite.com"/>
                {errors?.website && <span className="text-red-500 text-xs italic" >{errors.website}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="email" id="email" name="email"/>
                {errors?.email && <span className="text-red-500 text-xs italic" >{errors.email}</span>}
            </div>
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="country">Country</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="country"
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
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="county">County</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="county"
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
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="city">City</label>
                <div className="relative">
                    <select
                        name="city"
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
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
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="phone">Phone</label>
                {phoneCode &&
                    <div className="w-full  flex mb-4">
                        <input
                            className="
                                w-1/4
                                bg-gray-100
                                border border-gray-300
                                  text-gray-900 text-sm rounded-lg
                                   focus:ring-blue-500 focus:border-blue-500 
                                   block p-2.5 cursor-not-allowed
"
                            name="phoneCode" type="tel" value={phoneCode!.code} readOnly />
                        <input
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
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
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="industry">Industry</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="industry"
                    >
                        {industries && industries.map((industry) => (
                            <option
                                key={industry.id} value={industry.name}>{industry.name}</option>
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
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="foundationDate">Foundation Date</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="date" id="foundationDate" name="foundationDate"/>
                {errors?.foundationDate && <span className="text-red-500 text-xs italic" >{errors.foundationDate}</span>}
            </div>
            <div className="mb-4">
                <div className="flex items-start">
                    <div className="flex items-center h-5">
                        <input
                            name="terms"
                            type="checkbox"
                            className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300" />
                    </div>
                    <label
                        className="block text-gray-700 text-sm font-bold mb-2 ms-2"
                        htmlFor="terms">
                        I agree to the terms and conditions
                    </label>
                </div>
                {errors?.terms && <span className="text-red-500 text-xs italic">{errors.terms}</span>}
            </div>
            <div className="flex flex-col justify-center items-center">
                <PrimaryButton text="Sign up" />
            </div>
        </Form >
    );
}

export default RegisterEnterprise;
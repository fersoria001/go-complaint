import { useState } from "react";
import useCounties from "../lib/hooks/useCounties";
import useCities from "../lib/hooks/useCities";
import SelectIcon from "../components/icons/SelectIcon";
import PrimaryButton from "../components/buttons/PrimaryButton";
import { Route } from "../routes/sign-up";
import usePhonecode from "../lib/hooks/usePhonecode";
import { ErrorType, SignUpSchema } from "../lib/types";
import { syncParseSchema } from "../lib/parse_schema";
import { createUser } from "../lib/create_user";
import { useRouter } from "@tanstack/react-router";



//Check birthdate 
const SignUp: React.FC = () => {
    const countries = Route.useLoaderData();
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
        const form = document.getElementById('sign-up-form') as HTMLFormElement;
        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());
        const { data: parsed, errors } = syncParseSchema(data, SignUpSchema);
        if (Object.keys(errors).length > 0) {
            setErrors(errors);
            return;
        }
        const ok = await createUser(parsed);
        if (ok) {
            router.navigate({ to: '/sign-in' });
        } else {
            setErrors({ form: 'An error occurred while creating the user' });
        }
    }
    return (

        <form id="sign-up-form" className="block bg-white 
        shadow-md rounded px-8 pt-6 pb-8 mb-4  max-w-lg mx-auto">

            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="email"
                    type="email"
                    placeholder="Email"

                />
                {errors?.email && <span className="text-red-500 text-xs italic" >{errors.email}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="password">Password</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="password" type="password" placeholder="Password" />
                {errors?.password && <span className="text-red-500 text-xs italic">{errors.password}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="confirmPassword">Confirm Password</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="confirmPassword" type="password" placeholder="Confirm Password" />
                {errors?.confirmPassword && <span className="text-red-500 text-xs italic">{errors.confirmPassword}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="firstName">First Name</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                     text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="firstName" type="text" placeholder="First Name" />
                {errors?.firstName && <span className="text-red-500 text-xs italic">{errors.firstName}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="lastName">Last Name</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                     text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="lastName" type="text" placeholder="Last Name" />
                {errors?.lastName && <span className="text-red-500 text-xs italic">{errors.lastName}</span>}
            </div>
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="gender">Gender</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                         border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                          focus:outline-none focus:bg-white focus:border-gray-500"
                        name="gender">
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                        <option value="non-declared">Non-declared</option>
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
                {errors?.gender && <span className="text-red-500 text-xs italic">{errors.gender}</span>}
            </div>
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="pronoun">Pronoun</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                         border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                          focus:outline-none focus:bg-white focus:border-gray-500"
                        name="pronoun">
                        <option value={'she'}>She</option>
                        <option value={'he'}>He</option>
                        <option value={'they'}>They</option>
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
                {errors?.pronoun && <span className="text-red-500 text-xs italic">{errors.pronoun}</span>}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="birthDate">Birth Date</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3
                     text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="birthDate" type="date" />
                {errors?.birthDate && <span className="text-red-500 text-xs italic">{errors.birthDate}</span>}
            </div>

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
                        className="block appearance-none w-full bg-gray-200 border border-gray-200
                         text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
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
                                   block p-2.5 cursor-not-allowed"
                            name="phoneCode" type="tel" value={phoneCode} readOnly />
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
                {errors?.form && <span className="text-red-500 text-xs italic" >{errors.form}</span>}
            </div>

            <div  
            onMouseUp={handleSubmit}
            className="flex flex-col justify-center items-center">
                <PrimaryButton text="Sign up" />
            </div>
        </form>
    );
}
export default SignUp;
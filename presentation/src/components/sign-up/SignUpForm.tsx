'use client'
import { useState } from "react";
import SelectIcon from "../icons/SelectIcon";
import { useQuery, useSuspenseQuery } from "@tanstack/react-query";
import getGraphQLClient from "@/graphql/graphQLClient";
import countriesQuery from "@/graphql/queries/countriesQuery";
import countryStatesQuery from "@/graphql/queries/countryStatesQuery";
import citiesQuery from "@/graphql/queries/citiesQuery";
import { userSignUp } from "@/lib/actions/userSignUp";
import { useFormState, useFormStatus } from "react-dom";
import { z } from "zod";
import signUpSchema from "@/lib/validation/signUpSchema";
import InlineAlert from "../error/InlineAlert";

const initialState: Partial<z.inferFlattenedErrors<typeof signUpSchema>> | undefined = {}
const SignUpForm: React.FC = () => {
    const gqlClient = getGraphQLClient()
    const { data } = useSuspenseQuery({
        queryKey: ['countries'],
        queryFn: async () => gqlClient.request(countriesQuery),
    })
    const [countryId, setCountryId] = useState<number | undefined>(data?.countries?.at(0)?.id!)
    const [countryStateId, setCountryStateId] = useState<number | undefined>()

    const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setCountryId(parseInt(e.currentTarget.value))
    }
    const handleCountryStateChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setCountryStateId(parseInt(e.currentTarget.value))
    }
    const { data: countryStatesData } = useQuery({
        queryKey: ['countryStates', countryId],
        queryFn: async () => {
            const data = await gqlClient.request(countryStatesQuery, { id: countryId! })
            if (data.countryStates) {
                setCountryStateId(data.countryStates.at(0)?.id!)
            }
            return data
        },
        enabled: !!countryId,
    })
    const { data: citiesData } = useQuery({
        queryKey: ['cities', countryStateId],
        queryFn: async () => gqlClient.request(citiesQuery, { id: countryStateId! }),
        enabled: !!countryStateId,
    })
    const { pending } = useFormStatus()
    const [state, formAction] = useFormState(userSignUp, initialState)
    return (
        <form
            action={formAction}
            id="sign-up-form"
            className="block border-t bg-white rounded-md shadow-md px-8 pt-6 pb-8 my-4  max-w-lg mx-auto">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="userName">Email</label>
                <input
                    className="appearance-none focus:outline-none border rounded w-full py-2 px-3
                text-md lg:text-lg
                text-gray-700 leading-tight  focus:shadow-outline"
                    name="userName"
                    type="email"
                    placeholder="Email"
                    autoComplete="username"
                />
            </div>
            {state?.fieldErrors?.userName && <InlineAlert errors={state.fieldErrors.userName} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="password">Password</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    autoComplete="new-password"
                    name="password" type="password" placeholder="Password" />
            </div>
            {state?.fieldErrors?.password && <InlineAlert errors={state.fieldErrors.password} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="confirmPassword">Confirm Password</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    autoComplete="new-password"
                    name="confirmPassword" type="password" placeholder="Confirm password" />
            </div>
            {state?.fieldErrors?.confirmPassword && <InlineAlert errors={state.fieldErrors.confirmPassword} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="firstName">First Name</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="firstName" type="text" placeholder="First name" />

            </div>
            {state?.fieldErrors?.firstName && <InlineAlert errors={state.fieldErrors.firstName} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="lastName">Last Name</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="lastName" type="text" placeholder="Last name" />

            </div>
            {state?.fieldErrors?.lastName && <InlineAlert errors={state.fieldErrors.lastName} />}
            <div className="w-full  mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="genre">Genre</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                        text-md lg:text-lg
                     border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                      focus:outline-none focus:bg-white focus:border-gray-500"
                        name="genre"
                        defaultValue={"male"}>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                        <option value="non-declared">Non-declared</option>
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            {state?.fieldErrors?.genre && <InlineAlert errors={state.fieldErrors.genre} />}
            <div className="w-full  mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="pronoun">Pronoun</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                        text-md lg:text-lg
                     border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                      focus:outline-none focus:bg-white focus:border-gray-500"
                        name="pronoun"
                        defaultValue={"she"}>
                        <option value={'she'}>She</option>
                        <option value={'he'}>He</option>
                        <option value={'they'}>They</option>
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            {state?.fieldErrors?.pronoun && <InlineAlert errors={state.fieldErrors.pronoun} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="birthDate">Birth Date</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="birthDate" type="date" />
            </div>
            {state?.fieldErrors?.birthDate && <InlineAlert errors={state.fieldErrors.birthDate} />}
            <div className="w-full  mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="country">Country</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                        text-md lg:text-lg
                     border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                      focus:outline-none focus:bg-white focus:border-gray-500"
                        onChange={handleCountryChange}
                        name="countryId">
                        {
                            data.countries && data.countries.map((c) => {
                                if (c) {
                                    return (
                                        <option key={c.id} value={c.id!}>{c.name}</option>
                                    )
                                } else {
                                    return null
                                }
                            })
                        }
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            {state?.fieldErrors?.countryId && <InlineAlert errors={state.fieldErrors.countryId} />}
            <div className="w-full  mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="state">State</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200
                        text-md lg:text-lg text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        onChange={handleCountryStateChange}
                        name="countryStateId">
                        {
                            countryStatesData && countryStatesData.countryStates && countryStatesData.countryStates.map((c) => {
                                if (c) {
                                    return (
                                        <option key={c.id} value={c.id!}>{c.name}</option>
                                    )
                                } else {
                                    return null
                                }
                            })
                        }
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            {state?.fieldErrors?.countryStateId && <InlineAlert errors={state.fieldErrors.countryStateId} />}
            <div className="w-full  mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="city">City</label>
                <div className="relative">
                    <select
                        name="cityId"
                        className="block appearance-none w-full bg-gray-200 border border-gray-200
                        text-md lg:text-lg
                     text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                    >
                        {
                            citiesData && citiesData.cities && citiesData.cities.map((c) => {
                                if (c) {
                                    return (
                                        <option key={c.id} value={c.id!}>{c.name}</option>
                                    )
                                } else {
                                    return null
                                }
                            })
                        }
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            {state?.fieldErrors?.cityId && <InlineAlert errors={state.fieldErrors.cityId} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm lg:text-md font-bold mb-2"
                    htmlFor="phoneNumber">Phone</label>
                <div className="w-full  flex mb-4">
                    <input
                        className="appearance-none border rounded w-full py-2 px-3 text-gray-700 
                        text-md lg:text-lg
                        leading-tight focus:outline-none focus:shadow-outline"
                        name="phoneNumber" type="tel" placeholder="4321 567-8910" />
                </div>
            </div>
            {state?.fieldErrors?.phoneNumber && <InlineAlert errors={state.fieldErrors.phoneNumber} />}
            <div className="mb-4">
                <div className="flex items-start">
                    <div className="flex items-center h-5">
                        <input
                            name="terms"
                            type="checkbox"
                            className="w-4 h-4 lg:w-5 lg:h-5 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300" />
                    </div>
                    <label
                        className="block text-gray-700 text-sm lg:text-md font-bold mb-2 ms-2"
                        htmlFor="terms">
                        I agree to the terms and conditions
                    </label>
                </div>
            </div>
            {state?.fieldErrors?.terms && <InlineAlert errors={state.fieldErrors.terms} />}
            {state?.formErrors && <InlineAlert errors={state.formErrors} />}
            <div className="flex flex-col justify-center items-center">
                <button
                    className="bg-blue-500 px-5 py-3 rounded-md hover:bg-blue-700 text-white text-md lg:text-lg"
                    type="submit"
                    disabled={pending}>
                    Sign up
                </button>
            </div>
        </form>
    )
}
export default SignUpForm;
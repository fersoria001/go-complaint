'use client'
import getGraphQLClient from "@/graphql/graphQLClient";
import citiesQuery from "@/graphql/queries/citiesQuery";
import countryStatesQuery from "@/graphql/queries/countryStatesQuery";
import { useQuery, useSuspenseQueries } from "@tanstack/react-query";
import { useState } from "react";
import SelectIcon from "../icons/SelectIcon";
import countriesQuery from "@/graphql/queries/countriesQuery";
import industriesQuery from "@/graphql/queries/industriesQuery";
import { useFormState } from "react-dom";
import { registerEnterprise, RegisterEnterpriseFormState } from "@/lib/actions/graphqlActions";
import InlineAlert from "../error/InlineAlert";

const initialState: RegisterEnterpriseFormState = {}
const RegisterEnterpriseForm: React.FC = () => {
    const gqlClient = getGraphQLClient()
    const [countriesSuspenseQuery, industriesSuspenseQuery] = useSuspenseQueries({
        queries: [
            {
                queryKey: ['countries'],
                queryFn: async () => gqlClient.request(countriesQuery),
            },
            {
                queryKey: ['industries'],
                queryFn: async () => gqlClient.request(industriesQuery),
            }
        ]
    })

    const [countryId, setCountryId] = useState<number | undefined>(countriesSuspenseQuery.data?.countries?.at(0)?.id!)
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
    const [state, formAction] = useFormState(registerEnterprise, initialState)
    return (
        <form
            action={formAction}
            id="register-enterprise-form"
            className="block border-t bg-white rounded-md shadow-md px-8 pt-6 pb-8 my-4  max-w-lg mx-auto">
            {state?.formErrors && <InlineAlert errors={state.formErrors} />}
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="name">Name</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="name" name="name" />
                {state?.fieldErrors?.name && <InlineAlert errors={state.fieldErrors.name} />}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="website">Website</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                  text-md lg:text-lg
                  text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="text" id="website" name="website" placeholder="http://www.mywebsite.com" />
                {state?.fieldErrors?.website && <InlineAlert errors={state.fieldErrors.website} />}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="email" id="email" name="email" />
                {state?.fieldErrors?.email && <InlineAlert errors={state.fieldErrors.email} />}
            </div>
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
                            countriesSuspenseQuery.data.countries && countriesSuspenseQuery.data.countries.map((c) => {
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
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="phone">Phone</label>

                <input
                    className="appearance-none border rounded w-full py-2 px-3
                    text-md lg:text-lg
                    text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="phoneNumber" type="tel" placeholder="4321 567-8910" />
                {state?.fieldErrors?.phoneNumber && <InlineAlert errors={state.fieldErrors.phoneNumber} />}
            </div>
            <div className="w-full  mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="industryId">Industry</label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                         text-md lg:text-lg
                      border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                       focus:outline-none focus:bg-white focus:border-gray-500"
                        name="industryId"
                    >
                        {
                            industriesSuspenseQuery.data.industries && industriesSuspenseQuery.data.industries.map((c) => {
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
                {state?.fieldErrors?.industryId && <InlineAlert errors={state.fieldErrors.industryId} />}
            </div>
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm md:text-xl font-bold mb-2"
                    htmlFor="foundationDate">Foundation Date</label>
                <input
                    className="appearance-none border rounded w-full py-2 px-3
                   text-md lg:text-lg
                   text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    type="date" id="foundationDate" name="foundationDate" />
                {state?.fieldErrors?.foundationDate && <InlineAlert errors={state.fieldErrors.foundationDate} />}
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
                {state?.fieldErrors?.terms && <InlineAlert errors={state.fieldErrors.terms} />}
            </div>
            <div className="flex flex-col justify-center items-center">
                <button
                    className="bg-blue-500 px-5 py-3 rounded-md hover:bg-blue-700 text-white text-md lg:text-lg"
                    type="submit">
                    Register enterprise
                </button>
            </div>
        </form >
    )
}
export default RegisterEnterpriseForm;
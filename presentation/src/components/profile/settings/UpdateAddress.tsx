import InlineAlert from "@/components/error/InlineAlert"
import CheckIcon from "@/components/icons/CheckIcon"
import SelectIcon from "@/components/icons/SelectIcon"
import { Country, UpdateUserAddress, UserDescriptor } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import citiesQuery from "@/graphql/queries/citiesQuery"
import countryStatesQuery from "@/graphql/queries/countryStatesQuery"
import { updateUserAdress } from "@/lib/actions/graphqlActions"
import { useQuery, useMutation } from "@tanstack/react-query"
import { useState } from "react"
import { z } from "zod"

const addressSchema = z.object({
    country: z.number(),
    county: z.number(),
    city: z.number(),
})

interface Props {
    descriptor: UserDescriptor
    countries: Country[]
}

const UpdateAddress: React.FC<Props> = ({ descriptor, countries }: Props) => {
    const [errors, setErrors] = useState<z.inferFlattenedErrors<typeof addressSchema> | undefined>()
    const gqlClient = getGraphQLClient()
    const [countryId, setCountryId] = useState<number | undefined>(countries.at(0)?.id!)
    const [countryStateId, setCountryStateId] = useState<number | undefined>()
    const [cityId, setCityId] = useState<number | undefined>()

    const handleCountryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setCountryId(parseInt(e.currentTarget.value))
    }

    const handleCountryStateChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setCountryStateId(parseInt(e.currentTarget.value))
    }

    const handleCityChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setCityId(parseInt(e.currentTarget.value))
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

    const updateMutation = useMutation({
        mutationFn: async ({ newCountryId, newCountyId, newCityId }: Omit<UpdateUserAddress, "userId">) => {
            await updateUserAdress({
                userId: descriptor.id,
                newCountryId,
                newCountyId,
                newCityId
            })
        }
    })

    const handleUpdate = async () => {
        const { data, success, error } = addressSchema.safeParse({
            country: countryId,
            county: countryStateId,
            city: cityId
        })
        if (!success) {
            setErrors(error.flatten());
            return
        }
        updateMutation.mutate({
            newCountryId: data.country,
            newCountyId: data.county,
            newCityId: data.city
        })
    }

    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto  mb-6">
            <div className="w-full  md:mb-0">
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
                {errors?.fieldErrors.country && <InlineAlert errors={errors.fieldErrors.country} />}
            </div>

            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="county">County
                </label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 
                        px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="county"
                        onChange={handleCountryStateChange} defaultValue={"Select a country first"}>
                        {countryStatesData?.countryStates && countryStatesData.countryStates.map((county) => (
                            <option key={county.id} value={county.id}>{county.name}</option>
                        )) || <option disabled>Select a country first</option>}
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
                {errors?.fieldErrors.county && <InlineAlert errors={errors.fieldErrors.county} />}
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
                        {citiesData?.cities && citiesData.cities.map((city) => (
                            <option key={city.id} value={city.id}>{city.name}</option>
                        )) || <option disabled>Select a county first</option>}
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                    {errors?.fieldErrors.city && <InlineAlert errors={errors.fieldErrors.city} />}
                </div>
            </div>
            <div className="self-end flex mt-2">
                {updateMutation.isSuccess && <CheckIcon className="w-6 h-6 my-auto fill-blue-300" />}
                <button
                    type="button"
                    onClick={() => handleUpdate()}
                    className="px-6 py-3 bg-blue-500 hover:bg-blue-600 rounded-md text-white font-bold">
                    Submit
                </button>
            </div>
            {
                updateMutation.isError
                && <InlineAlert errors={[updateMutation.error.message]} />
            }
        </div>
    )
}

export default UpdateAddress;
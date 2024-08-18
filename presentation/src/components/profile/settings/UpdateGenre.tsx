'use client'
import CheckIcon from "@/components/icons/CheckIcon"
import SelectIcon from "@/components/icons/SelectIcon"
import { UserDescriptor } from "@/gql/graphql"
import { changeUserGenre, changeUserPronoun } from "@/lib/actions/graphqlActions"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { useState } from "react"

interface Props {
    descriptor: UserDescriptor
}

const UpdateGenre: React.FC<Props> = ({ descriptor }: Props) => {
    const [genre, setGenre] = useState<string>(descriptor.genre)
    const [pronoun, setPronoun] = useState<string>(descriptor.pronoun)
    const genders = ["Female", "Male", "Non-declared"]
    const pronouns = ["He", "She", "They"]
    const queryClient = useQueryClient()
    const updateMutation = useMutation({
        mutationFn: async () => {
            await changeUserGenre({
                newGenre: genre,
                userId: descriptor.id
            })
            await changeUserPronoun({
                newPronoun: pronoun,
                userId: descriptor.id
            })
        },
        onSuccess: () => queryClient.refetchQueries({ queryKey: ['userDescriptor'] })
    })
    const handleGenreChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setGenre(e.target.value)
    }
    const handlePronounChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setPronoun(e.target.value)
    }
    return (
        <div className="flex flex-col w-full md:w-2/3 mx-auto">
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="gender">
                    Gender
                </label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                 border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                  focus:outline-none focus:bg-white focus:border-gray-500"
                        name="gender"
                        onChange={handleGenreChange}
                        value={genre}>
                        {
                            genders.map((gender) => {
                                return <option
                                    key={gender} value={gender}>{gender}
                                </option>
                            })
                        }
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            <div className="w-full mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="pronoun">
                    Pronoun
                </label>
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border
                 border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight
                  focus:outline-none focus:bg-white focus:border-gray-500"
                        name="pronoun"
                        value={pronoun}
                        onChange={handlePronounChange}>
                        {
                            pronouns.map((pronoun) => {
                                return <option
                                    key={pronoun} value={pronoun}>{pronoun}
                                </option>
                            })
                        }
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
            <div className="mt-5 self-end flex">
                {updateMutation.isSuccess && <CheckIcon className="w-6 h-6 my-auto fill-blue-300" />}
                <button
                    type="button"
                    onClick={() => updateMutation.mutate()}
                    className="px-6 py-3 bg-blue-500 hover:bg-blue-600 rounded-md text-white font-bold">
                    Submit
                </button>
            </div>
        </div>
    )
}

export default UpdateGenre;
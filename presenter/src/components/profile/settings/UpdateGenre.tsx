import { useState } from "react";
import AcceptBtn from "../../buttons/AcceptBtn";
import SelectIcon from "../../icons/SelectIcon";
import { UserDescriptor } from "../../../lib/types";
import { updateGenre } from "./settings_lib";

interface Props {
    descriptor: UserDescriptor
}
const UpdateGenre: React.FC<Props> = ({ descriptor }: Props) => {
    const [genre, setGenre] = useState<string>(descriptor.gender)
    const [pronoun, setPronoun] = useState<string>(descriptor.pronoun)
    const [reset, setReset] = useState<boolean>(false)
    const genders = ["Female", "Male", "Non-declared"]
    const pronouns = ["He", "She", "They"]
    const handleUpdate = async (): Promise<boolean> => {
        return await updateGenre(genre, pronoun)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    const handleGenreChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setGenre(e.target.value)
        setReset(true)
    }
    const handlePronounChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        setPronoun(e.target.value)
        setReset(true)
    }   
    return (
        <div className="flex flex-col">
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="gender">Gender</label>
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
            <div className="w-full  mb-6 md:mb-0">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="pronoun">Pronoun</label>
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
            <div className="mt-5 self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset}  variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdateGenre;
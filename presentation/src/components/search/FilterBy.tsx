import { Key } from "react"
import SelectIcon from "../icons/SelectIcon"

interface Props {
    options: {
        id: Key
        name: string,
        value: string | number
    }[]
    callback?: (value: string) => void
}
const FilterBy: React.FC<Props> = ({ options, callback = (v: string) => { } }: Props) => {
    return (
        <div className="flex gap-5 overflow-x-auto max-w-sm pt-2 items-center">
            <p className="text-gray-700 text-md font-bold ps-2 shrink-0">Filter by:</p>
            <div className="w-full">
                <div className="relative">
                    <select
                        className="block appearance-none w-full bg-gray-200 border text-md lg:text-lg border-gray-200 text-gray-700 py-2 px-4 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                        name="filterBy"
                        onChange={(e) => { callback(e.currentTarget.value) }}
                    >
                        {
                            options.map((option) => {
                                return (
                                    <option key={option.id} value={option.value}>
                                        {option.name}
                                    </option>
                                )
                            })
                        }
                    </select>
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                        <SelectIcon />
                    </div>
                </div>
            </div>
        </div>
    )
}
export default FilterBy
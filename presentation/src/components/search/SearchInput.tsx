'use client'
interface Props {
    placeholder: string
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
}
const SearchInput: React.FC<Props> = ({ placeholder, onChange }) => {
    return (
        <div className="flex px-2.5 mb-2">
            <input
                className="bg-gray-50 border border-gray-300 text-gray-700 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full  p-2.5"
                type="text" placeholder={placeholder}
                onChange={onChange} />
        </div>
    )
}
export default SearchInput;
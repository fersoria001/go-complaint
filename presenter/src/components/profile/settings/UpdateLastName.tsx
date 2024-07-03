import { z } from "zod";
import { ErrorType } from "../../../lib/types";
import { useState } from "react";
import { updateLastName } from "./settings_lib";
import AcceptBtn from "../../buttons/AcceptBtn";

const validName = z
    .string()
    .min(2, { message: "Last name must be at least 2 characters long" })
    .max(50, { message: "Last name must be at most 50 characters long" })
const UpdateLastName: React.FC = () => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [lastName, setFirstName] = useState<string>("")
    const [reset, setReset] = useState<boolean>(false)
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setFirstName(e.target.value)
        setReset(true)
    }
    const handleUpdate = async () => {
        const result = validName.safeParse(lastName)
        if (!result.success) {
            setErrors({ lastName: result.error.errors[0].message })
            return false
        }
        return await updateLastName(lastName)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    return (
        <div className="flex flex-col">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="lastName">Last name</label>
                <input
                    onChange={handleChange}
                    value={lastName}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="lastName" type="text" />
                {errors?.lastName && <span className="text-red-500 text-xs italic">{errors.lastName}</span>}
            </div>
            <div className="self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdateLastName;
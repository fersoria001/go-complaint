import { z } from "zod";
import { ErrorType } from "../../../lib/types";
import { useState } from "react";
import { updateFirstName } from "./settings_lib";
import AcceptBtn from "../../buttons/AcceptBtn";

const validName = z
    .string()
    .min(2, { message: "First name must be at least 2 characters long" })
    .max(50, { message: "First name must be at most 50 characters long" })
const UpdateFirstName: React.FC = () => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [firstName, setFirstName] = useState<string>("")
    const [reset, setReset] = useState<boolean>(false)
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setFirstName(e.target.value)
        setReset(true)
    }
    const handleUpdate = async () => {
        const result = validName.safeParse(firstName)
        if (!result.success) {
            setErrors({ firstName: result.error.errors[0].message })
            return false
        }
        return await updateFirstName(firstName)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    return (
        <div className="flex flex-col">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="firstName">First Name</label>
                <input
                    onChange={handleChange}
                    value={firstName}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="firstName" type="text" />
                {errors?.firstName && <span className="text-red-500 text-xs italic">{errors.firstName}</span>}
            </div>
            <div className="self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdateFirstName;
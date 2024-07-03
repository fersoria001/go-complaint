import { z } from "zod";
import { ErrorType } from "../../../lib/types";
import { useState } from "react";

import AcceptBtn from "../../buttons/AcceptBtn";
import { updateEmail } from "./settings_lib";

const validEmail = z.string().email({ message: "Please enter a valid email" })
interface Props {
    enterpriseID: string
}
const UpdateEmail: React.FC<Props> = ({ enterpriseID }: Props) => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [email, setEmail] = useState<string>("")
    const [reset, setReset] = useState<boolean>(false)
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setEmail(e.target.value)
        setReset(true)
    }
    const handleUpdate = async () => {
        const result = validEmail.safeParse(email)
        if (!result.success) {
            setErrors({ email: result.error.errors[0].message })
            return false
        }
        return await updateEmail(enterpriseID, email)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    return (
        <div className="flex flex-col">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    onChange={handleChange}
                    value={email}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="email" type="text" />
                {errors?.email && <span className="text-red-500 text-xs italic">{errors.email}</span>}
            </div>
            <div className="self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdateEmail;
import { z } from "zod";
import { ErrorType } from "../../../lib/types";
import { useState } from "react";

import AcceptBtn from "../../buttons/AcceptBtn";
import { updateWebsite } from "./settings_lib";

const validWebsite = z.string().url({
    message: "Please enter a valid website e.g: http://www.mywebsite.com",
})
interface Props {
    enterpriseID: string
}
const UpdateWebsite: React.FC<Props> = ({ enterpriseID }: Props) => {
    const [errors, setErrors] = useState<ErrorType>({})
    const [website, setWebsite] = useState<string>("")
    const [reset, setReset] = useState<boolean>(false)
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrors({})
        setWebsite(e.target.value)
        setReset(true)
    }
    const handleUpdate = async () => {
        const result = validWebsite.safeParse(website)
        if (!result.success) {
            setErrors({ website: result.error.errors[0].message })
            return false
        }
        return await updateWebsite(enterpriseID, website)
    }
    const handleCleanUp = () => {
        setReset(false)
    }
    return (
        <div className="flex flex-col">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="website">Website</label>
                <input
                    onChange={handleChange}
                    value={website}
                    className="shadow appearance-none border rounded w-full py-2 px-3
             text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    name="website" type="text" />
                {errors?.website && <span className="text-red-500 text-xs italic">{errors.website}</span>}
            </div>
            <div className="self-end">
                <AcceptBtn cleanUp={handleCleanUp} reset={reset} variant="primary" text="Submit" callback={handleUpdate} />
            </div>
        </div>
    )
}

export default UpdateWebsite;
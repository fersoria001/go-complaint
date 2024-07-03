import { useState } from "react";
import { login } from "../../lib/login";
import { ErrorType } from "../../lib/types";
import PrimaryButton from "../buttons/PrimaryButton";
import { useRouter } from "@tanstack/react-router";

const LoginForm: React.FC = () => {
    const router = useRouter();
    const [errors, setErrors] = useState<ErrorType>({});
    const handleLogin = async () => {
        let confirmationCode = "";
        for (let i = 0; i < 7; i++) {
            const c = document.getElementById(`code${i}`) as HTMLInputElement;
            confirmationCode += c.value;
        }
        const errors = await login(confirmationCode);
        if (Object.keys(errors).length === 0) {
            router.navigate({ to: "/profile" });
        }
        setErrors(errors);
    }

    return <div className="h-screen px-2 md:px-0">
        <form className="bg-white shadow-md rounded px-4 pt-6 pb-8 mb-4 mt-24 border max-w-lg mx-auto">
            <div className="mb-4">
                {errors?.form && <span className="mx-auto text-red-500 text-xs italic" >{errors.form}</span>}
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="confirmationCode">Enter the confirmation code</label>
                <div className="flex gap-0.5">
                    {
                        Array.from({ length: 7 }).map((_, i) => (
                            <input
                                key={i}
                                className="min-w-[24px]
                            text-sm md:text-xl shadow appearance-none border rounded py-2 px-3
                         text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                                type="text" id={`code${i}`} name={`code${i}`} maxLength={1} />
                        ))
                    }
                </div>
                {errors?.confirmationCode && <span className="text-red-500 text-xs mx-auto italic" >{errors.confirmationCode}</span>}
            </div>
            <span className="text-center" onMouseUp={handleLogin}>
                <PrimaryButton text="Confirm" />
            </span>
        </form>
    </div>
}

export default LoginForm;
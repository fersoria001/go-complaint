import { useState } from "react";
import PrimaryButton from "../components/buttons/PrimaryButton";
import { ErrorType, } from "../lib/types";
import { signIn } from "../lib/sign_in";
import { useRouter } from "@tanstack/react-router";

export const SignInPage: React.FC = () => {
    const [errors, setErrors] = useState<ErrorType>({});
    const router = useRouter();
    const handleSignIn = async () => {
        const email = document.getElementById("email") as HTMLInputElement;
        const password = document.getElementById("password") as HTMLInputElement;
        const rememberMe = document.getElementById("rememberMe") as HTMLInputElement;
        const errors = await signIn(email.value, password.value, rememberMe.checked);
        if (Object.keys(errors).length === 0) {
            router.navigate({ to: "/confirmation" });
        }
        setErrors(errors);
    }
    return (
        <div className="h-screen px-2 md:px-0">
            <form className="bg-white border shadow-md rounded px-8 pt-6 pb-8 mt-24 max-w-lg mx-auto">
                <div className="mb-4">
                    {errors?.form && <span className="text-red-500 text-xs italic" >{errors.form}</span>}
                    <label
                        className="block text-gray-700 text-sm font-bold mb-2"
                        htmlFor="email">Email</label>
                    <input
                        className="shadow appearance-none border rounded w-full py-2 px-3
                         text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        type="email" id="email" name="email" />
                    {errors?.email && <span className="text-red-500 text-xs italic" >{errors.email}</span>}
                </div>
                <div className="mb-4">
                    <label
                        className="block text-gray-700 text-sm font-bold mb-2"
                        htmlFor="password">Password</label>
                    <input
                        className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        type="password" id="password" name="password" />
                    {errors?.password && <span className="text-red-500 text-xs italic">{errors.password}</span>}
                </div>
                <div className="mb-4">
                    <div className="flex items-start">
                        <div className="flex items-center h-5">
                            <input
                                className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300"
                                type="checkbox" id="rememberMe" name="rememberMe"
                                defaultValue={'off'}
                            />
                        </div>
                        {errors?.rememberMe && <span className="text-red-500 text-xs italic">{errors.rememberMe}</span>}
                        <label
                            className="block text-gray-700 text-sm font-bold mb-2 ms-2"

                            htmlFor="rememberMe">Remember me</label>
                    </div>
                </div>
                <span className="" onMouseUp={handleSignIn}>
                    <PrimaryButton text="Sign in" />
                </span>
            </form>
        </div>
    );
}

export default SignInPage;
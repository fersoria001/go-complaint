import { Form, useActionData } from "react-router-dom";
import PrimaryButton from "../components/buttons/PrimaryButton";

function SignIn() {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const errors: any = useActionData();
    return (
        <Form
            className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4  max-w-lg mx-auto"
            method="post">
            <div className="mb-4">
                <label
                    className="block text-gray-700 text-sm font-bold mb-2"
                    htmlFor="email">Email</label>
                <input
                    className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
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
            <PrimaryButton text="Sign in" />

        </Form>
    );
}

export default SignIn;
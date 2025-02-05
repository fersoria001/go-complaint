'use server'
import axios, { AxiosInstance } from "axios"
import confirmationCodeSchema from "../validation/confirmationCodeSchema"
import signInSchema from "../validation/signInSchema"
import { redirect } from "next/navigation"
import { z } from "zod"
import querystring from 'node:querystring'
import { cookies } from "next/headers"

const axiosInstance: AxiosInstance | undefined = undefined
function getAxiosInstance() {
    if (axiosInstance) {
        return axiosInstance
    }
    const instance = axios.create({
        withCredentials: true
    });
    return instance
}

export type SignInFormState = Partial<z.inferFlattenedErrors<typeof signInSchema>>
export async function userSignIn(prevState: SignInFormState, fd: FormData): Promise<SignInFormState> {
    const { data, success, error } = signInSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    const url = process.env.SIGN_IN_ENDPOINT
    if (!url) {
        throw new Error("sign in endpoint not defined in process env")
    }
    try {
        const response = await getAxiosInstance().post(url, data)
        if (!response.headers["set-cookie"] || !response.headers["set-cookie"][0]) {
            throw new Error("response to sign in: cookie is not present in the set-cookie header")
        }
        const strCookie = response.headers["set-cookie"][0]
        const parsed = querystring.parse(strCookie, "; ")
        const cookie: any = { ...parsed }
        console.log(cookie)
        cookies().set({
            name: 'jwt',
            value: cookie.jwt,
            domain: cookie.domain,
            maxAge: cookie.maxAge,
            secure: cookie.secure,
            httpOnly: cookie.HttpOnly,
            path: cookie.Path,
        })
    } catch (e: any) {
        if (e.response.data) {
            let message = e.response.data
            if (e.response.data.includes("crypto/bcrypt: hashedPassword is not the hash of the given password")) { message = "password did not match" }
            if (e.response.data.includes("no rows")) { message = "the account doesn't exist"}
            return {
                formErrors: [message],
                fieldErrors: {}
            }
        }
        console.error(e)
    }
    redirect("/confirm-sign-in")
}

export type ConfirmSignInFormState = Partial<z.inferFlattenedErrors<typeof confirmationCodeSchema>>
export async function confirmSignIn(prevState: ConfirmSignInFormState, fd: FormData): Promise<ConfirmSignInFormState> {
    const { data, success, error } = confirmationCodeSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    try {
        const url = process.env.CONFIRM_SIGN_IN_ENDPOINT
        if (!url) {
            throw new Error("sign in endpoint not defined in process env")
        }
        const jwtCookie = cookies().get("jwt")
        if (!jwtCookie) {
            throw new Error("sign in cookie is not previously stored", jwtCookie)
        }
        const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
        const response = await getAxiosInstance().post(url, data, { headers: { Cookie: strCookie } })
        if (!response.headers["set-cookie"] || !response.headers["set-cookie"][0]) {
            throw new Error("response to confirm sign-in: cookie is not present in the set-cookie header")
        }

        const responseCookie = response.headers["set-cookie"][0]
        const parsed = querystring.parse(responseCookie, "; ")
        const cookie: any = { ...parsed }
        const environment = process.env.NEXT_PUBLIC_END_MODE
        cookies().set({
            name: 'jwt',
            value: cookie.jwt,
            domain: cookie.domain,
            maxAge: cookie.maxAge,
            secure: environment == "PROD" ? true : false,
            httpOnly: environment == "PROD" ? false : true,
            path: cookie.Path,
        })
    } catch (e: any) {
        if (e.response?.data) {
            const msg: string = e.response.data
            return {
                formErrors: [msg],
                fieldErrors: {},
            }
        }
        console.error(e)
    }
    redirect("/")
}

export async function logout() {
    cookies().delete("jwt")
    cookies().delete("alias")
    redirect("/")
}



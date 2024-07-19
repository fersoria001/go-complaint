'use server'
import axios from "axios"
import confirmationCodeSchema from "../validation/confirmationCodeSchema"
import signInSchema from "../validation/signInSchema"
import { redirect } from "next/navigation"
import { z } from "zod"
import querystring from 'node:querystring'
import { cookies } from "next/headers"
import getGraphQLClient from "@/graphql/graphQLClient"
import userDescriptorQuery from "@/graphql/queries/userDescriptorQuery"
import { UserDescriptor } from "@/gql/graphql"

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
        const response = await axios.post(url, data, {
            withCredentials: true,
        })
        if (!response.headers["set-cookie"] || !response.headers["set-cookie"][0]) {
            throw new Error("response to sign in: cookie is not present in the set-cookie header")
        }
        const strCookie = response.headers["set-cookie"][0]
        const parsed = querystring.parse(strCookie, "; ")
        const cookie: any = { ...parsed }
        cookies().set({
            name: 'jwt',
            value: cookie.jwt,
            httpOnly: cookie.HttpOnly ? true : false,
            path: cookie.Path,
            expires: cookie.Expires ? new Date(cookie.Expires) : new Date()
        })
    } catch (e: any) {
        if (e.response.data) {
            let message = e.response.data
            if (e.response.data.includes("crypto/bcrypt: hashedPassword is not the hash of the given password")) { message = "password did not match" }
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
        const response = await axios.post(url, data, { withCredentials: true, headers: { Cookie: strCookie } })
        if (!response.headers["set-cookie"] || !response.headers["set-cookie"][0]) {
            throw new Error("response to confirm sign-in: cookie is not present in the set-cookie header")
        }
        const responseCookie = response.headers["set-cookie"][0]
        const parsed = querystring.parse(responseCookie, "; ")
        const cookie: any = { ...parsed }
        cookies().set({
            name: 'jwt',
            value: cookie.jwt,
            httpOnly: cookie.HttpOnly ? true : false,
            path: cookie.Path,
            expires: cookie.Expires ? new Date(cookie.Expires) : new Date()
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
    redirect("/profile")
}

export async function logout() { }

export async function recoverPassword(prevState: any, fd: FormData) {
    const justTheEmail = signInSchema.pick({ userName: true })
    const { data, success, error } = justTheEmail.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    return {
        formErrors: ["the provided email is not registered"],
        fieldErrors: {}
    }
}


export async function getSession(): Promise<UserDescriptor | undefined> {
    const jwtCookie = cookies().get("jwt")
    if (!jwtCookie) {
        return undefined
    }
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    try {
        const response = await gqlClient.request(userDescriptorQuery)
        return response.userDescriptor
    } catch (e: any) {
        return undefined
    }

}
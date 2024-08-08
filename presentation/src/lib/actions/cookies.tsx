'use server'

import { cookies } from "next/headers"

export async function getCookie(name: string): Promise<string | undefined> {
    const res = cookies().get(name)
    return res?.value
}

export async function setCookie(name: string, value: string): Promise<string | undefined> {
    const oneHour = 3600 * 8
    const res = cookies().set({
        name: name,
        value: value,
        httpOnly: process.env.ENV_MODE == "prod",
        path: "/",
        maxAge: oneHour
    })
    return res.get(name)?.value
}

export async function removeCookie(name: string): Promise<void> {
    cookies().delete(name)
}
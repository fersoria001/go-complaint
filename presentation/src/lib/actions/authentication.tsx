'use server'

import confirmationCodeSchema from "../validation/confirmationCodeSchema"
import signInSchema from "../validation/signInSchema"

export async function userSignIn(prevState: any, fd: FormData) {
    const { data, success, error } = signInSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    return {
        formErrors: ["username does not exists"],
        fieldErrors: {},
    }
}

export async function confirmSignIn(prevState: any, fd: FormData) {
    const { data, success, error } = confirmationCodeSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    return {
        formErrors: ["the confirmation code you have provided is wrong"],
        fieldErrors: {}
    }
}


export async function logout() { }

export async function recoverPassword(prevState: any, fd: FormData) {
    const justTheEmail = signInSchema.pick({ email: true })
    const { data, success, error } = justTheEmail.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    return {
        formErrors: ["the provided email is not registered"],
        fieldErrors: {}
    }
}
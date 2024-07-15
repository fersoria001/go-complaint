'use server'

import getGraphQLClient from "@/graphql/graphQLClient";
import signUpSchema from "../validation/signUpSchema";
import createUserMutation from "@/graphql/mutations/createUserMutation";


export async function userSignUp(prevState: any, fd: FormData) {
    const { data, success, error } = signUpSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    try {
        const response = await getGraphQLClient().request(
            createUserMutation, data)
    } catch (e: any) {
        return {
            formErrors: ["user already exists or server error"],
            fieldErrors: {}
        }
    }
}


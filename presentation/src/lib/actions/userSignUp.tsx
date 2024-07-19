'use server'

import getGraphQLClient from "@/graphql/graphQLClient";
import signUpSchema from "../validation/signUpSchema";
import { z } from "zod";
import createUserMutation from "@/graphql/mutations/createUserMutation";
import { CreateUser } from "@/gql/graphql";
import { redirect } from "next/navigation";


export async function userSignUp(prevState: Partial<z.inferFlattenedErrors<typeof signUpSchema>> | undefined, fd: FormData) {
    const { data, success, error } = signUpSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    try {
        const newUser: Partial<z.infer<typeof signUpSchema>> = { ...data }
        delete newUser.confirmPassword
        delete newUser.terms
        const userName = await getGraphQLClient().request(createUserMutation, { input: newUser as CreateUser })
    } catch (e: any) {
        let msg: string = e.message;
        if (e.message.includes("SQLSTATE 23505")) {
            msg = "user already exists"
        }
        return {
            formErrors: [msg],
            fieldErrors: {}
        }
    }
    redirect("/sign-up?success=1")
}


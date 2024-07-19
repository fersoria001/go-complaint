'use server'
import { ComplaintInfo, CreateEnterprise } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import complaintsInfoQuery from "@/graphql/queries/complaintsInfoQuery"
import { cookies } from "next/headers"
import { z } from "zod"
import registerEnterpriseSchema from "../validation/registerEnterpriseSchema"
import createEnterpriseMutation from "@/graphql/mutations/createEnterpriseMutation"
import { redirect } from "next/navigation"


function gqlClientWithCookie() {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    return gqlClient
}

export async function getComplaintsInfo(id: string): Promise<ComplaintInfo> {
    const result = await gqlClientWithCookie().request(complaintsInfoQuery, { id: id })
    return result.complaintsReceivedInfo
}

export type RegisterEnterpriseFormState = Partial<z.inferFlattenedErrors<typeof registerEnterpriseSchema>>
export async function registerEnterprise(state: RegisterEnterpriseFormState | undefined, fd: FormData) {
    const { data, success, error } = registerEnterpriseSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    try {
        const newEnterprise: Partial<z.infer<typeof registerEnterpriseSchema>> = { ...data }
        delete newEnterprise.terms
        const enterpriseId = await gqlClientWithCookie().request(createEnterpriseMutation, { input: newEnterprise as CreateEnterprise })
    } catch (e: any) {
        let msg: string = e.message;
        if (e.message.includes("SQLSTATE 23505")) {
            msg = "enterprise name is in use"
        }
        return {
            formErrors: [msg],
            fieldErrors: {}
        }
    }
    redirect("/enterprises")
}


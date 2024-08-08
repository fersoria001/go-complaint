'use server'

import { SendComplaint, DescribeComplaint, Complaint, CreateNewComplaint, ComplaintsInfo, CreateEnterprise } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import createEnterpriseMutation from "@/graphql/mutations/createEnterpriseMutation"
import createNewComplaintMutation from "@/graphql/mutations/createNewComplaintMutation"
import describeComplaintMutation from "@/graphql/mutations/describeComplaintMutation"
import sendComplaintMutation from "@/graphql/mutations/sendComplaintMutation"
import complaintsInfoQuery from "@/graphql/queries/complaintsInfoQuery"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import { z } from "zod"
import describeComplaintSchema from "../validation/describeComplaintSchema"
import registerEnterpriseSchema from "../validation/registerEnterpriseSchema"
import sendComplaintSchema from "../validation/sendComplaintSchema"



function gqlClientWithCookie() {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    return gqlClient
}

export type SendComplaintFormState = Partial<z.inferFlattenedErrors<typeof sendComplaintSchema>>
export async function sendComplaint(state: SendComplaintFormState | undefined, fd: FormData) {
    const { data, success, error } = sendComplaintSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    try {
       await gqlClientWithCookie().request(sendComplaintMutation, { input: data as SendComplaint })
    } catch (e: any) {
        let msg: string = e.message
        return {
            formErrors: [msg],
            fieldErrors: {}
        }
    }
    redirect(`/complaints`)
}

export type DescribeComplaintFormState = Partial<z.inferFlattenedErrors<typeof describeComplaintSchema>>
export async function describeComplaint(state: DescribeComplaintFormState | undefined, fd: FormData) {
    const { data, success, error } = describeComplaintSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return error.flatten()
    }
    try {
        await gqlClientWithCookie().request(describeComplaintMutation, { input: data as DescribeComplaint })
    } catch (e: any) {
        let msg: string = e.message
        return {
            formErrors: [msg],
            fieldErrors: {}
        }
    }
    redirect(`/complaints/send-complaint?step=3&id=${data.complaintId}`)
}
export async function createNewComplaint(receiverId: string): Promise<Complaint> {
    const authorId = cookies().get("alias")?.value as string
    const result = await gqlClientWithCookie().request(createNewComplaintMutation, { input: { authorId, receiverId } as CreateNewComplaint })
    return result.createNewComplaint
}

export async function getComplaintsInfo(id: string): Promise<ComplaintsInfo> {
    const result = await gqlClientWithCookie().request(complaintsInfoQuery, { id: id })
    return result.complaintsInfo
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


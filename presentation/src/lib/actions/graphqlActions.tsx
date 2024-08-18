'use server'

import { SendComplaint, DescribeComplaint, Complaint, CreateNewComplaint, CreateEnterprise, RateComplaint, InviteToProject, AcceptHiringInvitation, RejectHiringInvitation, HireEmployee, CancelHiringProcess, PromoteEmployee, FireEmployee, AddFeedbackComment, AddFeedbackReply, RemoveFeedbackReply, RemoveFeedbackComment, EndFeedback, FindEnterpriseChat, CreateEnterpriseChat, UserDescriptor, ChangePassword, ChangeUserGenre, ChangeUserPronoun, ChangeUserFirstName, ChangeUserLastName, ChangeUserPhone, UpdateUserAddress, ChangeEnterpriseAddress, ChangeEnterprisePhone, ChangeEnterpriseEmail, ChangeEnterpriseWebsite } from "@/gql/graphql"
import getGraphQLClient from "@/graphql/graphQLClient"
import createEnterpriseMutation from "@/graphql/mutations/createEnterpriseMutation"
import createNewComplaintMutation from "@/graphql/mutations/createNewComplaintMutation"
import describeComplaintMutation from "@/graphql/mutations/describeComplaintMutation"
import sendComplaintMutation from "@/graphql/mutations/sendComplaintMutation"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import { z } from "zod"
import describeComplaintSchema from "../validation/describeComplaintSchema"
import registerEnterpriseSchema from "../validation/registerEnterpriseSchema"
import sendComplaintSchema from "../validation/sendComplaintSchema"
import rateComplaintSchema from "../validation/rateComplaintSchema"
import rateComplaintMutation from "@/graphql/mutations/rateComplaintMutation"
import inviteToProjectMutation from "@/graphql/mutations/inviteToProjectMutation"
import inviteToProjectSchema from "../validation/inviteToProjectSchema"
import acceptHiringInvitationMutation from "@/graphql/mutations/acceptHiringInvitationMutation"
import rejectHiringInvitationMutation from "@/graphql/mutations/rejectHiringInvitationMutation"
import hireEmployeeMutation from "@/graphql/mutations/hireEmployeeMutation"
import cancelHiringProcessMutation from "@/graphql/mutations/cancelHiringProcessMutation"
import promoteEmployeeMutation from "@/graphql/mutations/promoteEmployeeMutation"
import fireEmployeeMutation from "@/graphql/mutations/fireEmployeeMutation"
import addFeedbackCommentMutation from "@/graphql/mutations/addFeedbackCommentMutation"
import addFeedbackReplyMutation from "@/graphql/mutations/addFeedbackReplyMutation"
import removeFeedbackReplyMutation from "@/graphql/mutations/removeFeedbackReplyMutation"
import removeFeedbackCommentMutation from "@/graphql/mutations/removeFeedbackCommentMutation"
import endFeedbackMutation from "@/graphql/mutations/endFeedbackMutation"
import createEnterpriseChatMutation from "@/graphql/mutations/createEnterpriseChatMutation"
import markNotificationAsReadMutation from "@/graphql/mutations/markNotificationAsReadMutation"
import updatePasswordMutation from "@/graphql/mutations/updatePasswordMutation"
import changeUserGenreMutation from "@/graphql/mutations/changeUserGenreMutation"
import changeUserPronounMutation from "@/graphql/mutations/changeUserPronounMutation"
import updateFirstNameMutation from "@/graphql/mutations/updateFirstNameMutation"
import updateLastNameMutation from "@/graphql/mutations/updateLastNameMutation"
import changeUserPhoneMutation from "@/graphql/mutations/changeUserPhoneMutation"
import updateUserAddressMutation from "@/graphql/mutations/updateUserAddressMutation"
import changeEnterpriseAddressMutation from "@/graphql/mutations/changeEnterpriseAddressMutation"
import changeEnterprisePhoneMutation from "@/graphql/mutations/changeEnterprisePhoneMutation"
import changeEnterpriseEmailMutation from "@/graphql/mutations/changeEnterpriseEmailMutation"
import changeEnterpriseWebsiteMutation from "@/graphql/mutations/changeEnterpriseWebsiteMutation"



function gqlClientWithCookie() {
    const jwtCookie = cookies().get("jwt")
    const strCookie = `${jwtCookie?.name}=${jwtCookie?.value}`
    const gqlClient = getGraphQLClient()
    gqlClient.setHeader("Cookie", strCookie)
    return gqlClient
}

export async function changeEnterpriseWebsite(input: ChangeEnterpriseWebsite) {
    await gqlClientWithCookie().request(changeEnterpriseWebsiteMutation, { input })
}

export async function changeEnterpriseEmail(input: ChangeEnterpriseEmail) {
    await gqlClientWithCookie().request(changeEnterpriseEmailMutation, { input })
}

export async function changeEnterprisePhone(input: ChangeEnterprisePhone) {
    await gqlClientWithCookie().request(changeEnterprisePhoneMutation, { input })
}

export async function changeEnterpriseAddress(input: ChangeEnterpriseAddress) {
    await gqlClientWithCookie().request(changeEnterpriseAddressMutation, { input })
}

export async function updateUserAdress(input: UpdateUserAddress) {
    await gqlClientWithCookie().request(updateUserAddressMutation, { input })
}

export async function changeUserPhone(input: ChangeUserPhone) {
    await gqlClientWithCookie().request(changeUserPhoneMutation, { input })
}

export async function changeLastName(input: ChangeUserLastName) {
    await gqlClientWithCookie().request(updateLastNameMutation, { input })
}

export async function changeFirstName(input: ChangeUserFirstName) {
    await gqlClientWithCookie().request(updateFirstNameMutation, { input })
}

export async function changeUserPronoun(input: ChangeUserPronoun) {
    await gqlClientWithCookie().request(changeUserPronounMutation, { input })
}

export async function changeUserGenre(input: ChangeUserGenre) {
    await gqlClientWithCookie().request(changeUserGenreMutation, { input })
}

export async function updatePassword(input: ChangePassword) {
    await gqlClientWithCookie().request(updatePasswordMutation, { input })
}

export async function markNotificationAsRead(id: string) {
    await gqlClientWithCookie().request(markNotificationAsReadMutation, { id })
}

export async function createEnterpriseChat(input: CreateEnterpriseChat) {
    await gqlClientWithCookie().request(createEnterpriseChatMutation, { input })
}

export async function endFeedback(input: EndFeedback) {
    await gqlClientWithCookie().request(endFeedbackMutation, { input })
}

export async function removeFeedbackComment(input: RemoveFeedbackComment) {
    await gqlClientWithCookie().request(removeFeedbackCommentMutation, { input })
}

export async function removeFeedbackReply(input: RemoveFeedbackReply) {
    await gqlClientWithCookie().request(removeFeedbackReplyMutation, { input })
}

export async function addFeedbackReply(input: AddFeedbackReply) {
    await gqlClientWithCookie().request(addFeedbackReplyMutation, { input })
}

export async function addFeedbackComment(input: AddFeedbackComment) {
    await gqlClientWithCookie().request(addFeedbackCommentMutation, { input })
}

export async function fireEmployee(input: FireEmployee) {
    await gqlClientWithCookie().request(fireEmployeeMutation, { input })
}

export async function promoteEmployee(input: PromoteEmployee) {
    await gqlClientWithCookie().request(promoteEmployeeMutation, { input })
}

export async function cancelHiringProcess(input: CancelHiringProcess) {
    await gqlClientWithCookie().request(cancelHiringProcessMutation, { input })
}

export async function hireEmployee(input: HireEmployee) {
    await gqlClientWithCookie().request(hireEmployeeMutation, { input })
}

export async function rejectHiringInvitation(input: RejectHiringInvitation) {
    await gqlClientWithCookie().request(rejectHiringInvitationMutation, { input })
}

export async function acceptHiringInvitation(input: AcceptHiringInvitation) {
    await gqlClientWithCookie().request(acceptHiringInvitationMutation, { input })
}

export async function inviteToProject(input: InviteToProject) {
    const { data, success, error } = inviteToProjectSchema.safeParse(input)
    if (!success) {
        const fieldErrors = error.flatten().fieldErrors
        let msg = Object.keys(fieldErrors).map((k) => fieldErrors[k as keyof typeof fieldErrors]?.join("\n"))
        throw new Error(msg.join())
    }
    await gqlClientWithCookie().request(inviteToProjectMutation, { input: data })
}

export type ReviewFormErrors = Partial<z.inferFlattenedErrors<typeof rateComplaintSchema>>
export async function rateComplaint(fd: FormData): Promise<{ success: boolean, errors: ReviewFormErrors }> {
    const { data, success, error } = rateComplaintSchema.safeParse(Object.fromEntries(fd))
    if (!success) {
        return { success, errors: error.flatten() }
    }
    try {
        await gqlClientWithCookie().request(rateComplaintMutation, { input: data as RateComplaint })
    } catch (e: any) {
        let msg: string = e.message
        return {
            success: false, errors: {
                formErrors: [msg],
                fieldErrors: {}
            }
        }
    }
    return {
        success: true, errors: {
            formErrors: [],
            fieldErrors: {}
        }
    }
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
    return result.createNewComplaint as Complaint
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


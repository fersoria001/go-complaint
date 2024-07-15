'use server'

import contactSchema from "../validation/contactSchema"

export default async function contact(prevState: any, fd: FormData) {
    const {data,error,success} = contactSchema.safeParse(Object.fromEntries(fd))
    if(!success){
        return error.flatten()
    }
    return {
        formErrors: ["error at sending contact message"],
        fieldErrors: {}
    }
}
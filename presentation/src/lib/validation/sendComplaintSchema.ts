import { z } from "zod";

const sendComplaintSchema = z.object({
    complaintId: z.string(),
    body: z
        .string({ message: "your complaint cannot be empty" })
        .min(50, { message: "write a minimum of 50 characters" })
        .max(250, { message: "the maximum characters length for the complaint is 250" })
})

export default sendComplaintSchema;
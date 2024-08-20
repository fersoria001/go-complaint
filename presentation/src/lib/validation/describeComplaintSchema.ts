import { z } from "zod";

const describeComplaintSchema = z.object({
  complaintId: z.string(),
  title: z
    .string({ message: "write a message for the title" })
    .min(10, "write a minimum of 3 characters")
    .max(79, "the max characters length is 80"),
  description: z
    .string({ message: "write a message for the description" })
    .min(3, "write a minimum of 3 characters")
    .max(119, "the max characters length is 80"),
});

export default describeComplaintSchema;
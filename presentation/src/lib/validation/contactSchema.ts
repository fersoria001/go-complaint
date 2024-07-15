import { z } from "zod";

const contactSchema = z.object({
  email: z.string().email({ message: "enter a valid email" }),
  text: z.string().min(20, { message: "write at least 20 characters" }),
});

export default contactSchema;

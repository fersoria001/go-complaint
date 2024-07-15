import { z } from "zod";

const signInSchema = z.object({
  email: z.string().email({ message: "enter a valid email" }),
  password: z.string().min(8, { message: "enter your password" }),
  rememberMe: z.boolean().optional()
});

export default signInSchema;

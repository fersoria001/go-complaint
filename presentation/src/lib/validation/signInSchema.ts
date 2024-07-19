import { z } from "zod";

const signInSchema = z.object({
  userName: z.string().email({ message: "enter a valid email" }),
  password: z.string().min(8, { message: "enter your password" }),
  rememberMe: z.string().transform(value => value === 'on').catch(false)
});

export default signInSchema;

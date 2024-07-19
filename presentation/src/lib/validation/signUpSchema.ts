import { z } from "zod";
import { passwordRegex } from "./regex";

const signUpSchema = z
  .object({
    userName: z.string().email({ message: "Please enter a valid email" }),
    password: z
      .string()
      .regex(
        passwordRegex,
        "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
      ),
    confirmPassword: z
      .string()
      .regex(
        passwordRegex,
        "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
      ),
    firstName: z
      .string()
      .min(2, { message: "First name must be at least 2 characters long" })
      .max(50, { message: "First name must be at most 50 characters long" }),
    lastName: z
      .string()
      .min(2, { message: "Last name must be at least 2 characters long" })
      .max(50, { message: "Last name must be at most 50 characters long" }),
    genre: z.enum(["male", "female", "non-declared"], {
      message: "Please select a gender from the list provided",
    }),
    pronoun: z.enum(["he", "she", "they"], {
      message: "Please select a pronoun from the list provided",
    }),
    birthDate: z
      .string()
      .date()
      .transform((val, ctx) => {
        const stringDate = Date.parse(val).toString();
        if (isNaN(parseInt(stringDate))) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please select a valid date",
          });
          return z.NEVER;
        }
        return stringDate;
      }),
    phoneNumber: z
      .string({ message: "We could not validate your phone number" })
      .min(10, { message: "We could not validate your phone number" })
      .transform((val, ctx) => {
        const parsed = parseInt(val);
        if (isNaN(parsed)) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Not a number",
          });
          return z.NEVER;
        }
        return val;
      }),
    countryId: z.string().transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must select a country",
        });
        return z.NEVER;
      }
      return parsed;
    }),
    countryStateId: z.string().transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must select a county",
        });
        return z.NEVER;
      }
      return parsed;
    }),
    cityId: z.string().transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must select a city",
        });
        return z.NEVER;
      }
      return parsed;
    }),
    terms: z.enum(["true", "on", "1"], {
      message: "You must accept the terms and conditions",
    }),
  })
  .superRefine(({ confirmPassword, password }, ctx) => {
    if (confirmPassword !== password) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        path: ["confirmPassword"],
        message: "The passwords did not match",
      });
    }
});
export default signUpSchema;
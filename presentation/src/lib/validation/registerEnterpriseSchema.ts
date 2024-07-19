import { z } from "zod";

const registerEnterpriseSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email" }),
  name: z
    .string()
    .min(3, "The enterprise name should be of at least 3 characters length")
    .max(120, "The enterprise name should be of at most 120 characters length"),
  //   .transform(async (val, ctx) => {
  //     if (
  //       !(await Query<boolean>(
  //         IsEnterpriseNameAvailableQuery,
  //         IsEnterpriseNameAvailable,
  //         [val]
  //       ))
  //     ) {
  //       ctx.addIssue({
  //         code: z.ZodIssueCode.custom,
  //         message:
  //           "Enterprise name is already taken, please choose a different one",
  //       });
  //     }
  //     return val;
  //   }),
  website: z.string().url({
    message: "Please enter a valid website e.g: http://www.mywebsite.com",
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
  industryId: z
    .string()
    .min(1, { message: "Please select an industry" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select an industry",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  countryId: z
    .string()
    .min(1, { message: "Please select a country" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select a country",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  countryStateId: z
    .string()
    .min(1, { message: "Please select a state" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select a state",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  cityId: z
    .string()
    .min(1, { message: "Please select a city" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select a city",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  foundationDate: z
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
  terms: z.enum(["true", "on", "1"], {
    message: "You must accept the terms and conditions",
  }),
});

export default registerEnterpriseSchema;
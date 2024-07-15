import { z } from "zod";

const confirmationCodeSchema = z.object({
  confirmationCode: z
    .string({message: "enter the confirmation code you have received in your email"})
    .length(7, {
      message: "the confirmation code you have provided is not a valid confirmation code",
    })
    .transform((val, ctx) => {
      const segments = val.split("");
      for (const segment of segments) {
        if (!segment.match(/^[0-9]+$/)) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "the confirmation code you have provided is not a valid confirmation code",
          });
          return z.NEVER;
        }
      }
      const parsed = parseInt(val);
      if (isNaN(parsed)) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "the confirmation code you have provided is not a valid confirmation code",
        });
        return z.NEVER;
      }
      return parsed;
    }),
});

export default confirmationCodeSchema;

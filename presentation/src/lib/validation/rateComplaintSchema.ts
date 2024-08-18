import { z } from "zod";

const rateComplaintSchema = z.object({
  userId: z.string(),
  complaintId: z.string(),
  rate: z
    .string()
    .min(1)
    .max(1)
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed)) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Not a number",
        });
        return z.NEVER;
      }
      if (parsed < 0 || parsed > 5) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "the rating score should be between 0 and 5",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  comment: z
    .string()
    .min(3, "write at least 3 characters for the comment or leave it empty")
    .max(250, "comments have a maximum of 250 characters")
    .optional()
    .or(z.literal("")),
});

export default rateComplaintSchema;

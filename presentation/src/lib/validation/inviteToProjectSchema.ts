import { z } from "zod";

const inviteToProjectSchema = z.object({
  enterpriseId: z.string(),
  role: z.enum(["ASSISTANT", "MANAGER"], {
    message: "select a role option from the list",
  }),
  proposeTo: z.string(),
  proposedBy: z.string(),
});

export default inviteToProjectSchema
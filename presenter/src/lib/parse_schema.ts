/* eslint-disable @typescript-eslint/no-explicit-any */
import { z } from "zod";
import { ErrorType } from "./types";

export async function parseSchema<T extends z.ZodTypeAny>(
  obj: object,
  schema: T,
  async: boolean = false
): Promise<{
  data: z.infer<T>;
  errors: ErrorType;
}> {
  let parsed: z.infer<T>;
  if (async) {
    parsed = (await schema.spa(obj)) as z.infer<T>;
  } else {
    parsed = schema.safeParse(obj) as z.infer<T>;
  }
  const errors: ErrorType = {};
  let errorPath: string;
  if (!parsed.success) {
    parsed.error.errors.forEach((error: any) => {
      errorPath = error.path.join("");
      errors[errorPath] = error.message;
    });
  }
  return { data: parsed.data, errors };
}

export function syncParseSchema<T extends z.ZodTypeAny>(
  obj: object,
  schema: T
): {
  data: z.infer<T>;
  errors: ErrorType;
} {
  const parsed = schema.safeParse(obj) as z.infer<T>;
  const errors: ErrorType = {};
  let errorPath: string;
  if (!parsed.success) {
    parsed.error.errors.forEach((error: any) => {
      errorPath = error.path.join("");
      errors[errorPath] = error.message;
    });
  }
  return { data: parsed.data, errors };
}

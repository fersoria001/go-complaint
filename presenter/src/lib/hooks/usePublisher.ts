import { useEffect, useState } from "react";
import { ErrorType } from "../types";
import { z } from "zod";
import { parseSchema } from "../parse_schema";
import { Publish } from "../publish";

/* eslint-disable @typescript-eslint/no-explicit-any */
export function usePublisher<T extends z.ZodTypeAny>(
  formData: FormData | null,
  schema: T,
  mutationFn: (data: z.infer<T>) => string,
  subscriptionID: string,
  async: boolean = false
) {
  const [success, setSuccess] = useState<boolean>(false);
  const [errors, setErrors] = useState<ErrorType>({});
  useEffect(() => {
    async function submitForm() {
      if (!formData) return;
      const obj = Object.fromEntries(formData.entries());
      const { data, errors } = await parseSchema(obj, schema, async);
      if (Object.keys(errors).length > 0) {
        setErrors(errors);
        return;
      }
      try {
        const result = await Publish<z.infer<T>>(mutationFn, data, subscriptionID);
        setSuccess(result);
      } catch (e: any) {
        console.error(e);
        setErrors({ form: e.message });
      }
    }
    submitForm();
  }, [formData, schema, mutationFn, async, subscriptionID]);
  return { success, errors };
}

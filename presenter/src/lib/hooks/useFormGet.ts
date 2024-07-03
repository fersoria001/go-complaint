/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState, useEffect } from "react";
import { Query } from "../queries";
import { ErrorType } from "../types";
import { z } from "zod";

export function useFormGet<T, S extends z.ZodTypeAny>(
  args: object,
  schema: S,
  queryFn: (...args: any[]) => string,
  toReturnType: (data: any) => T,
): { data: T; errors: ErrorType } {
  const [errors, setErrors] = useState<ErrorType>({});
  const [result, setResult] = useState<T | null>(null);
  useEffect(() => {
    if (Object.keys(args).length === 0) {
      return;
    }
    const parsed = schema.safeParse(args) as z.infer<S>;
    const errors: ErrorType = {};
    let errorPath: string;
    if (!parsed.success) {
      parsed.error.errors.forEach((error: any) => {
        errorPath = error.path.join("");
        errors[errorPath] = error.message;
      });
    }
    if (Object.keys(errors).length > 0) {
      setErrors(errors);
      return;
    }
    async function get() {
      const argsArray = [];
      for (const key in parsed.data) {
        argsArray.push(parsed.data[key]);
      }
      try {
        const result = await Query<T>(
          queryFn,
          toReturnType,
          [...argsArray],
        );
        setResult(result);
      } catch (e: any) {
        setErrors({ form: e.message });
      }
    }
    get();
  }, [args, schema, queryFn, toReturnType]);
  return { data: result as T, errors };
}

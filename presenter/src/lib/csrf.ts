export async function csrf(): Promise<string> {
  const response = await fetch(import.meta.env.VITE_CSRF_TOKEN_ENDPOINT, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const token = response.headers.get("x-csrf-token");
  return token ? token : "";
}

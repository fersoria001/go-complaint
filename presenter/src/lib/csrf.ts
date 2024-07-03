
export async function csrf(): Promise<string> {
  const response = await fetch("http://localhost:5170/csrf", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const token = response.headers.get("x-csrf-token");
  return token ? token : "";
}

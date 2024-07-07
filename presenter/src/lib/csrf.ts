export async function csrf(): Promise<string> {
  const response = await fetch("https://api.go-complaint.com/csrf", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const token = response.headers.get("x-csrf-token");
  return token ? token : "";
}

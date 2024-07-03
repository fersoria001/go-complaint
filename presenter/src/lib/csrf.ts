
export async function csrf(): Promise<string> {
  const response = await fetch("https://go-complaint-server-latest.onrender.com", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const token = response.headers.get("x-csrf-token");
  return token ? token : "";
}
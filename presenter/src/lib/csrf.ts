
export async function csrf(): Promise<string> {
  const response = await fetch("http://3.143.110.143:5555", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  const token = response.headers.get("x-csrf-token");
  return token ? token : "";
}

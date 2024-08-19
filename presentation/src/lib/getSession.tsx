export default async function getSession(strCookie: string) {
    const apiKey = process.env.NEXT_PUBLIC_API_KEY;
    if (!apiKey) {
      throw new Error("api key axios instance not defined in process env");
    }
    const res = await fetch(
        process.env.NEXT_PUBLIC_BASE_URL + "/session", {
        method: "GET",
        headers: {
            "Cookie": strCookie,
            "api_key": apiKey
        },
        credentials: 'include'
    })
    const ok = await res.json()
    if (!ok) {
        throw new Error("failed to get session")
    }
    return ok
}
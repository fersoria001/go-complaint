export default async function getSession(strCookie: string) {
    const res = await fetch(
        process.env.NEXT_PUBLIC_BASE_URL + "/session", {
        method: "GET",
        headers: {
            "Cookie": strCookie
        },
        credentials: 'include'
    })
    const ok = await res.json()
    if (!ok) {
        throw new Error("failed to get session")
    }
    return ok
}
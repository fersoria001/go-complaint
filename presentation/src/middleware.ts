import { NextResponse, type NextRequest } from "next/server";
import getSession from "./lib/getSession";
import { UserDescriptor } from "./gql/graphql";

export async function middleware(request: NextRequest) {
  const currentUser = request.cookies.get("jwt")?.value;
  if (!currentUser && request.nextUrl.pathname.startsWith("/profile")) {
    return Response.redirect(new URL("/sign-in", request.url));
  }
  if (request.nextUrl.pathname.startsWith("/complaint")) {
    const strCookie = `jwt=${currentUser}`;
    const session = (await getSession(strCookie)) as UserDescriptor;
    if (!session) {
      return Response.redirect(new URL("/sign-in", request.url));
    }
    const aliasCookie = request.cookies.get("alias");
    if (!aliasCookie) {
      const response = NextResponse.next();
      response.cookies.set("alias", session.id);
      return NextResponse.rewrite(request.nextUrl, response);
    }
  }
}

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|.*\\.png$).*)"],
};

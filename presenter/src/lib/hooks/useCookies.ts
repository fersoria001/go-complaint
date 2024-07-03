import Cookies from 'js-cookie';
export function useCookies() {
  function setCookie(name: string, value: string, days: number) {
    if (days) {
      const date = new Date();
      date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
      Cookies.set(name, value, { path: "/", expires: date });
      return;
    }
    throw new Error("days is required");
  }
  function removeCookie(name: string) {
    Cookies.remove(name);
  }
  return { setCookie, removeCookie };
}

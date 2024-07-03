export function daysAgoFilter(key: string) {
  const today = new Date();
  const oneDay = 24 * 60 * 60 * 1000;
  let result: [string, string] = ["", ""];
  switch (key) {
    case "Last day": {
      const aDayAgo = new Date(new Date().getTime() - oneDay);
      result = [aDayAgo.getTime().toString(), today.getTime().toString()];
      break;
    }
    case "Last 7 days": {
      const sevenDaysAgo = new Date(today.getTime() - 7 * oneDay);
      result = [sevenDaysAgo.getTime().toString(), today.getTime().toString()];
      break;
    }
    case "Last 30 days": {
      const thirtyDaysAgo = new Date(today.getTime() - 30 * oneDay);
      result = [thirtyDaysAgo.getTime().toString(), today.getTime().toString()];
      break;
    }
    case "Last month": {
      const daysAfterFirst = new Date(today.getFullYear(), today.getMonth(), 1);
      const sub = new Date(daysAfterFirst.getTime() - oneDay);
      const subOneMonth = sub.getMonth() - 1;
      const lastMonthDaysQtty =
        new Date(sub.getFullYear(), subOneMonth, 0).getDate() - 1;
      const daysToSubstract = lastMonthDaysQtty * oneDay;
      const before = new Date(sub.getTime() - daysToSubstract);
      const after = new Date(sub.getTime() + oneDay);
      result = [after.getTime().toString(), before.getTime().toString()];
      break;
    }
    case "Last year": {
      const aYearAgo = new Date(
        today.getFullYear() - 1,
        today.getMonth(),
        today.getDate()
      );
      result = [aYearAgo.getTime().toString(), today.getTime().toString()];
      break;
    }
  }
  return result;
}

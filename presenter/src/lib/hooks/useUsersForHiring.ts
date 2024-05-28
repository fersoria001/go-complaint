import { useEffect, useState } from "react";
import { Query, UsersForHiringQuery, UsersForHiringType } from "../queries";
import { UsersForHiring } from "../types";

function useUsersForHiring(
  id: string,
  page: string,
  query: string
): UsersForHiring | null {
  const [list, setList] = useState<UsersForHiring | null>(null);
  const offset = (parseInt(page) - 1) * 10;
  useEffect(() => {
    const fetch = async () => {
      const usersForHiring = await Query<UsersForHiring>(
        UsersForHiringQuery,
        UsersForHiringType,
        [id, 10, offset, query]
      );
      if (usersForHiring) {
        setList(usersForHiring);
      }
    };
    fetch();
  }, [id, offset, query]);
  return list;
}
export default useUsersForHiring;

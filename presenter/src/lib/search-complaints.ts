import { Query, SearchInDraftQuery, SearchInDraftTypeList } from "./queries"
import { ComplaintTypeList } from "./types"

export async function searchInInbox(id: string, query: string, afterBefore: [string, string]) : Promise<ComplaintTypeList> {
    const results = await Query<ComplaintTypeList>(SearchInDraftQuery,
        SearchInDraftTypeList, [id, query, ...afterBefore, 10, 0])
    return results
}
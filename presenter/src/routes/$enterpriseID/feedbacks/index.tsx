import { createFileRoute } from '@tanstack/react-router'
import { ComplaintQuery, ComplaintQueryType, FeedbackByIDQuery, FeedbackByIDType, Query, UserQuery, UserType } from '../../../lib/queries';
import { ComplaintType, FeedbackType, User } from '../../../lib/types';
import EmployeeFeedback from '../../../components/enterprise/feedback/EmployeeFeedback';
type FeedbacksSearchType = {
  id: string;
}
export const Route = createFileRoute('/$enterpriseID/feedbacks/')({
  validateSearch: (search: Record<string, unknown>): FeedbacksSearchType => {
    return {
      id: search.id
    } as FeedbacksSearchType
  },
  loaderDeps: ({ search: { id } }) => ({ id }),
  loader: async ({ context: { fetchUserDescriptor }, deps: { id } }) => {
    const descriptor = await fetchUserDescriptor()
    const reviewer = await Query<User>(
      UserQuery,
      UserType,
      [descriptor.email]
    )
    
    const feedback = await Query<FeedbackType>(
      FeedbackByIDQuery,
      FeedbackByIDType,
      [id]
    )
    const complaint = await Query<ComplaintType>(
      ComplaintQuery,
      ComplaintQueryType,
      [feedback.complaintID]
    )
    return {
      subscriptionID: null,
      reviewer,
      feedback,
      complaint
    }
  },
  component: EmployeeFeedback
})
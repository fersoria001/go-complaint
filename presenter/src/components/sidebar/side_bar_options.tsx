import DashboardIcon from "../icons/DashboardIcon";
import EmployeesIcon from "../icons/EmployeesIcon";
import HiringIcon from "../icons/HiringIcon";
import InboxIcon from "../icons/InboxIcon";
import SendComplaintIcon from "../icons/SendComplaintIcon";
import SentIcon from "../icons/SentIcon";
import { ComplaintTypeList, SideBarOptionsType, UserDescriptor } from "../../lib/types";
import RegisterEnterpriseIcon from "../icons/RegisterEnterpriseIcon";
import HiringInvitationsIcon from "../icons/HiringInvitationsIcon";
import HiringProccesesIcon from "../icons/HiringProccesesIcon";
import RateReviewIcon from "../icons/RateReviewIcon";
import HistoryIcon from "../icons/HistoryIcon";
import { Query, SearchInDraftQuery, SearchInDraftTypeList, SentSearchQuery, SentSearchTypeList } from "../../lib/queries";
import { hasPermission } from "../../lib/is_logged_in";
import FeedbackIcon from "../icons/FeedbackIcon";
export const countInboxUnreads = async (
  descriptorID: string,
  id: string,
): Promise<number> => {
  const inbox = await Query<ComplaintTypeList>(
    SearchInDraftQuery,
    SearchInDraftTypeList,
    [
      id,
      "",
      "",
      "",
      0,
      0,
    ])
  let unread = 0
  if (inbox.complaints) {
    for (const complaint of inbox.complaints) {
      const filtered = complaint.replies!.filter(reply => reply.senderID != descriptorID && !reply.read)
      unread += filtered.length
    }
  }
  return unread
}
export const countSentUnreads = async (
  descriptorID: string,
  id: string,
): Promise<number> => {
  const inbox = await Query<ComplaintTypeList>(
    SentSearchQuery,
    SentSearchTypeList,
    [
      id,
      "",
      "",
      "",
      0,
      0,
    ])
  let unread = 0
  if (inbox.complaints) {
    for (const complaint of inbox.complaints) {
      const filtered = complaint.replies!.filter(reply => reply.senderID != descriptorID && !reply.read)
      unread += filtered.length
    }
  }
  return unread
}
//className="flex-shrink-0 w-5 h-5 text-gray-700 transition duration-75 group-hover:text-gray-900" fill="#374151"
export const enterpriseOptions = async (descriptorID: UserDescriptor, id: string): Promise<SideBarOptionsType[]> => {
  const inbox = await countInboxUnreads(descriptorID.email, id)
  const sent = await countSentUnreads(descriptorID.email, id)
  const opts = [
    {
      link: `/${id}`,
      icon: <DashboardIcon />,
      title: "Dashboard",
    },
    {
      link: `/${id}/send-complaint`,
      icon: <SendComplaintIcon />,
      title: "Send a Complaint",
    },
    {
      link: `/${id}/inbox`,
      icon: <InboxIcon />,
      title: "Inbox",
      unread: inbox
    },
    {
      link: `/${id}/sent`,
      icon: <SentIcon />,
      title: "Sent",
      unread: sent
    },
    {
      link: `/${id}/reviews`,
      icon: <RateReviewIcon />,
      title: "Reviews",
    },
    {
      link: `/${id}/history`,
      icon: <HistoryIcon />,
      title: "History",
    }
  ]
  let ok = await hasPermission("OWNER", id)
  if (!ok) {
    opts.push({
      link: `/${id}/feedback`,
      icon: <FeedbackIcon />,
      title: "Feedback",
    },)
  }
  ok = await hasPermission("MANAGER", id)
  if (ok) {
    opts.push(...[{
      link: `/${id}/hiring`,
      icon: <HiringIcon />,
      title: "Hiring",
    },
    {
      link: `/${id}/hiring-procceses`,
      icon: <HiringProccesesIcon />,
      title: "Hiring Processes",
    },
    {
      link: `/${id}/employees`,
      icon: <EmployeesIcon />,
      title: "Employees",
    }])
  }

  return opts
}

export const profileOptions = async (descriptorID: string, id: string): Promise<SideBarOptionsType[]> => {
  const inbox = await countInboxUnreads(descriptorID, id)
  const sent = await countSentUnreads(descriptorID, id)
  return [
    {
      link: `/profile`,
      icon: < DashboardIcon />,
      title: "Dashboard",
    },
    {
      link: `/send-complaint`,
      icon: <SendComplaintIcon />,
      title: "Send a Complaint",
    },
    {
      link: `/inbox`,
      icon: <InboxIcon />,
      title: "Inbox",
      unread: inbox
    },
    {
      link: `/sent`,
      icon: <SentIcon />,
      title: "Sent",
      unread: sent
    },
    {
      link: `/register-enterprise`,
      icon: <RegisterEnterpriseIcon />,
      title: "Register an Enterprise",
    },
    {
      link: `/hiring-invitations`,
      icon: <HiringInvitationsIcon />,
      title: "Hiring Invitations",
    },
    {
      link: `/reviews`,
      icon: <RateReviewIcon />,
      title: "Reviews",
    },
    {
      link: `/history`,
      icon: <HistoryIcon />,
      title: "History",
    }
  ]
};

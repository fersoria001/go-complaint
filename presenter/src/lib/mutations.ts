import EmailAlreadyInUseError from "../components/error/EmailAlreadyInUseError";
import { ChangePasswordType } from "../components/profile/settings/settings_lib";
import { csrf } from "./csrf";
import { deleteLinebreaks } from "./delete_line_breaks";
import {
  ContactType,
  CreateAFeedback,
  CreateUser,
  DeclineHiringInvitation,
  EndHiringProcess,
  EnterpriseEventType,
  FireEmployeeType,
  InviteToProject,
  MarkComplaintRepliesAsSeenType,
  MarkReplyChatAsSeenType,
  PromoteEmployeeType,
  RateComplaint,
  RegisterEnterprise,
  ReplyComplaintType,
  ReplyEnterpriseChatType,
  SendComplaintType,
  UpdateEnterpriseType,
  UpdateUserType,
} from "./types";

export const ContactMutation = ({ email, text }: ContactType): string => {
  return `
  mutation {
    Contact(
        email: "${email}",
        text: "${text}"
    )
}
  `;
};

export const CreateUserMutation = ({
  email,
  password,
  firstName,
  lastName,
  gender,
  pronoun,
  birthDate,
  phoneCode,
  phone,
  country,
  county,
  city,
}: CreateUser): string =>
  `
        mutation {
            CreateUser(
                email: "${email}",
                password: "${password}",
                firstName: "${firstName}",
                lastName: "${lastName}",
                gender: "${gender}",
                pronoun: "${pronoun}",
                birthDate: "${birthDate}",
                phoneCode: "${phoneCode}",
                phone: "${phone}",
                country: ${country},
                county: ${county},
                city: ${city},
            )
        }
    `;

export const MarkNotificationAsRead = (id: string): string => `
mutation {
    MarkNotificationAsRead(
        id: "${id}"
    )
}
`;
export const UpdateUserMutation = ({ updateType, value }: UpdateUserType) => {
  return `mutation {
    UpdateUser(
        updateType: "${updateType}",
        value: "${value}"
    )
}
`;
};
export const UpdateUserMutation2 = ({
  updateType,
  numberValue,
}: UpdateUserType) => {
  return `mutation {
    UpdateUser(
        updateType: "${updateType}",
        numberValue: ${numberValue}
    )
}
`;
};
export const ChangePasswordMutation = ({
  oldPassword,
  newPassword,
}: ChangePasswordType): string => `
mutation {
    ChangePassword(
        oldPassword: "${oldPassword}",
        newPassword: "${newPassword}"
    )
}
`;
export const CreateEnterpriseMutation = ({
  name,
  email,
  website,
  phoneCode,
  phone,
  industryID,
  countryID,
  countryStateID,
  cityID,
  foundationDate,
}: RegisterEnterprise): string => `
mutation {
    CreateEnterprise(
        name: "${name}",
        email: "${email}",
        website: "${website}",
        phoneCode: "${phoneCode}",
        phone: "${phone}",
        industryID: ${industryID},
        countryID: ${countryID},
        countryStateID: ${countryStateID},
        cityID: ${cityID},
        foundationDate: "${foundationDate}",
    )
}
`;

export const UpdateEnterpriseMutation = ({
  updateType,
  enterpriseID,
  value,
}: UpdateEnterpriseType) => {
  return `mutation {
    UpdateEnterprise(
        updateType: "${updateType}",
        enterpriseID: "${enterpriseID}",
        value: "${value}"
    )
}
`;
};
export const UpdateEnterpriseMutation2 = ({
  updateType,
  enterpriseID,
  numberValue,
}: UpdateEnterpriseType) => {
  return `mutation {
    UpdateEnterprise(
        updateType: "${updateType}",
        enterpriseID: "${enterpriseID}",
        numberValue: ${numberValue}
    )
}
`;
};

export const HireEmployeeMutation = ({
  enterpriseName,
  eventID,
}: EnterpriseEventType) => {
  return `
  mutation {
    HireEmployee(
      enterpriseName: "${enterpriseName}",
      eventID: "${eventID}",
    )
  }
  `;
};
export const CancelHiringProccessMutation = ({
  enterpriseName,
  eventID,
}: EnterpriseEventType) => {
  return `
  mutation {
    CancelHiringProccess(
      enterpriseName: "${enterpriseName}",
      eventID: "${eventID}",
    )
  }
  `;
};
export const SendComplaintMutation = ({
  authorID,
  receiverID,
  receiverFullName,
  receiverProfileIMG,
  title,
  description,
  content,
}: SendComplaintType): string => `
mutation {
    SendComplaint(
        authorID: "${authorID}",
        receiverID: "${receiverID}",
        receiverFullName: "${receiverFullName}",
        receiverProfileIMG: "${receiverProfileIMG}",
        title: "${title}",
        description: "${description}",
        content: "${content}",
    )
}
`;

export const ReplyComplaintMutation = ({
  complaintID,
  replyAuthorID,
  replyBody,
  replyEnterpriseID,
}: ReplyComplaintType): string => `
mutation {
    ReplyComplaint(
        complaintID: "${complaintID}",
        replyAuthorID: "${replyAuthorID}",
        replyBody: "${replyBody}",
        replyEnterpriseID: "${replyEnterpriseID}",
    )
}
`;

export const MarkComplaintRepliesAsSeenMutation = ({
  complaintID,
  repliesID,
}: MarkComplaintRepliesAsSeenType): string => {
  return `
mutation {
    MarkAsSeen(
        complaintID: "${complaintID}",
        ids: "${repliesID.join(",")}",
    )
}
`;
};

export const InviteToProjectMutation = ({
  enterpriseName,
  proposedPosition,
  proposeTo,
}: InviteToProject): string => {
  return `
  mutation {
    InviteToEnterprise(
      enterpriseName: "${enterpriseName}",
      proposedPosition: "${proposedPosition}",
      proposeTo: "${proposeTo}"
    )
  }
  `;
};

export const AcceptHiringInvitationMutation = (id: string): string => {
  return `
  mutation {
    AcceptEnterpriseInvitation(
      id: "${id}"
    )
  }
  `;
};
export const DeclineHiringInvitationMutation = ({
  id,
  reason,
}: DeclineHiringInvitation): string => {
  return `
  mutation {
    RejectEnterpriseInvitation(
      id: "${id}"
      reason: "${reason}"
    )
  }
  `;
};
export const EndHiringProcessMutation = (ehp: EndHiringProcess): string => {
  return `mutation {
    EndHiringProcess(
      pendingEventID: "${ehp.pendingEventID}",
      enterpriseName: "${ehp.enterpriseID}",
      accepted: ${ehp.accepted}
    )
  }
  `;
};
export type LeaveEnterpriseType = {
  enterpriseName: string;
  employeeID: string;
};
export const LeaveEnterpriseMutation = ({
  enterpriseName,
  employeeID,
}: LeaveEnterpriseType): string => {
  return `mutation {
    LeaveEnterprise(
      enterpriseName: "${enterpriseName}",
      employeeID: "${employeeID}"
    )
  }
  `;
};

export const MarkAsReviewableMutation = (id: string): string => {
  return `mutation {
    SendForReviewing(
      id: "${id}",
    )
  }
  `;
};

export const RateComplaintMutation = (rate: RateComplaint): string => {
  return `mutation {
    RateComplaint(
      complaintId: "${rate.complaintId}",
      eventId: "${rate.eventId}",
      rate: ${rate.rate},
      comment: "${rate.comment}",
      )
    }`;
};

export const CreateFeedbackMutation = ({
  enterpriseID,
  complaintID,
  color,
}: CreateAFeedback) => {
  return `
  mutation{
    CreateFeedback(
      enterpriseID: "${enterpriseID}",
      complaintID: "${complaintID}",
      color: "${color}",
      )
    }`;
};
export const AddReplyMutation = ({
  enterpriseID,
  feedbackID,
  reviewerID,
  color,
  repliesID,
}: CreateAFeedback) => {
  return `
  mutation{
    AddReply(
      enterpriseID: "${enterpriseID}",
      feedbackID: "${feedbackID}",
      reviewerID: "${reviewerID}",
      color: "${color}",
      repliesID: "${repliesID}",
      )
      }`;
};

export const RemoveReplyMutation = ({
  enterpriseID,
  feedbackID,
  color,
  repliesID,
}: CreateAFeedback) => {
  return `
  mutation{
    RemoveReply(
      enterpriseID: "${enterpriseID}",
      feedbackID: "${feedbackID}",
      color: "${color}",
      repliesID: "${repliesID}",
      )
      }`;
};

export const AddCommentMutation = ({
  enterpriseID,
  feedbackID,
  color,
  comment,
}: CreateAFeedback) => {
  return `
  mutation{
    AddComment(
      enterpriseID: "${enterpriseID}",
      feedbackID: "${feedbackID}",
      color: "${color}",
      comment: "${comment}",
      )
      }`;
};

export const DeleteCommentMutation = ({
  enterpriseID,
  feedbackID,
  color,
}: CreateAFeedback) => {
  return `
  mutation{
    DeleteComment(
      enterpriseID: "${enterpriseID}",
      feedbackID: "${feedbackID}",
      color: "${color}",
      )
      }`;
};

export const EndFeedbackMutation = ({
  enterpriseID,
  feedbackID,
}: CreateAFeedback) => {
  return `
  mutation{
    EndFeedback(
      enterpriseID: "${enterpriseID}",
      feedbackID: "${feedbackID}",
      )
    }`;
};

export const ReplyChatMutation = ({
  id,
  enterpriseName,
  senderID,
  content,
}: ReplyEnterpriseChatType): string => `
mutation {
    ReplyChat(
        id: "${id}",
        enterpriseName: "${enterpriseName}",
        senderID: "${senderID}",
        content: "${content}",
    )
}
`;

export const MarkReplyChatAsSeenMutation = ({
  chatID,
  enterpriseName,
  repliesID,
}: MarkReplyChatAsSeenType): string => {
  return `
mutation {
    MarkReplyChatAsSeen(
        chatID: "${chatID}",
        enterpriseName: "${enterpriseName}",
        repliesID: "${repliesID.join(",")}",
    )
}
`;
};

export const PromoteEmployeeMutation = ({
  enterpriseName,
  employeeID,
  position,
}: PromoteEmployeeType): string => `
mutation{
  PromoteEmployee(
    enterpriseName:"${enterpriseName}",
    employeeID:"${employeeID}",
    position:"${position}"
  )
}
`;

export const FireEmployeeMutation = ({
  enterpriseName,
  employeeID,
}: FireEmployeeType): string => `
mutation{
  FireEmployee(
    enterpriseName:"${enterpriseName}",
    employeeID:"${employeeID}"
  )
}
`;

export const Mutation = async <T>(
  mutationFn: (data: T) => string,
  arg: T
): Promise<boolean> => {
  const token = await csrf();
  if (token != "") {
    const response = await fetch("https://api.go-complaint.com/graphql", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-csrf-token": token,
      },
      credentials: "include",
      body: JSON.stringify({ query: deleteLinebreaks(mutationFn(arg)) }),
    });
    const data = await response.json();
    if (data.errors) {
      const stringErr = data.errors[0].message;
      if (stringErr.includes("SQLSTATE 23505")) {
        throw new EmailAlreadyInUseError();
      }
      throw new Error(stringErr);
    }
    return true;
  }
  throw new Error("No CSRF token");
};

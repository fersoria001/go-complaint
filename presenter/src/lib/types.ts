/* eslint-disable @typescript-eslint/no-explicit-any */
import { z } from "zod";
import {
  Query,
  IsEnterpriseNameAvailableQuery,
  IsEnterpriseNameAvailable,
} from "./queries";

export type ContactType = {
  email: string;
  text: string;
};
export const contactSchema = z.object({
  email: z.string().email({ message: "enter a valid email" }),
  text: z.string().min(20).email({ message: "write at least 20 characters" }),
});

export type MarkReplyChatAsSeenType = {
  chatID: string;
  enterpriseName: string;
  repliesID: string[];
};

export type MarkComplaintRepliesAsSeenType = {
  complaintID: string;
  repliesID: string[];
};

export type PromoteEmployeeType = {
  enterpriseName: string;
  employeeID: string;
  position: string;
};
export type FireEmployeeType = {
  enterpriseName: string;
  employeeID: string;
};

export type SideBarOptionsType = {
  link: string;
  icon: React.ReactNode;
  title: string;
  unread?: number;
};
export type Subscription<T> = {
  connection_ack: ConnectionACK;
  subscription: string;
  subscriptionID: string;
  subscriptionReturnType: (data: any) => T;
};
export type UpdateUserType = {
  updateType: string;
  value?: string;
  numberValue?: number;
};
export type UpdateEnterpriseType = {
  updateType: string;
  enterpriseID: string;
  value?: string;
  numberValue?: number;
};
export type ConnectionACK = {
  type: string;
  payload: {
    query: string;
    subscription_id: string;
    token: string;
  };
};

export type StringID = {
  id: string;
};
export type DeclineHiringInvitation = {
  id: string;
  reason: string;
};
export type UserDescriptor = {
  email: string;
  fullName: string;
  profileIMG: string;
  gender: string;
  pronoun: string;
  loginDate: string;
  ip: string;
  device: string;
  geolocation: {
    latitude: number;
    longitude: number;
  };
  grantedAuthorities: GrantedAuthority[];
};
export type GrantedAuthority = {
  enterpriseID: string;
  authority: string;
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function isUserDescriptor(obj: any): obj is UserDescriptor {
  return (
    obj.email !== undefined &&
    obj.fullName !== undefined &&
    obj.profileIMG !== undefined &&
    obj.ip !== undefined
  );
}
export type Address = {
  country: string;
  county: string;
  city: string;
};

export type Country = {
  id: number;
  name: string;
  phoneCode: string;
};

export type CountryState = {
  id: number;
  name: string;
};

export type City = {
  id: number;
  name: string;
};

export type PhoneCode = {
  id: number;
  code: string;
};
export type EnterpriseChatType = {
  id: string;
  replies: EnterpriseChatReplyType[];
};
export type ReplyEnterpriseChatType = {
  id: string;
  enterpriseName: string;
  senderID: string;
  content: string;
};
export type EnterpriseChatReplyType = {
  id: string;
  chatID: string;
  user: User;
  content: string;
  seen: boolean;
  createdAt: string;
  updatedAt: string;
};
export type User = {
  profileIMG: string;
  email: string;
  firstName: string;
  lastName: string;
  gender: string;
  pronoun: string;
  age: number;
  phone: string;
  address: Address;
  status: string;
  msgs: number | 0;
};
export type UsersForHiring = {
  users: User[];
  count: number;
  currentLimit: number;
  currentOffset: number;
};
export type HiringProccessList = {
  hiringProccesses: HiringProccessType[];
  count: number;
  currentLimit: number;
  currentOffset: number;
};
export type HiringProccessType = {
  eventID: string;
  user: User;
  position: string;
  status: string;
  reason: string;
  lastUpdate: string;
  emitedBy: User;
  occurredOn: string;
};
export type Industry = {
  id: number;
  name: string;
};

export type Enterprise = {
  name: string;
  logoIMG: string;
  bannerIMG: string;
  email: string;
  website: string;
  phone: string;
  industry: string;
  address: Address;
  foundationDate: string;
  ownerID?: string;
  employees?: EmployeeType[];
};

export type HiringInvitationType = {
  eventID: string;
  enterpriseID: string;
  ownerID: string;
  proposedPosition: string;
  fullName: string;
  enterpriseEmail: string;
  enterprisePhone: string;
  enterpriseLogoIMG: string;
  occurredOn: string;
  seen: boolean;
  status: string;
  reason: string;
};
export type EnterpriseEventType = {
  enterpriseName: string;
  eventID: string;
  reason?: string;
};

export type EmployeeType = {
  id: string;
  profileIMG: string;
  firstName: string;
  lastName: string;
  age: number;
  email: string;
  phone: string;
  hiringDate: string;
  approvedHiring: boolean;
  approvedHiringAt: string;
  position: string;
  complaintsSolved: number;
  complaintsSolvedIds: string[];
  complaintsRated: number;
  complaintsRatedIDs: string[];
  complaintsFeedbacked: number;
  complaintsFeedbackedIDs: string[];
  feedbackReceived: number;
  feedbackReceivedIDs: string[];
  hireInvitationsSent: number;
  employeesHired: number;
  employeesFired: number;
};

export type Receiver = {
  id: string;
  fullName: string;
  thumbnail: string;
};
export type UserLog = {
  count: number;
  complaint_rated: ComplaintRated[];
};
export type ComplaintRated = {
  event_id: string;
  complaint_id: string;
  rated_by: string;
  assistant_user_id: string;
  occurred_on: string;
};

export type SolvedReview = {
  User: {
    firstName: string;
    lastName: string;
  };
  Complaint: {
    id: string;
    message: { title: string };
    rating: {
      rate: number;
      comment: string;
    };
  };
};
export type Notifications = {
  id: string;
  ownerID: string;
  thumbnail: string;
  title: string;
  content: string;
  link: string;
  seen: boolean;
  occurredOn: string;
};
export type UserNotificationType = "waiting_for_review" | "hiring_invitation";
export type EnterpriseNotificationType =
  | "employee_waiting_for_approval"
  | "waiting_for_review";
export type WaitingForReview = {
  event_id: string;
  complaint_id: string;
  receiver_id: string;
  triggered_by: string;
  author_id: string;
  occurred_on: string;
  seen: boolean;
};

export type ComplaintReviewType = {
  eventID: string;
  triggeredBy: User;
  complaint: ComplaintType;
  ratedBy: User;
  status: string;
};
export type RateComplaint = {
  complaintId: string;
  eventId: string;
  rate: number;
  comment: string;
};
export type EnterpriseNotifications = {
  count: number;
  employee_waiting_for_approval: EmployeeWaitingForApproval[];
  waiting_for_review: WaitingForReview[];
};
export type EmployeeWaitingForApproval = {
  id: string;
  enterprise_id: string;
  invited_user_id: string;
  proposed_position: string;
  occurred_on: string;
  invitation_id: string;
  seen: boolean;
};

export type EndHiringProcess = {
  pendingEventID: string;
  enterpriseID: string;
  accepted: boolean;
};

export type ComplaintTypeList = {
  complaints: ComplaintType[];
  count: number;
  currentLimit: number;
  currentOffset: number;
};

export type ComplaintType = {
  id: string;
  authorID: string;
  authorFullName: string;
  authorProfileIMG: string;
  receiverID: string;
  receiverFullName: string;
  receiverProfileIMG: string;
  status: string;
  message: Message;
  rating?: Rating;
  createdAt: string;
  updatedAt: string;
  replies?: Reply[];
  industry?: string;
};
export type Message = {
  title: string;
  description: string;
  body: string;
};
export type Rating = {
  rate: number;
  comment: string;
};

export type Reply = {
  id: string;
  complaintID: string;
  senderID: string;
  senderIMG: string;
  senderName: string;
  body: string;
  createdAt: string;
  read: boolean;
  readAt: string;
  updatedAt: string;
  isEnterprise: boolean;
  enterpriseID: string;
  complaintStatus: string;
};
export type ComplaintInfo = {
  complaintsReceived: number;
  complaintsResolved: number;
  complaintsReviewed: number;
  complaintsPending: number;
  averageRating: number;
};
export type CreateAFeedback = {
  enterpriseID: string;
  complaintID?: string;
  feedbackID?: string;
  reviewerID?: string;
  comment?: string;
  color?: string;
  repliesID?: string[];
};
export type FeedbackType = {
  id: string;
  complaintID: string;
  enterpriseID: string;
  replyReview: ReplyReviewType[];
  feedbackAnswers: FeedbackAnswerType[];
  reviewedAt: string;
  updatedAt: string;
  isDone: boolean;
};
export type FeedbackAnswerType = {
  id: string;
  feedbackID: string;
  senderID: string;
  senderIMG: string;
  senderName: string;
  body: string;
  createdAt: string;
  read: boolean;
  readAt: string;
  updatedAt: string;
  isEnterprise: boolean;
  enterpriseID: string;
};
export type ReplyReviewType = {
  id: string;
  feedbackID: string;
  replies: Reply[];
  review: ReviewType;
  reviewer: User;
  color: string;
  createdAt: string;
};
export type ReviewType = {
  replyRevieweID: string;
  comment: string;
};
export type AuthMsg = {
  content: "auth";
  jwt_token: string;
  enterprise_id?: string;
};

export type MarkAsReviewable = {
  complaintID: string;
  enterpriseID: string;
  assistantID: string;
};

export function newAuthMsg(bearer: string, enterpriseID?: string): AuthMsg {
  let jwtToken = bearer.split("20%");
  if (jwtToken.length < 2) {
    jwtToken = bearer.split(" ");
  }
  return {
    content: "auth",
    jwt_token: jwtToken[1],
    enterprise_id: enterpriseID ? enterpriseID : "",
  };
}
export type Sender = {
  thumbnail: string;
  fullName: string;
  isEnterprise: boolean;
  enterpriseID: string;
};

export type WebSocketData = {
  content: "auth" | "reply";
  reply?: Reply;
  success?: boolean;
};
export function newReply(
  senderIMG: string,
  senderName: string,
  body: string,
  isEnterprise: boolean,
  enterpriseID: string
): Reply {
  return {
    id: "",
    complaintID: "",
    senderID: "",
    senderIMG: senderIMG,
    senderName: senderName,
    body: body,
    createdAt: "",
    read: false,
    readAt: "",
    updatedAt: "",
    isEnterprise: isEnterprise,
    enterpriseID: enterpriseID,
    complaintStatus: "",
  };
}
export type ErrorType = {
  [key: string]: string;
};

export type Office = {
  employeeID: string;
  employeeFirstName: string;
  employeePosition: string;
  enterpriseLogoIMG: string;
  enterpriseName: string;
  enterpriseWebsite: string;
  enterprisePhone: string;
  enterpriseEmail: string;
  enterpriseIndustry: string;
  enterpriseAddress: Address;
  ownerFullName: string;
};
//REACT CONTEXT TYPES

//FORM VALIDATION TYPES
export const passwordRegex = new RegExp(
  /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/
);

export const InviteToProjectSchema = z.object({
  enterpriseName: z.string(),
  proposedPosition: z.enum(["Assistant", "Manager"]),
  proposeTo: z.string(),
});

export type InviteToProject = z.infer<typeof InviteToProjectSchema>;

export const SignUpSchema = z
  .object({
    email: z.string().email({ message: "Please enter a valid email" }),
    password: z
      .string()
      .regex(
        passwordRegex,
        "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
      ),
    confirmPassword: z
      .string()
      .regex(
        passwordRegex,
        "Password must contain at least 8 characters, one uppercase letter, one lowercase letter and one number"
      ),
    firstName: z
      .string()
      .min(2, { message: "First name must be at least 2 characters long" })
      .max(50, { message: "First name must be at most 50 characters long" }),
    lastName: z
      .string()
      .min(2, { message: "Last name must be at least 2 characters long" })
      .max(50, { message: "Last name must be at most 50 characters long" }),
    gender: z.enum(["male", "female", "non-declared"], {
      message: "Please select a gender from the list provided",
    }),
    pronoun: z.enum(["he", "she", "they"], {
      message: "Please select a pronoun from the list provided",
    }),
    birthDate: z
      .string()
      .date()
      .transform((val, ctx) => {
        const stringDate = Date.parse(val).toString();
        if (isNaN(parseInt(stringDate))) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please select a valid date",
          });
          return z.NEVER;
        }
        return stringDate;
      }),
    phoneCode: z.string().min(1, {
      message: "Select a country and the phonecode will be automatically added",
    }),
    phone: z
      .string({ message: "We could not validate your phone number" })
      .min(8, { message: "We could not validate your phone number" })
      .transform((val, ctx) => {
        const parsed = parseInt(val);
        if (isNaN(parsed)) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Not a number",
          });
          return z.NEVER;
        }
        return parsed;
      }),
    country: z.string().transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must select a country",
        });
        return z.NEVER;
      }
      return parsed;
    }),
    county: z.string().transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must select a county",
        });
        return z.NEVER;
      }
      return parsed;
    }),
    city: z.string().transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "You must select a city",
        });
        return z.NEVER;
      }
      return parsed;
    }),
    terms: z.enum(["true", "on", "1"], {
      message: "You must accept the terms and conditions",
    }),
  })
  .superRefine(({ confirmPassword, password }, ctx) => {
    if (confirmPassword !== password) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        path: ["confirmPassword"],
        message: "The passwords did not match",
      });
    }
  });
export type CreateUser = z.infer<typeof SignUpSchema>;
export const SignInSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email" }),
  password: z.string().min(1, { message: "Please enter your password" }),
  rememberMe: z.boolean(),
});
export type SignIn = z.infer<typeof SignInSchema>;
export const RegisterEnterpriseSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email" }),
  name: z
    .string()
    .min(3, "The enterprise name should be of at least 3 characters length")
    .max(120, "The enterprise name should be of at most 120 characters length")
    .transform(async (val, ctx) => {
      if (
        !(await Query<boolean>(
          IsEnterpriseNameAvailableQuery,
          IsEnterpriseNameAvailable,
          [val]
        ))
      ) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message:
            "Enterprise name is already taken, please choose a different one",
        });
      }
      return val;
    }),
  website: z.string().url({
    message: "Please enter a valid website e.g: http://www.mywebsite.com",
  }),
  phoneCode: z.string().min(1, {
    message: "Select a country and the phonecode will be automatically added",
  }),
  phone: z
    .string({ message: "We could not validate your phone number" })
    .min(8, { message: "We could not validate your phone number" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed)) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Not a number",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  industryID: z
    .string()
    .min(1, { message: "Please select an industry" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select an industry",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  countryID: z
    .string()
    .min(1, { message: "Please select a country" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select a country",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  countryStateID: z
    .string()
    .min(1, { message: "Please select a state" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select a state",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  cityID: z
    .string()
    .min(1, { message: "Please select a city" })
    .transform((val, ctx) => {
      const parsed = parseInt(val);
      if (isNaN(parsed) || parsed === 0) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please select a city",
        });
        return z.NEVER;
      }
      return parsed;
    }),
  foundationDate: z.coerce.date().transform((val, ctx) => {
    const stringDate = val.getMilliseconds().toString();
    if (isNaN(parseInt(stringDate))) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Please select a valid date",
      });
      return z.NEVER;
    }
    return stringDate;
  }),
  terms: z.enum(["true", "on", "1"], {
    message: "You must accept the terms and conditions",
  }),
});
export type RegisterEnterprise = z.infer<typeof RegisterEnterpriseSchema>;

export const SendComplaintValidationSchema = z.object({
  authorID: z.string(),
  receiverID: z.string({
    message: "Please select a receiver from the list and click it.",
  }),
  receiverFullName: z.string(),
  receiverProfileIMG: z.string(),
  title: z
    .string({ message: "Please provide a reason for the complaint" })
    .min(10, { message: "Reason must be at least 10 characters long" })
    .max(80, { message: "Reason must be at most 80 characters long" }),
  description: z
    .string({ message: "Please provide a description" })
    .min(30, {
      message: "Please describe the problem with at least 30 characters",
    })
    .max(120, {
      message:
        "Problem description must be at most 120 characters long, later you can tell us more about it",
    }),
  content: z
    .string()
    .min(50, {
      message:
        "If you reached this point complain about the problem in at least 50 characters",
    })
    .max(250, {
      message:
        "Hold on! 250 characters is the limit, you can still chat with him later",
    }),
});

export type SendComplaintType = z.infer<typeof SendComplaintValidationSchema>;

export const ReplyComplaintValidationSchema = z.object({
  complaintID: z.string(),
  replyAuthorID: z.string(),
  replyBody: z
    .string()
    .min(1, { message: "Please write at least 10 characters" })
    .max(120, { message: "Reply must be at most 250 characters long" }),
  replyEnterpriseID: z.string().optional(),
});

export type ReplyComplaintType = z.infer<typeof ReplyComplaintValidationSchema>;

export const RateValidationSchema = z.object({
  rate: z
    .number()
    .min(0, { message: "Please rate the attention" })
    .max(5, { message: "Rate must be at most 5" }),
  comment: z
    .string()
    .min(3, { message: "Please write at least one word about the attention" })
    .max(250, { message: "Comment must be at most 250 characters long" }),
});

export const ConfirmationCodeValidationSchema = z.object({
  confirmationCode: z
    .string()
    .length(7, { message: "Please enter a valid confirmation code" })
    .transform((val, ctx) => {
      const segments = val.split("");
      for (const segment of segments) {
        if (!segment.match(/^[0-9]+$/)) {
          ctx.addIssue({
            code: z.ZodIssueCode.custom,
            message: "Please enter a valid confirmation code",
          });
          return z.NEVER;
        }
      }
      const parsed = parseInt(val);
      if (isNaN(parsed)) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: "Please enter a valid confirmation code",
        });
        return z.NEVER;
      }
      return parsed;
    }),
});

export type ConfirmationCode = z.infer<typeof ConfirmationCodeValidationSchema>;

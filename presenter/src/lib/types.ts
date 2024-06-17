import { z } from "zod";
import {
  Query,
  IsEnterpriseNameAvailableQuery,
  IsEnterpriseNameAvailable,
  IsValidReceiverQuery,
  IsValidReceiver,
} from "./queries";
export type StringID = {
  id: string;
};
export type UserDescriptor = {
  email: string;
  fullName: string;
  profileIMG: string;
  ip: string;
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
};

export type County = {
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

export type CreateUser = {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  birthDate: string;
  phone: string;
  country: number;
  county: number;
  city: number;
};
export type User = {
  profileIMG: string;
  email: string;
  firstName: string;
  lastName: string;
  age: number;
  phone: string;
  address: Address;
};
export type UsersForHiring = {
  users: User[];
  count: number;
  currentLimit: number;
  currentOffset: number;
};

export type Industry = {
  id: number;
  name: string;
};
export type CreateEnterprise = {
  name: string;
  email: string;
  website: string;
  phone: string;
  industry: string;
  country: number;
  county: number;
  city: number;
  foundationDate: string;
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
};
export type Employee = {
  ID: string;
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
  complaintsSolved: string;
  complaintsSolvedIDs: string[];
};
export type InviteToProject = {
  enterpriseName: string;
  userEmail: string;
  userFullName: string;
  position: string;
};
export type Receiver = {
  ID: string;
  fullName: string;
  IMG: string;
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
export type UserNotifications = {
  count: number;
  hiring_invitation: HiringInvitation[];
  waiting_for_review: WaitingForReview[];
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
export type HiringInvitation = {
  event_id: string;
  enterprise_id: string;
  position_proposal: string;
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  age: number;
  occurred_on: string;
  seen: boolean;
};
export type InfoForReview = {
  User: {
    firstName: string;
    lastName: string;
  };
  Complaint: {
    id: string;
    receiverFullName: string;
    message: {
      title: string;
    };
  };
};
export type RateComplaint = {
  notificationID: string;
  complaintID: string;
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

export type SendComplaint = {
  fullName: string;
  senderID: string;
  receiverID: string;
  reason: string;
  description: string;
  body: string;
};
export type ComplaintTypeList = {
  complaints: Complaint[];
  count: number;
  currentLimit: number;
  currentOffset: number;
};
export type DeserializedComplaint = {
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
  replies?: DeserializedReply[];
  industry?: string;
};

export type Complaint = {
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
  replies?: DeserializedReply[];
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
  complaint_id: string;
  sender_id: string;
  sender_img: string;
  sender_name: string;
  body: string;
  created_at: string;
  read: boolean;
  read_at: string;
  updated_at: string;
  is_enterprise: boolean;
  enterprise_id: string;
};
export type DeserializedReply = {
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
};

export type CreateAFeedback = {
  complaintID: string;
  reviewerID: string;
  reviewedID: string;
  reviewerIMG: string;
  reviewerName: string;
  senderID: string;
  senderIMG: string;
  senderName: string;
  body: string;
  createdAt: string;
  read: boolean;
  readAt: string;
  updatedAt: string;
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
    complaint_id: "",
    sender_id: "",
    sender_img: senderIMG,
    sender_name: senderName,
    body: body,
    created_at: "",
    read: false,
    read_at: "",
    updated_at: "",
    is_enterprise: isEnterprise,
    enterprise_id: enterpriseID,
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
export type ComplaintState = {
  complaintData?: SendComplaint;
  updateState: (newState: Partial<ComplaintState>) => void;
};
export type UserState = {
  userSession: UserDescriptor | null;
  userNotifications: UserNotifications | null;
  updateState: (newState?: Partial<UserState>) => void;
};
//FORM VALIDATION TYPES
const passwordRegex = new RegExp(/^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/);
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
    birthDate: z.coerce.date(),
    phoneCode: z.string().min(1, {
      message: "Select a country and the phonecode will be automatically added",
    }),
    phone: z
      .string({ message: "We could not validate your phone number" })
      .min(6, { message: "We could not validate your phone number" })
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

export const SignInSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email" }),
  password: z.string().min(1, { message: "Please enter your password" }),
  rememberMe: z
    .enum(["true", "false", "on", "off", "1", "0"])
    .default("off")
    .transform((val) => val === "true" || val === "on" || val === "1"),
});

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
    .min(6, { message: "We could not validate your phone number" })
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
  industry: z.string().min(1, { message: "Please select an industry" }),
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
  foundationDate: z.coerce.date(),
  terms: z.enum(["true", "on", "1"], {
    message: "You must accept the terms and conditions",
  }),
});

export const ReceiverValidationSchema = z.object({
  term: z
    .string()
    .min(1, { message: "Please select a receiver" })
    .transform(async (val, ctx) => {
      const receivers = await Query<boolean>(
        IsValidReceiverQuery,
        IsValidReceiver,
        [val]
      );
      if (!receivers) {
        ctx.addIssue({
          code: z.ZodIssueCode.custom,
          message: `The receiver does not exist, please select a valid one or write the exact full name`,
        });
      }
      return val;
    }),
});

export const DescriptionValidationSchema = z.object({
  reason: z
    .string()
    .min(10, { message: "Reason must be at least 10 characters long" })
    .max(80, { message: "Reason must be at most 80 characters long" }),
  description: z
    .string()
    .min(30, {
      message: "Please describe the problem with at least 30 characters",
    })
    .max(120, {
      message:
        "Problem description must be at most 120 characters long, later you can tell us more about it",
    }),
});

export const ComplaintBodyValidationSchema = z.object({
  body: z
    .string()
    .min(50, {
      message:
        "If you reached this point complain about the problem with at least 50 characters",
    })
    .max(250, {
      message:
        "Hold on! 250 characters is the limit, you can still chat with him later",
    }),
});

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

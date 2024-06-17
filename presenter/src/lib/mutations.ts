import {
  CreateEnterprise,
  CreateUser,
  EndHiringProcess,
  InviteToProject,
  MarkAsReviewable,
  RateComplaint,
  SendComplaint,
  StringID,
} from "./types";

export const CreateUserMutation = ({
  email,
  password,
  firstName,
  lastName,
  birthDate,
  phone,
  country,
  county,
  city,
}: CreateUser): string => `
    
        mutation {
            CreateUser(
                email: "${email}",
                password: "${password}",
                firstName: "${firstName}",
                lastName: "${lastName}",
                birthDate: "${birthDate}",
                phone: "${phone}",
                country: ${country},
                county: ${county},
                city: ${city},
            )
        }
    `;

export const CreateEnterpriseMutation = ({
  name,
  email,
  website,
  phone,
  industry,
  country,
  county,
  city,
  foundationDate,
}: CreateEnterprise): string => `
mutation {
    CreateEnterprise(
        name: "${name}",
        email: "${email}",
        website: "${website}",
        phone: "${phone}",
        industry: "${industry}",
        country: ${country},
        county: ${county},
        city: ${city},
        foundationDate: "${foundationDate}",
    )
}
`;

export const SendComplaintMutation = ({
  senderID,
  receiverID,
  reason,
  description,
  body,
}: SendComplaint): string => `
mutation {
    SendComplaint(
        authorID: "${senderID}",
        receiverID: "${receiverID}",
        title: "${reason}",
        description: "${description}",
        content: "${body}",
    )
}
`;
export const InviteToProjectMutation = ({
  enterpriseName,
  userEmail,
  position,
}: InviteToProject): string => {
  return `
  mutation {
    InviteToProject(
      enterpriseName: "${enterpriseName}",
      userEmail: "${userEmail}",
      position: "${position}"
    )
  }
  `;
};

export const AcceptHiringInvitationMutation = (id: StringID): string => {
  return `
  mutation {
    AcceptHiringInvitation(
      ID: "${id.id}"
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

export const MarkAsReviewableMutation = (mar: MarkAsReviewable): string => {
  return `mutation {
    MarkAsReviewable(
      complaintID: "${mar.complaintID}",
      enterpriseID: "${mar.enterpriseID}",
      assistantID: "${mar.assistantID}",
    )
  }
  `;
};

export const RateComplaintMutation = (rate: RateComplaint): string => {
  return `mutation {
    RateComplaint(
      notificationID: "${rate.notificationID}",
      complaintID: "${rate.complaintID}",
      rate: ${rate.rate},
      comment: "${rate.comment}",
      )
    }`;
};

export const Mutation = async <T>(
  mutationFn: (data: T) => string,
  arg: T
): Promise<void> => {
  console.log(mutationFn(arg));
  return fetch(import.meta.env.VITE_GRAPHQL_ENDPOINT, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify({ query: deleteLinebreaks(mutationFn(arg)) }),
  })
    .then((res) => res.json())
    .then((data) => {
      const keys = Object.keys(data);
      if (keys.includes("errors")) {
        throw new Error(data.errors[0].message);
      }
    });
};

const deleteLinebreaks = (str: string) => str.replace(/(\r\n|\n|\r)/gm, "");

import NotificationType from "../lib/types/notificationType";
import Roles from "../lib/types/rolesEnum";
import UserDescriptor from "../lib/types/userDescriptorType";

export const userDescriptor: UserDescriptor = {
  email: "john_doe@email.com",
  fullName: "John Doe",
  profileImg: "/default.jpg",
  gender: "male",
  pronoun: "female",
  loginDate: "1720359265225",
  ip: "0.0.0.0",
  device: "userAgent",
  geolocation: {
    latitude: 0,
    longitude: 0,
  },
  grantedAuthorities: [
    {
      enterpriseId: "enterpriseId",
      authority: Roles.OWNER,
    },
  ],
};

export const notificationsList: NotificationType[] = [
  {
    id: "1",
    ownerId: userDescriptor.email,
    thumbnail: "/default.jpg",
    title: "You have received a new complaint",
    content: "Jane Doe has sent you a complaint",
    link: "/complaints-received?id=10297522",
    seen: false,
    occurredOn: "1720359265225",
  },
  {
    id: "2",
    ownerId: userDescriptor.email,
    thumbnail: "/default.jpg",
    title: "You have been ask to rate someone attention",
    content: "Jane Doe has ask you for a review on her attention",
    link: "/reviews?id=99302721",
    seen: false,
    occurredOn: "1720359265225",
  },
  {
    id: "3",
    ownerId: userDescriptor.email,
    thumbnail: "/default.jpg",
    title: "You have receive a hiring invitation",
    content: "Go Complaint invited you to join them",
    link: "/hiring-invitations?id=98712203121",
    seen: false,
    occurredOn: "1720359265225",
  },
];

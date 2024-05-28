import React from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import "./index.css";
import Root from "./Root";
import { ErrorPage } from "./pages/ErrorPage";
import SignIn from "./pages/SignIn";
import Home from "./pages/Home";
import SignUp from "./pages/SignUp";
import SuccessPage from "./pages/SuccessPage";
import { AcceptInvitationLoader, ComplaintLoader, EmployeesLoader, EnterpriseInboxLoader, EnterpriseLoader, EnterpriseSentComplaintLoader, EnterpriseSentLoader, FindReceiverLoader, HireLoader, HiringLoader, PendingHiresLoader, ProfileLoader, RegisterEnterpriseLoader, RootLoader, SignUpLoader, UserInboxLoader, UserSendComplaintLoader, UserSentLoader } from "./lib/loaders";
import { RegisterEnterpriseAction, SignInAction, SignUpAction } from "./lib/actions";
import ProfileLayout from "./ProfileLayout";
import RegisterEnterprise from "./pages/RegisterEnterprise";
import Profile from "./pages/Profile";
import DescribeComplaint from "./pages/DescribeComplaint";
import Complain from "./pages/Complain";
import Inbox from "./pages/Inbox";
import Sent from "./pages/Sent";
import ComplaintPage from "./pages/ComplaintPage";
import EnterpriseLayout from "./EnterpriseLayout";
import Enterprise from "./pages/Enterprise";

import Hiring from "./pages/Hiring";
import Hire from "./pages/Hire";
import AcceptInvitation from "./pages/AcceptInvitation";
import PendingHires from "./pages/PendingHires";
import Employees from "./pages/Employees";
import FindReceiver from "./pages/FindReceiver";
const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    loader: RootLoader,
    children: [
      {
        path: "/",
        element: <Home />,
      },
      {
        path: "/sign-in",
        element: <SignIn />,
        action: SignInAction,
      },
      {
        path: "/sign-up",
        element: <SignUp />,
        loader: SignUpLoader,
        action: SignUpAction,
      },
      {
        path: "/success/:subject",
        element: <SuccessPage />,
      },
      {
        path: "/invitation/:type/:id",
        element: <AcceptInvitation />,
        loader: AcceptInvitationLoader
      }
    ],
  },
  {
    path: "/profile",
    element: <ProfileLayout />,
    errorElement: <ErrorPage />,
    loader: ProfileLoader,
    children: [
      { index: true, element: <Profile /> },
      {
        path: "/profile/register-enterprise",
        element: <RegisterEnterprise />,
        loader: RegisterEnterpriseLoader,
        action: RegisterEnterpriseAction,
      },
      {
        path: "/profile/complaint",
        element: <FindReceiver />,
        loader: FindReceiverLoader,
      },
      {
        path: "/profile/complaint/describe",
        element: <DescribeComplaint />
      },
      {
        path: "/profile/complaint/complain",
        element: <Complain />,
        loader: UserSendComplaintLoader,
      },
      {
        path: "/profile/inbox",
        element: <Inbox />,
        loader: UserInboxLoader,
      },
      {
        path: "/profile/:complaintID",
        element: <ComplaintPage />,
        loader: ComplaintLoader,
      },
      {
        path: "/profile/sent",
        element: <Sent />,
        loader: UserSentLoader,
      }
    ],
  },
  {
    path: "/enterprises/:id",
    element: <EnterpriseLayout />,
    errorElement: <ErrorPage />,
    loader: EnterpriseLoader,
    children: [
      { index: true, element: <Enterprise /> },
      {
        path: "/enterprises/:id/complaint",
        element: <FindReceiver />,
        loader: FindReceiverLoader,
      },
      {
        path: "/enterprises/:id/complaint/describe",
        element: <DescribeComplaint />
      },
      {
        path: "/enterprises/:id/complaint/complain",
        element: <Complain />,
        loader: EnterpriseSentComplaintLoader
      },
      {
        path: "/enterprises/:id/inbox",
        element: <Inbox />,
        loader: EnterpriseInboxLoader,
      },
      {
        path: "/enterprises/:id/sent",
        element: <Sent />,
        loader: EnterpriseSentLoader,
      },
      {
        path: "/enterprises/:id/:complaintID",
        element: <ComplaintPage />,
        loader: ComplaintLoader,
      },
      {
        path: "/enterprises/:id/hiring",
        element: <Hiring />,
        loader: HiringLoader,
      },
      {
        path: "/enterprises/:id/employees",
        element: <Employees />,
        loader: EmployeesLoader,
      },
      {
        path: "/enterprises/:id/hiring/:userID",
        element: <Hire />,
        loader: HireLoader,
      },
      {
        path: "/enterprises/:id/pending",
        element: <PendingHires />,
        loader: PendingHiresLoader,
      }
    ],
  }
]);


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)

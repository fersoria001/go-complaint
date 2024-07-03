/* prettier-ignore-start */

/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file is auto-generated by TanStack Router

import { createFileRoute } from '@tanstack/react-router'

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as SignUpImport } from './routes/sign-up'
import { Route as LicensingImport } from './routes/licensing'
import { Route as ContactImport } from './routes/contact'
import { Route as ConfirmationImport } from './routes/confirmation'
import { Route as ProfileImport } from './routes/_profile'
import { Route as EnterpriseIDImport } from './routes/$enterpriseID'
import { Route as EnterpriseIDIndexImport } from './routes/$enterpriseID/index'
import { Route as ProfileSettingsImport } from './routes/_profile/settings'
import { Route as ProfileRegisterEnterpriseImport } from './routes/_profile/register-enterprise'
import { Route as ProfileProfileImport } from './routes/_profile/profile'
import { Route as ProfileHiringInvitationsImport } from './routes/_profile/hiring-invitations'
import { Route as ProfileComplaintSentImport } from './routes/_profile/complaint-sent'
import { Route as ProfileSendComplaintImport } from './routes/_profile/_send-complaint'
import { Route as EnterpriseIDSuccessImport } from './routes/$enterpriseID/success'
import { Route as EnterpriseIDSettingsImport } from './routes/$enterpriseID/settings'
import { Route as EnterpriseIDHiringProccesesImport } from './routes/$enterpriseID/hiring-procceses'
import { Route as EnterpriseIDHiringImport } from './routes/$enterpriseID/hiring'
import { Route as EnterpriseIDHireImport } from './routes/$enterpriseID/hire'
import { Route as EnterpriseIDComplaintSentImport } from './routes/$enterpriseID/complaint-sent'
import { Route as EnterpriseIDSendComplaintImport } from './routes/$enterpriseID/_send-complaint'
import { Route as ProfileSentIndexImport } from './routes/_profile/sent/index'
import { Route as ProfileReviewsIndexImport } from './routes/_profile/reviews/index'
import { Route as ProfileInboxIndexImport } from './routes/_profile/inbox/index'
import { Route as ProfileHistoryIndexImport } from './routes/_profile/history/index'
import { Route as EnterpriseIDSentIndexImport } from './routes/$enterpriseID/sent/index'
import { Route as EnterpriseIDReviewsIndexImport } from './routes/$enterpriseID/reviews/index'
import { Route as EnterpriseIDInboxIndexImport } from './routes/$enterpriseID/inbox/index'
import { Route as EnterpriseIDHistoryIndexImport } from './routes/$enterpriseID/history/index'
import { Route as EnterpriseIDFeedbacksIndexImport } from './routes/$enterpriseID/feedbacks/index'
import { Route as EnterpriseIDFeedbackIndexImport } from './routes/$enterpriseID/feedback/index'
import { Route as EnterpriseIDEmployeesIndexImport } from './routes/$enterpriseID/employees/index'
import { Route as ProfileSendComplaintSendComplaintImport } from './routes/_profile/_send-complaint/send-complaint'
import { Route as ProfileSendComplaintDescribeComplaintImport } from './routes/_profile/_send-complaint/describe-complaint'
import { Route as ProfileSendComplaintComplainImport } from './routes/_profile/_send-complaint/complain'
import { Route as EnterpriseIDSendComplaintSendComplaintImport } from './routes/$enterpriseID/_send-complaint/send-complaint'
import { Route as EnterpriseIDSendComplaintDescribeComplaintImport } from './routes/$enterpriseID/_send-complaint/describe-complaint'
import { Route as EnterpriseIDSendComplaintComplainImport } from './routes/$enterpriseID/_send-complaint/complain'
import { Route as ProfileSentComplaintIdIndexImport } from './routes/_profile/sent/$complaintId/index'
import { Route as ProfileInboxComplaintIdIndexImport } from './routes/_profile/inbox/$complaintId/index'
import { Route as EnterpriseIDSentComplaintIdIndexImport } from './routes/$enterpriseID/sent/$complaintId/index'
import { Route as EnterpriseIDInboxComplaintIdIndexImport } from './routes/$enterpriseID/inbox/$complaintId/index'
import { Route as EnterpriseIDFeedbackEmployeeIdIndexImport } from './routes/$enterpriseID/feedback/$employeeId/index'
import { Route as ProfileSentComplaintIdChatImport } from './routes/_profile/sent/$complaintId/chat'
import { Route as ProfileInboxComplaintIdChatImport } from './routes/_profile/inbox/$complaintId/chat'
import { Route as EnterpriseIDSentComplaintIdChatImport } from './routes/$enterpriseID/sent/$complaintId/chat'
import { Route as EnterpriseIDInboxComplaintIdChatImport } from './routes/$enterpriseID/inbox/$complaintId/chat'
import { Route as EnterpriseIDEmployeesSolvedEmployeeIdImport } from './routes/$enterpriseID/employees/solved/$employeeId'

// Create Virtual Routes

const SignInLazyImport = createFileRoute('/sign-in')()
const PrivacyLazyImport = createFileRoute('/privacy')()
const ErrorsLazyImport = createFileRoute('/errors')()
const AboutLazyImport = createFileRoute('/about')()
const IndexLazyImport = createFileRoute('/')()

// Create/Update Routes

const SignInLazyRoute = SignInLazyImport.update({
  path: '/sign-in',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/sign-in.lazy').then((d) => d.Route))

const PrivacyLazyRoute = PrivacyLazyImport.update({
  path: '/privacy',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/privacy.lazy').then((d) => d.Route))

const ErrorsLazyRoute = ErrorsLazyImport.update({
  path: '/errors',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/errors.lazy').then((d) => d.Route))

const AboutLazyRoute = AboutLazyImport.update({
  path: '/about',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/about.lazy').then((d) => d.Route))

const SignUpRoute = SignUpImport.update({
  path: '/sign-up',
  getParentRoute: () => rootRoute,
} as any)

const LicensingRoute = LicensingImport.update({
  path: '/licensing',
  getParentRoute: () => rootRoute,
} as any)

const ContactRoute = ContactImport.update({
  path: '/contact',
  getParentRoute: () => rootRoute,
} as any)

const ConfirmationRoute = ConfirmationImport.update({
  path: '/confirmation',
  getParentRoute: () => rootRoute,
} as any)

const ProfileRoute = ProfileImport.update({
  id: '/_profile',
  getParentRoute: () => rootRoute,
} as any)

const EnterpriseIDRoute = EnterpriseIDImport.update({
  path: '/$enterpriseID',
  getParentRoute: () => rootRoute,
} as any)

const IndexLazyRoute = IndexLazyImport.update({
  path: '/',
  getParentRoute: () => rootRoute,
} as any).lazy(() => import('./routes/index.lazy').then((d) => d.Route))

const EnterpriseIDIndexRoute = EnterpriseIDIndexImport.update({
  path: '/',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const ProfileSettingsRoute = ProfileSettingsImport.update({
  path: '/settings',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileRegisterEnterpriseRoute = ProfileRegisterEnterpriseImport.update({
  path: '/register-enterprise',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileProfileRoute = ProfileProfileImport.update({
  path: '/profile',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileHiringInvitationsRoute = ProfileHiringInvitationsImport.update({
  path: '/hiring-invitations',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileComplaintSentRoute = ProfileComplaintSentImport.update({
  path: '/complaint-sent',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileSendComplaintRoute = ProfileSendComplaintImport.update({
  id: '/_send-complaint',
  getParentRoute: () => ProfileRoute,
} as any)

const EnterpriseIDSuccessRoute = EnterpriseIDSuccessImport.update({
  path: '/success',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDSettingsRoute = EnterpriseIDSettingsImport.update({
  path: '/settings',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDHiringProccesesRoute =
  EnterpriseIDHiringProccesesImport.update({
    path: '/hiring-procceses',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

const EnterpriseIDHiringRoute = EnterpriseIDHiringImport.update({
  path: '/hiring',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDHireRoute = EnterpriseIDHireImport.update({
  path: '/hire',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDComplaintSentRoute = EnterpriseIDComplaintSentImport.update({
  path: '/complaint-sent',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDSendComplaintRoute = EnterpriseIDSendComplaintImport.update({
  id: '/_send-complaint',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const ProfileSentIndexRoute = ProfileSentIndexImport.update({
  path: '/sent/',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileReviewsIndexRoute = ProfileReviewsIndexImport.update({
  path: '/reviews/',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileInboxIndexRoute = ProfileInboxIndexImport.update({
  path: '/inbox/',
  getParentRoute: () => ProfileRoute,
} as any)

const ProfileHistoryIndexRoute = ProfileHistoryIndexImport.update({
  path: '/history/',
  getParentRoute: () => ProfileRoute,
} as any)

const EnterpriseIDSentIndexRoute = EnterpriseIDSentIndexImport.update({
  path: '/sent/',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDReviewsIndexRoute = EnterpriseIDReviewsIndexImport.update({
  path: '/reviews/',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDInboxIndexRoute = EnterpriseIDInboxIndexImport.update({
  path: '/inbox/',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDHistoryIndexRoute = EnterpriseIDHistoryIndexImport.update({
  path: '/history/',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDFeedbacksIndexRoute = EnterpriseIDFeedbacksIndexImport.update(
  {
    path: '/feedbacks/',
    getParentRoute: () => EnterpriseIDRoute,
  } as any,
)

const EnterpriseIDFeedbackIndexRoute = EnterpriseIDFeedbackIndexImport.update({
  path: '/feedback/',
  getParentRoute: () => EnterpriseIDRoute,
} as any)

const EnterpriseIDEmployeesIndexRoute = EnterpriseIDEmployeesIndexImport.update(
  {
    path: '/employees/',
    getParentRoute: () => EnterpriseIDRoute,
  } as any,
)

const ProfileSendComplaintSendComplaintRoute =
  ProfileSendComplaintSendComplaintImport.update({
    path: '/send-complaint',
    getParentRoute: () => ProfileSendComplaintRoute,
  } as any)

const ProfileSendComplaintDescribeComplaintRoute =
  ProfileSendComplaintDescribeComplaintImport.update({
    path: '/describe-complaint',
    getParentRoute: () => ProfileSendComplaintRoute,
  } as any)

const ProfileSendComplaintComplainRoute =
  ProfileSendComplaintComplainImport.update({
    path: '/complain',
    getParentRoute: () => ProfileSendComplaintRoute,
  } as any)

const EnterpriseIDSendComplaintSendComplaintRoute =
  EnterpriseIDSendComplaintSendComplaintImport.update({
    path: '/send-complaint',
    getParentRoute: () => EnterpriseIDSendComplaintRoute,
  } as any)

const EnterpriseIDSendComplaintDescribeComplaintRoute =
  EnterpriseIDSendComplaintDescribeComplaintImport.update({
    path: '/describe-complaint',
    getParentRoute: () => EnterpriseIDSendComplaintRoute,
  } as any)

const EnterpriseIDSendComplaintComplainRoute =
  EnterpriseIDSendComplaintComplainImport.update({
    path: '/complain',
    getParentRoute: () => EnterpriseIDSendComplaintRoute,
  } as any)

const ProfileSentComplaintIdIndexRoute =
  ProfileSentComplaintIdIndexImport.update({
    path: '/sent/$complaintId/',
    getParentRoute: () => ProfileRoute,
  } as any)

const ProfileInboxComplaintIdIndexRoute =
  ProfileInboxComplaintIdIndexImport.update({
    path: '/inbox/$complaintId/',
    getParentRoute: () => ProfileRoute,
  } as any)

const EnterpriseIDSentComplaintIdIndexRoute =
  EnterpriseIDSentComplaintIdIndexImport.update({
    path: '/sent/$complaintId/',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

const EnterpriseIDInboxComplaintIdIndexRoute =
  EnterpriseIDInboxComplaintIdIndexImport.update({
    path: '/inbox/$complaintId/',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

const EnterpriseIDFeedbackEmployeeIdIndexRoute =
  EnterpriseIDFeedbackEmployeeIdIndexImport.update({
    path: '/feedback/$employeeId/',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

const ProfileSentComplaintIdChatRoute = ProfileSentComplaintIdChatImport.update(
  {
    path: '/sent/$complaintId/chat',
    getParentRoute: () => ProfileRoute,
  } as any,
)

const ProfileInboxComplaintIdChatRoute =
  ProfileInboxComplaintIdChatImport.update({
    path: '/inbox/$complaintId/chat',
    getParentRoute: () => ProfileRoute,
  } as any)

const EnterpriseIDSentComplaintIdChatRoute =
  EnterpriseIDSentComplaintIdChatImport.update({
    path: '/sent/$complaintId/chat',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

const EnterpriseIDInboxComplaintIdChatRoute =
  EnterpriseIDInboxComplaintIdChatImport.update({
    path: '/inbox/$complaintId/chat',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

const EnterpriseIDEmployeesSolvedEmployeeIdRoute =
  EnterpriseIDEmployeesSolvedEmployeeIdImport.update({
    path: '/employees/solved/$employeeId',
    getParentRoute: () => EnterpriseIDRoute,
  } as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexLazyImport
      parentRoute: typeof rootRoute
    }
    '/$enterpriseID': {
      id: '/$enterpriseID'
      path: '/$enterpriseID'
      fullPath: '/$enterpriseID'
      preLoaderRoute: typeof EnterpriseIDImport
      parentRoute: typeof rootRoute
    }
    '/_profile': {
      id: '/_profile'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof ProfileImport
      parentRoute: typeof rootRoute
    }
    '/confirmation': {
      id: '/confirmation'
      path: '/confirmation'
      fullPath: '/confirmation'
      preLoaderRoute: typeof ConfirmationImport
      parentRoute: typeof rootRoute
    }
    '/contact': {
      id: '/contact'
      path: '/contact'
      fullPath: '/contact'
      preLoaderRoute: typeof ContactImport
      parentRoute: typeof rootRoute
    }
    '/licensing': {
      id: '/licensing'
      path: '/licensing'
      fullPath: '/licensing'
      preLoaderRoute: typeof LicensingImport
      parentRoute: typeof rootRoute
    }
    '/sign-up': {
      id: '/sign-up'
      path: '/sign-up'
      fullPath: '/sign-up'
      preLoaderRoute: typeof SignUpImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutLazyImport
      parentRoute: typeof rootRoute
    }
    '/errors': {
      id: '/errors'
      path: '/errors'
      fullPath: '/errors'
      preLoaderRoute: typeof ErrorsLazyImport
      parentRoute: typeof rootRoute
    }
    '/privacy': {
      id: '/privacy'
      path: '/privacy'
      fullPath: '/privacy'
      preLoaderRoute: typeof PrivacyLazyImport
      parentRoute: typeof rootRoute
    }
    '/sign-in': {
      id: '/sign-in'
      path: '/sign-in'
      fullPath: '/sign-in'
      preLoaderRoute: typeof SignInLazyImport
      parentRoute: typeof rootRoute
    }
    '/$enterpriseID/_send-complaint': {
      id: '/$enterpriseID/_send-complaint'
      path: ''
      fullPath: '/$enterpriseID'
      preLoaderRoute: typeof EnterpriseIDSendComplaintImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/complaint-sent': {
      id: '/$enterpriseID/complaint-sent'
      path: '/complaint-sent'
      fullPath: '/$enterpriseID/complaint-sent'
      preLoaderRoute: typeof EnterpriseIDComplaintSentImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/hire': {
      id: '/$enterpriseID/hire'
      path: '/hire'
      fullPath: '/$enterpriseID/hire'
      preLoaderRoute: typeof EnterpriseIDHireImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/hiring': {
      id: '/$enterpriseID/hiring'
      path: '/hiring'
      fullPath: '/$enterpriseID/hiring'
      preLoaderRoute: typeof EnterpriseIDHiringImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/hiring-procceses': {
      id: '/$enterpriseID/hiring-procceses'
      path: '/hiring-procceses'
      fullPath: '/$enterpriseID/hiring-procceses'
      preLoaderRoute: typeof EnterpriseIDHiringProccesesImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/settings': {
      id: '/$enterpriseID/settings'
      path: '/settings'
      fullPath: '/$enterpriseID/settings'
      preLoaderRoute: typeof EnterpriseIDSettingsImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/success': {
      id: '/$enterpriseID/success'
      path: '/success'
      fullPath: '/$enterpriseID/success'
      preLoaderRoute: typeof EnterpriseIDSuccessImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/_profile/_send-complaint': {
      id: '/_profile/_send-complaint'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof ProfileSendComplaintImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/complaint-sent': {
      id: '/_profile/complaint-sent'
      path: '/complaint-sent'
      fullPath: '/complaint-sent'
      preLoaderRoute: typeof ProfileComplaintSentImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/hiring-invitations': {
      id: '/_profile/hiring-invitations'
      path: '/hiring-invitations'
      fullPath: '/hiring-invitations'
      preLoaderRoute: typeof ProfileHiringInvitationsImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/profile': {
      id: '/_profile/profile'
      path: '/profile'
      fullPath: '/profile'
      preLoaderRoute: typeof ProfileProfileImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/register-enterprise': {
      id: '/_profile/register-enterprise'
      path: '/register-enterprise'
      fullPath: '/register-enterprise'
      preLoaderRoute: typeof ProfileRegisterEnterpriseImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/settings': {
      id: '/_profile/settings'
      path: '/settings'
      fullPath: '/settings'
      preLoaderRoute: typeof ProfileSettingsImport
      parentRoute: typeof ProfileImport
    }
    '/$enterpriseID/': {
      id: '/$enterpriseID/'
      path: '/'
      fullPath: '/$enterpriseID/'
      preLoaderRoute: typeof EnterpriseIDIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/_send-complaint/complain': {
      id: '/$enterpriseID/_send-complaint/complain'
      path: '/complain'
      fullPath: '/$enterpriseID/complain'
      preLoaderRoute: typeof EnterpriseIDSendComplaintComplainImport
      parentRoute: typeof EnterpriseIDSendComplaintImport
    }
    '/$enterpriseID/_send-complaint/describe-complaint': {
      id: '/$enterpriseID/_send-complaint/describe-complaint'
      path: '/describe-complaint'
      fullPath: '/$enterpriseID/describe-complaint'
      preLoaderRoute: typeof EnterpriseIDSendComplaintDescribeComplaintImport
      parentRoute: typeof EnterpriseIDSendComplaintImport
    }
    '/$enterpriseID/_send-complaint/send-complaint': {
      id: '/$enterpriseID/_send-complaint/send-complaint'
      path: '/send-complaint'
      fullPath: '/$enterpriseID/send-complaint'
      preLoaderRoute: typeof EnterpriseIDSendComplaintSendComplaintImport
      parentRoute: typeof EnterpriseIDSendComplaintImport
    }
    '/_profile/_send-complaint/complain': {
      id: '/_profile/_send-complaint/complain'
      path: '/complain'
      fullPath: '/complain'
      preLoaderRoute: typeof ProfileSendComplaintComplainImport
      parentRoute: typeof ProfileSendComplaintImport
    }
    '/_profile/_send-complaint/describe-complaint': {
      id: '/_profile/_send-complaint/describe-complaint'
      path: '/describe-complaint'
      fullPath: '/describe-complaint'
      preLoaderRoute: typeof ProfileSendComplaintDescribeComplaintImport
      parentRoute: typeof ProfileSendComplaintImport
    }
    '/_profile/_send-complaint/send-complaint': {
      id: '/_profile/_send-complaint/send-complaint'
      path: '/send-complaint'
      fullPath: '/send-complaint'
      preLoaderRoute: typeof ProfileSendComplaintSendComplaintImport
      parentRoute: typeof ProfileSendComplaintImport
    }
    '/$enterpriseID/employees/': {
      id: '/$enterpriseID/employees/'
      path: '/employees'
      fullPath: '/$enterpriseID/employees'
      preLoaderRoute: typeof EnterpriseIDEmployeesIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/feedback/': {
      id: '/$enterpriseID/feedback/'
      path: '/feedback'
      fullPath: '/$enterpriseID/feedback'
      preLoaderRoute: typeof EnterpriseIDFeedbackIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/feedbacks/': {
      id: '/$enterpriseID/feedbacks/'
      path: '/feedbacks'
      fullPath: '/$enterpriseID/feedbacks'
      preLoaderRoute: typeof EnterpriseIDFeedbacksIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/history/': {
      id: '/$enterpriseID/history/'
      path: '/history'
      fullPath: '/$enterpriseID/history'
      preLoaderRoute: typeof EnterpriseIDHistoryIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/inbox/': {
      id: '/$enterpriseID/inbox/'
      path: '/inbox'
      fullPath: '/$enterpriseID/inbox'
      preLoaderRoute: typeof EnterpriseIDInboxIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/reviews/': {
      id: '/$enterpriseID/reviews/'
      path: '/reviews'
      fullPath: '/$enterpriseID/reviews'
      preLoaderRoute: typeof EnterpriseIDReviewsIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/sent/': {
      id: '/$enterpriseID/sent/'
      path: '/sent'
      fullPath: '/$enterpriseID/sent'
      preLoaderRoute: typeof EnterpriseIDSentIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/_profile/history/': {
      id: '/_profile/history/'
      path: '/history'
      fullPath: '/history'
      preLoaderRoute: typeof ProfileHistoryIndexImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/inbox/': {
      id: '/_profile/inbox/'
      path: '/inbox'
      fullPath: '/inbox'
      preLoaderRoute: typeof ProfileInboxIndexImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/reviews/': {
      id: '/_profile/reviews/'
      path: '/reviews'
      fullPath: '/reviews'
      preLoaderRoute: typeof ProfileReviewsIndexImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/sent/': {
      id: '/_profile/sent/'
      path: '/sent'
      fullPath: '/sent'
      preLoaderRoute: typeof ProfileSentIndexImport
      parentRoute: typeof ProfileImport
    }
    '/$enterpriseID/employees/solved/$employeeId': {
      id: '/$enterpriseID/employees/solved/$employeeId'
      path: '/employees/solved/$employeeId'
      fullPath: '/$enterpriseID/employees/solved/$employeeId'
      preLoaderRoute: typeof EnterpriseIDEmployeesSolvedEmployeeIdImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/inbox/$complaintId/chat': {
      id: '/$enterpriseID/inbox/$complaintId/chat'
      path: '/inbox/$complaintId/chat'
      fullPath: '/$enterpriseID/inbox/$complaintId/chat'
      preLoaderRoute: typeof EnterpriseIDInboxComplaintIdChatImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/sent/$complaintId/chat': {
      id: '/$enterpriseID/sent/$complaintId/chat'
      path: '/sent/$complaintId/chat'
      fullPath: '/$enterpriseID/sent/$complaintId/chat'
      preLoaderRoute: typeof EnterpriseIDSentComplaintIdChatImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/_profile/inbox/$complaintId/chat': {
      id: '/_profile/inbox/$complaintId/chat'
      path: '/inbox/$complaintId/chat'
      fullPath: '/inbox/$complaintId/chat'
      preLoaderRoute: typeof ProfileInboxComplaintIdChatImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/sent/$complaintId/chat': {
      id: '/_profile/sent/$complaintId/chat'
      path: '/sent/$complaintId/chat'
      fullPath: '/sent/$complaintId/chat'
      preLoaderRoute: typeof ProfileSentComplaintIdChatImport
      parentRoute: typeof ProfileImport
    }
    '/$enterpriseID/feedback/$employeeId/': {
      id: '/$enterpriseID/feedback/$employeeId/'
      path: '/feedback/$employeeId'
      fullPath: '/$enterpriseID/feedback/$employeeId'
      preLoaderRoute: typeof EnterpriseIDFeedbackEmployeeIdIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/inbox/$complaintId/': {
      id: '/$enterpriseID/inbox/$complaintId/'
      path: '/inbox/$complaintId'
      fullPath: '/$enterpriseID/inbox/$complaintId'
      preLoaderRoute: typeof EnterpriseIDInboxComplaintIdIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/$enterpriseID/sent/$complaintId/': {
      id: '/$enterpriseID/sent/$complaintId/'
      path: '/sent/$complaintId'
      fullPath: '/$enterpriseID/sent/$complaintId'
      preLoaderRoute: typeof EnterpriseIDSentComplaintIdIndexImport
      parentRoute: typeof EnterpriseIDImport
    }
    '/_profile/inbox/$complaintId/': {
      id: '/_profile/inbox/$complaintId/'
      path: '/inbox/$complaintId'
      fullPath: '/inbox/$complaintId'
      preLoaderRoute: typeof ProfileInboxComplaintIdIndexImport
      parentRoute: typeof ProfileImport
    }
    '/_profile/sent/$complaintId/': {
      id: '/_profile/sent/$complaintId/'
      path: '/sent/$complaintId'
      fullPath: '/sent/$complaintId'
      preLoaderRoute: typeof ProfileSentComplaintIdIndexImport
      parentRoute: typeof ProfileImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren({
  IndexLazyRoute,
  EnterpriseIDRoute: EnterpriseIDRoute.addChildren({
    EnterpriseIDSendComplaintRoute: EnterpriseIDSendComplaintRoute.addChildren({
      EnterpriseIDSendComplaintComplainRoute,
      EnterpriseIDSendComplaintDescribeComplaintRoute,
      EnterpriseIDSendComplaintSendComplaintRoute,
    }),
    EnterpriseIDComplaintSentRoute,
    EnterpriseIDHireRoute,
    EnterpriseIDHiringRoute,
    EnterpriseIDHiringProccesesRoute,
    EnterpriseIDSettingsRoute,
    EnterpriseIDSuccessRoute,
    EnterpriseIDIndexRoute,
    EnterpriseIDEmployeesIndexRoute,
    EnterpriseIDFeedbackIndexRoute,
    EnterpriseIDFeedbacksIndexRoute,
    EnterpriseIDHistoryIndexRoute,
    EnterpriseIDInboxIndexRoute,
    EnterpriseIDReviewsIndexRoute,
    EnterpriseIDSentIndexRoute,
    EnterpriseIDEmployeesSolvedEmployeeIdRoute,
    EnterpriseIDInboxComplaintIdChatRoute,
    EnterpriseIDSentComplaintIdChatRoute,
    EnterpriseIDFeedbackEmployeeIdIndexRoute,
    EnterpriseIDInboxComplaintIdIndexRoute,
    EnterpriseIDSentComplaintIdIndexRoute,
  }),
  ProfileRoute: ProfileRoute.addChildren({
    ProfileSendComplaintRoute: ProfileSendComplaintRoute.addChildren({
      ProfileSendComplaintComplainRoute,
      ProfileSendComplaintDescribeComplaintRoute,
      ProfileSendComplaintSendComplaintRoute,
    }),
    ProfileComplaintSentRoute,
    ProfileHiringInvitationsRoute,
    ProfileProfileRoute,
    ProfileRegisterEnterpriseRoute,
    ProfileSettingsRoute,
    ProfileHistoryIndexRoute,
    ProfileInboxIndexRoute,
    ProfileReviewsIndexRoute,
    ProfileSentIndexRoute,
    ProfileInboxComplaintIdChatRoute,
    ProfileSentComplaintIdChatRoute,
    ProfileInboxComplaintIdIndexRoute,
    ProfileSentComplaintIdIndexRoute,
  }),
  ConfirmationRoute,
  ContactRoute,
  LicensingRoute,
  SignUpRoute,
  AboutLazyRoute,
  ErrorsLazyRoute,
  PrivacyLazyRoute,
  SignInLazyRoute,
})

/* prettier-ignore-end */

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/$enterpriseID",
        "/_profile",
        "/confirmation",
        "/contact",
        "/licensing",
        "/sign-up",
        "/about",
        "/errors",
        "/privacy",
        "/sign-in"
      ]
    },
    "/": {
      "filePath": "index.lazy.tsx"
    },
    "/$enterpriseID": {
      "filePath": "$enterpriseID.tsx",
      "children": [
        "/$enterpriseID/_send-complaint",
        "/$enterpriseID/complaint-sent",
        "/$enterpriseID/hire",
        "/$enterpriseID/hiring",
        "/$enterpriseID/hiring-procceses",
        "/$enterpriseID/settings",
        "/$enterpriseID/success",
        "/$enterpriseID/",
        "/$enterpriseID/employees/",
        "/$enterpriseID/feedback/",
        "/$enterpriseID/feedbacks/",
        "/$enterpriseID/history/",
        "/$enterpriseID/inbox/",
        "/$enterpriseID/reviews/",
        "/$enterpriseID/sent/",
        "/$enterpriseID/employees/solved/$employeeId",
        "/$enterpriseID/inbox/$complaintId/chat",
        "/$enterpriseID/sent/$complaintId/chat",
        "/$enterpriseID/feedback/$employeeId/",
        "/$enterpriseID/inbox/$complaintId/",
        "/$enterpriseID/sent/$complaintId/"
      ]
    },
    "/_profile": {
      "filePath": "_profile.tsx",
      "children": [
        "/_profile/_send-complaint",
        "/_profile/complaint-sent",
        "/_profile/hiring-invitations",
        "/_profile/profile",
        "/_profile/register-enterprise",
        "/_profile/settings",
        "/_profile/history/",
        "/_profile/inbox/",
        "/_profile/reviews/",
        "/_profile/sent/",
        "/_profile/inbox/$complaintId/chat",
        "/_profile/sent/$complaintId/chat",
        "/_profile/inbox/$complaintId/",
        "/_profile/sent/$complaintId/"
      ]
    },
    "/confirmation": {
      "filePath": "confirmation.tsx"
    },
    "/contact": {
      "filePath": "contact.tsx"
    },
    "/licensing": {
      "filePath": "licensing.tsx"
    },
    "/sign-up": {
      "filePath": "sign-up.tsx"
    },
    "/about": {
      "filePath": "about.lazy.tsx"
    },
    "/errors": {
      "filePath": "errors.lazy.tsx"
    },
    "/privacy": {
      "filePath": "privacy.lazy.tsx"
    },
    "/sign-in": {
      "filePath": "sign-in.lazy.tsx"
    },
    "/$enterpriseID/_send-complaint": {
      "filePath": "$enterpriseID/_send-complaint.tsx",
      "parent": "/$enterpriseID",
      "children": [
        "/$enterpriseID/_send-complaint/complain",
        "/$enterpriseID/_send-complaint/describe-complaint",
        "/$enterpriseID/_send-complaint/send-complaint"
      ]
    },
    "/$enterpriseID/complaint-sent": {
      "filePath": "$enterpriseID/complaint-sent.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/hire": {
      "filePath": "$enterpriseID/hire.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/hiring": {
      "filePath": "$enterpriseID/hiring.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/hiring-procceses": {
      "filePath": "$enterpriseID/hiring-procceses.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/settings": {
      "filePath": "$enterpriseID/settings.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/success": {
      "filePath": "$enterpriseID/success.tsx",
      "parent": "/$enterpriseID"
    },
    "/_profile/_send-complaint": {
      "filePath": "_profile/_send-complaint.tsx",
      "parent": "/_profile",
      "children": [
        "/_profile/_send-complaint/complain",
        "/_profile/_send-complaint/describe-complaint",
        "/_profile/_send-complaint/send-complaint"
      ]
    },
    "/_profile/complaint-sent": {
      "filePath": "_profile/complaint-sent.tsx",
      "parent": "/_profile"
    },
    "/_profile/hiring-invitations": {
      "filePath": "_profile/hiring-invitations.tsx",
      "parent": "/_profile"
    },
    "/_profile/profile": {
      "filePath": "_profile/profile.tsx",
      "parent": "/_profile"
    },
    "/_profile/register-enterprise": {
      "filePath": "_profile/register-enterprise.tsx",
      "parent": "/_profile"
    },
    "/_profile/settings": {
      "filePath": "_profile/settings.tsx",
      "parent": "/_profile"
    },
    "/$enterpriseID/": {
      "filePath": "$enterpriseID/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/_send-complaint/complain": {
      "filePath": "$enterpriseID/_send-complaint/complain.tsx",
      "parent": "/$enterpriseID/_send-complaint"
    },
    "/$enterpriseID/_send-complaint/describe-complaint": {
      "filePath": "$enterpriseID/_send-complaint/describe-complaint.tsx",
      "parent": "/$enterpriseID/_send-complaint"
    },
    "/$enterpriseID/_send-complaint/send-complaint": {
      "filePath": "$enterpriseID/_send-complaint/send-complaint.tsx",
      "parent": "/$enterpriseID/_send-complaint"
    },
    "/_profile/_send-complaint/complain": {
      "filePath": "_profile/_send-complaint/complain.tsx",
      "parent": "/_profile/_send-complaint"
    },
    "/_profile/_send-complaint/describe-complaint": {
      "filePath": "_profile/_send-complaint/describe-complaint.tsx",
      "parent": "/_profile/_send-complaint"
    },
    "/_profile/_send-complaint/send-complaint": {
      "filePath": "_profile/_send-complaint/send-complaint.tsx",
      "parent": "/_profile/_send-complaint"
    },
    "/$enterpriseID/employees/": {
      "filePath": "$enterpriseID/employees/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/feedback/": {
      "filePath": "$enterpriseID/feedback/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/feedbacks/": {
      "filePath": "$enterpriseID/feedbacks/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/history/": {
      "filePath": "$enterpriseID/history/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/inbox/": {
      "filePath": "$enterpriseID/inbox/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/reviews/": {
      "filePath": "$enterpriseID/reviews/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/sent/": {
      "filePath": "$enterpriseID/sent/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/_profile/history/": {
      "filePath": "_profile/history/index.tsx",
      "parent": "/_profile"
    },
    "/_profile/inbox/": {
      "filePath": "_profile/inbox/index.tsx",
      "parent": "/_profile"
    },
    "/_profile/reviews/": {
      "filePath": "_profile/reviews/index.tsx",
      "parent": "/_profile"
    },
    "/_profile/sent/": {
      "filePath": "_profile/sent/index.tsx",
      "parent": "/_profile"
    },
    "/$enterpriseID/employees/solved/$employeeId": {
      "filePath": "$enterpriseID/employees/solved/$employeeId.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/inbox/$complaintId/chat": {
      "filePath": "$enterpriseID/inbox/$complaintId/chat.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/sent/$complaintId/chat": {
      "filePath": "$enterpriseID/sent/$complaintId/chat.tsx",
      "parent": "/$enterpriseID"
    },
    "/_profile/inbox/$complaintId/chat": {
      "filePath": "_profile/inbox/$complaintId/chat.tsx",
      "parent": "/_profile"
    },
    "/_profile/sent/$complaintId/chat": {
      "filePath": "_profile/sent/$complaintId/chat.tsx",
      "parent": "/_profile"
    },
    "/$enterpriseID/feedback/$employeeId/": {
      "filePath": "$enterpriseID/feedback/$employeeId/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/inbox/$complaintId/": {
      "filePath": "$enterpriseID/inbox/$complaintId/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/$enterpriseID/sent/$complaintId/": {
      "filePath": "$enterpriseID/sent/$complaintId/index.tsx",
      "parent": "/$enterpriseID"
    },
    "/_profile/inbox/$complaintId/": {
      "filePath": "_profile/inbox/$complaintId/index.tsx",
      "parent": "/_profile"
    },
    "/_profile/sent/$complaintId/": {
      "filePath": "_profile/sent/$complaintId/index.tsx",
      "parent": "/_profile"
    }
  }
}
ROUTE_MANIFEST_END */

export const profileOptions = (descriptorID: string, id: string): { link: string, title: string, unread?: number }[] => {
    // const inbox = await countInboxUnreads(descriptorID, id)
    // const sent = await countSentUnreads(descriptorID, id)
    return [
        {
            link: `/send-complaint`,
            //icon: <SendComplaintIcon />,
            title: "Complain",
        },
        {
            link: `/inbox`,
            //icon: <InboxIcon />,
            title: "Inbox",
            unread: 0
        },
        {
            link: `/sent`,
            //icon: <SentIcon />,
            title: "Sent",
            unread: 0
        },
        {
            link: `/enterprises`,
            //icon: <RegisterEnterpriseIcon />,
            title: "Enterprises",
        },
        // {
        //     link: `/hiring`,
        //     //icon: <HiringInvitationsIcon />,
        //     title: "Hiring",
        // },
        {
            link: `/reviews`,
            //icon: <RateReviewIcon />,
            title: "Reviews",
        },
        {
            link: `/history`,
            //icon: <HistoryIcon />,
            title: "History",
        }
    ]
};

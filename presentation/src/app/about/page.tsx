const About: React.FC = () => {
    return (
        <div className="pt-12 px-3 md:p-12 flex flex-col">
            <p className="text-gray-700 text-md md:text-xl mb-4 p-5">
                We are a complaint management system, where you can receive them as an individual or
                register your enterprise. Registering an enterprise provide you with options to
                monitor the activity.
            </p>
            <p className="text-gray-700 text-md md:text-xl mb-4 p-5">
                Individuals can rate the received attention, it can be from anyone registered at Go Complaint.
                Everyone can receive invitations to enterprises, they can be accepted or rejected,
                and end in a hiring or not.
            </p>
            <p className="text-gray-700 text-md md:text-xl mb-4 p-5">
                All enterprises can search and find users to add to the enterprise user base.
                The fresh hired users can send complaints as the enterprise, while it can track
                users reviews and feedback them so that they can improve.
            </p>
            <p className="text-gray-700 text-md md:text-xl mb-4 p-5 md:mx-auto">
                Go-Complaint it&apos;s our solution.
            </p>
        </div>
    )
}
export default About;
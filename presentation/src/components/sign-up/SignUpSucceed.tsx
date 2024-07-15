const SignUpSucceed: React.FC = () => {
    return (
        <div className="bg-white shadow rounded-md
            my-16 lg:my-24 mx-auto border-t
            max-w-lg self-center py-12 flex flex-col justify-center">
            <h3 className="
                text-gray-700
                text-lg lg:text-xl 
                font-bold mb-4 px-4">
                You have successfully registered at Go Complaint</h3>
            <p className="text-gray-700 text-md lg:text-lg  mb-4 px-4">
                Validation is needed before being able to sign in.
                We sent you an email with a confirmation link to verify your email address
                and validate your account.
            </p>
        </div>
    )
}
export default SignUpSucceed;
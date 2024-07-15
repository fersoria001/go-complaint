import PageProps from "@/app/pageProps";
import PasswordRecoveryForm from "@/components/authentication/PasswordRecoveryForm";
import PasswordRecoverySucceed from "@/components/authentication/PasswordRecoverySucceed";

const PasswordRecovery: React.FC<PageProps> = ({ searchParams }: PageProps) => {
    if (searchParams?.success) {
        return <PasswordRecoverySucceed />
    }
    return (
        <PasswordRecoveryForm />
    )
}
export default PasswordRecovery;
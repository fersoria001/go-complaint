import ExclamationIcon from "../icons/ExclamationIcon";

interface Props {
    errors: string[]
    className?: string
}
const InlineAlert: React.FC<Props> = ({
    errors,
    className = "flex items-center p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50",
}: Props) => {
    return (
        <>
            {
                errors.map((error) => {
                    if (error) {
                        return (
                            <div key={error} className={className}>
                                <ExclamationIcon className="shrink-0 px-0.5 w-6 h-6 lg:w-7 lg:h-7 self-start" fill="#991b1b" />
                                <p>{error}</p>
                            </div >
                        )
                    }
                })
            }
        </>
    )
}
export default InlineAlert;
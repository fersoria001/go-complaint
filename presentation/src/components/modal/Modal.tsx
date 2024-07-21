import Link from "next/link"
import CloseIcon from "../icons/CloseIcon"

interface Props {
    children: React.ReactNode
    className: string
}

const Modal: React.FC<Props> = ({ children, className }: Props) => {
    return (
        <div className={className}>
            <Link href={"/complaints"}>
                <CloseIcon width={20} height={20} className="fill-gray-700 p-0.5 ml-auto"/>
            </Link>
            <div className="flex flex-col h-full">
                {children}
            </div>
        </div>
    )
}

export default Modal
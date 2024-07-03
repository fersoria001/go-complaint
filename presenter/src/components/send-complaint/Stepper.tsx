type Props = { step: number } & typeof defaultProps;
const defaultProps = {
    step: 1,
}
function Stepper({ step }: Props) {
    return (<ol
        className="items-center w-full space-y-4 
sm:flex sm:space-x-8 sm:space-y-0 rtl:space-x-reverse">
        <li
            className={step == 1 ? "flex items-center text-cyan-600 space-x-2.5 rtl:space-x-reverse" :
                "flex items-center text-gray-500 space-x-2.5 rtl:space-x-reverse"
            }>
            <span
                className={step == 1 ? "flex items-center justify-center w-8 h-8 border border-cyan-600 rounded-full shrink-0" :
                    "flex items-center justify-center w-8 h-8 border border-gray-500 rounded-full shrink-0"
                }>
                1
            </span>
            <span>
                <h3
                    className="font-medium leading-tight">Complaint to</h3>
                <p
                    className="text-sm">Choose an individual you want to complaint to</p>
            </span>
        </li>
        <li
            className={step == 2 ? "flex items-center text-cyan-600 space-x-2.5 rtl:space-x-reverse" :
                "flex items-center text-gray-500 space-x-2.5 rtl:space-x-reverse"
            }>
            <span
                className={step == 2 ? "flex items-center justify-center w-8 h-8 border border-cyan-600 rounded-full shrink-0" :
                    "flex items-center justify-center w-8 h-8 border border-gray-500 rounded-full shrink-0"
                }>
                2
            </span>
            <span>
                <h3
                    className="font-medium leading-tight">Describe</h3>
                <p
                    className="text-sm">Write a title and a short description</p>
            </span>
        </li>
        <li
            className={step == 3 ? "flex items-center text-cyan-600 space-x-2.5 rtl:space-x-reverse" :
                "flex items-center text-gray-500 space-x-2.5 rtl:space-x-reverse"
            }>
            <span
                className={step == 3 ? "flex items-center justify-center w-8 h-8 border border-cyan-600 rounded-full shrink-0" :
                    "flex items-center justify-center w-8 h-8 border border-gray-500 rounded-full shrink-0"
                }>
                3
            </span>
            <span>
                <h3
                    className="font-medium leading-tight">Complaint</h3>
                <p
                    className="text-sm">Complaint about it</p>
            </span>
        </li>
    </ol >)
}

export default Stepper;
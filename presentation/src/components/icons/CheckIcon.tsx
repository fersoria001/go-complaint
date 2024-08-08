import IconsProps from "./IconsProps";

const CheckIcon: React.FC<IconsProps> = ({ height = 0, width = 0, fill = "#5f6368", className = "" }: IconsProps) => {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height={`${height}px`}
            viewBox="0 -960 960 960"
            width={`${width}px`}
            className={className}
            fill={fill}>
            <path d="M382-240 154-468l57-57 171 171 367-367 57 57-424 424Z" />
        </svg>
    )
}
export default CheckIcon;



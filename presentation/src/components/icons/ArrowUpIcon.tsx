import IconsProps from "./IconsProps";

const ArrowUpIcon: React.FC<IconsProps> = ({
    height = 32,
    width = 32,
    fill = "#5f6368",
    className = ""
}) => {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height={`${height}px`}
            className={className}
            viewBox="0 -960 960 960"
            width={`${width}px`}
            fill={fill}>
            <path d="M440-160v-487L216-423l-56-57 320-320 320 320-56 57-224-224v487h-80Z" />
        </svg>
    )
}
export default ArrowUpIcon;
import IconsProps from "./IconsProps";

const KeyboardArrowLeftIcon: React.FC<IconsProps> = ({
    onClick = () => { },
    height = 32,
    width = 32,
    fill = "#5f6368",
    className = ""
}) => {
    return (
        <svg
            onClick={onClick}
            xmlns="http://www.w3.org/2000/svg"
            height={`${height}px`}
            className={className}
            viewBox="0 -960 960 960"
            width={`${width}px`}
            fill={fill}>
            <path d="M560-240 320-480l240-240 56 56-184 184 184 184-56 56Z" />
        </svg>
    )
}
export default KeyboardArrowLeftIcon;
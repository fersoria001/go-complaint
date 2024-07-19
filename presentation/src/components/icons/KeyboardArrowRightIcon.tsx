import IconsProps from "./IconsProps";

const KeyboardArrowRightIcon: React.FC<IconsProps> = ({
    onClick = () => { },
    height = 32,
    width = 32,
    fill = "#5f6368",
    className = ""
}) => {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            height={`${height}px`}
            className={className}
            viewBox="0 -960 960 960"
            width={`${width}px`}
            fill={fill}>
            <path d="M504-480 320-664l56-56 240 240-240 240-56-56 184-184Z" />
        </svg>
    )
}
export default KeyboardArrowRightIcon;
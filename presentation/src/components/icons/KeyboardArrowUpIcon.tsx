import IconsProps from "./IconsProps";

const KeyboardArrowUpIcon: React.FC<IconsProps> = ({ onClick = () => { }, height = 32, width = 32, fill = "#5f6368", className = "" }: IconsProps) => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            height={height}
            viewBox="0 -960 960 960"
            width={width}
            className={className}
            fill={fill}>
            <path d="M480-528 296-344l-56-56 240-240 240 240-56 56-184-184Z" />
        </svg>
    )
}
export default KeyboardArrowUpIcon;
import IconsProps from "./IconsProps";

const KeyboardArrowDownIcon: React.FC<IconsProps> = ({ onClick = () => { }, height = 32, width = 32, fill = "#5f6368", className = "" }: IconsProps) => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            height={height}
            viewBox="0 -960 960 960"
            width={width}
            className={className}
            fill={fill}>
            <path d="M480-344 240-584l56-56 184 184 184-184 56 56-240 240Z" />
        </svg>
    )
}
export default KeyboardArrowDownIcon;
import IconsProps from "./IconsProps";

const StarIcon: React.FC<IconsProps> = ({ onClick = () => { }, height = 32, width = 32, fill = "#5f6368", className = "" }: IconsProps) => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg"
            onClick={onClick}
            height={height}
            viewBox="0 -960 960 960"
            width={width}
            className={className}
            fill={fill}>
            <path d="m233-120 65-281L80-590l288-25 112-265 112 265 288 25-218 189 65 281-247-149-247 149Z" />
        </svg>
    )
}
export default StarIcon;